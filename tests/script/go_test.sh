#!/bin/bash

GO_BINARY=~/.local/share/nvim/lazy/nvim-dbeer/bin/dbeer
QUERY="select * from prueba"

# $GO_BINARY -engine "postgres" \
#     -conn-str "host=localhost port=5432 dbname=db_dummy user=admin password=admin sslmode=disable" \
#     -dbname "db_dummy" \
#     -queries "$QUERY" \
#     -dest-folder /tmp \
#     -border-style 3 \
#     && cat /tmp/dbeer

$GO_BINARY -engine "mongo" \
    -conn-str "mongodb://admin:admin@127.0.0.1:27017" \
    -dbname "db_dummy" \
    -queries "dummies" \
    -dest-folder /tmp \
    -dbeer-log-file /tmp/caca \
    -border-style 3 \
    -option 1 \
