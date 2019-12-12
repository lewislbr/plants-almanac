import express from 'express';
import { ApolloServer } from 'apollo-server-express';
import mongoose from 'mongoose';

import { typeDefs } from './graphql/typeDefs';
import { resolvers } from './graphql/resolvers';

const app = express();

const server = new ApolloServer({
  typeDefs,
  resolvers,
});

server.applyMiddleware({ app, path: '/graphql' });

const startServer = async (): Promise<void> => {
  await mongoose
    .connect(
      `mongodb+srv://${process.env.MONGODB_USER}:${process.env.MONGODB_PASSWORD}@cluster0-ovl3w.mongodb.net/${process.env.MONGODB_DB}?retryWrites=true&w=majority`,
      { useNewUrlParser: true, useUnifiedTopology: true }
    )
    .then(() => console.log('MongoDB connected'))
    .catch((error) => console.log(error));

  const PORT = process.env.PORT || 4040;

  app.listen(PORT, () => {
    console.log(
      `Server ready at http://localhost:${PORT}${server.graphqlPath} 🚀`
    );
  });
};

startServer();
