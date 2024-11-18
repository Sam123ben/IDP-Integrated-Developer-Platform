# Stage 1: Build
FROM node:16 AS build

WORKDIR /app

COPY package.json ./

RUN npm install

COPY . .

RUN npm run build

# Stage 2: Run
FROM node:16-slim

RUN npm install -g serve

WORKDIR /app

# Copy the build artifacts from the build stage
COPY --from=build /app/dist ./dist

EXPOSE 3000

# Start the application
CMD ["serve", "-s", "dist", "-l", "3000"]