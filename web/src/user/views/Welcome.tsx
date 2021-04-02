import React from "react"
import {useHistory} from "react-router-dom"
import {Button, Typography} from "@material-ui/core"
import {CREATE_ACCOUNT, LOG_IN, WELCOME} from "../constants/copy"

export function Welcome(): JSX.Element {
  const history = useHistory()

  function goToLogIn(): void {
    history.push("/log-in")
  }

  function goToCreateAccount(): void {
    history.push("/create-account")
  }

  return (
    <>
      <Typography variant="h1">{WELCOME}</Typography>
      <section style={{marginTop: "50vh"}}>
        <Button
          color="primary"
          fullWidth
          onClick={goToLogIn}
          style={{marginTop: "30px"}}
          variant="contained"
        >
          {LOG_IN}
        </Button>
        <Button
          color="secondary"
          fullWidth
          onClick={goToCreateAccount}
          style={{marginTop: "30px"}}
          variant="contained"
        >
          {CREATE_ACCOUNT}
        </Button>
      </section>
    </>
  )
}
