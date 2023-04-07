FROM golang:1.20

# Set destination for COPY
WORKDIR /app

#rr Download Go modules
ENV GO111MODULE on
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY server.go ./

# Maintain build directory structure
COPY cmd ./cmd
COPY web ./web

# Build
RUN CGO_ENABLED=0 GO111MODULE=on GOOS=linux go build -o /server

# Use a lightweight Alpine image as the base image for the final container
FROM alpine:latest

# Copy the Go binary from the previous stage to the final container
COPY --from=0  /server .
COPY --from=0  /app/web ./web

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080

# Run
CMD [ "/server" ]