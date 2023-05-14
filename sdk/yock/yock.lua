---@meta _


---@param tbl table
function pull(tbl)
end

---@param plugin string
---@return string,string
function parse_plugin(plugin)
    return "", ""
end

---@param path string
---@return boolean
function is_url(path)
    return false
end

---@param file string
---@return string
function export_builder(file)
    return ""
end

---@class pluginlist
plugin_list = {}

---@param path string
---@return boolean
function plugin_list:IsExist(path)
    return false
end

---@param pid string
---@param path string
function plugin_list:AddPlugin(pid, path)

end

---@class exportOpt
---@field update fun()
---@field install fun()
---@field uninstall fun()
---@field init fun(env: any)
local exportOpt = {}

---
---{{.export}}
---
---@param opt exportOpt
---
function export(opt)

end

---
---{{.optional}}
---
---@param cases table
---@param bad_case fun()
function optional(cases, bad_case)

end

---
---{{.case}}
---
function case(...)

end

---
---@class rmOpt
---@field safe boolean
---@field pattern string
---
---{{.rm_opt}}
---
---
local rmOpt = {}

---
--- {{.rm}}
---
---@param opt rmOpt
---@vararg string
---
function rm(opt, ...)

end

---
---{{.driver}}
---
---@param callback fun(...): ...:any
---
function driver(callback)

end

---
---{{.uninit_driver}}
---
---@param fn string
---@return fun(opt: table, ...:any)
---
function uninit_driver(fn)
    ---@param opt table
    ---@vararg any
    return function(opt, ...)

    end
end

---
---{{.set_driver}}
---
---@param driver string
---@param name string
---@return string
---
function set_driver(driver, name)
    return ""
end

---
---{{.exec_driver}}
---
---@param did string
---@param opt table
---@vararg any
---
function exec_driver(did, opt, ...)

end

---
---{{.unzip}}
---
---@param opt table
---@vararg string
---
function unzip(opt, ...)

end

---
---{{.table_dump}}
---
---@param tbl table
---
function table.dump(tbl)

end

---
---{{.exec}}
---
---@param opt table
---@vararg string
---
function exec(opt, ...)

end

---
---@class platform
---@field OS string
---@field Ver string
---@field Arch string
---
---{{.platform}}
---
local platform = {}

---
---{{.env}}
---
---@class env
---@field args table
---@field platform platform
---@field flags table
---@field job string
---@field workdir string
---
env = {}

--
---{{.cmdf}}
--
---@vararg string
---
function cmdf(...)

end

---
---{{.ldns}}
---
---@class ldns
---
ldns = {}

---
---{{.lsdn_get_driver}}
---
---@param domain string
---
function ldns:GetDriver(domain)

end

---
---{{.ldns_get_plugin}}
---
---@param domain string
---
function ldns:GetPlugin(domain)

end

---
---{{.lsdn_put_plugin}}
---
---@param domain string
---@param url string
---@param path string
---
function ldns:PutPlugin(domain, url, path)

end

---
---{{.lsdn_put_driver}}
---
---@param domain string
---@param url string
---@param path string
---
function ldns:PutDriver(domain, url, path)

end

---
---{{.lsdn_alias_driver}}
---
---@param domain string
---@param alias string
---
function ldns:AliasDriver(domain, alias)

end

---
---{{.lsdn_alias_plugin}}
---
---@param domain string
---@param alias string
---
function ldns:AliasPlugin(domain, alias)

end

---
---{{.gdns}}
---
---@class gdns
---
gdns = {}

---
---{{.gdns_get_driver}}
---
---@param domain string
---
function gdns:GetDriver(domain)

end

---
---{{.gdns_get_plugin}}
---
---@param domain string
---@return table
---
function gdns:GetPlugin(domain)
    return {}
end

---@param domain string
---@param url string
---@param path string
function gdns:UpdatePlugin(domain, url, path)

end

---
---{{.go}}
---
---@param callback fun()
---@async
---
function go(callback)

end

---
---{{.wait}}
---
---@param sig string
---
function wait(sig)

end

---
---{{.waits}}
---
---@vararg string
---
function waits(...)

end

---
---{{.notify}}
---
---@param sig string
---
function notify(sig)

end

---
---{{.http}}
---
---@param opt table
---@vararg string
---
function http(opt, ...)

end

---
---{{.pathf}}
---
---@param path string
---@return string
---
function pathf(path)
    return ""
end

---
---{{.path}}
---
---@class path
---
path = {}

---
---{{.path_filename}}
---
---@param filepath string
---@return string
---
function path.filename(filepath)
    return ""
end

---
---{{.path_exist}}
---
---@param filepath string
---@return boolean
---
function path.exist(filepath)
    return false
end

---@vararg string
---@return string
function path.join(...)
    return ""
end

---@param path string
---@return string
function path.dir(path)
    return ""
end

---@param path string
---@return string
function path.base(path)
    return ""
end

---@param path string
---@return string
function path.clean(path)
    return ""
end

---@param path string
---@return string
function path.ext(path)
    return ""
end

---@param path string
---@return string, string
function path.abs(path)
    return "", ""
end

---
---@class random
---
---{{.random}}
---
random = {}

---
---{{.random_str}}
---
---@return string
---
function random.str()
    return ""
end

---
---{{.job}}
---
---@param name string
---@param callback fun(cenv: table):boolean
---
function job(name, callback)

end

---
---{{.jobs}}
---
---@param name string
---@vararg string
---
function jobs(name, ...)

end

---
---{{.job_option}}
---
---@param opt table
---
function job_option(opt)

end

---
---{{.time}}
---
---@class time
---@field microsecond number
---@field millisecond number
---@field second number
---
time = {}

---
---{{.time_sleep}}
---
---@param sec number
---
function time.sleep(sec)

end

---
---@class waitGroup
---
---{{.wait_group}}
---
local waitGroup = {}

---
---{{.wait_group_add}}
---
---@param delta number
---
function waitGroup:Add(delta)

end

---
---{{.wait_group_done}}
---
function waitGroup:Done()

end

---
---{{.wait_group_wait}}
---
function waitGroup:Wait()

end

---
---{{.sync}}
---
sync = {}

---
---{{.sync_new}}
---
---@return waitGroup
---
function sync.new()
    return {}
end

---@return boolean
function Windows()
    return false
end

---@return boolean
function Darwin()
    return false
end

---@return boolean
function Linux()
    return false
end

---
---@param want_os string
---@param want_ver string
---@return boolean
function OS(want_os, want_ver)
    return false
end

---@class flag_type
---@field string_type number
---@field number_type number
---@field array_type number
---@field bool_type number
flag_type = {}

---@param env env
---@param todo table
function parse_flags(env, todo)

end

---@class strings
strings = {}

---@param s string
---@param prefix string
---@return boolean
function strings.HasPrefix(s, prefix)
    return false
end

---@param s string
---@param suffix string
---@return boolean
function strings.HasSuffix(s, suffix)
    return false
end

---@param s string
---@param sep string
---@return string, string, boolean
function strings.Cut(s, sep)
    return "", "", false
end

---@param s string
---@param prefix string
---@return string, boolean
function strings.CutPrefix(s, prefix)
    return "", false
end

---@param s string
---@param suffix string
---@return string, boolean
function strings.CutSuffix(s, suffix)
    return "", false
end

---@param s string
---@param substr string
---@return boolean
function strings.Contains(s, substr)
    return false
end

---@param s string
---@param chars string
---@return boolean
function strings.ContainsAny(s, chars)
    return false
end

---@param s string
---@param r string
---@return boolean
function strings.ContainsRune(s, r)
    return false
end

---@param s string
---@param substr string
---@return number
function strings.Count(s, substr)
    return 0
end

---@param s string
---@param old string
---@param new string
---@param n number
---@return string
function strings.Replace(s, old, new, n)
    return ""
end

---@param s string
---@param old string
---@param new string
---@return string
function strings.ReplaceAll(s, old, new)
    return ""
end

---@param s string
---@return string
function strings.Clone(s)
    return ""
end

---@param a string
---@param b string
---@return number
function strings.Compare(a, b)
    return 0
end

---@param url string
---@return boolean
function is_localhost(url)
    return false
end

---@class sshClient
local sshClient = {}

---@param cmds table
function sshClient:Exec(cmds)

end

function sshClient:Shell()

end

---@class sshOpt
---@field user string
---@field pwd string
---@field ip string
---@field network string
---@field redirect boolean
local sshOpt = {}

---@param opt sshOpt
---@param cb fun(client: sshClient)
---@return sshClient
function ssh(opt, cb)
    return {}
end

---@param dir string
function mkdir(dir)

end

---@param src string
---@param dst string
function cp(src, dst)

end

---@param src string
---@param dst string
function mv(src, dst)

end

---@param opt table
---@vararg string
function installs(opt, ...)

end

---@param plugin string
---@param opt table
function install(plugin, opt)

end

---@param file string
---@return string
function load_plugin(file)
    return ""
end

plugins = {}

---@param opt table
function plugin(opt)

end

---@return string, boolean
function pwd()
    return "", false
end

---@param module string
---@return any
function import(module)

end
