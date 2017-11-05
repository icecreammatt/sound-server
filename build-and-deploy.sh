#!/bin/bash

GOOS=linux GOARCH=arm go build
scp sound-server root@192.168.40.18:
