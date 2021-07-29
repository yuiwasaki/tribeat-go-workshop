#!/bin/bash

oapi-codegen -generate "server" -package oapi sample.yml > oapi/openapi.gen.go
oapi-codegen -generate "types" -package oapi sample.yml > oapi/types.gen.go
