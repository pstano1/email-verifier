#!/bin/bash

go run ./cmd & 

cd frontend && npm run start

sleep 5

open bin/index.html