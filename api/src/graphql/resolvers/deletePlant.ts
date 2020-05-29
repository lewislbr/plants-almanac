export async function deletePlant(
  {_id}: {_id: string},
  {mongodb}: {mongodb: any},
): Promise<object | null> {
  return await mongodb.plants.deleteOne({_id});
}
