local util = require 'tabula.util'

local host = "127.0.0.1"
local default_posgres_port = "5432"
local default_mongo_port = "27017"
local default_mysql_port = "3306"
local default_mssql_port = "1433"
local default_informix_port = "1526"
local go_executor = util.tabula_root_path .. "bin/tabula"
local java_executor = string.format("java -jar %s%s", util.tabula_root_path, " bin/tabula.jar")

return {
    db = {
        postgres = {
            title = "PostgreSQL",
            default_port = default_posgres_port,
            default_host = host,
            executor = go_executor,
            get_connection_string = function(connection)
                return string.format("host=%s port=%s dbname=%s %s %s sslmode=disable",
                    connection.host or host,
                    connection.port or default_posgres_port,
                    connection.dbname,
                    connection.user and "user=" .. connection.user or "",
                    connection.password and "password=" .. connection.password or ""
                )
            end
        },
        mysql = {
            title = "MySQL",
            default_port = default_mysql_port,
            default_host = host,
            executor = go_executor,
            get_connection_string = function(connection)
                return string.format("%s%stcp(%s:%s)/%s",
                    connection.user and connection.user .. ":" or "",
                    connection.password and connection.password .. "@" or "",
                    connection.host or host,
                    connection.port or default_mysql_port,
                    connection.dbname
                )
            end
        },
        mongo = {
            title = "MongoDB",
            default_port = default_mongo_port,
            default_host = host,
            executor = go_executor,
            get_connection_string = function(connection)
                return string.format("mongodb://%s%s%s:%s",
                    connection.user and connection.user or "",
                    connection.password and ":" .. connection.password .. "@" or "",
                    connection.host or host,
                    connection.port or default_mongo_port
                )
            end
        },
        mssql = {
            title = "MS-SQL",
            default_port = default_mssql_port,
            default_host = host,
            executor = go_executor,
            get_connection_string = function(connection)
                return string.format("sqlserver://%s%s%s:%s?database=%s",
                    connection.user and connection.user or "",
                    connection.password and ":" .. connection.password .. "@" or "",
                    connection.host or host,
                    connection.port or default_mssql_port,
                    connection.dbname
                )
            end
        },
        informix = {
            title = "Informix",
            default_port = default_informix_port,
            default_host = host,
            executor = java_executor,
            get_connection_string = function(connection)
                return string.format("jdbc:informix-sqli://%s:%s/%s%s%s",
                    connection.host or host,
                    connection.port or default_informix_port,
                    connection.dbname,
                    connection.user and ":USER=" .. connection.user or "",
                    connection.password and ";PASSWORD=" .. connection.password .. ";" or ""
                )
            end
        },
    }
}
