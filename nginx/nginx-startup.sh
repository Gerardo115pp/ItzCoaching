#!/bin/sh
echo "HeliOS Nginx Startup Script";

for file in /etc/nginx/sites-available/*; do
  ln -sf $file /etc/nginx/sites-enabled/$(basename $file);
done

nginx -g "daemon off;"
