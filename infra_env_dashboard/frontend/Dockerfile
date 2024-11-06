# Stage 1: Build
FROM node:16

# Set working directory
WORKDIR /app

# Copy package.json and package-lock.json to the working directory
COPY package.json package-lock.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

# Build the application
RUN npm run build

# Start the app directly (remove nginx stage)
# Assuming the app runs on port 3000; adjust as needed
EXPOSE 3000

CMD ["npm", "start"]