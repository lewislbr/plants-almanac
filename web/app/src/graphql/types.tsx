import gql from "graphql-tag"
import * as ApolloReactCommon from "@apollo/client"
import * as ApolloReactHooks from "@apollo/client"
export type Maybe<T> = T | null
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string
  String: string
  Boolean: boolean
  Int: number
  Float: number
}

export type Mutation = {
  __typename?: "Mutation"
  /** Adds a plant */
  add?: Maybe<Scalars["ID"]>
  /** Deletes a plant */
  delete?: Maybe<Scalars["Int"]>
  /** Edits a plant */
  edit?: Maybe<Scalars["Int"]>
}

export type MutationAddArgs = {
  prune_season?: Maybe<Scalars["String"]>
  tips?: Maybe<Scalars["String"]>
  name: Scalars["String"]
  other_names?: Maybe<Scalars["String"]>
  description?: Maybe<Scalars["String"]>
  plant_season?: Maybe<Scalars["String"]>
  harvest_season?: Maybe<Scalars["String"]>
}

export type MutationDeleteArgs = {
  id: Scalars["ID"]
}

export type MutationEditArgs = {
  id: Scalars["ID"]
  name?: Maybe<Scalars["String"]>
  other_names?: Maybe<Scalars["String"]>
  description?: Maybe<Scalars["String"]>
  plant_season?: Maybe<Scalars["String"]>
  harvest_season?: Maybe<Scalars["String"]>
  prune_season?: Maybe<Scalars["String"]>
  tips?: Maybe<Scalars["String"]>
}

export type Plant = {
  __typename?: "Plant"
  description?: Maybe<Scalars["String"]>
  harvest_season?: Maybe<Scalars["String"]>
  id?: Maybe<Scalars["ID"]>
  name?: Maybe<Scalars["String"]>
  other_names?: Maybe<Scalars["String"]>
  plant_season?: Maybe<Scalars["String"]>
  prune_season?: Maybe<Scalars["String"]>
  tips?: Maybe<Scalars["String"]>
}

export type Query = {
  __typename?: "Query"
  /** Returns a plant */
  plant?: Maybe<Plant>
  /** Returns all plants */
  plants?: Maybe<Array<Maybe<Plant>>>
}

export type QueryPlantArgs = {
  id: Scalars["ID"]
}

export type AddMutationVariables = {
  name: Scalars["String"]
  other_names?: Maybe<Scalars["String"]>
  description?: Maybe<Scalars["String"]>
  plant_season?: Maybe<Scalars["String"]>
  harvest_season?: Maybe<Scalars["String"]>
  prune_season?: Maybe<Scalars["String"]>
  tips?: Maybe<Scalars["String"]>
}

export type AddMutation = {__typename?: "Mutation"} & Pick<Mutation, "add">

export type DeleteMutationVariables = {
  id: Scalars["ID"]
}

export type DeleteMutation = {__typename?: "Mutation"} & Pick<
  Mutation,
  "delete"
>

export type PlantQueryVariables = {
  id: Scalars["ID"]
}

export type PlantQuery = {__typename?: "Query"} & {
  plant?: Maybe<
    {__typename?: "Plant"} & Pick<
      Plant,
      | "id"
      | "name"
      | "other_names"
      | "description"
      | "plant_season"
      | "harvest_season"
      | "prune_season"
      | "tips"
    >
  >
}

export type PlantsQueryVariables = {}

export type PlantsQuery = {__typename?: "Query"} & {
  plants?: Maybe<
    Array<Maybe<{__typename?: "Plant"} & Pick<Plant, "id" | "name">>>
  >
}

export const AddDocument = gql`
  mutation Add(
    $name: String!
    $other_names: String
    $description: String
    $plant_season: String
    $harvest_season: String
    $prune_season: String
    $tips: String
  ) {
    add(
      name: $name
      other_names: $other_names
      description: $description
      plant_season: $plant_season
      harvest_season: $harvest_season
      prune_season: $prune_season
      tips: $tips
    )
  }
`
export type AddMutationFn = ApolloReactCommon.MutationFunction<
  AddMutation,
  AddMutationVariables
>

/**
 * __useAddMutation__
 *
 * To run a mutation, you first call `useAddMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useAddMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [addMutation, { data, loading, error }] = useAddMutation({
 *   variables: {
 *      name: // value for 'name'
 *      other_names: // value for 'other_names'
 *      description: // value for 'description'
 *      plant_season: // value for 'plant_season'
 *      harvest_season: // value for 'harvest_season'
 *      prune_season: // value for 'prune_season'
 *      tips: // value for 'tips'
 *   },
 * });
 */
export function useAddMutation(
  baseOptions?: ApolloReactHooks.MutationHookOptions<
    AddMutation,
    AddMutationVariables
  >,
) {
  return ApolloReactHooks.useMutation<AddMutation, AddMutationVariables>(
    AddDocument,
    baseOptions,
  )
}
export type AddMutationHookResult = ReturnType<typeof useAddMutation>
export type AddMutationResult = ApolloReactCommon.MutationResult<AddMutation>
export type AddMutationOptions = ApolloReactCommon.BaseMutationOptions<
  AddMutation,
  AddMutationVariables
>
export const DeleteDocument = gql`
  mutation Delete($id: ID!) {
    delete(id: $id)
  }
`
export type DeleteMutationFn = ApolloReactCommon.MutationFunction<
  DeleteMutation,
  DeleteMutationVariables
>

/**
 * __useDeleteMutation__
 *
 * To run a mutation, you first call `useDeleteMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDeleteMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [deleteMutation, { data, loading, error }] = useDeleteMutation({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useDeleteMutation(
  baseOptions?: ApolloReactHooks.MutationHookOptions<
    DeleteMutation,
    DeleteMutationVariables
  >,
) {
  return ApolloReactHooks.useMutation<DeleteMutation, DeleteMutationVariables>(
    DeleteDocument,
    baseOptions,
  )
}
export type DeleteMutationHookResult = ReturnType<typeof useDeleteMutation>
export type DeleteMutationResult = ApolloReactCommon.MutationResult<
  DeleteMutation
>
export type DeleteMutationOptions = ApolloReactCommon.BaseMutationOptions<
  DeleteMutation,
  DeleteMutationVariables
>
export const PlantDocument = gql`
  query Plant($id: ID!) {
    plant(id: $id) {
      id
      name
      other_names
      description
      plant_season
      harvest_season
      prune_season
      tips
    }
  }
`

/**
 * __usePlantQuery__
 *
 * To run a query within a React component, call `usePlantQuery` and pass it any options that fit your needs.
 * When your component renders, `usePlantQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = usePlantQuery({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function usePlantQuery(
  baseOptions?: ApolloReactHooks.QueryHookOptions<
    PlantQuery,
    PlantQueryVariables
  >,
) {
  return ApolloReactHooks.useQuery<PlantQuery, PlantQueryVariables>(
    PlantDocument,
    baseOptions,
  )
}
export function usePlantLazyQuery(
  baseOptions?: ApolloReactHooks.LazyQueryHookOptions<
    PlantQuery,
    PlantQueryVariables
  >,
) {
  return ApolloReactHooks.useLazyQuery<PlantQuery, PlantQueryVariables>(
    PlantDocument,
    baseOptions,
  )
}
export type PlantQueryHookResult = ReturnType<typeof usePlantQuery>
export type PlantLazyQueryHookResult = ReturnType<typeof usePlantLazyQuery>
export type PlantQueryResult = ApolloReactCommon.QueryResult<
  PlantQuery,
  PlantQueryVariables
>
export const PlantsDocument = gql`
  query Plants {
    plants {
      id
      name
    }
  }
`

/**
 * __usePlantsQuery__
 *
 * To run a query within a React component, call `usePlantsQuery` and pass it any options that fit your needs.
 * When your component renders, `usePlantsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = usePlantsQuery({
 *   variables: {
 *   },
 * });
 */
export function usePlantsQuery(
  baseOptions?: ApolloReactHooks.QueryHookOptions<
    PlantsQuery,
    PlantsQueryVariables
  >,
) {
  return ApolloReactHooks.useQuery<PlantsQuery, PlantsQueryVariables>(
    PlantsDocument,
    baseOptions,
  )
}
export function usePlantsLazyQuery(
  baseOptions?: ApolloReactHooks.LazyQueryHookOptions<
    PlantsQuery,
    PlantsQueryVariables
  >,
) {
  return ApolloReactHooks.useLazyQuery<PlantsQuery, PlantsQueryVariables>(
    PlantsDocument,
    baseOptions,
  )
}
export type PlantsQueryHookResult = ReturnType<typeof usePlantsQuery>
export type PlantsLazyQueryHookResult = ReturnType<typeof usePlantsLazyQuery>
export type PlantsQueryResult = ApolloReactCommon.QueryResult<
  PlantsQuery,
  PlantsQueryVariables
>
