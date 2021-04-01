export interface Plant {
  id: string
  created_at: string
  edited_at: string
  name: string
  other_names: string
  description: string
  plant_season: string
  harvest_season: string
  prune_season: string
  tips: string
}

export type NewPlant = Pick<
  Plant,
  | "name"
  | "other_names"
  | "description"
  | "plant_season"
  | "harvest_season"
  | "prune_season"
  | "tips"
>
