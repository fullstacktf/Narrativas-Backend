FROM node:14.14.0
RUN apt-get update

WORKDIR /app
COPY src/* /app

EXPOSE 80/tcp
CMD ["node", "app.js"]

