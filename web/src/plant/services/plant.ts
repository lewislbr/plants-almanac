import {NewPlant, Plant} from "../interfaces/plant"

export async function addOne(plant: NewPlant): Promise<void> {
  const response = await fetch(`${process.env.PLANTS_URL as string}/add`, {
    body: JSON.stringify(plant),
    credentials: "include",
    headers: {"Content-Type": "application/json"},
    method: "POST",
  })

  if (!response.ok) {
    throw new Error(await response.text())
  }
}

export async function listAll(): Promise<Plant[]> {
  const response = await fetch(`${process.env.PLANTS_URL as string}/list`, {
    credentials: "include",
  })

  if (!response.ok) {
    throw new Error(await response.text())
  }

  return await response.json()
}

export async function listOne(id: Plant["id"]): Promise<Plant> {
  const response = await fetch(`${process.env.PLANTS_URL as string}/list/${id}`, {
    credentials: "include",
  })

  if (!response.ok) {
    throw new Error(await response.text())
  }

  return await response.json()
}

export async function editOne(id: Plant["id"], plant: NewPlant): Promise<void> {
  const response = await fetch(`${process.env.PLANTS_URL as string}/edit/${id}`, {
    body: JSON.stringify(plant),
    credentials: "include",
    headers: {"Content-Type": "application/json"},
    method: "PUT",
  })

  if (!response.ok) {
    throw new Error(await response.text())
  }
}

export async function deleteOne(id: Plant["id"]): Promise<void> {
  const response = await fetch(`${process.env.PLANTS_URL as string}/delete/${id}`, {
    credentials: "include",
    method: "DELETE",
  })

  if (!response.ok) {
    throw new Error(await response.text())
  }
}
