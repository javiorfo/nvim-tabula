local logger = require'tabula.util'.logger
local no_sql_augroup = vim.api.nvim_create_augroup("NoSqlFileSettings", { clear = true })

vim.api.nvim_create_autocmd({ "BufNewFile", "VimEnter" }, {
    group = no_sql_augroup,
    pattern = "*.js",
    callback = function()
        local db = require'tabula'.SETTINGS.db
        if db.connections then
            local connection = db.connections[require'tabula'.default_db]
            if connection.engine == "mongo" then
                logger:info(string.format("Database set to [%s]", connection.name))
            end
        end
    end,
})
