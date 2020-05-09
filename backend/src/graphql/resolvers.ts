import {ApolloError} from 'apollo-server';
import {IResolvers} from 'graphql-tools';

export const resolvers: IResolvers = {
  Query: {
    async getPlant(_, args, {mongodb: {plants}}): Promise<any | null> {
      try {
        const requestedPlant = await plants.findOne(args);
        return requestedPlant;
      } catch (error) {
        throw new ApolloError(error);
      }
    },
    async getPlants(_, __, {mongodb: {plants}}): Promise<any | null> {
      try {
        const allPlants = await plants.find().toArray();
        return allPlants;
      } catch (error) {
        throw new ApolloError(error);
      }
    },
  },

  Mutation: {
    async addPlant(_, args, {mongodb: {plants}}): Promise<any> {
      try {
        const addedPlant = await plants.insertOne({
          _id: String(Date.now()),
          ...args,
        });
        return addedPlant;
      } catch (error) {
        throw new ApolloError(error);
      }
    },
    async deletePlant(_, args, {mongodb: {plants}}): Promise<any | null> {
      try {
        const deletedPlant = await plants.deleteOne(args);
        return deletedPlant;
      } catch (error) {
        throw new ApolloError(error);
      }
    },
  },
};
