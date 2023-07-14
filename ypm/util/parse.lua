return function(module, handle)
    ---@return string|nil
    local parseProxy = function(p, m)
        for _, v in pairs(p["filter"]) do
            if v == m then
                return nil
            end
        end
        local rurl = p["redirect"][m]
        if rurl ~= nil then
            return rurl
        end
        return strf(p["url"], {
            ver = m
        })
    end
    local proxies, err = find({
        dir = false,
        pattern = "\\.lua"
    }, pathf("~/ypm/proxy"))
    if err ~= nil or #proxies == 0 then
        cp(cat(pathf("~/ypm/template/defaultSource.tpl")), pathf("~/ypm/proxy"))
        proxies, err = find({
            dir = false,
            pattern = "\\.lua"
        }, pathf("~/ypm/proxy"))
        if err ~= nil then
            yassert("system error")
        end
    end

    local defaultProxy
    local candidates = {}
    if type(proxies) == "table" then
        for _, proxy in ipairs(proxies) do
            local filename = path.filename(proxy)
            if filename == config:rawget("defaultSource") then
                defaultProxy = import(proxy)
            else
                table.insert(candidates, import(proxy))
            end
        end
    end

    local urls = {}
    if defaultProxy ~= nil then
        table.insert(urls, parseProxy(defaultProxy, module))
    end
    for _, proxy in ipairs(candidates) do
        table.insert(urls, parseProxy(proxy, module))
    end

    loadbalance({ maxRetry = #urls }, multi_bind(urls, handle))
end
