FROM node:22-alpine

WORKDIR /var/www/html

COPY package*.json ./
RUN npm install

# Copy source and build
COPY . .

EXPOSE 5173

RUN npm run dev
