import {Document} from 'mongoose';
import {IResolvers} from 'graphql-tools';

import {Plant} from '../models/Plant';

export const resolvers: IResolvers = {
  Query: {
    async getPlant(_parent, {name}): Promise<Document | null> {
      const plant = await Plant.findOne({name});
      return plant;
    },
    async getPlants(): Promise<Document[]> {
      const plants = await Plant.find();
      return plants;
    },
  },

  Mutation: {
    async addPlant(
      _parent,
      {
        name,
        otherNames,
        description,
        plantSeason,
        harvestSeason,
        pruneSeason,
        tips,
      },
    ): Promise<Document> {
      const newPlant = new Plant({
        name,
        otherNames,
        description,
        plantSeason,
        harvestSeason,
        pruneSeason,
        tips,
      });
      const addedPlant = await newPlant.save();
      return addedPlant;
    },
    async deletePlant(_parent, {_id}): Promise<Document | null> {
      const deletedPlant = await Plant.findByIdAndDelete({_id});
      return deletedPlant;
    },
  },
};
