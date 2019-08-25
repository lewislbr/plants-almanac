import express from 'express';
import graphqlHTTP from 'express-graphql';
import mongoose from 'mongoose';

import schema from '../graphql/schemas';
import resolver from '../graphql/resolvers';

const app = express();
const port = 4040;

mongoose
  .connect(
    `mongodb+srv://${process.env.MONGODB_USER}:${process.env.MONGODB_PASSWORD}@cluster0-ovl3w.mongodb.net/${process.env.MONGODB_DB}?retryWrites=true&w=majority`
  )
  .then(() => {
    app.get('/', (req, res) => {
      res.status(200).send('Server working');
    });

    app.listen(port, () => {
      console.log(`App listening on port: ${port}`);
    });
  })
  .catch((error) => {
    console.log(error);
  });

app.use(
  '/graphql',
  graphqlHTTP({
    schema: schema,
    rootValue: resolver,
    graphiql: true,
  })
);
