import express from 'express';
import graphqlHTTP from 'express-graphql';
import { buildSchema } from 'graphql';
import mongoose from 'mongoose';

import Plant from '../models/plant';

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
    schema: buildSchema(`
      type Plant {
        _id: ID!
        name: String!
        description: String
        plantSeason: [String!]
        harvestSeason: [String!]
        pruneSeason: [String!]
        tips: String
      }

      input PlantInput {
        name: String!
        description: String
        plantSeason: [String!]
        harvestSeason: [String!]
        pruneSeason: [String!]
        tips: String
      }

      type RootQuery {
        plants: [Plant!]!
      }

      type RootMutation {
        createPlant(plantInput: PlantInput): Plant
      }
      
      schema {
        query: RootQuery
        mutation: RootMutation
      }
    `),
    rootValue: {
      plants: (): Promise<any> => {
        return Plant.find()
          .then((plants) => {
            return plants.map((plant) => {
              return { ...plant._doc };
            });
          })
          .catch((error) => {
            throw error;
          });
      },
      createPlant: ({ args }: any): Promise<any> => {
        const plant = new Plant({
          name: args.plantInput.name,
          description: args.plantInput.description,
          plantSeason: args.plantInput.plantSeason,
          harvestSeason: args.plantInput.harvestSeason,
          pruneSeason: args.plantInput.pruneSeason,
          tips: args.plantInput.tips,
        });
        return plant
          .save()
          .then((result) => {
            console.log(result);
            return { ...result._doc };
          })
          .catch((error) => {
            console.log(error);
            throw error;
          });
      },
    },
    graphiql: true,
  })
);
