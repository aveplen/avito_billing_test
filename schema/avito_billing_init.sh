#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE DATABASE avito_billing;
EOSQL

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "avito_billing" <<-EOSQL
    CREATE TABLE user_balance (
        user_id     BIGSERIAL   NOT NULL    PRIMARY KEY,
        balance     MONEY       NOT NULL    CHECK (balance::numeric >= 0)
    );

EOSQL