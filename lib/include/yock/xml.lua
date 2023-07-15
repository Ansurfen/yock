-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

---{{.xmlDocWriteSettings}}
---@class xmlDocWriteSettings
---@field UseCRLF boolean
local xmlDocWriteSettings = {}

---{{.xmlDoc}}
---@class xmlDoc
---@field WriteSettings xmlDocWriteSettings
local xmlDoc = {}

---{{.xml}}
---@return xmlDoc
function xml()
    return {}
end

---{{.xmlDoc_ReadFromBytes}}
---@param b string
---@return userdata
function xmlDoc:ReadFromBytes(b)
end

---@param s string
---@return userdata
function xmlDoc:ReadFromString(s)
end

---@param file string
---@return userdata
function xmlDoc:ReadFromFile(file)
end

---@param e string
---@return xmlDoc
function xmlDoc:SelectElement(e)
    return {}
end

---@param e string
---@return userdata
function xmlDoc:FindElements(e)
end

---@return string
function xmlDoc:Text()
    return ""
end

---@param e string
---@return xmlDoc
function xmlDoc:CreateElement(e)
    return {}
end

---@param v string
function xmlDoc:SetText(v)

end

---@param file string
---@return userdata
function xmlDoc:WriteToFile(file)
end

function xmlDoc:IndentTabs()

end

---@param e userdata
function xmlDoc:RemoveChild(e)

end

---@param e string
---@return userdata[]
function xmlDoc:SelectElements(e)
end

---@param b userdata
function xmlDoc:WriteTo(b)
end

---@class xmlFile
xmlFile = {}

---@param file string
---@return xmlFile
function xmlFile:open(file)
end

---@param str string
---@return xmlFile
function xmlFile:read(str)
end

---@param k string
---@return xmlDoc
function xmlFile:select(k)
end

---@param k string
---@return xmlDoc[]
function xmlFile:selects(k)
end

---@param k string
---@return xmlDoc
function xmlFile:create_element(k)
end

function xmlFile:dump()
end
