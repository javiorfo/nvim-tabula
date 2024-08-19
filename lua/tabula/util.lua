local logger = require 'tabula.logger'

local M = {}

M.logger = logger:new()
M.tabula_log_file = vim.fn.stdpath('log') .. "/tabula.log"
M.debug_header = string.format("[DEBUG][%s]:", os.date("%m/%d/%Y %H:%M:%S"))
M.lua_tabula_path = debug.getinfo(1).source:match("@?(.*/)")
M.lua_tabula_path = debug.getinfo(1).source:match("@?(.*/)")
M.tabula_root_path = M.lua_tabula_path:gsub("/lua/tabula", "")
M.tabula_bin_path = M.tabula_root_path .. "bin/tabula"

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

local function is_valid_UTF8(byte)
    return byte >= 0 and byte <= 127
end

function M.count_UTF8_characters(str)
    local count = 0
    local i = 1
    local length = #str

    while i <= length do
        local byte = string.byte(str, i)

        if is_valid_UTF8(byte) then
            count = count + 1
            i = i + 1
        elseif byte >= 192 and byte <= 223 then
            count = count + 1
            i = i + 2
        elseif byte >= 224 and byte <= 239 then
            count = count + 1
            i = i + 3
        elseif byte >= 240 and byte <= 247 then
            count = count + 1
            i = i + 4
        else
            i = i + 1
        end
    end

    return count
end

function M.get_numeral_sprinner()
    local numbers = {}

    for i = 1, 5000 do
        table.insert(numbers, string.format("[%.2f secs]", i * 0.2))
    end
    return numbers
end

return M
