FROM alpine:3.16.2

# Installs nodejs and npm
RUN apk add --update nodejs npm

# Creates a app folder in the container
RUN mkdir -p /app

# Change our directory to the app folder
WORKDIR /app

# Copy the files from frontend to the app folder
COPY frontend/ .

# Installs the dependencies of the web application
RUN npm install

# Builds the app
RUN npm run build

# Starts the application
CMD ["npm", "start"]