FROM node:18.7-alpine

WORKDIR /chat_app_build

COPY package.json .
COPY package-lock.json .

RUN npm install && \
    npm cache rm --force

WORKDIR /chat_app

COPY . .

EXPOSE 8080

ENTRYPOINT ["/bin/sh", "./entrypoint.sh"]

CMD ["npm", "run", "serve"]
