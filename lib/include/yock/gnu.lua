---@diagnostic disable: duplicate-doc-field
-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---@alias echo_mode string
---|> "c" # create a new file if none exists.
---| "t" # truncate regular writable file when opened.
---| "r" # open the file read-only.
---| "w" # open the file write-only.
---| "rw" # open the file read-write.
---| "a" # append data to the file when writing.
---| "e" # used with `c`, file must not exist.
---| "s" # open for synchronous I/O.

---@class echo_opt
---@field mode? echo_mode # indicates how files are opened
---@field fd? string[] # fd is short for file descriptor, which used for indicating where stream outputs. You can use filename as fd to write file, or print terminal by stdout, stderr.

---echo prints variable string argument on terminal and
---returns an array that saves every result of print.
---Except primitive string, you also print environment variable
---corresponding value.
---### Example:
---```lua
---local data, err = echo("Hello", "World")
---yassert(#data == 2 and not err)
---
---local data = echo("$Path")
---if #data > 0 then
--- print("Path: ", data[1]) -- effect like environ("Path")
---end
---```
---@vararg string
---@return string[], err
function echo(...) end

---echo prints variable string argument on terminal and
---returns an array that saves every result of print.
---Except primitive string, you also print environment variable
---corresponding value.
---### Option:
---* mode, echo_mode (string), indicates how files are opened
---* fd, string[], fd is short for file descriptor, which used for indicating where stream outputs. You can use filename as fd to write file, or print terminal by stdout, stderr.
---
---### Example:
---```lua
---# append write
---echo({ fd = { "stdout", "test.txt" }, mode = "c|a|rw" }, "Hello World!")
---
---# truncate write
---echo({ fd = { "stdout", "test.txt" }, mode = "c|t|rw" }, "Hello World!")
---```
---@param opt echo_opt
---@vararg string
---@return string[], err
function echo(opt, ...) end

---whoami returns hostname
---@return string, err
function whoami() end

---clear clears outputs on terminal.
function clear() end

---cd changes the current working directory to the named directory.
---@param dir string
---@return err
function cd(dir) end

---touch creates an empty file when file isn't exist.
---@param file string
---@return err
function touch(file) end

---cat reads content from specified file.
---@param file string
---@return string, err
function cat(file) end

---ls lists the information of directory or file according to specified directory.
---
---`NOTE`: results to be returned by ls isn't like other gun command. In order to
---save memory, it's set array format to store information. You can see detail in
---the following.
---
---### Format of info:
---[1] permission, string, e.g. -rwxrwxrwx, -rw-rw-rw-
---
---[2] size, number
---
---[3] mod_time, string, e.g. Aug  6 15:26
---
---[4] filename, string
---@param dir string
---@return string[][], err
function ls(dir) end

---ls recurses given directory, and can set
---callback that recives visited path and 
-- path's information while walking dir.
---### Example:
---```lua
---ls(".", function(path, info)
---     print(path)
---end)
---```
---@param dir string
---@param callback fun(path: string, info: fileinfo)
---@return nil, err
function ls(dir, callback) end

---chmod changes the mode of the named file to mode.
---If the file is a symbolic link, it changes the mode of the link's target.
---@param name string
---@param mode number
---@return err
function chmod(name, mode) end

---chown changes the numeric uid and gid of the named file.
---If the file is a symbolic link, it changes the uid and gid of the link's target.
---A uid or gid of -1 means to not change that value.
---@param name string
---@param uid number
---@param gid number
---@return err
function chown(name, uid, gid) end

---mkdir recurses to create directory.
---@vararg string
---@return err
function mkdir(...) end

---cp copies file or directory from src to dst.
---
---`NOTE`: It isn't support recurse, and see other overload
---when you want to do it.
---### Example:
---```lua
---cp("a", "b")
---```
---@param src string
---@param dst string
function cp(src, dst) end

---@class cp_opt
---@field recurse? boolean # recurses to copy specified file or directory
---@field force? boolean # covers file with the same name when set true

---cp copies file or directory from src to dst
---### Option:
---* recurse?, boolean (default false), recurses to copy specified file or directory
---* force?, boolean (default false), covers file with the same name when set true
---
---### Example:
---```lua
---cp({ recurse = true }, {
---     ["a"] = "b",
---     ["c"] = "d",
---}
---```
---@param opt cp_opt
---@param path table<string, string>
function cp(opt, path) end

---mv moves directory or file from src to dst, supporting recurse
---@param src string
---@param dst string
function mv(src, dst) end

---pwd returns working directory for current program
---@return string, err
function pwd() end

---@class rm_opt
---@field safe? boolean # is the same with recurse field. Please use recurse instead of it, and it'll be deprecated in the future.
---@field pattern? string # remove directory or file to be matched if pattern's length is more than 0, and obeys golang's regular expressions.
---@field recurse? boolean # recurses to remove directory when set true

---rm removes specified directories or files, and has one function
---overload of which capacity just like rmdir command on bash.
---### Option
---* recurse?, boolean (default false), recurses to remove directory when set true
---* pattern?, string, remove directory or file to be matched if pattern's length is more than 0, and obeys golang's regular expressions.
---* safe?, boolean (default false), is the same with recurse field. Please use recurse instead of it, and it'll be deprecated in the future.
---
---### Example:
---```lua
---# just like rmdir, which only removes empty directory or single files
---rm("/a", "/b")
---
---# delete file with recuse
---rm({ recurse = true }, "/a", "/b")
---
---# remove with recurse and pattern
---rm({ recurse = true, pattern = ".exe$" }, "/a")
---```
---@param opt rm_opt
---@vararg string
---@return err
function rm(opt, ...) end

---rm removes empty directories or single files to be specified,
---and just like rmdir command on bash. If you want to remove
---directory with recurse, see its function overload.
---### Example:
-- ```lua
-- rm("/a", "/b")
-- ```
---@vararg string
---@return err
function rm(...) end

---rename resets filename from old to new
---@param old string
---@param new string
function rename(old, new) end

---sudo runs command with administrator permission
---@param cmd string
function sudo(cmd) end

---@class grep_opt
---@field case? boolean # determine whether case sensitivity ignored
---@field color? string # set the color format for output
---@field pattern string # indicates a pattern to match string
---@field file? string[] # searches matched results from files and its priority is more than str field, which means str field will be unavailable when set file field
---@field str? string[] # searches matched results from string array and is unavailable when set file field, no supporting string with line break ('\n')

---grep binds for [ripgrep](https://github.com/BurntSushi/ripgrep) to implement cross platform,
---which means it's different with native grep on bash, and can search string with fast, easy, convenient.
---### Option:
---* case?, boolean (default false), determine whether case sensitivity ignored
---* color?, string, set the color format for output
---* pattern, string, indicates a pattern to match string
---* file?, string[], searches matched results from files and its priority is more than str field, which means str field will be unavailable when set file field
---* str?, string[], searches matched results from string array and is unavailable when set file field, no supporting string with line break ('\n')
---
---### Example:
---```lua
---# queries according to string
---local res, err = grep({
---    pattern = "abc",
---    str = { "abcd", "bcd", "abbc" }
---})
---yassert(err)
---table.dump(strings.Split(res, "\n"))
---
---# queries files
---write("./test.txt", "get\n get abc\n getGeT\nGET")
---local res, err = grep({
---    case = true,
---    color = "never",
---    pattern = "get",
---    file = { "./test.txt" }
---})
---yassert(err)
---print(res)
---```
---@param opt grep_opt
---@return string, err
function grep(opt) end

---@class awk_opt
---@field prog string|string[] # indicates single or multiple rules to extract string. `NOTE`: the single rule (string) only supports to write explicit prog string, but multiple rules (string[]) only support loading from .awk files.
---@field var? table<string, string|number|integer> # defines keyed variable and can use its in prog through key.
---@field file? string[] # extracts matched results from files and its priority is more than str field, which means str field will be unavailable when set file field
---@field str? string[] # extracts matched results from string array and is unavailable when set file field, no supporting string with line break ('\n')

---awk binds for [goawk](https://github.com/benhoyt/goawk) to implement cross platform,
---which means it's different with native awk on bash, and can handle or extract string
---with fast, easy, convenient.
---### Option:
---* prog, string|string[], indicates single or multiple rules to extract string. `NOTE`: the single rule (string) only supports to write explicit prog string, but multiple rules (string[]) only support loading from .awk files.
---* var?, table<string, any>, defines keyed variable and can use its in prog through key.
---* file?, string[], extracts matched results from files and its priority is more than str field, which means str field will be unavailable when set file field.
---* str?, string[], extracts matched results from string array and is unavailable when set file field, no supporting string with line break ('\n')
---
---### Example:
---```lua
---# extracts from string
---local new, err = awk({
---    prog = "{ print $1 + $3 }",
---    str = { "1 2 3" }
---})
---yassert(err)
---table.dump(strings.Split(new, "\n"))
---
---# extracts and tests to define variable
---local new, err = awk({
---    prog = "{ print $1, name }",
---    str = { "'Hello World'" },
---    var = {
---        name = "yock"
---    }
---})
---yassert(err)
---table.dump(strings.Split(new, "\n"))
---
---# extracts based-on rule from prog files
---local res, err = awk({
---    prog = {
---        "./rule.awk",
---        "./rule2.awk"
---    },
---    file = {
---        "./test.txt"
---    },
---    var = {
---        name = "yock",
---        age = 20
---    }
---})
---yassert(err)
---print(res)
---```
---@param opt awk_opt
---@return string, err
function awk(opt) end

---@class sed_opt
---@field old string # old string to be replaced
---@field new string # new string that replace old string
---@field file? string[] # replaces matched results from files and its priority is more than str field, which means str field will be unavailable when set file field
---@field str? string[] # replaces matched results from string array and is unavailable when set file field, no supporting string with line break ('\n')

---sed binds for [sd](https://github.com/chmln/sd)  to implement cross platform,
---which means it's different with native sed on bash, and can handle or replace
---string with fast, easy, convenient.
---### Option:
---* old, string, old string to be replaced
---* new, string, new string that replace old string
---* file?, string[], replaces matched results from files and its priority is more than str field, which means str field will be unavailable when set file field
---* str? string[], replaces matched results from string array and is unavailable when set file field, no supporting string with line break ('\n')
---
---### Example:
---```lua
---# replaces content in string
---local res, err = sed({
---    old = "((([])))",
---    new = " ",
---    str = { "lots((([]))) of special chars" }
---})
---yassert(err)
---print(res)
---
---# replaces content in files
---res, err = sed({
---    old = "(.*)",
---    new = "//$1",
---    file = { "./test.txt" },
---})
---print(out, err)
---```
---@param opt sed_opt
---@return string, err
function sed(opt) end

---@class find_opt
---@field pattern? string # indicates rule to match directories or files, and writing format same as golang's regular expressions.
---@field dir? boolean # match directory when set true
---@field file? boolean # match file when set true

---find scans specified directory with recurse and returns
---matched results.
---### Option:
---* pattern?, string, indicates rule to match directories or files, and writing format same as golang's regular expressions.
---* dir?, boolean (default true), match directory when set true
---* file?, boolean (default true), match file when set true
---
---### Example:
---```lua
-- find({
--     dir = false,
--     pattern = "\\.lua"
-- }, "/script")
---```
---@param opt find_opt
---@param path string
---@return table, err
function find(opt, path) end

---find returns whether file or directory exist according to path
---@param path string
---@return boolean
function find(path) end

---whereis returns absolute path for given environment variable.
---@param k string
---@return string, err
function whereis(k) end

---alias takes an alias v for k, and it isn't directly call alias command
---on terminal or shell but saves it in the program's memory for mapping
---commands the `sh` function call.
---### Example:
---```lua
---alias("CC", "go")
---sh("$CC version")
---```
---@param k string
---@param v string
function alias(k, v) end

---unalias remove mapping relationship from alias.
---### Example:
---```lua
---alias("CC", "go")
---unalias("CC")
---sh("$CC version")
---```
---@vararg string
function unalias(...) end

---nohup launches backend process hidden window.
---@param cmd string
---@return err
function nohup(cmd) end

---@class pgrep_info
---@field name string
---@field pid integer

---pgrep returns results according to process's name.
---@param name string
---@return pgrep_info[]
function pgrep(name) end

---@class ps_info
---@field name string
---@field cmd string
---@field cpu? number
---@field start? number
---@field mem? any
---@field user? string

---@class ps_opt
---@field user? boolean # includes process's launcher when set true
---@field cpu? boolean # includes cpu usage ratio at calling moment when set true
---@field time? boolean # includes process's start time when set true
---@field mem? boolean # includes process's memory usage ratio when set true

---ps lists all process state when opt is nil.
---It's worthy of noting that there only are
---cmd (command) and name (process's name) field
---in default. If you want to get detailed info,
---try to use ps_opt (table) format to make it.
---
---Except above two method introduced, there are
---two way to query, and the one is indicated
---pid, and the other passes by string to do
---fuzzy matching according to cmd.
---
---### Option:
---* user?, boolean (default false), includes process's launcher when set true
---* cpu?, boolean (default false) , includes cpu usage ratio at calling moment when set true
---* time?, boolean (default false), includes process's start time when set true
---* mem?, boolean (default false), includes process's memory usage ratio when set true
---
---### Example:
---```lua
---local info, err = ps() -- fetches all
---yassert(err)
---table.dump(info)
---ps({ mem = true, user = true }) -- gets all with launcher and memory usage ratio info
---ps(20) -- queries process of pid 20
---ps("yock") -- fuzzy queries process of command with yock
---```
---@param opt ps_opt|string|integer|nil
---@return table<integer, ps_info>
function ps(opt) end

---kill kills process according to pid
---or process's name.
---@param k integer|string
---@return err
function kill(k) end

---tarc compresses src to dst base on tar.gz algorithm.
---Directly using it isn't recommended, you can use
---compress to instead of it. The compress function
---abstract tarc and zipc to fit in different platform
---default format.
---@param src string
---@param dst string
function tarc(src, dst) end

---zipc compresses src to dst base on zip algorithm.
---Directly using it isn't recommended, you can use
---compress to instead of it. The compress function
---abstract tarc and zipc to fit in different platform
---default format.
---@param src string
---@param dst string
function zipc(src, dst) end

---untar uncompress src to dst base on tar.gz algorithm.
---Directly using it isn't recommended, you can use
---uncompress to instead of it. The uncompress function
---abstract untar and unzip to fit in different platform
---default format.
---@param src string
---@param dst string
function untar(src, dst) end

---unzip uncompress src to dst base on zip algorithm.
---Directly using it isn't recommended, you can use
---uncompress to instead of it. The uncompress function
---abstract unzip and unzip to fit in different platform
---default format.
---@param src string
---@param dst string
function unzip(src, dst) end

---compress compresses src to dst base on
---tar.gz or zip according to filename extension.
---### Example:
---```lua
---compress("./test", "test.zip")
---compress("./test", "test.tar.gz")
---```
---@param src string
---@param dst string
function compress(src, dst) end

---uncompress uncompress src to dst base on
---tar.gz or zip according to filename extension,
---and returns an absolute path combining dst and
---root directory of compress package when uncompressed successfully.
---
---### Example:
---```lua
---uncompress("./test.zip", "./test")
---uncompress("./test.tar.gz", "./test")
---```
---@param src string
---@param dst string
---@return string, err
function uncompress(src, dst) end

---export sets user's environment variable
---for ever. If you only want to set temporary
---or local variable, see the `exportl` function.
---
---`NOTE`: In current overload, write is overwrite format.
---If you want to write by append, see other overload function.
---
---### Example:
---```lua
---# hardly hurt
---export("PATH:/bin/yock")
---# append write into PATH when value isn't exist, and it's available on windows,
---# which meant that it isn't required using Path instead of PATH.
---
---# please keep cautious!!!
---export("PATH", "/bin/yock") -- it'll overwrite entire PATH's value
---```
---@param k string
---@param v string
---@return err
function export(k, v) end

---export sets user's environment variable
---for ever. If you only want to set temporary
---or local variable, see the `exportl` function.
---
---Comparing with `export(k, v)`, the overload function
---is conservative, and write value by append. If you want
---to overwrite entire value, see other overload function.
---
---### Example:
---```lua
---export("PATH:/bin/yock")
---```
---@param kv string
---@return err
function export(kv) end

---unset removes specified environment variable for ever.
---If you only want to set temporary, or local variable,
---see the `unsetl` function.
---
---Just like export, the unset function supports entire
---delete and deletes one of values for specified key.
---
---### Example:
---```lua
---# entire deletes
---unset("PATH")
---
---# deletes one of values for specified key
---unset("PATH:/bin/yock")
---```
---@param k string
function unset(k) end

---@class ifconfig_addr
---@field addr string

---@alias ifconfig_flag string|"up"|"broadcast"|"multicast"|"loopback"

---@class ifconfig_result
---@field index integer
---@field mtu integer
---@field name string
---@field hardwareAddr string
---@field flags ifconfig_flag[]
---@field addrs ifconfig_addr[]

---ifconfig returns information about net interface.
---@return ifconfig_result[]
function ifconfig() end

---@class systemctl_opt

systemctl = {}

---@alias sys_state string
---|> "all"
---| "active"
---| "inactive"

---@alias sys_type string
---|> "target"
---|"service"

---@param t? sys_type
---@param s? sys_state
---@return sys_service[]
function systemctl.list(t, s) end

---@param name string
---@return err
function systemctl.restart(name) end

---@param name string
---@return err
function systemctl.start(name) end

---@param name string
---@return err
function systemctl.stop(name) end

---@param name string
---@return err
function systemctl.delete(name) end

---@param name string
---@return err
function systemctl.disable(name) end

---@param name string
---@return err
function systemctl.enable(name) end

---@class sc_create_opt_unit
---@field description? string
---@field before? string
---@field after? string

---@class sc_create_opt_service
---@field type? "simple"|"exec"|"forking"|"oneshot"|"dbus"|"notify"|"idle"
---@field execStart? string
---@field execStop? string
---@field privateTmp? boolean
---@field restartSec? integer
---@field restart? string

---@class sc_create_opt_install
---@field wantedBy string

---@class sc_create_opt_spec

---@class sc_create_opt
---@field unit? sc_create_opt_unit
---@field service? sc_create_opt_service
---@field install? sc_create_opt_install
---@field spec? sc_create_opt_spec

---@param name string
---@param opt sc_create_opt
---@return err
function systemctl.create(name, opt) end

---@alias service_status integer
---|> "running"
---| "stopped"
---| "unknown"

---@class sys_service
---@field pid integer
---@field name string
---@field status service_status

---@param name string
---@return sys_service, err
function systemctl.status(name) end

iptables = {}

---@class iptables_list_opt
---@field name string
---@field chain string
---@field legacy boolean

---@class fireware_rule
---@field name string
---@field proto string
---@field src string
---@field dst string
---@field action string

---### Option:
---* legacy: determine to use iptables or iptables-legacy (except windows)
---* name: returns service to be specified and all services when the length of name is empty/zero.
---@param opt iptables_list_opt
---@return fireware_rule[]|fireware_rule, err
function iptables.list(opt) end

---@class iptables_op_opt
---@field chain? string
---@field name? string
---@field protocol? string
---@field action? string
---@field destination? string
---@field legacy? boolean

-- chain: INPUT(linux), in(windows)
--
-- action: drop(linux), block(windows)
---@param opt iptables_op_opt
---@return err
function iptables.add(opt) end

---@param opt iptables_op_opt
---@return err
function iptables.del(opt) end

---@class lsof_info
---@field pid string
---@field state string
---@field proto string
---@field Local string
---@field foreign string

---lsof returns information about port occupancy
---status according to given port.
---@param port? integer
---@return lsof_info[]|lsof_info
function lsof(port) end

---@class curl_opt
---@field header? table<string, string> # header contains the request header fields either received by the server or to be sent by the client.
---@field method? string|"GET"|"POST"|"HEAD"|"PUT"|"DELETE"|"CONNECT"|"OPTIONS"|"TRACE"|"PATCH" # method specifies the HTTP method (GET, POST, PUT, etc.)
---@field data? string # data is the request's body.
---@field save? boolean # write body into specified file when set true.
---@field dir? string # set root directory of file to be saved.
---@field filename? fun(url: string): string # returns filename that will be saved according to url.
---@field async? boolean #

---curl receives urls and ranges its to send
---request one by one, and collects all response.body
---to contact and return according to double '\n'.
---
---### Option:
---* header?, table<string, string>, header contains the request header fields either received by the server or to be sent by the client.
---* method?, string (default GET), method specifies the HTTP method (GET, POST, PUT, etc.)
---* data?, string, data is the request's body.
---* save?, boolean (default true), write body into specified file when set true.
---* dir?, string, set root directory of file to be saved.
---* filename?, fun(url: string): string, returns filename that will be saved according to url.
---
---### Example:
---```lua
---# curl fetches url and saves it into specified path combing dir and filename.
---# `NOTE`: if save is true, body will not write into return's string.
---curl({
---    save = true,
---    dir = "./",
---    filename = function(url)
---        return path.base(url)
---    end
---}, "https://www.github.com/ansurfen/yock")
---
---local data = curl({
---     method = "POST",
---     data = json.encode({ username = "yock" })
---}, "")
---print(data)
---```
---
---@param opt curl_opt
---@vararg string
---@return string, err
function curl(opt, ...) end

---curl receives variable string argument
---and ranges its to send GET request one by one,
---and collects all response.body to
---contact and return according to double ‘\n’.
---
---### Example:
---```lua
---curl("https://www.github.com/ansurfen/yock")
---```
---@vararg string
---@return string, err
function curl(...) end

---write writes data to file and creates file
---when file isn't exist.
---@param file string
---@param data string
---@return err
function write(file, data) end

---read blocks and wait for inputting from user,
---and builds mapping relationship from given
---name to input's value.
---@param name string
function read(name) end

---exportl set temporary or local environment
---variable, and less like the `export` function
---it fails to write by append, and only supports to
---overwrite entire key, which meant you may need two step
---to complete append operation.
---@param k string
---@param v string
---@return err
function exportl(k, v) end

---unsetl removes temporary or local environment
---variable, and less like the `export` function
---it fails to removes one of values, and only supports to
---remove entire key, which meant you may need two step
---to complete append operation.
---@param k string
---@return err
function unsetl(k) end

---environ returns values of environment variables k,
---and if v includes multiple values (e.g. PATH), then
---it'll be split into string array.
---@param k string
---@return string[]
function environ(k) end

---environ returns all environment variables
---@return table<string, string>
function environ() end
