#!/bin/sh
command -v $1 >/dev/null 2>&1 && echo 1
exit 0
