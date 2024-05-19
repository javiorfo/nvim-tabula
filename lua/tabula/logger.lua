local M = {}

local function logger(msg)
    return function(level)
        msg = string.format("  Tabula   %s", msg)
        vim.notify(msg, level)
    end
end

function M:new()
    local table = {}
    self.__index = self
    setmetatable(table, self)
    return table
end

function M:warn(msg)
    logger(msg)(vim.log.levels.WARN)
end

function M:error(msg)
    logger(msg)(vim.log.levels.ERROR)
end

function M:info(msg)
    logger(msg)(vim.log.levels.INFO)
end

function M:debug(msg)
    local util = require 'tabula.util'
    if require'tabula'.SETTINGS.internal.log_debug then
        local file = io.open(util.tabula_log_file, "a")
        if file then
            file:write(string.format("%s %s\n", util.debug_header, msg))
            file:close()
        end
    end
end

return M
