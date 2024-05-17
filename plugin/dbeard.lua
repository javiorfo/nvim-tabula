if vim.g.dbeard then
    return
end

vim.g.dbeard = 1

vim.api.nvim_create_user_command('DBeard', function()
    require("dbeard.core").execute()
end, {})
