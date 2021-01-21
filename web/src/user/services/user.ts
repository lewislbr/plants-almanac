import {Credentials, NewUser} from "../interfaces/user"

export async function signUp(user: NewUser): Promise<void> {
  const newUserDTO = {
    name: user.name,
    email: user.email,
    password: user.password,
  }
  const response = await fetch(process.env.USERS_SIGNUP_URL as string, {
    body: JSON.stringify(newUserDTO),
    headers: {"Content-Type": "application/json"},
    method: "POST",
  })

  if (!response.ok) {
    throw new Error(await response.text())
  }
}

export async function logIn(user: Credentials): Promise<void> {
  const credentialsDTO = {
    email: user.email,
    password: user.password,
  }
  const response = await fetch(process.env.USERS_LOGIN_URL as string, {
    body: JSON.stringify(credentialsDTO),
    credentials: "include",
    headers: {"Content-Type": "application/json"},
    method: "POST",
  })

  if (!response.ok) {
    throw new Error(await response.text())
  }
}

export function isAuthenticated(): boolean {
  return document.cookie.split("; ").includes("te=true")
}
