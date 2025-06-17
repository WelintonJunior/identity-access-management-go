#!/bin/bash

./bin/api_server initApiServer &
cd frontend && npm run dev
wait
