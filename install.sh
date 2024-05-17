#! /bin/bash

ROOT=$1

(cd $ROOT && cargo build --release --quiet)

if [ $? -ne 0 ]; then
    exit 1
fi

cp $ROOT/target/release/libdbeard_rs.so $ROOT
mv $ROOT/libdbeard_rs.so $ROOT/dbeard_rs.so
mv $ROOT/dbeard_rs.so $ROOT/lua/

(cd $ROOT && cargo clean)
