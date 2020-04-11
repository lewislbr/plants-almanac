import mongoose, {Schema} from 'mongoose';

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

export const Plant = mongoose.model('Plant', plantSchema);
