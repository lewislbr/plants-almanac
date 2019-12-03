import { ApolloServer } from 'apollo-server';
import mongoose from 'mongoose';

import { typeDefs } from './graphql/schemas';
import { resolvers } from './graphql/resolvers';

const server = new ApolloServer({
  cors: {
    origin: '*',
    methods: 'POST, GET, OPTIONS',
  },
  typeDefs,
  resolvers,
  engine: {
    apiKey: process.env.ENGINE_API_KEY,
  },
});

mongoose
  .connect(
    `mongodb+srv://${process.env.MONGODB_USER}:${process.env.MONGODB_PASSWORD}@cluster0-ovl3w.mongodb.net/${process.env.MONGODB_DB}?retryWrites=true&w=majority`,
    { useNewUrlParser: true, useUnifiedTopology: true }
  )
  .catch((error) => console.log(error));

server.listen({ port: 4040 }).then(({ url }) => {
  console.log(`Server ready at ${url} 🚀`);
});
