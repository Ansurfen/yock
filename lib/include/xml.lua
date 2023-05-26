---@meta _

---@class xmlDocWriteSettings
---@field UseCRLF boolean
local xmlDocWriteSettings = {}

---@class xmlDoc
---@field WriteSettings xmlDocWriteSettings
local xmlDoc = {}

---@return xmlDoc
function xml()
    return {}
end

---@param b string
---@return userdata
function xmlDoc:ReadFromBytes(b)

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
