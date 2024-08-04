local setup = require'tabula'.SETTINGS
local engines = require'tabula.engines'
local util = require'tabula.util'
local spinetta = require 'spinetta'
local M = {}

function M.get_connection_string()
    local db = setup.db
    if db and db.default and db.connections then
        local connection = db.connections[db.default]
        return engines.db[connection.engine].get_connection_string(connection)
    end
end

function M.build()
    local root_path = util.lua_tabula_path:gsub("/lua/tabula", "")
    local script = string.format(
    "%sscript/build.sh %s 2> >( while read line; do echo \"[ERROR][$(date '+%%m/%%d/%%Y %%T')]: ${line}\"; done >> %s)", root_path,
        root_path, util.tabula_log_file)
    local spinner = spinetta:new {
        main_msg = "  Tabula   Building Go binary... ",
        speed_ms = 100,
        on_success = function()
            util.logger:info("  Tabula is ready to be used!")
        end,
        on_interrupted = function()
            vim.cmd("redraw")
            local msg = "Process interrupted!"
            util.logger:info(msg)
        end
    }

    spinner:start(spinetta.job_to_run(script))
end

return M
