#!/bin/bash

docker run --mount src=`pwd`,target=/mnt,type=bind wasm-go $@