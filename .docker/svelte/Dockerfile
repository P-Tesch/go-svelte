FROM node:22-alpine

WORKDIR /var/www/html

COPY ./src/frontend/package.json ./

RUN npm install

EXPOSE 5173

CMD [ "npm", "run", "dev", "--", "--host" ]
