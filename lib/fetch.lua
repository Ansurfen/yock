fetch = {}

function fetch.file(url, file_type)
    local tmp_path = path.join(env.workdir, "..", "yock_tmp")
    local file = ypm:get_cache(url)
    if not (type(file) == "string" and #file > 0) then
        file = random.str(8)
        yassert(http({
            debug = true,
            save = true,
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
    return fetch.file(url, ".zip")
end

function fetch.script(url)
    return fetch.file(url, ".lua")
end
