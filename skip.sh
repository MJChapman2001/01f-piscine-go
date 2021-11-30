#! /bin/bash

awk 'NR % 2 {print} !(NR % 2) continue'