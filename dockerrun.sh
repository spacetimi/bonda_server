#!/bin/bash
docker run --restart always -d -e app_environment=Test -e app_name=bonda -e app_dir_path=/go/src/github.com/spacetimi/bonda_server --publish 8000:8000 bonda-server | xargs -I containerId docker logs -f containerId
