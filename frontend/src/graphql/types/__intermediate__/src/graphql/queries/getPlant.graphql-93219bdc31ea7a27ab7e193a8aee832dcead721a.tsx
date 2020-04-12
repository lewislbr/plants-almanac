import gql from 'graphql-tag';
import * as React from 'react';
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
  Public = 'PUBLIC',
  Private = 'PRIVATE'
}


export type GetPlantQueryVariables = {
  name: Scalars['String'];
};


export type GetPlantQuery = (
  { __typename?: 'Query' }
  & { getPlant?: Maybe<(
    { __typename?: 'Plant' }
    & Pick<Plant, '_id' | 'name' | 'otherNames' | 'description' | 'plantSeason' | 'harvestSeason' | 'pruneSeason' | 'tips'>
  )> }
);


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
export type GetPlantComponentProps = Omit<ApolloReactComponents.QueryComponentOptions<GetPlantQuery, GetPlantQueryVariables>, 'query'> & ({ variables: GetPlantQueryVariables; skip?: boolean; } | { skip: boolean; });

    export const GetPlantComponent = (props: GetPlantComponentProps) => (
      <ApolloReactComponents.Query<GetPlantQuery, GetPlantQueryVariables> query={GetPlantDocument} {...props} />
    );
    

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
export function useGetPlantQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<GetPlantQuery, GetPlantQueryVariables>) {
        return ApolloReactHooks.useQuery<GetPlantQuery, GetPlantQueryVariables>(GetPlantDocument, baseOptions);
      }
export function useGetPlantLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<GetPlantQuery, GetPlantQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<GetPlantQuery, GetPlantQueryVariables>(GetPlantDocument, baseOptions);
        }
export type GetPlantQueryHookResult = ReturnType<typeof useGetPlantQuery>;
export type GetPlantLazyQueryHookResult = ReturnType<typeof useGetPlantLazyQuery>;
export type GetPlantQueryResult = ApolloReactCommon.QueryResult<GetPlantQuery, GetPlantQueryVariables>;