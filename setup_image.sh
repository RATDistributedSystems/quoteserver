#!/bin/bash

CGO_ENABLED=0 go build -a --installsuffix cgo --ldflags="-s" -o quoteserver

# Build the image
docker build -t quoteserver .

# Remove remnants
rm -f quoteserver
