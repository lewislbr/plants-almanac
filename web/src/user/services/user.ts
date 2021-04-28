import {Credentials, NewUser, UserInfo} from "../interfaces/user"

export async function signUp(user: NewUser): Promise<void> {
  const response = await fetch(`${process.env.USERS_URL as string}/registration`, {
    body: JSON.stringify(user),
    headers: {"Content-Type": "application/json"},
    method: "POST",
  })

  if (!response.ok) {
    throw new Error(await response.text())
  }
}

export async function logIn(credentials: Credentials): Promise<void> {
  const response = await fetch(`${process.env.USERS_URL as string}/login`, {
    body: JSON.stringify(credentials),
    credentials: "include",
    headers: {"Content-Type": "application/json"},
    method: "POST",
  })

  if (!response.ok) {
    throw new Error(await response.text())
  }
}

export async function refreshToken(): Promise<void> {
  const response = await fetch(`${process.env.USERS_URL as string}/refresh`, {
    credentials: "include",
  })

  if (!response.ok) {
    throw new Error(await response.text())
  }
}

export function isAuthenticated(): boolean {
  return document.cookie.split("; ").includes("te=1")
}

export async function logOut(): Promise<void> {
  const response = await fetch(`${process.env.USERS_URL as string}/logout`, {
    credentials: "include",
  })

  if (!response.ok) {
    throw new Error(await response.text())
  }
}

export async function getInfo(): Promise<UserInfo> {
  const response = await fetch(`${process.env.USERS_URL as string}/info`, {
    credentials: "include",
  })

  if (!response.ok) {
    throw new Error(await response.text())
  }

  return await response.json()
}
