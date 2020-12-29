import {Credentials, NewUser} from "../interfaces/user"
import {JWT} from "../constants/user"
import * as storageService from "../../shared/services/storage"

export async function signUp(user: Record<string, unknown>): Promise<void> {
  const dto: NewUser = {
    name: user.name as string,
    email: user.email as string,
    password: user.password as string,
  }
  const response = await fetch(
    process.env.NODE_ENV === "production"
      ? (process.env.USERS_SIGNUP_PRODUCTION_URL as string)
      : (process.env.USERS_SIGNUP_DEVELOPMENT_URL as string),
    {
      body: JSON.stringify(dto),
      headers: {"Content-Type": "application/json"},
      method: "POST",
    },
  )

  if (!response.ok) {
    throw new Error(await response.text())
  }
}

export async function logIn(user: Record<string, unknown>): Promise<void> {
  const dto: Credentials = {
    email: user.email as string,
    password: user.password as string,
  }
  const response = await fetch(
    process.env.NODE_ENV === "production"
      ? (process.env.USERS_LOGIN_PRODUCTION_URL as string)
      : (process.env.USERS_LOGIN_DEVELOPMENT_URL as string),
    {
      body: JSON.stringify(dto),
      headers: {"Content-Type": "application/json"},
      method: "POST",
    },
  )

  if (!response.ok) {
    throw new Error(await response.text())
  }

  const data = await response.text()

  storageService.store(JWT, data)
}

export function isAuthenticated(): boolean {
  const jwt = storageService.retrieve(JWT)

  if (!jwt) {
    return false
  }

  return true
}
