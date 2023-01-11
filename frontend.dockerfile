FROM alpine:3.16.2

# Installs nodejs and npm
RUN apk add --update nodejs npm

# Change our directory to the app folder
WORKDIR /app

# Copy the files from frontend to the app folder
COPY ./frontend/ /app

# Installs the dependencies and builds it
RUN npm install && npm run build
