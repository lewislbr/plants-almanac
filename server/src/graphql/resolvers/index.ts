import { Plant } from '../../models/plant';
import { IResolvers } from 'graphql-tools';

interface Plant {
  name: string;
  description: string;
  plantSeason: string;
  harvestSeason: string;
  pruneSeason: string;
  tips: string;
}

export const resolvers: IResolvers = {
  Query: {
    getPlant: async (_: any, { name }): Promise<any> => {
      try {
        const result = await Plant.findOne({ name: name });
        return result;
      } catch (error) {
        console.log(error);
        throw error;
      }
    },
    getPlants: async (): Promise<any> => {
      try {
        const plants = await Plant.find();
        return plants.map((plant) => {
          return { ...plant._doc };
        });
      } catch (error) {
        console.log(error);
        throw error;
      }
    },
  },

  Mutation: {
    createPlant: async (_: any, args: Plant): Promise<any> => {
      try {
        const plant = new Plant({
          name: args.name,
          description: args.description,
          plantSeason: args.plantSeason,
          harvestSeason: args.harvestSeason,
          pruneSeason: args.pruneSeason,
          tips: args.tips,
        });
        const result = await plant.save();
        return { ...result._doc };
      } catch (error) {
        console.log(error);
        throw error;
      }
    },
  },
};
