import {ApolloServer} from 'apollo-server';
import mongoose from 'mongoose';
import dotenv from 'dotenv';

import {typeDefs} from './graphql/typeDefs';
import {resolvers} from './graphql/resolvers';

dotenv.config();

const server = new ApolloServer({
  typeDefs,
  resolvers,
  context: ({res}): void => {
    res.header(
      'Access-Control-Allow-Origin',
      process.env.NODE_ENV === 'production'
        ? 'https://plants-almanac.netlify.app'
        : 'http://localhost:7001',
    );
  },
  engine: {
    apiKey: process.env.APOLLO_KEY,
  },
});

async function startServer(): Promise<void> {
  await mongoose
    .connect(
      `mongodb+srv://${process.env.MONGODB_USER}:${process.env.MONGODB_PASSWORD}@cluster0-ovl3w.mongodb.net/${process.env.MONGODB_DB}?retryWrites=true&w=majority`,
      {useNewUrlParser: true, useUnifiedTopology: true},
    )
    .then(() => console.log('MongoDB connected'))
    .catch(error => console.log(error));

  server.listen({port: process.env.PORT || 4000}).then(({url}) => {
    console.log(`🚀 Server ready at ${url}`);
  });
}

startServer();
