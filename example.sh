#!/bin/sh

example1() {
	echo 'FROM alpine:latest' |
		wasmtime \
			run \
			./dockerfile2json.wasm |
		jq
}

example2() {
	curl \
		--location \
		--show-error \
		--fail \
		--silent \
		https://raw.githubusercontent.com/rabbitmq/cluster-operator/refs/heads/main/Dockerfile |
		wasmtime \
			run \
			./dockerfile2json.wasm |
		dasel --read=json --write=toml
}

example2
