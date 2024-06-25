#!/bin/bash

# Directory containing SQL files
SCHEMA="database/schema"

# Load environment variables from .env file
export $(grep -v '^#' .env | xargs)

# Connect to PostgreSQL and run each SQL file
for SQL_FILE in $SCHEMA/*.sql
do
  PGPASSWORD=$POSTGRES_PASSWORD psql -h localhost -p 5432 \
  -U "$POSTGRES_USER" -d "$POSTGRES_DB" \
   -f "$SQL_FILE"
done