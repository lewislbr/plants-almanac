import mongoose, { Schema, Document } from 'mongoose';

interface Plant extends Document {
  _doc: any;
  name: string;
  description: string;
  plantSeason: string;
  harvestSeason: string;
  pruneSeason: string;
  tips: string;
}

const plantSchema: Schema = new Schema({
  name: {
    type: String,
    required: true,
  },
  otherNames: {
    type: String,
    required: false,
  },
  description: {
    type: String,
    required: false,
  },
  plantSeason: {
    type: String,
    required: false,
  },
  harvestSeason: {
    type: String,
    required: false,
  },
  pruneSeason: {
    type: String,
    required: false,
  },
  tips: {
    type: String,
    required: false,
  },
});

export const Plant = mongoose.model<Plant>('Plant', plantSchema);
