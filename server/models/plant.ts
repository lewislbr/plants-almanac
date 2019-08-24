import mongoose from 'mongoose';

const Schema = mongoose.Schema;

const plantSchema = new Schema({
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

const Plant = mongoose.model('Plant', plantSchema);

export default Plant;
