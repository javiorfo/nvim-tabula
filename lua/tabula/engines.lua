local host = "localhost"
local default_posgres_port = "5432"
local default_mongo_port = "27017"
local default_mysql_port = "3306"

return {
    db = {
        postgres = {
            title = "PostgreSQL",
            default_port = default_posgres_port,
            default_host = host,
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
            get_connection_string = function(connection)
                return string.format("%s%stcp(%s:%s)/%s",
                    connection.user and connection.user .. ":" or "",
                    connection.password and connection.password .. "@" or "",
                    connection.host or host,
                    connection.port or default_posgres_port,
                    connection.dbname
                )
            end
        },
        mongo = {
            title = "MongoDB",
            default_port = default_mongo_port,
            default_host = host,
            get_connection_string = function(connection)
                return string.format("mongodb://%s%s%s:%s/%s",
                    connection.user and connection.user or "",
                    connection.password and  ":" .. connection.password .. "@" or "",
                    connection.host or host,
                    connection.port or default_mongo_port,
                    connection.dbname
                )
            end
        }
    }
}
