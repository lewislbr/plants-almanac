export async function getPlants(
  _: Record<string, unknown>,
  {mongodb}: {mongodb: any},
): Promise<Record<string, unknown>[] | null> {
  return await mongodb.plants.find().toArray();
}
