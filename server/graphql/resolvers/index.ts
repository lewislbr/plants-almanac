import { Plant } from '../../models/plant';

export const resolver = {
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

  createPlant: async (args: {
    input: {
      name: string;
      description: string;
      plantSeason: string[];
      harvestSeason: string[];
      pruneSeason: string[];
      tips: string;
    };
  }): Promise<any> => {
    try {
      const plant = new Plant({
        name: args.input.name,
        description: args.input.description,
        plantSeason: args.input.plantSeason,
        harvestSeason: args.input.harvestSeason,
        pruneSeason: args.input.pruneSeason,
        tips: args.input.tips,
      });
      const result = await plant.save();
      console.log(result);
      return { ...result._doc };
    } catch (error) {
      console.log(error);
      throw error;
    }
  },
};
