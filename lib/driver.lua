function pull(opt)
    if opt ~= nil then
        if opt.drivers ~= nil then
            if type(opt.drivers) == "table" then
                for _, url in ipairs(opt.drivers) do
                    -- {url, alias}
                    if type(url) == "table" and #url == 2 then
                        if is_url(url[1]) then
                            local filename = random.str(8)
                            http({
                                save = true,
                                dir = pathf("@/driver/local/"),
                                fn = function(path)
                                    return filename
                                end
                            }, url[1])
                            ldns:PutDriver(url[1], url[1], "local/" .. filename)
                        end
                    else
                        if is_url(url) then
                            http({
                                save = true,
                                dir = pathf("@/driver/local/")
                            }, url)
                            ldns:PutDriver(url, url, "local/" .. path.filename(url))
                        end
                    end
                end
            end
        end
        if opt.plugins ~= nil then
            if type(opt.plugins) == "table" then
                for _, domain in ipairs(opt.plugins) do
                    if is_url(domain) then
                        if ldns:GetPlugin(domain).URL ~= "" then
                            print("module was be downloaded")
                        else
                            local filename = random.str(8)
                            http({
                                save = true,
                                dir = pathf("@/plugin/local/"),
                                fn = function(path)
                                    return filename
                                end
                            }, domain)
                            ldns:PutPlugin(domain, domain, "local/" .. filename)
                        end
                    else
                        local _plugin = gdns:GetPlugin(domain)
                        if _plugin.URL ~= "" then
                            if _plugin.Path == "" then
                                http({
                                    save = true,
                                    dir = pathf("@/plugin/"),
                                    fn = function(path)
                                        return domain
                                    end
                                }, _plugin.URL)
                                gdns:UpdatePlugin(domain, _plugin.URL, domain)
                            else
                                print("module was installed")
                            end
                        end
                    end
                end
            end
        end
    end
end

-- uninit_driver is a shell that wrap exist and non-exist driver
-- if driver is non-exist, it'll be set in null_driver
---@param fn string
function uninit_driver(fn)
    ---@param opt table
    ---@vararg string
    return function(opt, ...)
        -- try to find and execute driver or fetch driver from remote
        if opt ~= nil and opt.driver ~= nil then
            local _driver = ""
            local set_alias = false
            if type(opt.driver) == "table" then
                _driver = opt.driver[1]
                set_alias = true
            else
                _driver = opt.driver
            end

            -- find and execute driver from local cache
            -- priority of local dns is more than global dns like general scope of programm language
            if ldns:GetDriver(_driver).URL ~= "" then
                local did = set_driver(fn, ldns:GetDriver(_driver).Path)
                exec_driver(did, opt, ...)
                return
            end
            if gdns:GetDriver(_driver).URL ~= "" then
                set_driver(fn, _driver)
                return
            end

            -- try to fetch driver from specify url when driver isn't exist in localhost
            if is_url(_driver) then
                if set_alias then
                    pull({
                        drivers = { { _driver, opt.driver[2] } }
                    })
                else
                    pull({
                        drivers = { _driver }
                    })
                end

                -- fetch with success
                if ldns:GetDriver(_driver).URL ~= "" then
                    ldns:AliasDriver(_driver, opt.driver[2])
                    local did = set_driver(fn, ldns:GetDriver(_driver).Path)
                    exec_driver(did, opt, ...)
                    return
                end
                if gdns:GetDriver(_driver).URL ~= "" then
                    set_driver(fn, _driver)
                    return
                end
                return
            end
        end

        -- the worst condition, not url and local cache
        -- it'll set null_driver (@/driver/null.lua) as default driver that haven't implement.
        set_driver(fn, "null")
    end
end

---@see installs
function installs(opt, ...)
    if opt ~= nil and opt.plugins ~= nil then
        pull({ plugins = opt.plugins })
        for _, plugin in ipairs(opt.plugins) do
            if ldns:GetPlugin(plugin).URL ~= "" then
                local uid = load_plugin(pathf("@/plugin/") .. ldns:GetPlugin(plugin).Path .. ".lua")
                plugins[uid].install()
            end
            if gdns:GetPlugin(plugin).URL ~= "" then
                load_plugin(pathf("@/plugin/") .. gdns:GetPlugin(plugin).Path .. ".lua")
            end
        end
    end
end

---@see install
function install(plugin, opt)
    if ldns:GetPlugin(plugin).URL ~= "" then
        local uid = load_plugin(pathf("@/plugin/") .. ldns:GetPlugin(plugin).Path .. ".lua")
        plugins[uid].install(opt)
        return
    end
    pull({ plugins = { plugin } })
    if ldns:GetPlugin(plugin).URL ~= "" then
        local uid = load_plugin(pathf("@/plugin/") .. ldns:GetPlugin(plugin).Path .. ".lua")
        plugins[uid].install(opt)
    end
    if gdns:GetPlugin(plugin).URL ~= "" then
        load_plugin(pathf("@/plugin/") .. gdns:GetPlugin(plugin).Path .. ".lua")
    end
end
