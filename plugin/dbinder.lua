if vim.g.dbinder then
    return
end

vim.g.dbinder = 1

vim.api.nvim_create_user_command('DBinder', function()
    require("dbinder.core").execute()
end, {})
