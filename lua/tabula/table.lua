local util = require'tabula.util'
local core = require'tabula.core'
local setup = require'tabula'.SETTINGS
local M = {}

function M.show_table_info(args)
    local table_selected = args[1]
    print("Value: " .. table_selected)
end

function M.get_tables()
    local engine = (setup.db and setup.db.connections and setup.db.connections[require'tabula'.default_db].engine) or ""
    vim.fn.system(string.format("%s -option 2 -engine %s -conn-str %s -dest-folder %s", util.tabula_bin_path, engine, core.get_connection_string(), util.lua_tabula_path))

    local ok, tables = pcall(dofile, require 'tabula.util'.lua_tabula_path .. 'tables.lua')
    local names = {}
    if ok then
        if tables then
            for _, v in pairs(tables) do
                table.insert(names, v)
            end
        end
    else
        table.insert(names, "no tables in DB")
    end

    return names
end

return M
