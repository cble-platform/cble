import type { CodegenConfig } from "@graphql-codegen/cli";

const config: CodegenConfig = {
  overwrite: true,
  schema: "src/api/graphql/schema.graphqls",
  documents: "src/api/graphql/**/*.graphql",
  generates: {
    "src/api/graphql/generated/index.tsx": {
      plugins: ["typescript", "typescript-operations", "typescript-react-apollo"],
    },
  },
};

export default config;
