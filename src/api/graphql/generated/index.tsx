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
  Time: { input: any; output: any; }
};

export type Blueprint = {
  __typename?: 'Blueprint';
  blueprintTemplate: Scalars['String']['output'];
  deployments: Array<Maybe<Deployment>>;
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  parentGroup: Group;
  provider: Provider;
};

export type BlueprintInput = {
  blueprintTemplate: Scalars['String']['input'];
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
  id: Scalars['ID']['output'];
  requester: User;
};

export type Group = {
  __typename?: 'Group';
  blueprints?: Maybe<Array<Maybe<Blueprint>>>;
  children?: Maybe<Array<Maybe<Group>>>;
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
  parent?: Maybe<Group>;
  permissionPolicies?: Maybe<Array<Maybe<PermissionPolicy>>>;
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
  id: Scalars['ID']['output'];
  key?: Maybe<Scalars['String']['output']>;
  permissionPolicies?: Maybe<Array<Maybe<PermissionPolicy>>>;
};

export type PermissionPolicy = {
  __typename?: 'PermissionPolicy';
  group: Group;
  id: Scalars['ID']['output'];
  permission: Permission;
  type: PermissionPolicyType;
};

export enum PermissionPolicyType {
  Allow = 'ALLOW',
  Deny = 'DENY'
}

export type Provider = {
  __typename?: 'Provider';
  blueprints?: Maybe<Array<Maybe<Blueprint>>>;
  configBytes: Scalars['String']['output'];
  displayName: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  isLoaded: Scalars['Boolean']['output'];
  providerGitUrl: Scalars['String']['output'];
  providerVersion: Scalars['String']['output'];
};

export type ProviderCommand = {
  __typename?: 'ProviderCommand';
  commandType: CommandType;
  endTime?: Maybe<Scalars['Time']['output']>;
  error: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  output: Scalars['String']['output'];
  startTime?: Maybe<Scalars['Time']['output']>;
  status: CommandStatus;
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
  deployments: Array<Maybe<Deployment>>;
  email: Scalars['String']['output'];
  firstName: Scalars['String']['output'];
  groups: Array<Maybe<Group>>;
  id: Scalars['ID']['output'];
  lastName: Scalars['String']['output'];
  username: Scalars['String']['output'];
};

export type UserInput = {
  email: Scalars['String']['input'];
  firstName: Scalars['String']['input'];
  lastName: Scalars['String']['input'];
  username: Scalars['String']['input'];
};

export type UserFragmentFragment = { __typename?: 'User', id: string, username: string, email: string, firstName: string, lastName: string };

export type MeQueryVariables = Exact<{ [key: string]: never; }>;


export type MeQuery = { __typename?: 'Query', me: { __typename?: 'User', id: string, username: string, email: string, firstName: string, lastName: string } };

export type ListUsersQueryVariables = Exact<{ [key: string]: never; }>;


export type ListUsersQuery = { __typename?: 'Query', users: Array<{ __typename?: 'User', id: string, username: string, email: string, firstName: string, lastName: string }> };

export const UserFragmentFragmentDoc = gql`
    fragment UserFragment on User {
  id
  username
  email
  firstName
  lastName
}
    `;
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