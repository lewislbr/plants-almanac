export async function getPlants(
  _: object,
  {mongodb}: {mongodb: any},
): Promise<object[] | null> {
  return await mongodb.plants.find().toArray();
}
