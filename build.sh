#!/bin/bash

go clean -cache ./...
go build -o . -v ./...
