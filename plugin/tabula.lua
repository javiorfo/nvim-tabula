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

vim.api.nvim_create_user_command('TabulaShowTables', function()
    require("tabula.core").show_tables()
end, {})

vim.api.nvim_create_user_command('TabulaShowTableInfo', function(opts)
    require 'tabula.core'.show_table_info(opts.fargs)
end, {
    nargs = 1,
    complete = function(_, _)
        -- TODO call to Go
        local specials = require'tabula'.DEFAULTS.special
        local names = {}
        if specials then
           for _, v in pairs(specials) do
                table.insert(names, v.name)
           end
        end
        return names
    end
})

vim.api.nvim_create_user_command('TabulaShowLogs', function()
    require("tabula.core").show_logs()
end, {})

vim.api.nvim_create_user_command('TabulaShowDBInfo', function()
    require("tabula.core").show_db_info()
end, {})

vim.api.nvim_create_user_command('TabulaShowConnections', function()
    require("tabula.core").show_connections()
end, {})
