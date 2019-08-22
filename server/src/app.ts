const express = require('express');
const graphqlHTTP = require('express-graphql');
const { buildSchema } = require('graphql');

const app = express();

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
app.listen(4040);
console.log('Running a GraphQL API server at localhost:4040/graphql');
