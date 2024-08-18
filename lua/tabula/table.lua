local util = require 'tabula.util'
local core = require 'tabula.core'
local setup = require 'tabula'.SETTINGS
local popcorn = require 'popcorn'
local borders = require 'popcorn.borders'
local M = {}

local function get_file_line_info(filePath)
    local file = io.open(filePath, "r")
    if not file then
        print("Error: Could not open file.")
        return nil, nil
    end

    local first_line_length = nil
    local total_lines = 0

    for line in file:lines() do
        total_lines = total_lines + 1
        if total_lines == 1 then
            first_line_length = util.count_UTF8_characters(line)
        end
    end

    file:close()

    return first_line_length, total_lines
end

function M.show_table_info(args)
    local table_selected = args[1]
    local engine = (setup.db and setup.db.connections and setup.db.connections[require 'tabula'.default_db].engine) or ""
    local result = vim.fn.system(string.format(
        "%s -option 3 -engine %s -conn-str \"%s\" -queries %s -border-style %d -header-style-link %s",
        util.tabula_bin_path, engine, core.get_connection_string(), table_selected, setup.output.border_style,
        setup.output.header_style_link))

    local line_1, tabula_file

    for line in string.gmatch(result, "[^\r\n]+") do
        if not line_1 then
            line_1 = line
        elseif not tabula_file then
            tabula_file = line
            break
        end
    end

    if string.sub(line_1, 1, 7) ~= "[ERROR]" then
        if tabula_file then
            local line_len, row_len = get_file_line_info(tabula_file)
            local opts = {
                width = line_len + 4,
                height = row_len + 2,
                border = borders.simple_thick_border,
                title = { "  Tabula - Table Info", "Boolean" },
                footer = { "Table: " .. string.upper(table_selected), "String" },
                content = tabula_file,
                do_after = function()
                    vim.cmd [[ setlocal nowrap ]]
                    vim.cmd [[ setl noma ]]
                    vim.cmd(line_1)
                end
            }

            popcorn:new(opts):pop()
        else
            util.logger:error("Problem ocurred opening popup with table info.")
        end
    else
        util.logger:error(line_1)
    end
end

function M.get_tables()
    local engine = (setup.db and setup.db.connections and setup.db.connections[require 'tabula'.default_db].engine) or ""
    local result = vim.fn.system(string.format("%s -option 2 -engine %s -conn-str \"%s\"", util.tabula_bin_path, engine,
        core.get_connection_string()))

    local str = result:gsub("%[", ""):gsub("%]", ""):gsub("^%s*(.-)%s*$", "%1")

    local table_names = {}
    for word in str:gmatch("%S+") do
        table.insert(table_names, word)
    end
    return table_names
end

return M
