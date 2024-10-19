if vim.g.dbeer then
    return
end

vim.g.dbeer = 1

vim.api.nvim_create_user_command('DBeerBuild', function()
    require("dbeer.core").build()
end, {})

vim.api.nvim_create_user_command('DBeerTables', function()
    require("dbeer.table").show()
end, {})

vim.api.nvim_create_user_command('DBeerLogs', function()
    require("dbeer.core").show_logs()
end, {})

vim.api.nvim_create_user_command('DBeerDB', function()
    require("dbeer.database").show()
end, {})
