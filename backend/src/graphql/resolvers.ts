import {ApolloError} from 'apollo-server';
import {IResolvers} from 'graphql-tools';

export const resolvers: IResolvers = {
  Query: {
    async getPlant(_, args, context): Promise<any | null> {
      try {
        const requestedPlant = await context.mongodb.plants.findOne(args);
        return requestedPlant;
      } catch (error) {
        throw new ApolloError(error);
      }
    },
    async getPlants(_, __, context): Promise<any | null> {
      try {
        const allPlants = await context.mongodb.plants.find().toArray();
        return allPlants;
      } catch (error) {
        throw new ApolloError(error);
      }
    },
  },

  Mutation: {
    async addPlant(_, args, context): Promise<any> {
      try {
        const addedPlant = await context.mongodb.plants.insertOne(args);
        return addedPlant;
      } catch (error) {
        throw new ApolloError(error);
      }
    },
    async deletePlant(_, args, context): Promise<any | null> {
      try {
        const deletedPlant = await context.mongodb.plants.deleteOne(args);
        return deletedPlant;
      } catch (error) {
        throw new ApolloError(error);
      }
    },
  },
};
