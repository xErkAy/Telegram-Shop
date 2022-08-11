#!/bin/sh

cp -r /chat_app_build/node_modules /chat_app

chown -R node:node /chat_app/node_modules

exec "$@"