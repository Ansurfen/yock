argsparse(env, {
    repo = flag_type.bool
})

local repo = env.flags["repo"] or false

sh("ypm install")

if env.platform.OS == "windows" then
    local ark = import("opencmd/ark")
    ark.install("golang@1.20", "chocolatey", "git", "mingw")

    ---@type chocolatey
    local choco = import("opencmd/installer/chocolatey")

    ---@type conan
    local conan = import("opencmd/pm/conan")

    if not conan.exist() then
        choco.install("conan")
    end

    ---@type conanfile
    local conanfile = import("opencmd/pm/conan/conanfile")

    local c = conanfile.create("conanfile.txt")
    c:add_deps("libffi/3.4.4")
    c:add_gen("CMakeDeps", "CMakeToolchain")
    c:save()

    conan.install({
        conanfile = ".",
        build = "missing",
        ["output-folder"] = "build"
    })
else
    ---@type installer
    local ins
    local todo = {}

    ---@type apt
    local apt = import("opencmd/installer/apt")
    if ins == nil and apt.exist() then
        ---@diagnostic disable-next-line: cast-local-type
        ins = apt
        todo = { "libffi8", "libffi-dev" }
    end

    if ins == nil then
        yassert("no found matched installer")
    else
        for _, pack in ipairs(todo) do
            ins.install(pack)
        end
    end
end

if repo then
    ---@type git
    local git = import("opencmd/git")

    if not find("tmp") then
        git.clone("git@github.com:Ansurfen/yock.git", "tmp")
        cd("tmp")
        git.checkout("main")
        cd("..")
    end

    cd("tmp/ctl")
else
    cd("../ctl")
end

if find(pathf("$/../scheduler/yockf.go.txt")) then
    rename(pathf("$/../scheduler/yockf.go.txt"), pathf("$/../scheduler/yockf.go"))
end

---@type golang
local golang = import("opencmd/lang/golang")

golang.mod.tidy()
golang.env("GOOS", env.platform.OS)
golang.build({
    packages = { "." },
    output = wrapexf("yock"),
    ldflags = {
        X = {
            ["github.com/ansurfen/yock/util.YockBuild"] = "release",
            ["github.com/ansurfen/yock/util.YockVersion"] = import("opencmd/lang/yock").version(),
        }
    },
})
