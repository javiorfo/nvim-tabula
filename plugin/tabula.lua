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
--     require 'tabula.core'.show_table_info(opts.fargs)
end, {
    nargs = 1,
    complete = function(_, _)
        -- TODO call to Go
        local ok, tables = pcall(dofile, require'tabula.util'.lua_tabula_path .. 'tables.lua')
        local names = {}
        if ok then
            if tables then
               for _, v in pairs(tables) do
                    table.insert(names, v)
               end
            end
        else
            table.insert(names, "no tables")
        end
        return names
    end
})

vim.api.nvim_create_user_command('TabulaShowLogs', function()
    require("tabula.core").show_logs()
end, {})

vim.api.nvim_create_user_command('TabulaShowDBInfo', function()
    require("tabula.database").show_info()
end, {})

vim.api.nvim_create_user_command('TabulaSelectDB', function()
    require("tabula.database").select()
end, {})
