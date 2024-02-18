#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE USER shopus ENCRYPTED PASSWORD 'shopus' LOGIN;
	CREATE DATABASE shopus OWNER shopus;
EOSQL

psql -v ON_ERROR_STOP=1 --username "shopus" --dbname "shopus" -f /app/sql/init-db.sql