#!/bin/bash

GO_BINARY=~/.local/share/nvim/lazy/nvim-tabula/bin/tabula
QUERY="select info from dummies"

$GO_BINARY -engine "postgres" \
    -conn-str "host=localhost port=5432 dbname=db_dummy user=admin password=admin sslmode=disable" \
    -queries "$QUERY" \
    -dest-folder /tmp \
    -border-style 3 \
    && cat /tmp/tabula
