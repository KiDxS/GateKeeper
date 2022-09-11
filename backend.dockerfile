FROM golang:1.19

# Creates a app folder
RUN mkdir -p /app

# Changes our current directory to the app folder
WORKDIR /app

# Copies the dependencies
COPY backend/go.mod backend/go.sum ./

# Installs the dependencies
RUN go mod download && go mod verify

# Copies the whole codebase of the backend to the app folder.
COPY backend/ .

# Builds a executable file
RUN go build -o cmd/server cmd/main.go

# Starts the server
CMD ["cmd/server"]