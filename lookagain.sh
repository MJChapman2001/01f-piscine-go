#! /bin/bash

find . -name '*.sh' -printf "%p\n" | sort -n | cut -d "." -f2 | sed 's|/||'