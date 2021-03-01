import React, {createContext, Dispatch, ReactNode, SetStateAction, useState} from "react"
import * as userService from "../services/user"

export const AuthContext = createContext<{
  authenticatedUser: boolean
  setAuthenticatedUser: Dispatch<SetStateAction<boolean>>
}>({authenticatedUser: false, setAuthenticatedUser: () => true})

export function AuthProvider({children}: {children: ReactNode}): JSX.Element {
  const [authenticatedUser, setAuthenticatedUser] = useState(userService.isAuthenticated())

  return (
    <AuthContext.Provider value={{authenticatedUser, setAuthenticatedUser}}>
      {children}
    </AuthContext.Provider>
  )
}
