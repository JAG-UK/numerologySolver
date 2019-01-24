#!/bin/bash

#Server and DB env vars (over-ridable)
export NUMEROLOGYSERVER_HOST="127.0.0.1:8080"
export NUMEROLOGYSERVER_WORDLISTPATH="words_alpha.txt"
export NUMEROLOGYSERVERDB_USER="postgres"
export NUMEROLOGYSERVERDB_PASSWORD="postgresAdmin"
export NUMEROLOGYSERVERDB_DBNAME="postgres"
export NUMEROLOGYSERVERDB_HOST="127.0.0.1"
export UMEROLOGYSERVERDB_PORT="5432"

go run main.go &

#Post-run tester
sleep 5
echo Testing Numerology Server bring-up...
http --print=HhBb PUT ${NUMEROLOGYSERVER_HOST}/word Word="JONATHAN"

