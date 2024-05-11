local logger = require 'coagula.logger'

local M = {}

M.logger = logger:new()
M.coagula_log_file = vim.fn.stdpath('log') .. "/coagula.log"
M.debug_header = string.format("[DEBUG][%s]:", os.date("%m/%d/%Y %H:%M:%S"))
M.lua_coagula_path = debug.getinfo(1).source:match("@?(.*/)")

function M.dinamcally_get_rust_module()
    local rust_library_path = M.lua_coagula_path:gsub("/coagula", "") .. "coagula_rs.so"
    local rust_module = package.loadlib(rust_library_path, "luaopen_coagula_rs")
    if rust_module then
        return rust_module()
    else
        return nil
    end
end

return M
