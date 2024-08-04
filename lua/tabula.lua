local M = {}

M.SETTINGS = {
    output = {
        dest_folder = "/tmp"
    },
    db = {
        default = 2,
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

function M.setup(opts)
end

return M
