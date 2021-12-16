#!/bin/sh

#set DATABASE_URL=postgresql://postgres:root@127.0.0.1:5432/fiber_pgsql
set APP_NAME=SimpleBookManagement
set DB_USER=postgres
set SERVER_PORT=5432
set SERVER_DB_PASS=root
set SERVER_DB_NAME=fiber_pgsql
go run app.go