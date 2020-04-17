import {ApolloError} from 'apollo-server';
import {Document} from 'mongoose';
import {IResolvers} from 'graphql-tools';

import {Plant} from '../models/Plant';

export const resolvers: IResolvers = {
  Query: {
    async getPlant(_, args): Promise<Document | null> {
      try {
        const plant = await Plant.findOne({...args});
        return plant;
      } catch (error) {
        throw new ApolloError(error);
      }
    },
    async getPlants(): Promise<Document[] | null> {
      try {
        const plants = await Plant.find();
        return plants;
      } catch (error) {
        throw new ApolloError(error);
      }
    },
  },

  Mutation: {
    async addPlant(_, args): Promise<Document> {
      try {
        const newPlant = new Plant({...args});
        const addedPlant = await newPlant.save();
        return addedPlant;
      } catch (error) {
        throw new ApolloError(error);
      }
    },
    async deletePlant(_, args): Promise<Document | null> {
      try {
        const deletedPlant = await Plant.findOneAndDelete({...args});
        return deletedPlant;
      } catch (error) {
        throw new ApolloError(error);
      }
    },
  },
};
