-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

user = {}

---{{.userLookupGroup}}
---@param name string
---@return userGroup, err
function user.LookupGroup(name)
end

---{{.userLookupGroupId}}
---@param gid string
---@return userGroup, err
function user.LookupGroupId(gid)
end

---{{.userCurrent}}
---@return userUser, err
function user.Current()
end

---{{.userLookup}}
---@param username string
---@return userUser, err
function user.Lookup(username)
end

---{{.userLookupId}}
---@param uid string
---@return userUser, err
function user.LookupId(uid)
end

---@class userGroup
---@field Gid string
---@field Name string
local userGroup = {}

---@class userUnknownUserIdError
local userUnknownUserIdError = {}

---{{.userUnknownUserIdErrorError}}
---@return string
function userUnknownUserIdError:Error()
end

---@class userUser
---@field Uid string
---@field Gid string
---@field Username string
---@field Name string
---@field HomeDir string
local userUser = {}

---@class userUnknownUserError
local userUnknownUserError = {}

---{{.userUnknownUserErrorError}}
---@return string
function userUnknownUserError:Error()
end

---@class userUnknownGroupIdError
local userUnknownGroupIdError = {}

---{{.userUnknownGroupIdErrorError}}
---@return string
function userUnknownGroupIdError:Error()
end

---@class userUnknownGroupError
local userUnknownGroupError = {}

---{{.userUnknownGroupErrorError}}
---@return string
function userUnknownGroupError:Error()
end
