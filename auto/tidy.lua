local bin = {
    ripgrep = {
        windows = {
            ["*"] = "13.0.0/ripgrep-13.0.0-x86_64-pc-windows-gnu.zip"
        },
        linux = {
            ["*"] = "13.0.0/ripgrep-13.0.0-x86_64-unknown-linux-musl.tar.gz"
        },
        darwin = {
            ["*"] = "13.0.0/ripgrep-13.0.0-x86_64-apple-darwin.tar.gz"
        },
        url = "https://github.com/BurntSushi/ripgrep/releases/download/"
    },
    goawk = {
        windows = {
            i386 = "v1.23.3/goawk_v1.23.3_windows_386.zip",
            amd64 = "v1.23.3/goawk_v1.23.3_windows_amd64.zip",
        },
        linux = {
            i386 = "v1.23.3/goawk_v1.23.3_linux_386.tar.gz",
            amd64 = "v1.23.3/goawk_v1.23.3_linux_amd64.tar.gz",
            arm64 = "v1.23.3/goawk_v1.23.3_linux_arm64.tar.gz"
        },
        darwin = {
            arm64 = "v1.23.3/goawk_v1.23.3_darwin_arm64.tar.gz",
            amd64 = "v1.23.3/goawk_v1.23.3_darwin_amd64.tar.gz"
        },
        url = "https://github.com/benhoyt/goawk/releases/download/"
    },
    sd = {
        windows = {
            ["*"] = "v0.7.5/sd.0.7.5-.x86_64-pc-windows-msvc.zip"
        },
        linux = {
            ["*"] = "v0.7.6/sd-v0.7.6-x86_64-unknown-linux-gnu",
        },
        darwin = {
            ["*"] = "v0.7.6/sd-v0.7.6-x86_64-apple-darwin"
        },
        url = "https://github.com/chmln/sd/releases/download/"
    }
}

for name, todo in pairs(bin) do
    local fn = todo[env.platform.OS]
    if fn ~= nil then
        local target
        target = fn[env.platform.Arch]
        if target == nil then
            if fn["*"] ~= nil then
                target = fn["*"]
            else
                yassert(name .. " no support the platform")
            end
        end
        target = todo["url"] .. target
        local suffix
        if env.platform.OS == "windows" then
            suffix = ".zip"
        else
            suffix = ".tar.gz"
        end
        local fd = fetch.file(target, suffix)
        uncompress(path.join(env.yock_tmp, fd), "../bin")
    end
end
