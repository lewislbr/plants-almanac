export async function getPlant(
  {name}: {name: string},
  {mongodb}: {mongodb: any},
): Promise<Record<string, unknown> | null> {
  return await mongodb.plants.findOne({name});
}
