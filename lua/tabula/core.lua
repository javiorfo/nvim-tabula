local setup = require'tabula'.SETTINGS
local engines = require'tabula.engines'
local util = require'tabula.util'
local spinetta = require 'spinetta'
local M = {}

function M.get_connection_string()
    local db = setup.db
    if db.connections then
        local connection = db.connections[require'tabula'.default_db]
        return engines.db[connection.engine].get_connection_string(connection)
    end
end

function M.build()
    local root_path = util.tabula_root_path
    local script = string.format(
    "%sscript/build.sh %s 2> >( while read line; do echo \"[ERROR][$(date '+%%m/%%d/%%Y %%T')]: ${line}\"; done >> %s)", root_path,
        root_path, util.tabula_log_file)
    local spinner = spinetta:new {
        main_msg = "  Tabula   Building Go binary... ",
        speed_ms = 100,
        on_success = function()
            util.logger:info("  Tabula is ready to be used!")
        end
    }

    spinner:start(spinetta.job_to_run(script))
end

function M.show_logs()
    vim.cmd(string.format("vsp %s | normal G", util.tabula_log_file))
end

return M
