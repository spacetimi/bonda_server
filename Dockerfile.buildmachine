FROM golang

# Copy app source
ADD . /go/src/github.com/spacetimi/bonda_server

# Override any localhost URLs to host.docker.internal so we can use services running on host machine
RUN sed -i -e 's/localhost/host.docker.internal/g' /go/src/github.com/spacetimi/bonda_server/config/environment_config.local.json

# Fetch shared source
RUN git clone https://github.com/spacetimi/timi_shared_server /go/src/github.com/spacetimi/timi_shared_server

# Fetch required dependencies
RUN go get github.com/gorilla/mux
RUN go get go.mongodb.org/mongo-driver/mongo
RUN go get github.com/go-redis/redis
RUN go get github.com/dgrijalva/jwt-go

# Build and install the server binary
RUN go install github.com/spacetimi/bonda_server/main

ENTRYPOINT /go/bin/main
EXPOSE 8000
