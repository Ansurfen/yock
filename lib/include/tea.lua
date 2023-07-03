-- Copyright 2023 The Yock Authors. All rights reserved.
-- Use of this source code is governed by a MIT-style
-- license that can be found in the LICENSE file.

---@meta _

tea = {}

---@param model teaModel
---@return teaProgram
function tea.NewProgram(model)
end

---@return teaModel
function tea.NewModel()
end

function tea.Quit()
end

---@class teaModel
---@field InitCallback function
---@field UpdateCallback function
---@field ViewCallback function
local teaModel = {}

---@class teaCmd
local teaCmd = {}

---@class teaProgram
local teaProgram = {}

---@return teaModel, err
function teaProgram:Run()
end
