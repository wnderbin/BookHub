#!/bin/bash

go mod init scripts

go get github.com/gorilla/mux
go get github.com/go-sql-driver/mysql # installing dependencies

cd ../cmd/web_books/

echo -n -e "\nRun/Build? >> "
read command

echo -n -e "\nEnter the port on which the server will run >> "
read port

if [ $command = "run" ] || [ $command = "r" ]; then
    go run main.go $port
else
    go build main.go 
fi
