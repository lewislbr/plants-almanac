import {Plants_plants} from "../graphql/Plants"

export function asc(
  key: keyof Plants_plants,
): (a: Plants_plants | null, b: Plants_plants | null) => number {
  return (a, b): number => a?.[key].localeCompare(b?.[key])
}

export function desc(
  key: keyof Plants_plants,
): (a: Plants_plants | null, b: Plants_plants | null) => number {
  return (a, b): number => b?.[key].localeCompare(a?.[key])
}
