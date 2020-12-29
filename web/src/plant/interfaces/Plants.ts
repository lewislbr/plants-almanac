/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: Plants
// ====================================================

export interface Plants_plants {
  __typename: "Plant"
  id: string
  created_at: any
  edited_at: any | null
  name: string
}

export interface Plants {
  /**
   * Lists all plants, returning an array of objects with the existing plants, or an empty array if there are none.
   */
  plants: (Plants_plants | null)[]
}
