#!/bin/bash

ROOT=$1

## GO ##
(cd $ROOT/go && go build -o tabula main.go)

if [ $? -ne 0 ]; then
    exit 1
fi

BIN=$ROOT/bin
mkdir -p $BIN

mv $ROOT/go/tabula $BIN

## JAVA ##
# This section is optional
# If Java is installed (21 or newer) and maven is installed

# Function to check Java version
check_java_version() {
    version_output=$(java -version 2>&1)
    if [[ $version_output =~ ([0-9]+)\.([0-9]+) ]]; then
        major_version=${BASH_REMATCH[1]}
        minor_version=${BASH_REMATCH[2]}
        
        if (( major_version > 21 )) || { (( major_version == 21 )) && (( minor_version >= 0 )); }; then
            return 0  # Java version is 21 or newer
        fi
    fi
    return 1
}

check_command() {
    command -v "$1" &> /dev/null
}

# Then build with Maven
if check_command java; then
    if check_java_version; then
        if check_command mvn; then
            $(cd $ROOT/java && mvn clean install)
            cp $ROOT/java/target/tabula.jar $BIN
            $(cd $ROOT/java && mvn clean)
        fi
    fi
fi
