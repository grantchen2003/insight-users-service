#!/bin/bash

cd ..

export ENV=dev

nodemon --exec "go run cmd/users/main.go" --watch . -e go