export async function getPlant(
  {name}: {name: string},
  {mongodb}: {mongodb: any},
): Promise<object | null> {
  return await mongodb.plants.findOne({name});
}
