import {ApolloServer} from 'apollo-server';
import {MongoClient} from 'mongodb';
import dotenv from 'dotenv';

import {typeDefs} from './graphql/typeDefs';
import {resolvers} from './graphql/resolvers';

dotenv.config();

async function connectDatabase(): Promise<any> {
  const uri = `mongodb+srv://${process.env.MONGODB_USER}:${process.env.MONGODB_PASSWORD}@cluster0-ovl3w.mongodb.net/test?retryWrites=true&w=majority`;
  const cluster = await MongoClient.connect(uri, {
    useNewUrlParser: true,
    useUnifiedTopology: true,
  });
  const database = cluster.db('plants-almanac');

  return {plants: database.collection('plants')};
}

async function startServer(): Promise<void> {
  try {
    const mongodb = await connectDatabase();

    const server = new ApolloServer({
      cors: {
        origin: true,
      },
      typeDefs,
      resolvers,
      context: {mongodb},
      engine: {
        apiKey: process.env.APOLLO_KEY,
      },
    });

    server.listen({port: process.env.PORT || 4000}).then(({url}) => {
      console.log(`🚀 Server ready at ${url}`);
    });
  } catch (error) {
    console.log(error);
  }
}

startServer();
