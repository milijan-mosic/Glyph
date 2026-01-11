FROM node:24-bookworm

WORKDIR /app
# COPY package.json package-lock.json ./
COPY . ./
RUN npm ci

EXPOSE 3000
CMD ["npm", "run", "dev"]
