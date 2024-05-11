if vim.g.coagula then
    return
end

vim.g.coagula = 1

vim.api.nvim_create_user_command('Coagula', function()
    require("coagula.core").execute()
end, {})
