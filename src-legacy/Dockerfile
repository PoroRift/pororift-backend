# For development use only!

FROM node

RUN mkdir /usr/workspace

RUN mkdir /usr/app

RUN apt-get update

RUN apt-get install -y git

WORKDIR /usr/app

COPY . .

RUN touch api_key

RUN npm install

RUN ./updateData.sh

CMD ["npm", "run", "start"]

# CMD ["nodemon", "--exec", "./node_modules/.bin/babel", "src/index.js", "--ignore",  "data/"]

