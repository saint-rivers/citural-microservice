#!/bin/bash

# migrate create -ext sql -dir db/migration -seq init_schema

migrate -source file://db/migration -database "postgres://tinkerer:password@localhost:4666/tinker?sslmode=disable" up