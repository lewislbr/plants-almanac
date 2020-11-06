/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

// ====================================================
// GraphQL query operation: Plant
// ====================================================

export interface Plant_plant {
  __typename: "Plant"
  id: string | null
  name: string | null
  other_names: string | null
  description: string | null
  plant_season: string | null
  harvest_season: string | null
  prune_season: string | null
  tips: string | null
}

export interface Plant {
  /**
   * Returns a plant
   */
  plant: Plant_plant | null
}

export interface PlantVariables {
  id: string
}
