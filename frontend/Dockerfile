FROM node:17
WORKDIR /app
COPY . .

RUN npm ci
RUN npm run build
ENV NODE_ENV production

CMD ["npm", "run", "start"]
