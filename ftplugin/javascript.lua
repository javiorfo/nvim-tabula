local logger = require 'tabula.util'.logger

local db = require 'tabula'.SETTINGS.db
if db.connections then
    local connection = db.connections[require 'tabula'.default_db]
    if connection.name and connection.dbname and connection.engine and require 'tabula.engines'.db[connection.engine] and connection.engine == "mongo" then
        logger:info(string.format("Database set to [%s]", connection.name))
    end
end
