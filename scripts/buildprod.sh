#!/bin/bash

# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o notely
GOOS=linux GOARCH=arm64 go build -o notely .
