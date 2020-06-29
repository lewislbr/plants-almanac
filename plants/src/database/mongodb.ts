import {MongoClient} from "mongodb";
import dotenv from "dotenv";

dotenv.config();

export async function connectDatabase(): Promise<
  Record<string, unknown> | undefined
> {
  try {
    const uri = String(process.env.MONGODB_URI);
    const cluster = await MongoClient.connect(uri, {
      useNewUrlParser: true,
      useUnifiedTopology: true,
    });
    const database = cluster.db("plants");

    console.log("MongoDB database connected ✅");

    return {plants: database.collection("plants")};
  } catch (error) {
    console.log(error);
  }
}
