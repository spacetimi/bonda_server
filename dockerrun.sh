#!/bin/bash
# Run the docker container for bonda against $1 environment
if [ $1 = "Local" ] || [ $1 = "Test" ]; then
   docker run --restart always -d -e app_environment=$1 -e app_name=bonda -e app_dir_path=/go/src/github.com/spacetimi/bonda_server --publish 8000:8000 --publish 8001:8001 bonda-server | xargs -I containerId docker logs -f containerId
else
   echo "Usage: dockerrun.sh <Local|Test>"
fi

