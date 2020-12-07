/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL mutation operation: Edit
// ====================================================

export interface Edit {
  /**
   * Edits a plant by using its ID, and adding any new values for the existing fields, returning an integer with the numbers of plants edited.
   */
  edit: number | null
}

export interface EditVariables {
  id: string
  name: string
  other_names?: string | null
  description?: string | null
  plant_season?: string | null
  harvest_season?: string | null
  prune_season?: string | null
  tips?: string | null
}
