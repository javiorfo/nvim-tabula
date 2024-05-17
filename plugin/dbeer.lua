if vim.g.dbeer then
    return
end

vim.g.dbeer = 1

vim.api.nvim_create_user_command('DBeer', function()
    require("dbeer.core").execute()
end, {})
