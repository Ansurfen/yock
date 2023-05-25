parse_flags(env, {
    wd = flag_type.string_type,
    g = flag_type.bool_type,
    m = flag_type.string_type
})

ypm:open()
local module = env.flags["m"]
local g = env.flags["g"]
local wd = env.flags["wd"]
local yock_json = path.join(wd, "modules.json")
local tmp_path = path.join(env.workdir, "..", "yock_tmp")
local yock_modules = path.join(wd, "yock_modules")

if g then
    yock_modules = path.join(env.workdir, "..", "yock_modules")
end

if not is_exist(yock_json) then
    safe_write(yock_json, [[{"dependency":{}}]])
end

local fp

if g then
    fp = ypm.modules
else
    fp = jsonfile:open(yock_json)
    mkdir(yock_modules)
end

if #module > 0 then
    fp.buf.dependency[module] = ""
    fp:write()
end

for name, version in pairs(fp.buf["dependency"]) do
    local file = fetch.file(ypm:get(name))
    print(file)
    local boot = import(path.join(tmp_path, file))
    local todo = {
        version = version,
        modules_path = yock_modules
    }
    boot(todo)
    if #module > 0 and #fp.buf.dependency[module] == 0 then
        fp.buf.dependency[module] = todo["version"]
        fp:write()
    end
end

ypm:close()
