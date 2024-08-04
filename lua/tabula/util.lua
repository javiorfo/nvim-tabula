local logger = require 'tabula.logger'

local M = {}

M.logger = logger:new()
M.tabula_log_file = vim.fn.stdpath('log') .. "/tabula.log"
M.debug_header = string.format("[DEBUG][%s]:", os.date("%m/%d/%Y %H:%M:%S"))
M.lua_tabula_path = debug.getinfo(1).source:match("@?(.*/)")

return M
