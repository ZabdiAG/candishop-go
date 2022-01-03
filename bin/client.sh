#!/bin/bash

curl -XPOST http://localhost:8080/graphql \
-H 'Content-Type: application/graphql' \
-d 'query Root{ latestPost }'
