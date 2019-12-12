import { Plant } from '../models/plant';
import { IResolvers } from 'graphql-tools';
import { Document } from 'mongoose';

export const resolvers: IResolvers = {
  Query: {
    getPlant: async (_, { name }): Promise<Document | null> => {
      const plant = await Plant.findOne({ name });
      return plant;
    },
    getPlants: async (): Promise<Document[]> => {
      const plants = await Plant.find();
      return plants;
    },
  },

  Mutation: {
    createPlant: async (
      _,
      {
        name,
        otherNames,
        description,
        plantSeason,
        harvestSeason,
        pruneSeason,
        tips,
      }
    ): Promise<Document> => {
      const newPlant = new Plant({
        name,
        otherNames,
        description,
        plantSeason,
        harvestSeason,
        pruneSeason,
        tips,
      });
      const savedPlant = await newPlant.save();
      return savedPlant;
    },
  },
};
