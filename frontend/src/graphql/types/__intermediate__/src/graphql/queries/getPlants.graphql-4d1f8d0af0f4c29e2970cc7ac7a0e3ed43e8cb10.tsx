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


export type GetPlantsQueryVariables = {};


export type GetPlantsQuery = (
  { __typename?: 'Query' }
  & { getPlants?: Maybe<Array<(
    { __typename?: 'Plant' }
    & Pick<Plant, '_id' | 'name'>
  )>> }
);


export const GetPlantsDocument = gql`
    query GetPlants {
  getPlants {
    _id
    name
  }
}
    `;
export type GetPlantsComponentProps = Omit<ApolloReactComponents.QueryComponentOptions<GetPlantsQuery, GetPlantsQueryVariables>, 'query'>;

    export const GetPlantsComponent = (props: GetPlantsComponentProps) => (
      <ApolloReactComponents.Query<GetPlantsQuery, GetPlantsQueryVariables> query={GetPlantsDocument} {...props} />
    );
    

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
export function useGetPlantsQuery(baseOptions?: ApolloReactHooks.QueryHookOptions<GetPlantsQuery, GetPlantsQueryVariables>) {
        return ApolloReactHooks.useQuery<GetPlantsQuery, GetPlantsQueryVariables>(GetPlantsDocument, baseOptions);
      }
export function useGetPlantsLazyQuery(baseOptions?: ApolloReactHooks.LazyQueryHookOptions<GetPlantsQuery, GetPlantsQueryVariables>) {
          return ApolloReactHooks.useLazyQuery<GetPlantsQuery, GetPlantsQueryVariables>(GetPlantsDocument, baseOptions);
        }
export type GetPlantsQueryHookResult = ReturnType<typeof useGetPlantsQuery>;
export type GetPlantsLazyQueryHookResult = ReturnType<typeof useGetPlantsLazyQuery>;
export type GetPlantsQueryResult = ApolloReactCommon.QueryResult<GetPlantsQuery, GetPlantsQueryVariables>;