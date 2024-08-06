local logger = require 'tabula.logger'

local M = {}

M.logger = logger:new()
M.tabula_log_file = vim.fn.stdpath('log') .. "/tabula.log"
M.debug_header = string.format("[DEBUG][%s]:", os.date("%m/%d/%Y %H:%M:%S"))
M.lua_tabula_path = debug.getinfo(1).source:match("@?(.*/)")
M.lua_tabula_path = debug.getinfo(1).source:match("@?(.*/)")
M.tabula_root_path = M.lua_tabula_path:gsub("/lua/tabula", "")

function M.disable_editing_popup()
    -- Disable editing
    vim.cmd [[setl noma]]

    function Nothing()
        logger:info("Visual Mode is disabled in this window.")
        return ''
    end

    -- Disable Visual Mode
    vim.api.nvim_buf_set_keymap(0, 'n', 'v', 'v:lua.Nothing()', { noremap = true, expr = true })
    vim.api.nvim_buf_set_keymap(0, 'n', '<C-v>', 'v:lua.Nothing()', { noremap = true, expr = true })
    vim.api.nvim_buf_set_keymap(0, 'n', 'V', 'v:lua.Nothing()', { noremap = true, expr = true })
end

return M
