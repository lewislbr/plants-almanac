import express from 'express';
import graphqlHTTP from 'express-graphql';
import { buildSchema } from 'graphql';

const app = express();
const port = 4040;

const schema = buildSchema(`
  type Query {
    plant: String
  }
`);

const root = {
  plant: () => {
    return 'ROSEMARY';
  },
};

app.use(
  '/graphql',
  graphqlHTTP({
    schema: schema,
    rootValue: root,
    graphiql: true,
  })
);

app.get('/', (req, res) => {
  res.status(200).send('Server working');
});

app.listen(port, function() {
  console.log('App listening on port: ' + port);
});
