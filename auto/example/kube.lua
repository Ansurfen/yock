yassert(env.platform.OS == "linux")
argsparse(env, {
    master = flag_type.bool,
    id = flag_type.num,
})

local master = env.flags["master"] or false
yassert(not master and env.flags["id"] == nil)

local hostname
if master then
    hostname = "k8s-master"
else
    hostname = string.format("k8s-node%d", env.flags["id"])
end

---@type installer
local yum = import("opencmd/installer/yum")

for _, pack in ipairs({ "curl", "wget", "systemd", "bash-completion", "lrzsz" }) do
    yum.install(pack)
end

timedatectl.set({
    timezone = "Asia/Shanghai",
    ["local-rtc"] = 0
})
-- timedatectl.show()
systemctl.restart("rsyslog")
systemctl.restart("crond")
hostnamectl.set(hostname)

_ = file("/etc/hosts") <= stream([[
192.168.56.101    k8s-master
192.168.56.102    k8s-node1
192.168.56.103    k8s-node2]])

systemctl.disable("firewalld.service")
systemctl.stop("firewalld.service")

local modprobe = import("opencmd/modprobe")
modprobe("overlay", "br_netfilter")

_ = file("/etc/modules-load.d/k8s.conf") <= stream([[
overlay
br_netfilter
]])

_ = file("/etc/sysctl.d/k8s.conf") <= stream([[
net.bridge.bridge-nf-call-iptables=1
net.bridge.bridge-nf-call-ip6tables=1
net.ipv4.ip_forward=1
]])

sh([[
sysctl --system
# 通过运行以下指令确认 br_netfilter 和 overlay 模块被加载
lsmod | egrep 'overlay|br_netfilter'
# 通过运行以下指令确认 net.bridge.bridge-nf-call-iptables、net.bridge.bridge-nf-call-ip6tables 系统变量在你的 sysctl 配置中被设置为 1
sysctl net.bridge.bridge-nf-call-iptables net.bridge.bridge-nf-call-ip6tables net.ipv4.ip_forward
]])

yum.install("yum-utils")

sh("yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo")

mkdir("/etc/docker")

_ = file("/etc/docker/daemon.json") < stream(json.encode({
    ["exec-opts"] = { "native.cgroupdriver=systemd" },
    ["log-driver"] = "json-file",
    ["log-opts"] = {
        ["max-size"] = "100m",
    },
    ["storage-driver"] = "overlay2",
    ["storage-opts"] = {
        "overlay2.override_kernel_check=true"
    },
    ["registry-mirrors"] = { "https://hub-mirror.c.163.com", "https://docker.mirrors.ustc.edu.cn",
        "https://registry.docker-cn.com" },
}))

sh([[
yum makecache fast
yum install -y docker-ce-20.10.23 docker-ce-cli-20.10.23 containerd.io
systemctl daemon-reload
]])

systemctl.enable("docker")
systemctl.restart("docker")

sh([[
swapoff -a && sed -ri 's/.*swap.*/#&/' /etc/fstab
setenforce 0 && sed -i 's/^SELINUX=enforcing$/SELINUX=permissive/' /etc/selinux/config
]])

-- TODO: ini.encode() / ini.decode()
_ = file("/etc/yum.repos.d/kubernetes.repo") <= stream([[
[kubernetes]
name=Kubernetes
baseurl=https://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64/
enabled=1
gpgcheck=0
repo_gpgcheck=0
gpgkey=https://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg https://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
]])


for _, pack in ipairs({ "kubelet-1.23.17", "kubeadm-1.23.17", "kubectl-1.23.17" }) do
    ---@diagnostic disable-next-line: redundant-parameter
    yum.install(pack, "--disableexcludes=kubernetes")
end

_ = file("/etc/sysconfig/kubelet") <= stream([[
KUBELET_EXTRA_ARGS="--cgroup-driver=systemd"
]])

sh([[
crictl config runtime-endpoint unix:///var/run/containerd/containerd.sock
crictl config image-endpoint unix:///var/run/containerd/containerd.sock
sed -i '/KUBELET_KUBEADM_ARGS/s/"$/ --container-runtime=remote --container-runtime-endpoint=unix:\/\/\/run\/containerd\/containerd.sock"/' /var/lib/kubelet/kubeadm-flags.env]])

sed({

})

systemctl.enable("kubelet") -- systemctl enable --now kubelet

local info, err = systemctl.status("kubelet")
yassert(err == nil)
table.dump(info)

sh("journalctl -xe")

mkdir("/k8sdata/log/")

local kubeadm = import("opencmd/kube/adm")

_ = file("/k8sdata/log/kubeadm-init.log") <= stream(kubeadm.init({
    ["apiserver-advertise-address"] = "192.168.56.101",
    ["image-repository"] = "registry.cn-hangzhou.aliyuncs.com/google_containers",
    ["kubernetes-version"] = "v1.23.17",
    ["service-cidr"] = "10.96.0.0/12",
    ["pod-network-cidr"] = "10.244.0.0/16"
}))

mkdir(pathf("@/../.kube"))

cp("/etc/kubernetes/admin.conf", pathf("@/../.kube/config"))

mkdir("/k8sdata/network")

sh([[
chown "$(id -u)":"$(id -g)" "$HOME"/.kube/config

wget --no-check-certificate -O /k8sdata/network/flannelkube-flannel.yml https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml
kubectl create -f /k8sdata/network/flannelkube-flannel.yml

wget --no-check-certificate -O /k8sdata/network/calico.yml https://docs.projectcalico.org/manifests/calico.yaml
kubectl create -f /k8sdata/network/calico.yml

! grep -q kubectl "$HOME/.bashrc" && echo "source /usr/share/bash-completion/bash_completion" >>"$HOME/.bashrc"
! grep -q kubectl "$HOME/.bashrc" && echo "source <(kubectl completion bash)" >>"$HOME/.bashrc"
! grep -q kubeadm "$HOME/.bashrc" && echo "source <(kubeadm completion bash)" >>"$HOME/.bashrc"
! grep -q crictl "$HOME/.bashrc" && echo "source <(crictl completion bash)" >>"$HOME/.bashrc"
source "$HOME/.bashrc"
]])
