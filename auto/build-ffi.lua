sh([[ypm install github.tag/ansurfen/OpenCmd@0.0.1
ypm install github.release/ansurfen/ark@0.0.1]])

local ark = import("OpenCmd/ark")
ark.install("golang@1.20", "conan", "git", "gcc")

---@type conanfile
local conanfile = import("OpenCmd/conan/conanfile")

local c = conanfile.create("conanfile.txt")
c:add_deps("libffi/3.4.4")
c:add_gen("CMakeDeps", "CMakeToolchain")
c:save()

---@type conan
local conan = import("OpenCmd/conan")

conan.install({
    conanfile = ".",
    build = "missing",
    ["output-folder"] = "build"
})

---@type git
local git = import("OpenCmd/git")

git.clone("git@github.com:Ansurfen/yock.git", "tmp")
cd("tmp")
git.checkout("main")

rename(pathf("$/scheduler/yockf.go.txt"), pathf("$/scheduler/yockf.go"))

---@type golang
local golang = import("OpenCmd/golang")

golang.mod.tidy()
golang.env("GOOS", env.platform.OS)
golang.build({
    packages = { "." },
    output = wrapexf("yock"),
    ldflags = {
        X = {
            ["github.com/ansurfen/yock/util.YockBuild"] = "release",
            ["github.com/ansurfen/yock/util.YockVersion"] = import("OpenCmd/yock").version(),
        }
    },
})
