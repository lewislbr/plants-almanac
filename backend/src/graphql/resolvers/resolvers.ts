export const resolvers = {
  async getPlant(args: object, context: any): Promise<object | null> {
    return await context.mongodb.plants.findOne(args);
  },
  async getPlants(args: object, context: any): Promise<object[] | null> {
    return await context.mongodb.plants.find().toArray();
  },
  async addPlant(args: object, context: any): Promise<object> {
    return await context.mongodb.plants.insertOne({
      _id: String(Date.now()),
      ...args,
    });
  },
  async deletePlant(args: object, context: any): Promise<object | null> {
    return await context.mongodb.plants.deleteOne(args);
  },
};
