#!/bin/sh

PORT=33060
echo "bind to $PORT"
docker run -p $PORT:3306  --name local-mysql -e MYSQL_ROOT_PASSWORD=secret -d mysql:5.7.22
