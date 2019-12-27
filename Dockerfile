FROM golang

# Copy source
ADD . /go/src/github.com/spacetimi/server/

# Override any localhost URLs to host.docker.internal so we can use services running on host machine
RUN sed -i -e 's/localhost/host.docker.internal/g' /go/src/github.com/spacetimi/server/bonda/config/shared_config.local.json

# Fetch required dependencies
RUN go get github.com/gorilla/mux
RUN go get go.mongodb.org/mongo-driver/mongo
RUN go get github.com/go-redis/redis

# Build and install the server binary
RUN go install github.com/spacetimi/server/bonda/main

ENTRYPOINT /go/bin/main
EXPOSE 8000
