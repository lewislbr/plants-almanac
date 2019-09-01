import { ApolloServer } from 'apollo-server';
import mongoose from 'mongoose';

import { typeDefs } from '../graphql/schemas';
import { resolvers } from '../graphql/resolvers';

const server = new ApolloServer({
  cors: {
    origin: '*',
    methods: 'POST, GET, OPTIONS',
  },
  typeDefs,
  resolvers,
});

mongoose.connect(
  `mongodb+srv://${process.env.MONGODB_USER}:${process.env.MONGODB_PASSWORD}@cluster0-ovl3w.mongodb.net/${process.env.MONGODB_DB}?retryWrites=true&w=majority`
);

server.listen().then(({ url }) => {
  console.log(`Server ready at ${url} 🚀`);
});
