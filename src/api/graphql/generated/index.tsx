import { gql } from '@apollo/client';
import * as Apollo from '@apollo/client';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
const defaultOptions = {} as const;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  Map: { input: any; output: any; }
  StrMap: { input: any; output: any; }
  Time: { input: any; output: any; }
};

export type Blueprint = {
  __typename?: 'Blueprint';
  blueprintTemplate: Scalars['String']['output'];
  createdAt: Scalars['Time']['output'];
  deployments: Array<Maybe<Deployment>>;
  description: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  parentGroup: Group;
  provider: Provider;
  updatedAt: Scalars['Time']['output'];
};

export type BlueprintInput = {
  blueprintTemplate: Scalars['String']['input'];
  description: Scalars['String']['input'];
  name: Scalars['String']['input'];
  parentGroupId: Scalars['ID']['input'];
  providerId: Scalars['ID']['input'];
};

export enum CommandStatus {
  Failed = 'FAILED',
  Inprogress = 'INPROGRESS',
  Queued = 'QUEUED',
  Succeeded = 'SUCCEEDED'
}

export enum CommandType {
  Configure = 'CONFIGURE',
  Deploy = 'DEPLOY',
  Destroy = 'DESTROY'
}

export type Deployment = {
  __typename?: 'Deployment';
  blueprint: Blueprint;
  createdAt: Scalars['Time']['output'];
  deploymentState: Scalars['StrMap']['output'];
  deploymentVars: Scalars['Map']['output'];
  description: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  requester: User;
  templateVars: Scalars['Map']['output'];
  updatedAt: Scalars['Time']['output'];
};

export type DeploymentInput = {
  name: Scalars['String']['input'];
};

export type Group = {
  __typename?: 'Group';
  blueprints?: Maybe<Array<Maybe<Blueprint>>>;
  children?: Maybe<Array<Maybe<Group>>>;
  createdAt: Scalars['Time']['output'];
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  parent?: Maybe<Group>;
  permissionPolicies?: Maybe<Array<Maybe<PermissionPolicy>>>;
  updatedAt: Scalars['Time']['output'];
  users?: Maybe<Array<Maybe<User>>>;
};

export type Mutation = {
  __typename?: 'Mutation';
  /** Applies the stored configuration to the provider (requires permission `com.cble.providers.configure`) */
  configureProvider: Provider;
  /** Create a blueprint (requires permission `com.cble.blueprints.create`) */
  createBlueprint: Blueprint;
  /** Create a provider (requires permission `com.cble.providers.create`) */
  createProvider: Provider;
  /** Create a user (requires permission `com.cble.users.create`) */
  createUser: User;
  /** Delete a blueprint (requires permission `com.cble.blueprints.delete`) */
  deleteBlueprint: Scalars['Boolean']['output'];
  /** Delete a provider (requires permission `com.cble.providers.delete`) */
  deleteProvider: Scalars['Boolean']['output'];
  /** Delete a user (requires permission `com.cble.users.delete`) */
  deleteUser: Scalars['Boolean']['output'];
  /** Deploy a blueprint (requires permission `com.cble.blueprints.deploy`) */
  deployBlueprint: Deployment;
  /** Destroy a deployment (requires permission `com.cble.deployments.destroy`) */
  destroyDeployment: Deployment;
  /** Load a provider to connect it to CBLE (requires permission `com.cble.providers.load`) */
  loadProvider: Provider;
  /** Change current user's password */
  selfChangePassword: Scalars['Boolean']['output'];
  /** Unload a provider to disconnect it from CBLE (requires permission `com.cble.providers.unload`) */
  unloadProvider: Provider;
  /** Update a blueprint (requires permission `com.cble.blueprints.update`) */
  updateBlueprint: Blueprint;
  /** Update a deployment (requires permission `com.cble.deployments.update`) */
  updateDeployment: Deployment;
  /** Update a provider (requires permission `com.cble.providers.update`) */
  updateProvider: Provider;
  /** Update a user (requires permission `com.cble.users.update`) */
  updateUser: User;
};


export type MutationConfigureProviderArgs = {
  id: Scalars['ID']['input'];
};


export type MutationCreateBlueprintArgs = {
  input: BlueprintInput;
};


export type MutationCreateProviderArgs = {
  input: ProviderInput;
};


export type MutationCreateUserArgs = {
  input: UserInput;
};


export type MutationDeleteBlueprintArgs = {
  id: Scalars['ID']['input'];
};


export type MutationDeleteProviderArgs = {
  id: Scalars['ID']['input'];
};


export type MutationDeleteUserArgs = {
  id: Scalars['ID']['input'];
};


export type MutationDeployBlueprintArgs = {
  id: Scalars['ID']['input'];
};


export type MutationDestroyDeploymentArgs = {
  id: Scalars['ID']['input'];
};


export type MutationLoadProviderArgs = {
  id: Scalars['ID']['input'];
};


export type MutationSelfChangePasswordArgs = {
  currentPassword: Scalars['String']['input'];
  newPassword: Scalars['String']['input'];
};


export type MutationUnloadProviderArgs = {
  id: Scalars['ID']['input'];
};


export type MutationUpdateBlueprintArgs = {
  id: Scalars['ID']['input'];
  input: BlueprintInput;
};


export type MutationUpdateDeploymentArgs = {
  id: Scalars['ID']['input'];
  input: DeploymentInput;
};


export type MutationUpdateProviderArgs = {
  id: Scalars['ID']['input'];
  input: ProviderInput;
};


export type MutationUpdateUserArgs = {
  id: Scalars['ID']['input'];
  input: UserInput;
};

export type Permission = {
  __typename?: 'Permission';
  createdAt: Scalars['Time']['output'];
  id: Scalars['ID']['output'];
  key?: Maybe<Scalars['String']['output']>;
  permissionPolicies?: Maybe<Array<Maybe<PermissionPolicy>>>;
  updatedAt: Scalars['Time']['output'];
};

export type PermissionPolicy = {
  __typename?: 'PermissionPolicy';
  createdAt: Scalars['Time']['output'];
  group: Group;
  id: Scalars['ID']['output'];
  permission: Permission;
  type: PermissionPolicyType;
  updatedAt: Scalars['Time']['output'];
};

export enum PermissionPolicyType {
  Allow = 'ALLOW',
  Deny = 'DENY'
}

export type Provider = {
  __typename?: 'Provider';
  blueprints?: Maybe<Array<Maybe<Blueprint>>>;
  configBytes: Scalars['String']['output'];
  createdAt: Scalars['Time']['output'];
  displayName: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  isLoaded: Scalars['Boolean']['output'];
  providerGitUrl: Scalars['String']['output'];
  providerVersion: Scalars['String']['output'];
  updatedAt: Scalars['Time']['output'];
};

export type ProviderCommand = {
  __typename?: 'ProviderCommand';
  commandType: CommandType;
  createdAt: Scalars['Time']['output'];
  endTime?: Maybe<Scalars['Time']['output']>;
  error: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  output: Scalars['String']['output'];
  startTime?: Maybe<Scalars['Time']['output']>;
  status: CommandStatus;
  updatedAt: Scalars['Time']['output'];
};

export type ProviderInput = {
  configBytes: Scalars['String']['input'];
  displayName: Scalars['String']['input'];
  providerGitUrl: Scalars['String']['input'];
  providerVersion: Scalars['String']['input'];
};

export type Query = {
  __typename?: 'Query';
  /** Get a blueprint (requires permission `com.cble.blueprints.read`) */
  blueprint: Blueprint;
  /** List blueprints (requires permission `com.cble.blueprints.list`) */
  blueprints: Array<Blueprint>;
  /** Get a deployment (requires permission `com.cble.deployments.read`) */
  deployment: Deployment;
  /** List deployments (requires permission `com.cble.deployments.list`) */
  deployments: Array<Deployment>;
  /** Get a group (requires permission `com.cble.groups.read`) */
  group: Group;
  /** List groups (requires permission `com.cble.groups.list`) */
  groups: Array<Group>;
  /** Get current user */
  me: User;
  /** Retrieves if the current user has a given permission */
  meHasPermission: Scalars['Boolean']['output'];
  /** Get a provider (requires permission `com.cble.providers.read`) */
  provider: Provider;
  /** Get a provider command (requires permission `com.cble.providercommands.read`) */
  providerCommand: ProviderCommand;
  /** List provider commands (requires permission `com.cble.providercommands.list`) */
  providerCommands: Array<ProviderCommand>;
  /** List providers (requires permission `com.cble.providers.list`) */
  providers: Array<Provider>;
  /** Get a user (requires permission `com.cble.users.read`) */
  user: User;
  /** List users (requires permission `com.cble.users.list`) */
  users: Array<User>;
};


export type QueryBlueprintArgs = {
  id: Scalars['ID']['input'];
};


export type QueryDeploymentArgs = {
  id: Scalars['ID']['input'];
};


export type QueryGroupArgs = {
  id: Scalars['ID']['input'];
};


export type QueryMeHasPermissionArgs = {
  key: Scalars['String']['input'];
};


export type QueryProviderArgs = {
  id: Scalars['ID']['input'];
};


export type QueryProviderCommandArgs = {
  id: Scalars['ID']['input'];
};


export type QueryUserArgs = {
  id: Scalars['ID']['input'];
};

export type User = {
  __typename?: 'User';
  createdAt: Scalars['Time']['output'];
  deployments: Array<Maybe<Deployment>>;
  email: Scalars['String']['output'];
  firstName: Scalars['String']['output'];
  groups: Array<Maybe<Group>>;
  id: Scalars['ID']['output'];
  lastName: Scalars['String']['output'];
  updatedAt: Scalars['Time']['output'];
  username: Scalars['String']['output'];
};

export type UserInput = {
  email: Scalars['String']['input'];
  firstName: Scalars['String']['input'];
  lastName: Scalars['String']['input'];
  username: Scalars['String']['input'];
};

export type BlueprintFragementFragment = { __typename?: 'Blueprint', id: string, createdAt: any, updatedAt: any, name: string, description: string, blueprintTemplate: string, parentGroup: { __typename?: 'Group', id: string, name: string }, provider: { __typename?: 'Provider', id: string, displayName: string, isLoaded: boolean }, deployments: Array<{ __typename?: 'Deployment', id: string } | null> };

export type BlueprintsQueryVariables = Exact<{ [key: string]: never; }>;


export type BlueprintsQuery = { __typename?: 'Query', blueprints: Array<{ __typename?: 'Blueprint', id: string, createdAt: any, updatedAt: any, name: string, description: string, blueprintTemplate: string, parentGroup: { __typename?: 'Group', id: string, name: string }, provider: { __typename?: 'Provider', id: string, displayName: string, isLoaded: boolean }, deployments: Array<{ __typename?: 'Deployment', id: string } | null> }> };

export type GetBlueprintQueryVariables = Exact<{
  id: Scalars['ID']['input'];
}>;


export type GetBlueprintQuery = { __typename?: 'Query', blueprint: { __typename?: 'Blueprint', id: string, createdAt: any, updatedAt: any, name: string, description: string, blueprintTemplate: string, parentGroup: { __typename?: 'Group', id: string, name: string }, provider: { __typename?: 'Provider', id: string, displayName: string, isLoaded: boolean }, deployments: Array<{ __typename?: 'Deployment', id: string } | null> } };

export type CreateBlueprintMutationVariables = Exact<{
  input: BlueprintInput;
}>;


export type CreateBlueprintMutation = { __typename?: 'Mutation', createBlueprint: { __typename?: 'Blueprint', id: string, createdAt: any, updatedAt: any, name: string, description: string, blueprintTemplate: string, parentGroup: { __typename?: 'Group', id: string, name: string }, provider: { __typename?: 'Provider', id: string, displayName: string, isLoaded: boolean }, deployments: Array<{ __typename?: 'Deployment', id: string } | null> } };

export type UpdateBlueprintMutationVariables = Exact<{
  id: Scalars['ID']['input'];
  input: BlueprintInput;
}>;


export type UpdateBlueprintMutation = { __typename?: 'Mutation', updateBlueprint: { __typename?: 'Blueprint', id: string, createdAt: any, updatedAt: any, name: string, description: string, blueprintTemplate: string, parentGroup: { __typename?: 'Group', id: string, name: string }, provider: { __typename?: 'Provider', id: string, displayName: string, isLoaded: boolean }, deployments: Array<{ __typename?: 'Deployment', id: string } | null> } };

export type DeployBlueprintMutationVariables = Exact<{
  id: Scalars['ID']['input'];
}>;


export type DeployBlueprintMutation = { __typename?: 'Mutation', deployBlueprint: { __typename?: 'Deployment', id: string, createdAt: any, updatedAt: any, name: string, templateVars: any, deploymentVars: any, deploymentState: any, blueprint: { __typename?: 'Blueprint', id: string, createdAt: any, updatedAt: any, name: string, description: string, blueprintTemplate: string, parentGroup: { __typename?: 'Group', id: string, name: string }, provider: { __typename?: 'Provider', id: string, displayName: string, isLoaded: boolean }, deployments: Array<{ __typename?: 'Deployment', id: string } | null> }, requester: { __typename?: 'User', id: string, createdAt: any, updatedAt: any, username: string, email: string, firstName: string, lastName: string } } };

export type DeploymentFragmentFragment = { __typename?: 'Deployment', id: string, createdAt: any, updatedAt: any, name: string, templateVars: any, deploymentVars: any, deploymentState: any, blueprint: { __typename?: 'Blueprint', id: string, createdAt: any, updatedAt: any, name: string, description: string, blueprintTemplate: string, parentGroup: { __typename?: 'Group', id: string, name: string }, provider: { __typename?: 'Provider', id: string, displayName: string, isLoaded: boolean }, deployments: Array<{ __typename?: 'Deployment', id: string } | null> }, requester: { __typename?: 'User', id: string, createdAt: any, updatedAt: any, username: string, email: string, firstName: string, lastName: string } };

export type ListDeploymentsQueryVariables = Exact<{ [key: string]: never; }>;


export type ListDeploymentsQuery = { __typename?: 'Query', deployments: Array<{ __typename?: 'Deployment', id: string, createdAt: any, updatedAt: any, name: string, templateVars: any, deploymentVars: any, deploymentState: any, blueprint: { __typename?: 'Blueprint', id: string, createdAt: any, updatedAt: any, name: string, description: string, blueprintTemplate: string, parentGroup: { __typename?: 'Group', id: string, name: string }, provider: { __typename?: 'Provider', id: string, displayName: string, isLoaded: boolean }, deployments: Array<{ __typename?: 'Deployment', id: string } | null> }, requester: { __typename?: 'User', id: string, createdAt: any, updatedAt: any, username: string, email: string, firstName: string, lastName: string } }> };

export type GetDeploymentQueryVariables = Exact<{
  id: Scalars['ID']['input'];
}>;


export type GetDeploymentQuery = { __typename?: 'Query', deployment: { __typename?: 'Deployment', id: string, createdAt: any, updatedAt: any, name: string, templateVars: any, deploymentVars: any, deploymentState: any, blueprint: { __typename?: 'Blueprint', id: string, createdAt: any, updatedAt: any, name: string, description: string, blueprintTemplate: string, parentGroup: { __typename?: 'Group', id: string, name: string }, provider: { __typename?: 'Provider', id: string, displayName: string, isLoaded: boolean }, deployments: Array<{ __typename?: 'Deployment', id: string } | null> }, requester: { __typename?: 'User', id: string, createdAt: any, updatedAt: any, username: string, email: string, firstName: string, lastName: string } } };

export type UpdateDeploymentMutationVariables = Exact<{
  id: Scalars['ID']['input'];
  input: DeploymentInput;
}>;


export type UpdateDeploymentMutation = { __typename?: 'Mutation', updateDeployment: { __typename?: 'Deployment', id: string, createdAt: any, updatedAt: any, name: string, templateVars: any, deploymentVars: any, deploymentState: any, blueprint: { __typename?: 'Blueprint', id: string, createdAt: any, updatedAt: any, name: string, description: string, blueprintTemplate: string, parentGroup: { __typename?: 'Group', id: string, name: string }, provider: { __typename?: 'Provider', id: string, displayName: string, isLoaded: boolean }, deployments: Array<{ __typename?: 'Deployment', id: string } | null> }, requester: { __typename?: 'User', id: string, createdAt: any, updatedAt: any, username: string, email: string, firstName: string, lastName: string } } };

export type DestroyDeploymentMutationVariables = Exact<{
  id: Scalars['ID']['input'];
}>;


export type DestroyDeploymentMutation = { __typename?: 'Mutation', destroyDeployment: { __typename?: 'Deployment', id: string, createdAt: any, updatedAt: any, name: string, templateVars: any, deploymentVars: any, deploymentState: any, blueprint: { __typename?: 'Blueprint', id: string, createdAt: any, updatedAt: any, name: string, description: string, blueprintTemplate: string, parentGroup: { __typename?: 'Group', id: string, name: string }, provider: { __typename?: 'Provider', id: string, displayName: string, isLoaded: boolean }, deployments: Array<{ __typename?: 'Deployment', id: string } | null> }, requester: { __typename?: 'User', id: string, createdAt: any, updatedAt: any, username: string, email: string, firstName: string, lastName: string } } };

export type GroupFragmentFragment = { __typename?: 'Group', id: string, createdAt: any, updatedAt: any, name: string };

export type ListGroupsQueryVariables = Exact<{ [key: string]: never; }>;


export type ListGroupsQuery = { __typename?: 'Query', groups: Array<{ __typename?: 'Group', id: string, createdAt: any, updatedAt: any, name: string }> };

export type ProviderFragmentFragment = { __typename?: 'Provider', id: string, createdAt: any, updatedAt: any, displayName: string, providerGitUrl: string, providerVersion: string, isLoaded: boolean };

export type ProvidersQueryVariables = Exact<{ [key: string]: never; }>;


export type ProvidersQuery = { __typename?: 'Query', providers: Array<{ __typename?: 'Provider', id: string, createdAt: any, updatedAt: any, displayName: string, providerGitUrl: string, providerVersion: string, isLoaded: boolean }> };

export type UserFragmentFragment = { __typename?: 'User', id: string, createdAt: any, updatedAt: any, username: string, email: string, firstName: string, lastName: string };

export type MeQueryVariables = Exact<{ [key: string]: never; }>;


export type MeQuery = { __typename?: 'Query', me: { __typename?: 'User', id: string, createdAt: any, updatedAt: any, username: string, email: string, firstName: string, lastName: string } };

export type ListUsersQueryVariables = Exact<{ [key: string]: never; }>;


export type ListUsersQuery = { __typename?: 'Query', users: Array<{ __typename?: 'User', id: string, createdAt: any, updatedAt: any, username: string, email: string, firstName: string, lastName: string }> };

export type MeHasPermissionQueryVariables = Exact<{
  key: Scalars['String']['input'];
}>;


export type MeHasPermissionQuery = { __typename?: 'Query', meHasPermission: boolean };

export const BlueprintFragementFragmentDoc = gql`
    fragment BlueprintFragement on Blueprint {
  id
  createdAt
  updatedAt
  name
  description
  blueprintTemplate
  parentGroup {
    id
    name
  }
  provider {
    id
    displayName
    isLoaded
  }
  deployments {
    id
  }
}
    `;
export const UserFragmentFragmentDoc = gql`
    fragment UserFragment on User {
  id
  createdAt
  updatedAt
  username
  email
  firstName
  lastName
}
    `;
export const DeploymentFragmentFragmentDoc = gql`
    fragment DeploymentFragment on Deployment {
  id
  createdAt
  updatedAt
  name
  templateVars
  deploymentVars
  deploymentState
  blueprint {
    ...BlueprintFragement
  }
  requester {
    ...UserFragment
  }
}
    ${BlueprintFragementFragmentDoc}
${UserFragmentFragmentDoc}`;
export const GroupFragmentFragmentDoc = gql`
    fragment GroupFragment on Group {
  id
  createdAt
  updatedAt
  name
}
    `;
export const ProviderFragmentFragmentDoc = gql`
    fragment ProviderFragment on Provider {
  id
  createdAt
  updatedAt
  displayName
  providerGitUrl
  providerVersion
  isLoaded
}
    `;
export const BlueprintsDocument = gql`
    query Blueprints {
  blueprints {
    ...BlueprintFragement
  }
}
    ${BlueprintFragementFragmentDoc}`;

/**
 * __useBlueprintsQuery__
 *
 * To run a query within a React component, call `useBlueprintsQuery` and pass it any options that fit your needs.
 * When your component renders, `useBlueprintsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useBlueprintsQuery({
 *   variables: {
 *   },
 * });
 */
export function useBlueprintsQuery(baseOptions?: Apollo.QueryHookOptions<BlueprintsQuery, BlueprintsQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<BlueprintsQuery, BlueprintsQueryVariables>(BlueprintsDocument, options);
      }
export function useBlueprintsLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<BlueprintsQuery, BlueprintsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<BlueprintsQuery, BlueprintsQueryVariables>(BlueprintsDocument, options);
        }
export function useBlueprintsSuspenseQuery(baseOptions?: Apollo.SuspenseQueryHookOptions<BlueprintsQuery, BlueprintsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useSuspenseQuery<BlueprintsQuery, BlueprintsQueryVariables>(BlueprintsDocument, options);
        }
export type BlueprintsQueryHookResult = ReturnType<typeof useBlueprintsQuery>;
export type BlueprintsLazyQueryHookResult = ReturnType<typeof useBlueprintsLazyQuery>;
export type BlueprintsSuspenseQueryHookResult = ReturnType<typeof useBlueprintsSuspenseQuery>;
export type BlueprintsQueryResult = Apollo.QueryResult<BlueprintsQuery, BlueprintsQueryVariables>;
export const GetBlueprintDocument = gql`
    query GetBlueprint($id: ID!) {
  blueprint(id: $id) {
    ...BlueprintFragement
  }
}
    ${BlueprintFragementFragmentDoc}`;

/**
 * __useGetBlueprintQuery__
 *
 * To run a query within a React component, call `useGetBlueprintQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetBlueprintQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetBlueprintQuery({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useGetBlueprintQuery(baseOptions: Apollo.QueryHookOptions<GetBlueprintQuery, GetBlueprintQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetBlueprintQuery, GetBlueprintQueryVariables>(GetBlueprintDocument, options);
      }
export function useGetBlueprintLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetBlueprintQuery, GetBlueprintQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetBlueprintQuery, GetBlueprintQueryVariables>(GetBlueprintDocument, options);
        }
export function useGetBlueprintSuspenseQuery(baseOptions?: Apollo.SuspenseQueryHookOptions<GetBlueprintQuery, GetBlueprintQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useSuspenseQuery<GetBlueprintQuery, GetBlueprintQueryVariables>(GetBlueprintDocument, options);
        }
export type GetBlueprintQueryHookResult = ReturnType<typeof useGetBlueprintQuery>;
export type GetBlueprintLazyQueryHookResult = ReturnType<typeof useGetBlueprintLazyQuery>;
export type GetBlueprintSuspenseQueryHookResult = ReturnType<typeof useGetBlueprintSuspenseQuery>;
export type GetBlueprintQueryResult = Apollo.QueryResult<GetBlueprintQuery, GetBlueprintQueryVariables>;
export const CreateBlueprintDocument = gql`
    mutation CreateBlueprint($input: BlueprintInput!) {
  createBlueprint(input: $input) {
    ...BlueprintFragement
  }
}
    ${BlueprintFragementFragmentDoc}`;
export type CreateBlueprintMutationFn = Apollo.MutationFunction<CreateBlueprintMutation, CreateBlueprintMutationVariables>;

/**
 * __useCreateBlueprintMutation__
 *
 * To run a mutation, you first call `useCreateBlueprintMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateBlueprintMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createBlueprintMutation, { data, loading, error }] = useCreateBlueprintMutation({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useCreateBlueprintMutation(baseOptions?: Apollo.MutationHookOptions<CreateBlueprintMutation, CreateBlueprintMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<CreateBlueprintMutation, CreateBlueprintMutationVariables>(CreateBlueprintDocument, options);
      }
export type CreateBlueprintMutationHookResult = ReturnType<typeof useCreateBlueprintMutation>;
export type CreateBlueprintMutationResult = Apollo.MutationResult<CreateBlueprintMutation>;
export type CreateBlueprintMutationOptions = Apollo.BaseMutationOptions<CreateBlueprintMutation, CreateBlueprintMutationVariables>;
export const UpdateBlueprintDocument = gql`
    mutation UpdateBlueprint($id: ID!, $input: BlueprintInput!) {
  updateBlueprint(id: $id, input: $input) {
    ...BlueprintFragement
  }
}
    ${BlueprintFragementFragmentDoc}`;
export type UpdateBlueprintMutationFn = Apollo.MutationFunction<UpdateBlueprintMutation, UpdateBlueprintMutationVariables>;

/**
 * __useUpdateBlueprintMutation__
 *
 * To run a mutation, you first call `useUpdateBlueprintMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateBlueprintMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateBlueprintMutation, { data, loading, error }] = useUpdateBlueprintMutation({
 *   variables: {
 *      id: // value for 'id'
 *      input: // value for 'input'
 *   },
 * });
 */
export function useUpdateBlueprintMutation(baseOptions?: Apollo.MutationHookOptions<UpdateBlueprintMutation, UpdateBlueprintMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<UpdateBlueprintMutation, UpdateBlueprintMutationVariables>(UpdateBlueprintDocument, options);
      }
export type UpdateBlueprintMutationHookResult = ReturnType<typeof useUpdateBlueprintMutation>;
export type UpdateBlueprintMutationResult = Apollo.MutationResult<UpdateBlueprintMutation>;
export type UpdateBlueprintMutationOptions = Apollo.BaseMutationOptions<UpdateBlueprintMutation, UpdateBlueprintMutationVariables>;
export const DeployBlueprintDocument = gql`
    mutation DeployBlueprint($id: ID!) {
  deployBlueprint(id: $id) {
    ...DeploymentFragment
  }
}
    ${DeploymentFragmentFragmentDoc}`;
export type DeployBlueprintMutationFn = Apollo.MutationFunction<DeployBlueprintMutation, DeployBlueprintMutationVariables>;

/**
 * __useDeployBlueprintMutation__
 *
 * To run a mutation, you first call `useDeployBlueprintMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDeployBlueprintMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [deployBlueprintMutation, { data, loading, error }] = useDeployBlueprintMutation({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useDeployBlueprintMutation(baseOptions?: Apollo.MutationHookOptions<DeployBlueprintMutation, DeployBlueprintMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<DeployBlueprintMutation, DeployBlueprintMutationVariables>(DeployBlueprintDocument, options);
      }
export type DeployBlueprintMutationHookResult = ReturnType<typeof useDeployBlueprintMutation>;
export type DeployBlueprintMutationResult = Apollo.MutationResult<DeployBlueprintMutation>;
export type DeployBlueprintMutationOptions = Apollo.BaseMutationOptions<DeployBlueprintMutation, DeployBlueprintMutationVariables>;
export const ListDeploymentsDocument = gql`
    query ListDeployments {
  deployments {
    ...DeploymentFragment
  }
}
    ${DeploymentFragmentFragmentDoc}`;

/**
 * __useListDeploymentsQuery__
 *
 * To run a query within a React component, call `useListDeploymentsQuery` and pass it any options that fit your needs.
 * When your component renders, `useListDeploymentsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useListDeploymentsQuery({
 *   variables: {
 *   },
 * });
 */
export function useListDeploymentsQuery(baseOptions?: Apollo.QueryHookOptions<ListDeploymentsQuery, ListDeploymentsQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<ListDeploymentsQuery, ListDeploymentsQueryVariables>(ListDeploymentsDocument, options);
      }
export function useListDeploymentsLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<ListDeploymentsQuery, ListDeploymentsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<ListDeploymentsQuery, ListDeploymentsQueryVariables>(ListDeploymentsDocument, options);
        }
export function useListDeploymentsSuspenseQuery(baseOptions?: Apollo.SuspenseQueryHookOptions<ListDeploymentsQuery, ListDeploymentsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useSuspenseQuery<ListDeploymentsQuery, ListDeploymentsQueryVariables>(ListDeploymentsDocument, options);
        }
export type ListDeploymentsQueryHookResult = ReturnType<typeof useListDeploymentsQuery>;
export type ListDeploymentsLazyQueryHookResult = ReturnType<typeof useListDeploymentsLazyQuery>;
export type ListDeploymentsSuspenseQueryHookResult = ReturnType<typeof useListDeploymentsSuspenseQuery>;
export type ListDeploymentsQueryResult = Apollo.QueryResult<ListDeploymentsQuery, ListDeploymentsQueryVariables>;
export const GetDeploymentDocument = gql`
    query GetDeployment($id: ID!) {
  deployment(id: $id) {
    ...DeploymentFragment
  }
}
    ${DeploymentFragmentFragmentDoc}`;

/**
 * __useGetDeploymentQuery__
 *
 * To run a query within a React component, call `useGetDeploymentQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetDeploymentQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetDeploymentQuery({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useGetDeploymentQuery(baseOptions: Apollo.QueryHookOptions<GetDeploymentQuery, GetDeploymentQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetDeploymentQuery, GetDeploymentQueryVariables>(GetDeploymentDocument, options);
      }
export function useGetDeploymentLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetDeploymentQuery, GetDeploymentQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetDeploymentQuery, GetDeploymentQueryVariables>(GetDeploymentDocument, options);
        }
export function useGetDeploymentSuspenseQuery(baseOptions?: Apollo.SuspenseQueryHookOptions<GetDeploymentQuery, GetDeploymentQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useSuspenseQuery<GetDeploymentQuery, GetDeploymentQueryVariables>(GetDeploymentDocument, options);
        }
export type GetDeploymentQueryHookResult = ReturnType<typeof useGetDeploymentQuery>;
export type GetDeploymentLazyQueryHookResult = ReturnType<typeof useGetDeploymentLazyQuery>;
export type GetDeploymentSuspenseQueryHookResult = ReturnType<typeof useGetDeploymentSuspenseQuery>;
export type GetDeploymentQueryResult = Apollo.QueryResult<GetDeploymentQuery, GetDeploymentQueryVariables>;
export const UpdateDeploymentDocument = gql`
    mutation UpdateDeployment($id: ID!, $input: DeploymentInput!) {
  updateDeployment(id: $id, input: $input) {
    ...DeploymentFragment
  }
}
    ${DeploymentFragmentFragmentDoc}`;
export type UpdateDeploymentMutationFn = Apollo.MutationFunction<UpdateDeploymentMutation, UpdateDeploymentMutationVariables>;

/**
 * __useUpdateDeploymentMutation__
 *
 * To run a mutation, you first call `useUpdateDeploymentMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateDeploymentMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateDeploymentMutation, { data, loading, error }] = useUpdateDeploymentMutation({
 *   variables: {
 *      id: // value for 'id'
 *      input: // value for 'input'
 *   },
 * });
 */
export function useUpdateDeploymentMutation(baseOptions?: Apollo.MutationHookOptions<UpdateDeploymentMutation, UpdateDeploymentMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<UpdateDeploymentMutation, UpdateDeploymentMutationVariables>(UpdateDeploymentDocument, options);
      }
export type UpdateDeploymentMutationHookResult = ReturnType<typeof useUpdateDeploymentMutation>;
export type UpdateDeploymentMutationResult = Apollo.MutationResult<UpdateDeploymentMutation>;
export type UpdateDeploymentMutationOptions = Apollo.BaseMutationOptions<UpdateDeploymentMutation, UpdateDeploymentMutationVariables>;
export const DestroyDeploymentDocument = gql`
    mutation DestroyDeployment($id: ID!) {
  destroyDeployment(id: $id) {
    ...DeploymentFragment
  }
}
    ${DeploymentFragmentFragmentDoc}`;
export type DestroyDeploymentMutationFn = Apollo.MutationFunction<DestroyDeploymentMutation, DestroyDeploymentMutationVariables>;

/**
 * __useDestroyDeploymentMutation__
 *
 * To run a mutation, you first call `useDestroyDeploymentMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDestroyDeploymentMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [destroyDeploymentMutation, { data, loading, error }] = useDestroyDeploymentMutation({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useDestroyDeploymentMutation(baseOptions?: Apollo.MutationHookOptions<DestroyDeploymentMutation, DestroyDeploymentMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<DestroyDeploymentMutation, DestroyDeploymentMutationVariables>(DestroyDeploymentDocument, options);
      }
export type DestroyDeploymentMutationHookResult = ReturnType<typeof useDestroyDeploymentMutation>;
export type DestroyDeploymentMutationResult = Apollo.MutationResult<DestroyDeploymentMutation>;
export type DestroyDeploymentMutationOptions = Apollo.BaseMutationOptions<DestroyDeploymentMutation, DestroyDeploymentMutationVariables>;
export const ListGroupsDocument = gql`
    query ListGroups {
  groups {
    ...GroupFragment
  }
}
    ${GroupFragmentFragmentDoc}`;

/**
 * __useListGroupsQuery__
 *
 * To run a query within a React component, call `useListGroupsQuery` and pass it any options that fit your needs.
 * When your component renders, `useListGroupsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useListGroupsQuery({
 *   variables: {
 *   },
 * });
 */
export function useListGroupsQuery(baseOptions?: Apollo.QueryHookOptions<ListGroupsQuery, ListGroupsQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<ListGroupsQuery, ListGroupsQueryVariables>(ListGroupsDocument, options);
      }
export function useListGroupsLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<ListGroupsQuery, ListGroupsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<ListGroupsQuery, ListGroupsQueryVariables>(ListGroupsDocument, options);
        }
export function useListGroupsSuspenseQuery(baseOptions?: Apollo.SuspenseQueryHookOptions<ListGroupsQuery, ListGroupsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useSuspenseQuery<ListGroupsQuery, ListGroupsQueryVariables>(ListGroupsDocument, options);
        }
export type ListGroupsQueryHookResult = ReturnType<typeof useListGroupsQuery>;
export type ListGroupsLazyQueryHookResult = ReturnType<typeof useListGroupsLazyQuery>;
export type ListGroupsSuspenseQueryHookResult = ReturnType<typeof useListGroupsSuspenseQuery>;
export type ListGroupsQueryResult = Apollo.QueryResult<ListGroupsQuery, ListGroupsQueryVariables>;
export const ProvidersDocument = gql`
    query Providers {
  providers {
    ...ProviderFragment
  }
}
    ${ProviderFragmentFragmentDoc}`;

/**
 * __useProvidersQuery__
 *
 * To run a query within a React component, call `useProvidersQuery` and pass it any options that fit your needs.
 * When your component renders, `useProvidersQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useProvidersQuery({
 *   variables: {
 *   },
 * });
 */
export function useProvidersQuery(baseOptions?: Apollo.QueryHookOptions<ProvidersQuery, ProvidersQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<ProvidersQuery, ProvidersQueryVariables>(ProvidersDocument, options);
      }
export function useProvidersLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<ProvidersQuery, ProvidersQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<ProvidersQuery, ProvidersQueryVariables>(ProvidersDocument, options);
        }
export function useProvidersSuspenseQuery(baseOptions?: Apollo.SuspenseQueryHookOptions<ProvidersQuery, ProvidersQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useSuspenseQuery<ProvidersQuery, ProvidersQueryVariables>(ProvidersDocument, options);
        }
export type ProvidersQueryHookResult = ReturnType<typeof useProvidersQuery>;
export type ProvidersLazyQueryHookResult = ReturnType<typeof useProvidersLazyQuery>;
export type ProvidersSuspenseQueryHookResult = ReturnType<typeof useProvidersSuspenseQuery>;
export type ProvidersQueryResult = Apollo.QueryResult<ProvidersQuery, ProvidersQueryVariables>;
export const MeDocument = gql`
    query Me {
  me {
    ...UserFragment
  }
}
    ${UserFragmentFragmentDoc}`;

/**
 * __useMeQuery__
 *
 * To run a query within a React component, call `useMeQuery` and pass it any options that fit your needs.
 * When your component renders, `useMeQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useMeQuery({
 *   variables: {
 *   },
 * });
 */
export function useMeQuery(baseOptions?: Apollo.QueryHookOptions<MeQuery, MeQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<MeQuery, MeQueryVariables>(MeDocument, options);
      }
export function useMeLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<MeQuery, MeQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<MeQuery, MeQueryVariables>(MeDocument, options);
        }
export function useMeSuspenseQuery(baseOptions?: Apollo.SuspenseQueryHookOptions<MeQuery, MeQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useSuspenseQuery<MeQuery, MeQueryVariables>(MeDocument, options);
        }
export type MeQueryHookResult = ReturnType<typeof useMeQuery>;
export type MeLazyQueryHookResult = ReturnType<typeof useMeLazyQuery>;
export type MeSuspenseQueryHookResult = ReturnType<typeof useMeSuspenseQuery>;
export type MeQueryResult = Apollo.QueryResult<MeQuery, MeQueryVariables>;
export const ListUsersDocument = gql`
    query ListUsers {
  users {
    ...UserFragment
  }
}
    ${UserFragmentFragmentDoc}`;

/**
 * __useListUsersQuery__
 *
 * To run a query within a React component, call `useListUsersQuery` and pass it any options that fit your needs.
 * When your component renders, `useListUsersQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useListUsersQuery({
 *   variables: {
 *   },
 * });
 */
export function useListUsersQuery(baseOptions?: Apollo.QueryHookOptions<ListUsersQuery, ListUsersQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<ListUsersQuery, ListUsersQueryVariables>(ListUsersDocument, options);
      }
export function useListUsersLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<ListUsersQuery, ListUsersQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<ListUsersQuery, ListUsersQueryVariables>(ListUsersDocument, options);
        }
export function useListUsersSuspenseQuery(baseOptions?: Apollo.SuspenseQueryHookOptions<ListUsersQuery, ListUsersQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useSuspenseQuery<ListUsersQuery, ListUsersQueryVariables>(ListUsersDocument, options);
        }
export type ListUsersQueryHookResult = ReturnType<typeof useListUsersQuery>;
export type ListUsersLazyQueryHookResult = ReturnType<typeof useListUsersLazyQuery>;
export type ListUsersSuspenseQueryHookResult = ReturnType<typeof useListUsersSuspenseQuery>;
export type ListUsersQueryResult = Apollo.QueryResult<ListUsersQuery, ListUsersQueryVariables>;
export const MeHasPermissionDocument = gql`
    query MeHasPermission($key: String!) {
  meHasPermission(key: $key)
}
    `;

/**
 * __useMeHasPermissionQuery__
 *
 * To run a query within a React component, call `useMeHasPermissionQuery` and pass it any options that fit your needs.
 * When your component renders, `useMeHasPermissionQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useMeHasPermissionQuery({
 *   variables: {
 *      key: // value for 'key'
 *   },
 * });
 */
export function useMeHasPermissionQuery(baseOptions: Apollo.QueryHookOptions<MeHasPermissionQuery, MeHasPermissionQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<MeHasPermissionQuery, MeHasPermissionQueryVariables>(MeHasPermissionDocument, options);
      }
export function useMeHasPermissionLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<MeHasPermissionQuery, MeHasPermissionQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<MeHasPermissionQuery, MeHasPermissionQueryVariables>(MeHasPermissionDocument, options);
        }
export function useMeHasPermissionSuspenseQuery(baseOptions?: Apollo.SuspenseQueryHookOptions<MeHasPermissionQuery, MeHasPermissionQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useSuspenseQuery<MeHasPermissionQuery, MeHasPermissionQueryVariables>(MeHasPermissionDocument, options);
        }
export type MeHasPermissionQueryHookResult = ReturnType<typeof useMeHasPermissionQuery>;
export type MeHasPermissionLazyQueryHookResult = ReturnType<typeof useMeHasPermissionLazyQuery>;
export type MeHasPermissionSuspenseQueryHookResult = ReturnType<typeof useMeHasPermissionSuspenseQuery>;
export type MeHasPermissionQueryResult = Apollo.QueryResult<MeHasPermissionQuery, MeHasPermissionQueryVariables>;