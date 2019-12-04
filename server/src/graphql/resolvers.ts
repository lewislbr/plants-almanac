import { Plant } from '../models/plant';
import { IResolvers } from 'graphql-tools';

export const resolvers: IResolvers = {
  Query: {
    getPlant: async (_, { name }): Promise<any> => {
      const plant = await Plant.findOne({ name });
      return plant;
    },
    getPlants: async (): Promise<any> => {
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
    ): Promise<any> => {
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
