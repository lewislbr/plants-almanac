import React, {useContext, useState} from "react"
import {useHistory} from "react-router-dom"
import {Button, Typography} from "@material-ui/core"
import {logOut} from "../services/user"
import {AuthContext} from "../contexts/auth"
import {ACCOUNT, LOG_OUT} from "../constants/copy"
import {CANCEL} from "../../shared/constants/copy"
import {HTTPStatus} from "../../shared/constants/http"

export function Account(): JSX.Element {
  const [errors, setErrors] = useState({
    http: "",
  })
  const [status, setStatus] = useState(HTTPStatus.IDLE)
  const {setAuthenticatedUser} = useContext(AuthContext)
  const history = useHistory()

  async function handleLogOut(): Promise<void> {
    setStatus(HTTPStatus.LOADING)

    try {
      await logOut()

      setAuthenticatedUser(false)
      setStatus(HTTPStatus.SUCCESS)

      history.push("/welcome")
    } catch (error) {
      setErrors((errors) => ({...errors, http: String(error)}))
      setStatus(HTTPStatus.ERROR)

      console.error(error)
    }
  }

  function cancel(): void {
    history.push("/")
  }

  return (
    <>
      <Typography variant="h1">{ACCOUNT}</Typography>
      <section style={{marginTop: "50vh"}}>
        <Button
          color="primary"
          fullWidth
          onClick={handleLogOut}
          style={{marginTop: "30px"}}
          variant="contained"
        >
          {LOG_OUT}
        </Button>
        <Button
          color="secondary"
          fullWidth
          onClick={cancel}
          style={{marginTop: "30px"}}
          variant="contained"
        >
          {CANCEL}
        </Button>
      </section>
    </>
  )
}
