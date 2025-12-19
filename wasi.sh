#!/bin/sh

# Use the standard Go compiler instead of tinygo as of now.
# tinygo has limited reflection support, which causes panics when using
# the standard `encoding/json` package on the complex AST from the moby parser.
# The standard Go compiler's WASM target has better reflection support.
# This may be resolved in future versions of tinygo.

outname=./dockerfile2json.wasm
mainpat=./cmd/dockerfile2json/main.go

GOOS=wasip1 GOARCH=wasm go \
	build \
	-o "${outname}" \
	-ldflags="-s -w" \
	"${mainpat}"
