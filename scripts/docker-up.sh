#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")/.."
[ -f docker-compose.yml ] && [ -f .env ] || exit 1

dc() { docker compose -f docker-compose.yml --env-file .env "$@"; }

dc up -d db

dc exec -T db sh -c "/opt/mssql-tools18/bin/sqlcmd -S localhost -U sa -P \"\$MSSQL_SA_PASSWORD\" -C -Q \"IF NOT EXISTS (SELECT 1 FROM sys.databases WHERE name = N'thermometer') CREATE DATABASE [thermometer];\""
dc up -d --build
