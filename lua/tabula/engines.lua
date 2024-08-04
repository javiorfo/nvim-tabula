local host = "localhost"

return {
    db = {
        postgres = {
            get_connection_string = function(connection)
                local conn_str = string.format("host=%s port=%s dbname=%s %s %s sslmode=disable",
                    connection.host or host,
                    connection.port or "5432",
                    connection.dbname,
                    connection.user and "user=" .. connection.user or "",
                    connection.password and "password=" .. connection.password or ""
                )
                return conn_str
            end
        },
        mongo = {
            get_connection_string = function(connection)
                local conn_str = string.format("mongodb://%s%s%s:%s/%s",
                    connection.user and connection.user or "",
                    connection.password and  ":" .. connection.password .. "@" or "",
                    connection.host or host,
                    connection.port or "27017",
                    connection.dbname
                )
                return conn_str
            end
        }
    }
}
