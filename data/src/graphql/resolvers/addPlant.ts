export async function addPlant(
  args: object,
  {mongodb}: {mongodb: any},
): Promise<object> {
  return await mongodb.plants.insertOne({
    _id: String(Date.now()),
    ...args,
  });
}
