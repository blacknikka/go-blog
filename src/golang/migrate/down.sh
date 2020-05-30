#!/bin/bash

USER=$POSTGRES_USER
PASSWORD=$POSTGRES_PASSWORD
HOST=db
DB=db

CONNECTION=postgres://$USER:$PASSWORD@$HOST:5432/$DB?sslmode=disable
migrate.linux-amd64 -source file://sql -database $CONNECTION down
