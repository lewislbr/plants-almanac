import {Credentials, NewUser} from "../interfaces/user"

export async function signUp(user: Record<string, unknown>): Promise<void> {
  const newUserDTO: NewUser = {
    name: user.name as string,
    email: user.email as string,
    password: user.password as string,
  }
  const response = await fetch(
    process.env.NODE_ENV === "production"
      ? (process.env.USERS_SIGNUP_PRODUCTION_URL as string)
      : (process.env.USERS_SIGNUP_DEVELOPMENT_URL as string),
    {
      body: JSON.stringify(newUserDTO),
      headers: {"Content-Type": "application/json"},
      method: "POST",
    },
  )

  if (!response.ok) {
    throw new Error(await response.text())
  }
}

export async function logIn(user: Record<string, unknown>): Promise<void> {
  const credentialsDTO: Credentials = {
    email: user.email as string,
    password: user.password as string,
  }
  const response = await fetch(
    process.env.NODE_ENV === "production"
      ? (process.env.USERS_LOGIN_PRODUCTION_URL as string)
      : (process.env.USERS_LOGIN_DEVELOPMENT_URL as string),
    {
      body: JSON.stringify(credentialsDTO),
      credentials: "include",
      headers: {"Content-Type": "application/json"},
      method: "POST",
    },
  )

  if (!response.ok) {
    throw new Error(await response.text())
  }
}

export function isAuthenticated(): boolean {
  return document.cookie.split("; ").includes("te=true")
}
