FROM golang:alpine AS build

# Update the repository and install git
RUN apk --no-cache add gcc g++ make git

# Switches to tmp/app as the working directory
WORKDIR /go/src/gotodo-api
COPY . .

# Builds the current project to a binary called gotodo-api
RUN go get ./...
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/gotodo-api .

# Now that the project has been successfully built, we will use
# alpine image to run the server
FROM alpine:3.9

# Add CA certificates to the image
RUN apk --no-cache add ca-certificates

# Switch working directory to /usr/bin
WORKDIR /usr/bin
RUN mkdir db

# Copies the binary file from the BUILD container to /usr/bin
COPY --from=build /go/src/gotodo-api/bin /go/bin

# Exposes port 80 from the container
EXPOSE 80

ENTRYPOINT /go/bin/gotodo-api --port 3200