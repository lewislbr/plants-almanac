import {Plant} from "../interfaces/plant"

export function asc(key: keyof Plant): (a: Plant, b: Plant) => number {
  return (a, b): number => a[key].localeCompare(b[key])
}

export function desc(key: keyof Plant): (a: Plant, b: Plant) => number {
  return (a, b): number => b[key].localeCompare(a[key])
}
