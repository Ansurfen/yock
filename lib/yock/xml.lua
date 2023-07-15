--  Copyright 2023 The Yock Authors. All rights reserved.
--  Use of this source code is governed by a MIT-style
--  license that can be found in the LICENSE file.

---@diagnostic disable: duplicate-set-field

xmlFile = {}

function xmlFile:open(file)
    local obj = {
        file = file,
        doc = xml()
    }
    setmetatable(obj, { __index = self })
    local err = obj.doc:ReadFromFile(file)
    yassert(err)
    return obj
end

function xmlFile:read(str)
    local obj = {
        doc = xml()
    }
    setmetatable(obj, { __index = self })
    local err = obj.doc:ReadFromString(str)
    yassert(err)
    return obj
end

function xmlFile:select(k)
    local keys = strings.Split(k, ".")
    local doc = self.doc
    for _, kk in ipairs(keys) do
        doc = doc:SelectElement(kk)
    end
    return doc
end

function xmlFile:selects(k)
    local keys = strings.Split(k, ".")
    local doc = self.doc
    if #keys > 1 then
        for i, kk in ipairs(keys) do
            if i == #keys then
                break
            end
            doc = doc:SelectElement(kk)
        end
    end
    return doc:FindElements(keys[#keys])
end

function xmlFile:create_element(k)
    local keys = strings.Split(k, ".")
    local doc = self.doc
    if #keys > 1 then
        for i, kk in ipairs(keys) do
            if i == #keys then
                break
            end
            doc = doc:SelectElement(kk)
        end
    end
    return doc:CreateElement(keys[#keys])
end

function xmlFile:dump()
    self.doc:IndentTabs()
    self.doc:WriteTo(os.Stdout)
end
