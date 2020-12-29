/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: Plant
// ====================================================

export interface Plant_plant {
  __typename: "Plant"
  id: string
  created_at: any
  edited_at: any | null
  name: string
  other_names: string | null
  description: string | null
  plant_season: string | null
  harvest_season: string | null
  prune_season: string | null
  tips: string | null
}

export interface Plant {
  /**
   * Lists a plant by using its ID, returning an object with the plant, or null if it does not exist.
   */
  plant: Plant_plant | null
}

export interface PlantVariables {
  id: string
}
