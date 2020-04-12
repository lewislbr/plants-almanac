declare module '*/deletePlant.graphql' {
  /// <reference types="react" />
  import * as ApolloReactCommon from '@apollo/react-common';
  import * as ApolloReactComponents from '@apollo/react-components';
  import * as ApolloReactHooks from '@apollo/react-hooks';
  export type Maybe<T> = T | null;
  export type Omit<T, K extends keyof T> = Pick<T, Exclude<keyof T, K>>;
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
      __typename?: 'Plant';
      _id: Scalars['ID'];
      name: Scalars['String'];
      otherNames?: Maybe<Scalars['String']>;
      description?: Maybe<Scalars['String']>;
      plantSeason?: Maybe<Scalars['String']>;
      harvestSeason?: Maybe<Scalars['String']>;
      pruneSeason?: Maybe<Scalars['String']>;
      tips?: Maybe<Scalars['String']>;
  };
  export type Query = {
      __typename?: 'Query';
      getPlants?: Maybe<Array<Plant>>;
      getPlant?: Maybe<Plant>;
  };
  export type QueryGetPlantArgs = {
      name: Scalars['String'];
  };
  export type Mutation = {
      __typename?: 'Mutation';
      addPlant?: Maybe<Plant>;
      deletePlant?: Maybe<Plant>;
  };
  export type MutationAddPlantArgs = {
      name: Scalars['String'];
      otherNames?: Maybe<Scalars['String']>;
      description?: Maybe<Scalars['String']>;
      plantSeason?: Maybe<Scalars['String']>;
      harvestSeason?: Maybe<Scalars['String']>;
      pruneSeason?: Maybe<Scalars['String']>;
      tips?: Maybe<Scalars['String']>;
  };
  export type MutationDeletePlantArgs = {
      _id: Scalars['ID'];
  };
  export enum CacheControlScope {
      Public = "PUBLIC",
      Private = "PRIVATE"
  }
  export type DeletePlantMutationVariables = {
      _id: Scalars['ID'];
  };
  export type DeletePlantMutation = ({
      __typename?: 'Mutation';
  } & {
      deletePlant?: Maybe<({
          __typename?: 'Plant';
      } & Pick<Plant, '_id' | 'name' | 'otherNames' | 'description' | 'plantSeason' | 'harvestSeason' | 'pruneSeason' | 'tips'>)>;
  });
  export const DeletePlantDocument: import("graphql").DocumentNode;
  export type DeletePlantMutationFn = ApolloReactCommon.MutationFunction<DeletePlantMutation, DeletePlantMutationVariables>;
  export type DeletePlantComponentProps = Omit<ApolloReactComponents.MutationComponentOptions<DeletePlantMutation, DeletePlantMutationVariables>, 'mutation'>;
  export const DeletePlantComponent: (props: Pick<ApolloReactComponents.MutationComponentOptions<DeletePlantMutation, DeletePlantMutationVariables>, "client" | "update" | "children" | "variables" | "onCompleted" | "onError" | "fetchPolicy" | "errorPolicy" | "notifyOnNetworkStatusChange" | "context" | "optimisticResponse" | "refetchQueries" | "awaitRefetchQueries" | "ignoreResults">) => JSX.Element;
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
  export function useDeletePlantMutation(baseOptions?: ApolloReactHooks.MutationHookOptions<DeletePlantMutation, DeletePlantMutationVariables>): ApolloReactHooks.MutationTuple<DeletePlantMutation, DeletePlantMutationVariables>;
  export type DeletePlantMutationHookResult = ReturnType<typeof useDeletePlantMutation>;
  export type DeletePlantMutationResult = ApolloReactCommon.MutationResult<DeletePlantMutation>;
  export type DeletePlantMutationOptions = ApolloReactCommon.BaseMutationOptions<DeletePlantMutation, DeletePlantMutationVariables>;
}