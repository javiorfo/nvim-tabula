#! /bin/bash

ROOT=$1

(cd $ROOT && cargo build --release --quiet)

if [ $? -ne 0 ]; then
    exit 1
fi

cp $ROOT/target/release/libdbeer_rs.so $ROOT
mv $ROOT/libdbeer_rs.so $ROOT/dbeer_rs.so
mv $ROOT/dbeer_rs.so $ROOT/lua/

(cd $ROOT && cargo clean)
