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

local function get_buffer_content()
    local mode = vim.api.nvim_get_mode().mode

    if mode == 'v' or mode == 'V' or mode == '\22' then
        vim.cmd("normal " .. vim.api.nvim_replace_termcodes("<esc>", true, true, true))
        local start_pos = vim.fn.getpos("'<")
        local end_pos = vim.fn.getpos("'>")

        local lines = vim.api.nvim_buf_get_lines(0, start_pos[2] - 1, end_pos[2], false)
        if #lines == 0 then
            return ""
        end

        if start_pos[2] == end_pos[2] then
            return lines[1]:sub(start_pos[3], end_pos[3])
        else
            lines[1] = lines[1]:sub(start_pos[3])
            lines[#lines] = lines[#lines]:sub(1, end_pos[3])
            return table.concat(lines, "")
        end
    else
        local buf_number = vim.api.nvim_get_current_buf()
        local lines = vim.api.nvim_buf_get_lines(buf_number, 0, -1, false)
        local content = table.concat(lines, "")
        return content
    end
end

function M.run()
    local queries = get_buffer_content()
    local engine = (setup.db and setup.db.connections and setup.db.connections[require'tabula'.default_db].engine) or ""
    vim.fn.system(string.format("%s -engine %s -conn-str \"%s\" -queries \"%s\" -dest-folder %s -border-style %d", util.tabula_bin_path, engine, M.get_connection_string(), queries, setup.output.dest_folder, setup.output.border_style))

    local orientation = "sp"
    vim.cmd(string.format("%d%s %s", 20, orientation, "/tmp/tabula"))

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
