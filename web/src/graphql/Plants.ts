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
  created_at: any | null
  edited_at: any | null
  name: string | null
}

export interface Plants {
  /**
   * Lists all plants, returning an array of objects with the existing plants, or an empty array if there are none.
   */
  plants: (Plants_plants | null)[] | null
}
