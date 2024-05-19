#! /bin/bash

ROOT=$1

(cd $ROOT && cargo build --release --quiet)

if [ $? -ne 0 ]; then
    exit 1
fi

cp $ROOT/target/release/libtabula_rs.so $ROOT
mv $ROOT/libtabula_rs.so $ROOT/tabula_rs.so
mv $ROOT/tabula_rs.so $ROOT/lua/

(cd $ROOT && cargo clean)
