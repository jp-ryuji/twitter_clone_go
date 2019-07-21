#!/bin/bash

CONFIG_FILE="../config/db/sqlboiler.toml"
GENERATED_FILES_DIR="../internal/infrastructure/orm"

cd `dirname $0`

sqlboiler -c $CONFIG_FILE --wipe psql
go test $GENERATED_FILES_DIR -c $CONFIG_FILE
