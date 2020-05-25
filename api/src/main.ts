import express from "express";
import {readFileSync} from "fs";
import {buildSchema} from "graphql";
import graphqlHTTP from "express-graphql";
import expressPlayground from "graphql-playground-middleware-express";
import dotenv from "dotenv";

import {connectDatabase} from "./repository/mongodb";
import {resolvers} from "./graphql/resolvers/resolvers";

dotenv.config();

const server = express();

server.disable("x-powered-by");

server.use((req, res, next) => {
  res.header("Access-Control-Allow-Headers", "Content-Type, Origin");
  res.header(
    "Access-Control-Allow-Origin",
    process.env.NODE_ENV === "production"
      ? process.env.PRODUCTION_URL
      : process.env.DEVELOPMENT_URL,
  );

  if (process.env.NODE_ENV === "production") {
    res.header("Content-Security-Policy", "default-src 'self'");
    res.header(
      "Strict-Transport-Security",
      "max-age=63072000; includeSubDomains; preload",
    );
  }

  if (req.method === "OPTIONS") {
    res.sendStatus(204);
  } else {
    next();
  }
});

async function startServer(): Promise<void> {
  const mongodb = await connectDatabase();

  const schema = readFileSync(
    __dirname + "/graphql/schema/schema.graphql",
    "utf8",
  );

  server.use(
    "/graphql",
    graphqlHTTP({
      schema: buildSchema(schema),
      rootValue: resolvers,
      context: {mongodb},
      graphiql: false,
      customFormatErrorFn: error => ({
        message: error.message,
        locations: error.locations,
        path: error.path,
        stack: error.stack ? error.stack.split("\n") : [],
      }),
    }),
  );

  if (process.env.NODE_ENV === "development") {
    server.get("/playground", expressPlayground({endpoint: "/graphql"}));
  }

  const port = process.env.PORT || 4040;

  server.listen(port, () => console.log(`API ready at localhost:${port} ✅`));
}

startServer();
