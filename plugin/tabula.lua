if vim.g.tabula then
    return
end

vim.g.tabula = 1

vim.api.nvim_create_user_command('TabulaBuild', function()
    require("tabula.core").build()
end, {})

vim.api.nvim_create_user_command('TabulaRun', function()
    require("tabula.core").run()
end, {})

vim.api.nvim_create_user_command('TabulaClose', function()
    require("tabula.core").close()
end, {})

vim.api.nvim_create_user_command('TabulaTableInfo', function(opts)
    require 'tabula.table'.show_table_info(opts.fargs)
end, {
    nargs = 1,
    complete = function(_, _)
        return require'tabula.table'.get_tables()
    end
})

vim.api.nvim_create_user_command('TabulaLogs', function()
    require("tabula.core").show_logs()
end, {})

vim.api.nvim_create_user_command('TabulaDBInfo', function()
    require("tabula.database").show_info()
end, {})

vim.api.nvim_create_user_command('TabulaSelectDB', function()
    require("tabula.database").select()
end, {})
