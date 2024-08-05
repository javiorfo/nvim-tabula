local M = {}

M.SETTINGS = {
    output = {
        dest_folder = "/tmp",
        border_style = 1,
        header_style_link = "Boolean",
    },
    db = {
        default = 1,
        -- connections not in this settings
        connections = {
            -- Mandatory name, engine, dbname
            {
                name = "Postgres 1",
                engine = "postgres",
                host = "localhost",
                port = "5432",
                dbname = "db_dummy",
                user = "admin",
                password = "admin",
            },
            {
                name = "Mongo 1",
                engine = "mongo",
                host = "localhost",
                port = "27017",
                dbname = "db_dummy",
                user = "admin",
                password = "admin",
            }
        }
    },
    internal = {
        log_debug = false
    }
}

M.default_db = M.SETTINGS.db.default

function M.setup(opts)
end

return M
