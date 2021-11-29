#! /bin/bash

curl -s https://learn.01founders.co/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"mjc2001\"}}){id}}"}' | sed 's/[^0-9]//g'