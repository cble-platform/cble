import { ApolloClient, InMemoryCache, split } from "@apollo/client";
import { getMainDefinition } from "@apollo/client/utilities";
import { GraphQLWsLink } from "@apollo/client/link/subscriptions";
import { createClient } from "graphql-ws";
import createUploadLink from "apollo-upload-client/createUploadLink.mjs";
import { Kind, OperationTypeNode } from "graphql";

// eslint-disable-next-line @typescript-eslint/no-unsafe-call
const uploadLink = createUploadLink({
  uri: new URL("/api/graphql/query", import.meta.env.VITE_API_BASE_URL).toString(),
  credentials: "include",
});

const wsLink = new GraphQLWsLink(
  createClient({
    url: new URL("/api/graphql/query", import.meta.env.VITE_API_BASE_URL).toString(),
  })
);

const splitLink = split(
  ({ query }) => {
    const definition = getMainDefinition(query);
    return definition.kind === Kind.OPERATION_DEFINITION && definition.operation === OperationTypeNode.SUBSCRIPTION;
  },
  wsLink,
  uploadLink
);

export const client = new ApolloClient({
  link: splitLink,
  cache: new InMemoryCache(),
});
