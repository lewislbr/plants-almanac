import mongoose, { Schema, Document } from 'mongoose';

export interface Plant extends Document {
  _doc: any;
  name: string;
  description: string;
  plantSeason: string[];
  harvestSeason: string[];
  pruneSeason: string[];
  tips: string;
}

const plantSchema: Schema = new Schema({
  name: {
    type: String,
    required: true,
  },
  description: {
    type: String,
    required: false,
  },
  plantSeason: {
    type: [String],
    required: false,
  },
  harvestSeason: {
    type: [String],
    required: false,
  },
  pruneSeason: {
    type: [String],
    required: false,
  },
  tips: {
    type: String,
    required: false,
  },
});

export default mongoose.model<Plant>('Plant', plantSchema);
