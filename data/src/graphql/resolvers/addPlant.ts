export async function addPlant(
  args: Record<string, unknown>,
  {mongodb}: {mongodb: any},
): Promise<Record<string, unknown>> {
  return await mongodb.plants.insertOne({
    _id: Date.now().toString(),
    ...args,
  });
}
