local logger = require 'tabula.util'.logger
local setup = require 'tabula'.SETTINGS
local db = setup.db

if db.connections then
    local connection = db.connections[require 'tabula'.default_db]
    if connection.name and connection.dbname and connection.engine and require 'tabula.engines'.db[connection.engine] and connection.engine == "mongo" then
        logger:info(string.format("Database set to [%s]", connection.name))
    end
end

vim.api.nvim_set_keymap('v', setup.commands.execute, '<cmd>lua require("tabula.core").run()<CR>',
    { noremap = true, silent = true })
vim.api.nvim_set_keymap('n', setup.commands.execute, '<cmd>lua require("tabula.core").run()<CR>',
    { noremap = true, silent = true })
vim.api.nvim_set_keymap('n', setup.commands.close, '<cmd>lua require("tabula.core").close()<CR>',
    { noremap = true, silent = true })
