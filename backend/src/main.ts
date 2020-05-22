import express from "express";
import {buildSchema} from "graphql";
import {importSchema} from "graphql-import";
import graphqlHTTP from "express-graphql";
import expressPlayground from "graphql-playground-middleware-express";

import {connectDatabase} from "./repository/mongodb";
import {resolvers} from "./graphql/resolvers/resolvers";

const server = express();

server.disable("x-powered-by");

server.use((req, res, next) => {
  res.header("Access-Control-Allow-Headers", "Content-Type, Origin");
  res.header(
    "Access-Control-Allow-Origin",
    process.env.NODE_ENV === "development"
      ? process.env.DEVELOPMENT_URL
      : process.env.PRODUCTION_URL,
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

  server.use(
    "/graphql",
    graphqlHTTP({
      schema: buildSchema(importSchema("**/*.graphql")),
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

  const port = process.env.PORT || 4000;

  server.listen(port, () =>
    console.log(`Server ready at localhost:${port} ✅`),
  );
}

startServer();
