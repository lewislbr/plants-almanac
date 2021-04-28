export interface NewUser {
  name: string
  email: string
  password: string
}

export interface Credentials {
  email: string
  password: string
}

export interface UserInfo {
  name: string
  email: string
  created_at: Date
}
