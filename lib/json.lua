---@diagnostic disable: duplicate-set-field

jsonfile = {}

function jsonfile:open(filename)
    local obj = {}
    local fp, err = io.open(filename, "r+")
    yassert(err)
    obj.fp = fp
    obj.filename = filename
    if type(fp) ~= "nil" then
        obj.buf = json.decode(fp:read("*a"))
    else
        yassert("invalid file")
    end
    setmetatable(obj, self)
    self.__index = self
    return obj
end

function jsonfile:read()
    return self.fp:read("*a")
end

function jsonfile:write()
    yassert(write_file(self.filename, json.encode(self.buf)))
end

function jsonfile:set(k, v)
    self.buf[k] = v
end

function jsonfile:get(k)
    return self.buf[k]
end

function jsonfile:close()
    self.fp:close()
end
