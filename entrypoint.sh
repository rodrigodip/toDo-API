#!/bin/sh
chown -R 1001:1001 /task-data
exec /todo-api
