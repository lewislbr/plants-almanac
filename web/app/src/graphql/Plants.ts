/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: Plants
// ====================================================

export interface Plants_plants {
  __typename: "Plant"
  id: string | null
  name: string | null
}

export interface Plants {
  /**
   * Returns all plants
   */
  plants: (Plants_plants | null)[] | null
}
