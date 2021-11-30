#! /bin/bash

find . -name '*.sh' -type f | cut -d "." -f2 | sed 's|/||'