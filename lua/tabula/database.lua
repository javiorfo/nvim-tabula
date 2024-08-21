local popcorn = require 'popcorn'
local borders = require 'popcorn.borders'
local constants = require 'tabula.constants'
local engines = require 'tabula.engines'
local setup = require 'tabula'.SETTINGS
local util = require 'tabula.util'

local M = {}

function M.select()
    local content = {}
    local db = setup.db
    local default = require 'tabula'.default_db
    if db.connections then
        content = { { "󰆼 Database", "Type" } }
        for i, v in pairs(db.connections) do
            local name = constants.UNCHECKED_ICON .. v.name
            if i == default then
                name = constants.CHECKED_ICON .. v.name
            end
            table.insert(content, { name })
        end
    end

    if #content == 0 then
        content = { { "No databases available", "ErrorMsg" } }
    end

    local popup_opts = {
        width = 40,
        height = #content + 3,
        border = borders.rounded_corners_border,
        title = { "  TABULA - Select DB", "Boolean" },
        footer = { "<Ctrl-space> to select", "String" },
        content = content,
        do_after = function()
            util.disable_editing_popup()

            if #content > 0 then
                vim.api.nvim_win_set_cursor(0, { default + 1, 0 })
                vim.api.nvim_buf_set_keymap(0, 'n', '<C-space>',
                    '<cmd>lua require("tabula.database").set()<CR>', { noremap = true, silent = true })

                vim.api.nvim_create_autocmd({ "CursorMoved" }, {
                    pattern = { "<buffer>" },
                    callback = function()
                        local pos = vim.api.nvim_win_get_cursor(0)
                        if pos[1] < 2 then
                            vim.api.nvim_win_set_cursor(0, { 2, 0 })
                        end
                        if pos[2] > 0 then
                            vim.api.nvim_win_set_cursor(0, { pos[1], 0 })
                        end
                    end
                })
            end
        end
    }
    popcorn:new(popup_opts):pop()
end

local function select_or_unselect(lines, line_nr)
    for _, v in pairs(lines) do
        if v == line_nr then
            local selected = vim.fn.getline('.')
            local final = tostring(selected):gsub(constants.UNCHECKED_ICON, constants.CHECKED_ICON)
            vim.fn.setline(line_nr, final)
            require 'tabula'.default_db = v - 1
            local connection = setup.db.connections[v - 1]
            util.logger:info(string.format("Database set to [%s]", connection.name))
        else
            local unselected = vim.fn.getline(v)
            local final = tostring(unselected):gsub(constants.CHECKED_ICON, constants.UNCHECKED_ICON)
            vim.fn.setline(v, final)
        end
    end
end

function M.set()
    vim.cmd [[setl ma]]
    local line_nr = vim.fn.line('.')
    local len = #setup.db.connections + 2

    if line_nr > 1 and line_nr < len then
        local lines = {}
        for i = 2, len do table.insert(lines, i) end
        select_or_unselect(lines, line_nr)
    end

    vim.cmd [[setl noma]]
end

function M.show_info()
    local content = {}
    local footer = { "" }
    local db = setup.db
    if db.connections then
        local connection = db.connections[require 'tabula'.default_db]
        local db_const_data = engines.db[connection.engine]
        footer = { db_const_data.title, "String" }
        content = {
            { "󱘖 Connection Data", "Type" },
            { "NAME     󰁕 " .. connection.name },
            { "HOST     󰁕 " .. (connection.host or db_const_data.default_host) },
            { "PORT     󰁕 " .. (connection.port or db_const_data.default_port) },
            { "DB NAME  󰁕 " .. connection.dbname },
            { "USER     󰁕 " .. (((connection.user and setup.view.show_user and connection.user) or connection.user and "********") or "-") },
            { "PASSWORD 󰁕 " .. (((connection.password and setup.view.show_password and connection.password) or connection.password and "********") or "-") },
        }
    end

    if #content == 0 then
        content = { { "No database available", "ErrorMsg" } }
    end

    local popup_opts = {
        width = 40,
        height = #content + 3,
        border = borders.rounded_corners_border,
        title = { "  TABULA - DB Info", "Boolean" },
        footer = footer,
        content = content,
        do_after = function()
            util.disable_editing_popup()
        end
    }
    popcorn:new(popup_opts):pop()
end

return M
