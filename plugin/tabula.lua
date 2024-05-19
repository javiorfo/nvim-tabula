if vim.g.tabula then
    return
end

vim.g.tabula = 1

vim.api.nvim_create_user_command('Tabula', function()
    require("tabula.core").execute()
end, {})
