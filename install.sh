#! /bin/bash

ROOT=$1

(cd $ROOT && cargo build --release --quiet)

if [ $? -ne 0 ]; then
    exit 1
fi

cp $ROOT/target/release/libcoagula_rs.so $ROOT
mv $ROOT/libcoagula_rs.so $ROOT/coagula_rs.so
mv $ROOT/coagula_rs.so $ROOT/lua/
