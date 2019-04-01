#!/usr/bin/env bash

/wait-for-it.sh db:5432 || exit 1
./app config.json