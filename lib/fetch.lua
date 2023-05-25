fetch = {}

function fetch.file(url, file_type)
    local tmp_path = path.join(env.yock_path, "yock_tmp")
    local file = ypm:get_cache(url)
    if not (type(file) == "string" and #file > 0) then
        file = random.str(8)
        yassert(http({
            debug = true,
            save = true,
            strict = true,
            dir = tmp_path,
            filename = function(s)
                return file .. file_type
            end
        }, url))
        ypm:set_cache(url, file)
    end
    return file
end

function fetch.zip(url)
    local suffix
    if env.platform.OS == "windows" then
        suffix = ".zip"
    else
        suffix = ".tar.gz"
    end
    return fetch.file(url, suffix)
end

function fetch.script(url)
    return fetch.file(url, ".lua")
end
