#! /bin/bash

ls -1 find . -name '*.sh' | sed 's/.sh//g'