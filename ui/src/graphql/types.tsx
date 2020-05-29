import gql from "graphql-tag";
import * as ApolloReactCommon from "@apollo/client";
import * as ApolloReactHooks from "@apollo/client";
export type Maybe<T> = T | null;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  /** The `Upload` scalar type represents a file upload. */
  Upload: any;
};

export type Plant = {
  __typename?: "Plant";
  _id: Scalars["ID"];
  name: Scalars["String"];
  otherNames?: Maybe<Scalars["String"]>;
  description?: Maybe<Scalars["String"]>;
  plantSeason?: Maybe<Scalars["String"]>;
  harvestSeason?: Maybe<Scalars["String"]>;
  pruneSeason?: Maybe<Scalars["String"]>;
  tips?: Maybe<Scalars["String"]>;
};

export type AddPlantResponse = {
  __typename?: "AddPlantResponse";
  insertedId: Scalars["ID"];
};

export type DeletePlantResponse = {
  __typename?: "DeletePlantResponse";
  deletedCount?: Maybe<Scalars["Int"]>;
};

export type Query = {
  __typename?: "Query";
  getPlants?: Maybe<Array<Plant>>;
  getPlant?: Maybe<Plant>;
};

export type QueryGetPlantArgs = {
  name: Scalars["String"];
};

export type Mutation = {
  __typename?: "Mutation";
  addPlant: AddPlantResponse;
  deletePlant?: Maybe<DeletePlantResponse>;
};

export type MutationAddPlantArgs = {
  name: Scalars["String"];
  otherNames?: Maybe<Scalars["String"]>;
  description?: Maybe<Scalars["String"]>;
  plantSeason?: Maybe<Scalars["String"]>;
  harvestSeason?: Maybe<Scalars["String"]>;
  pruneSeason?: Maybe<Scalars["String"]>;
  tips?: Maybe<Scalars["String"]>;
};

export type MutationDeletePlantArgs = {
  _id: Scalars["ID"];
};

export enum CacheControlScope {
  Public = "PUBLIC",
  Private = "PRIVATE",
}

export type AddPlantMutationVariables = {
  name: Scalars["String"];
  otherNames?: Maybe<Scalars["String"]>;
  description?: Maybe<Scalars["String"]>;
  plantSeason?: Maybe<Scalars["String"]>;
  harvestSeason?: Maybe<Scalars["String"]>;
  pruneSeason?: Maybe<Scalars["String"]>;
  tips?: Maybe<Scalars["String"]>;
};

export type AddPlantMutation = {__typename?: "Mutation"} & {
  addPlant: {__typename?: "AddPlantResponse"} & Pick<
    AddPlantResponse,
    "insertedId"
  >;
};

export type DeletePlantMutationVariables = {
  _id: Scalars["ID"];
};

export type DeletePlantMutation = {__typename?: "Mutation"} & {
  deletePlant?: Maybe<
    {__typename?: "DeletePlantResponse"} & Pick<
      DeletePlantResponse,
      "deletedCount"
    >
  >;
};

export type GetPlantQueryVariables = {
  name: Scalars["String"];
};

export type GetPlantQuery = {__typename?: "Query"} & {
  getPlant?: Maybe<
    {__typename?: "Plant"} & Pick<
      Plant,
      | "_id"
      | "name"
      | "otherNames"
      | "description"
      | "plantSeason"
      | "harvestSeason"
      | "pruneSeason"
      | "tips"
    >
  >;
};

export type GetPlantsQueryVariables = {};

export type GetPlantsQuery = {__typename?: "Query"} & {
  getPlants?: Maybe<
    Array<{__typename?: "Plant"} & Pick<Plant, "_id" | "name">>
  >;
};

export const AddPlantDocument = gql`
  mutation AddPlant(
    $name: String!
    $otherNames: String
    $description: String
    $plantSeason: String
    $harvestSeason: String
    $pruneSeason: String
    $tips: String
  ) {
    addPlant(
      name: $name
      otherNames: $otherNames
      description: $description
      plantSeason: $plantSeason
      harvestSeason: $harvestSeason
      pruneSeason: $pruneSeason
      tips: $tips
    ) {
      insertedId
    }
  }
`;
export type AddPlantMutationFn = ApolloReactCommon.MutationFunction<
  AddPlantMutation,
  AddPlantMutationVariables
>;

/**
 * __useAddPlantMutation__
 *
 * To run a mutation, you first call `useAddPlantMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useAddPlantMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [addPlantMutation, { data, loading, error }] = useAddPlantMutation({
 *   variables: {
 *      name: // value for 'name'
 *      otherNames: // value for 'otherNames'
 *      description: // value for 'description'
 *      plantSeason: // value for 'plantSeason'
 *      harvestSeason: // value for 'harvestSeason'
 *      pruneSeason: // value for 'pruneSeason'
 *      tips: // value for 'tips'
 *   },
 * });
 */
export function useAddPlantMutation(
  baseOptions?: ApolloReactHooks.MutationHookOptions<
    AddPlantMutation,
    AddPlantMutationVariables
  >,
) {
  return ApolloReactHooks.useMutation<
    AddPlantMutation,
    AddPlantMutationVariables
  >(AddPlantDocument, baseOptions);
}
export type AddPlantMutationHookResult = ReturnType<typeof useAddPlantMutation>;
export type AddPlantMutationResult = ApolloReactCommon.MutationResult<
  AddPlantMutation
>;
export type AddPlantMutationOptions = ApolloReactCommon.BaseMutationOptions<
  AddPlantMutation,
  AddPlantMutationVariables
>;
export const DeletePlantDocument = gql`
  mutation DeletePlant($_id: ID!) {
    deletePlant(_id: $_id) {
      deletedCount
    }
  }
`;
export type DeletePlantMutationFn = ApolloReactCommon.MutationFunction<
  DeletePlantMutation,
  DeletePlantMutationVariables
>;

/**
 * __useDeletePlantMutation__
 *
 * To run a mutation, you first call `useDeletePlantMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDeletePlantMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [deletePlantMutation, { data, loading, error }] = useDeletePlantMutation({
 *   variables: {
 *      _id: // value for '_id'
 *   },
 * });
 */
export function useDeletePlantMutation(
  baseOptions?: ApolloReactHooks.MutationHookOptions<
    DeletePlantMutation,
    DeletePlantMutationVariables
  >,
) {
  return ApolloReactHooks.useMutation<
    DeletePlantMutation,
    DeletePlantMutationVariables
  >(DeletePlantDocument, baseOptions);
}
export type DeletePlantMutationHookResult = ReturnType<
  typeof useDeletePlantMutation
>;
export type DeletePlantMutationResult = ApolloReactCommon.MutationResult<
  DeletePlantMutation
>;
export type DeletePlantMutationOptions = ApolloReactCommon.BaseMutationOptions<
  DeletePlantMutation,
  DeletePlantMutationVariables
>;
export const GetPlantDocument = gql`
  query GetPlant($name: String!) {
    getPlant(name: $name) {
      _id
      name
      otherNames
      description
      plantSeason
      harvestSeason
      pruneSeason
      tips
    }
  }
`;

/**
 * __useGetPlantQuery__
 *
 * To run a query within a React component, call `useGetPlantQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetPlantQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetPlantQuery({
 *   variables: {
 *      name: // value for 'name'
 *   },
 * });
 */
export function useGetPlantQuery(
  baseOptions?: ApolloReactHooks.QueryHookOptions<
    GetPlantQuery,
    GetPlantQueryVariables
  >,
) {
  return ApolloReactHooks.useQuery<GetPlantQuery, GetPlantQueryVariables>(
    GetPlantDocument,
    baseOptions,
  );
}
export function useGetPlantLazyQuery(
  baseOptions?: ApolloReactHooks.LazyQueryHookOptions<
    GetPlantQuery,
    GetPlantQueryVariables
  >,
) {
  return ApolloReactHooks.useLazyQuery<GetPlantQuery, GetPlantQueryVariables>(
    GetPlantDocument,
    baseOptions,
  );
}
export type GetPlantQueryHookResult = ReturnType<typeof useGetPlantQuery>;
export type GetPlantLazyQueryHookResult = ReturnType<
  typeof useGetPlantLazyQuery
>;
export type GetPlantQueryResult = ApolloReactCommon.QueryResult<
  GetPlantQuery,
  GetPlantQueryVariables
>;
export const GetPlantsDocument = gql`
  query GetPlants {
    getPlants {
      _id
      name
    }
  }
`;

/**
 * __useGetPlantsQuery__
 *
 * To run a query within a React component, call `useGetPlantsQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetPlantsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetPlantsQuery({
 *   variables: {
 *   },
 * });
 */
export function useGetPlantsQuery(
  baseOptions?: ApolloReactHooks.QueryHookOptions<
    GetPlantsQuery,
    GetPlantsQueryVariables
  >,
) {
  return ApolloReactHooks.useQuery<GetPlantsQuery, GetPlantsQueryVariables>(
    GetPlantsDocument,
    baseOptions,
  );
}
export function useGetPlantsLazyQuery(
  baseOptions?: ApolloReactHooks.LazyQueryHookOptions<
    GetPlantsQuery,
    GetPlantsQueryVariables
  >,
) {
  return ApolloReactHooks.useLazyQuery<GetPlantsQuery, GetPlantsQueryVariables>(
    GetPlantsDocument,
    baseOptions,
  );
}
export type GetPlantsQueryHookResult = ReturnType<typeof useGetPlantsQuery>;
export type GetPlantsLazyQueryHookResult = ReturnType<
  typeof useGetPlantsLazyQuery
>;
export type GetPlantsQueryResult = ApolloReactCommon.QueryResult<
  GetPlantsQuery,
  GetPlantsQueryVariables
>;
