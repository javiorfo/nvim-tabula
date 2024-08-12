local logger = require'tabula.util'.logger
local sql_augroup = vim.api.nvim_create_augroup("SqlFileSettings", { clear = true })

vim.api.nvim_create_autocmd({ "BufNewFile", "VimEnter" }, {
    group = sql_augroup,
    pattern = "*.sql",
    callback = function()
        local db = require'tabula'.SETTINGS.db
        if db.connections then
            local connection = db.connections[require'tabula'.default_db]
            logger:info(string.format("Database set to [%s]", connection.name))
        else
            logger:info("No database configured.")
        end
    end,
})
