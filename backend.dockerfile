FROM golang:1.19-alpine AS build


# Changes our current directory to the app folder
WORKDIR /app

# Copies the files to the current directory
COPY ./backend/ .

# Builds the binary
RUN CGO_ENABLED=0 go build -v -o /bin/api-server ./cmd/main.go

# Multibuild
FROM busybox
COPY --from=build /bin/api-server /bin/

# Starts the server
CMD ["/bin/api-server"]