#! /bin/bash

find . -name '*.sh' -printf "%p\n" | sort -k6M -k7n -r | cut -d "." -f2 | sed 's|/||'