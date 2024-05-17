#! /bin/bash

ROOT=$1

(cd $ROOT && cargo build --release --quiet)

if [ $? -ne 0 ]; then
    exit 1
fi

cp $ROOT/target/release/libdbinder_rs.so $ROOT
mv $ROOT/libdbinder_rs.so $ROOT/dbinder_rs.so
mv $ROOT/dbinder_rs.so $ROOT/lua/
