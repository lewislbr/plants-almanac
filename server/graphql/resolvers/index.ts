import { Plant } from '../../models/plant';
import { IResolvers } from 'graphql-tools';

export const resolvers: IResolvers = {
  Query: {
    plants: async (): Promise<any> => {
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
    createPlant: async (
      _: any,
      args: {
        name: string;
        description: string;
        plantSeason: string[];
        harvestSeason: string[];
        pruneSeason: string[];
        tips: string;
      }
    ): Promise<any> => {
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
        console.log(result);
        return { ...result._doc };
      } catch (error) {
        console.log(error);
        throw error;
      }
    },
  },
};
