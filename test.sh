#!/usr/bin/env bash

go clean -testcache
go test ./fisttp/*.go -v
