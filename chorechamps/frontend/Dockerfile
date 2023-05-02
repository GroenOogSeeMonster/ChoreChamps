# Use the official Node.js image as the base image
FROM node:14 as build

# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json files
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application files
COPY . .

# Build the application
RUN npm run build

# Use a minimal image for deployment
FROM nginx:1.21-alpine

# Copy the build output to the Nginx html directory
COPY --from=build /app/build /usr/share/nginx/html

# Expose the Nginx server port
EXPOSE 80

# Start the Nginx server
CMD ["nginx", "-g", "daemon off;"]