-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

user = {}

--- Current returns the current user.
---
--- The first call will cache the current user information.
--- Subsequent calls will return the cached value and will not reflect
--- changes to the current user.
---@return userUser, err
function user.Current() end

--- Lookup looks up a user by username. If the user cannot be found, the
--- returned error is of type UnknownUserError.
---@param username string
---@return userUser, err
function user.Lookup(username) end

--- LookupId looks up a user by userid. If the user cannot be found, the
--- returned error is of type UnknownUserIdError.
---@param uid string
---@return userUser, err
function user.LookupId(uid) end

--- LookupGroup looks up a group by name. If the group cannot be found, the
--- returned error is of type UnknownGroupError.
---@param name string
---@return userGroup, err
function user.LookupGroup(name) end

--- LookupGroupId looks up a group by groupid. If the group cannot be found, the
--- returned error is of type UnknownGroupIdError.
---@param gid string
---@return userGroup, err
function user.LookupGroupId(gid) end

--- UnknownUserError is returned by Lookup when
--- a user cannot be found.
---@class userUnknownUserError
local userUnknownUserError = {}


---@return string
function userUnknownUserError:Error() end

--- UnknownGroupIdError is returned by LookupGroupId when
--- a group cannot be found.
---@class userUnknownGroupIdError
local userUnknownGroupIdError = {}


---@return string
function userUnknownGroupIdError:Error() end

--- UnknownGroupError is returned by LookupGroup when
--- a group cannot be found.
---@class userUnknownGroupError
local userUnknownGroupError = {}


---@return string
function userUnknownGroupError:Error() end

--- User represents a user account.
---@class userUser
---@field Uid string
---@field Gid string
---@field Username string
---@field Name string
---@field HomeDir string
local userUser = {}

--- Group represents a grouping of users.
---
--- On POSIX systems Gid contains a decimal number representing the group ID.
---@class userGroup
---@field Gid string
---@field Name string
local userGroup = {}

--- UnknownUserIdError is returned by LookupId when a user cannot be found.
---@class userUnknownUserIdError
local userUnknownUserIdError = {}


---@return string
function userUnknownUserIdError:Error() end
