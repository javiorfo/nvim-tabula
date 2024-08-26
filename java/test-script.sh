#!/bin/bash

CONN_STR="jdbc:informix-sqli://ol_testing.psa.com.ar:1526/psa_dev_ipsaem:USER=jboss-dev;PASSWORD=UaC0TN2hqH;";
DB_NAME="psa_dev_ipsaem"
# QUERIES="select * from status_garantias"
QUERIES="co__promociones"

mvn clean install

java -jar target/tabula.jar -engine "informix" \
    -conn-str "$CONN_STR" \
    -dbname $DB_NAME \
    -queries "$QUERIES" \
    -dest-folder /tmp \
    -tabula-log-file /tmp/papa.log \
    -border-style 4 \
    -option 3 \
