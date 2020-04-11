import {ApolloServer} from 'apollo-server';
import mongoose from 'mongoose';
import dotenv from 'dotenv';

import {typeDefs} from './graphql/typeDefs';
import {resolvers} from './graphql/resolvers';

dotenv.config();

const server = new ApolloServer({
  typeDefs,
  resolvers,
});

async function startServer(): Promise<void> {
  await mongoose
    .connect(
      `mongodb+srv://${process.env.MONGODB_USER}:${process.env.MONGODB_PASSWORD}@cluster0-ovl3w.mongodb.net/${process.env.MONGODB_DB}?retryWrites=true&w=majority`,
      {useNewUrlParser: true, useUnifiedTopology: true},
    )
    .then(() => console.log('MongoDB connected'))
    .catch(error => console.log(error));

  server.listen().then(({url}) => {
    console.log(`🚀 Server ready at ${url}`);
  });
}

startServer();
