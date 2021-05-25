import React, {ChangeEvent, useContext, useEffect, useState} from "react"
import {useHistory} from "react-router-dom"
import {Button, IconButton, InputAdornment, TextField, Typography} from "@material-ui/core"
import {Visibility, VisibilityOff} from "@material-ui/icons"
import {Error, Loading} from "../../shared/components"
import {logIn} from "../services/user"
import {LOG_IN} from "../constants/copy"
import {EMAIL_PATTERN, PASSWORD_PATTERN} from "../constants/user"
import {CANCEL} from "../../shared/constants/copy"
import {HTTPStatus} from "../../shared/constants/http"
import {AuthContext} from "../contexts/auth"

export function LogIn(): JSX.Element {
  const [errors, setErrors] = useState({
    email: false,
    password: false,
    http: "",
  })
  const [buttonDisabled, setButtonDisabled] = useState(true)
  const [status, setStatus] = useState(HTTPStatus.IDLE)
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [showPassword, setShowPassword] = useState(false)
  const {setAuthenticatedUser} = useContext(AuthContext)
  const history = useHistory()
  const missingFields = !email || !password
  const activeErrors = Object.values(errors).includes(true)

  useEffect(() => {
    if (missingFields || activeErrors) {
      setButtonDisabled(true)
    } else {
      setButtonDisabled(false)
    }
  }, [missingFields, activeErrors])

  function updateEmail(event: ChangeEvent<HTMLInputElement>): void {
    if (!EMAIL_PATTERN.test(event.target.value)) {
      setErrors((errors) => ({...errors, email: true}))
    } else {
      setErrors((errors) => ({...errors, email: false}))
    }

    setEmail(event.target.value)
  }

  function updatePassword(event: ChangeEvent<HTMLTextAreaElement>): void {
    if (!PASSWORD_PATTERN.test(event.target.value)) {
      setErrors((errors) => ({...errors, password: true}))
    } else {
      setErrors((errors) => ({...errors, password: false}))
    }

    setPassword(event.target.value)
  }

  function toggleShowPassword(): void {
    setShowPassword(!showPassword)
  }

  async function handleLogIn(): Promise<void> {
    setStatus(HTTPStatus.LOADING)

    try {
      await logIn({
        email,
        password,
      })

      setAuthenticatedUser(true)
      setStatus(HTTPStatus.SUCCESS)

      history.push("/plants")
    } catch (error) {
      setErrors((errors) => ({...errors, http: error.message}))
      setStatus(HTTPStatus.ERROR)

      console.error(error)
    }
  }

  function cancel(): void {
    history.push("/")
  }

  return (
    <>
      <Typography variant="h1">{LOG_IN}</Typography>
      {status === HTTPStatus.LOADING ? (
        <Loading />
      ) : (
        <section style={{marginTop: "30px"}}>
          <TextField
            error={errors.email}
            fullWidth
            label="Email"
            onChange={updateEmail}
            required
            type="email"
            value={email}
            variant="outlined"
          />
          <TextField
            error={errors.password}
            fullWidth
            InputProps={{
              endAdornment: (
                <InputAdornment position="end">
                  <IconButton edge="end" onClick={toggleShowPassword}>
                    {showPassword ? <Visibility /> : <VisibilityOff />}
                  </IconButton>
                </InputAdornment>
              ),
            }}
            label="Password"
            onChange={updatePassword}
            required
            type={showPassword ? "text" : "password"}
            value={password}
            variant="outlined"
          />
          <Button
            color="primary"
            disabled={buttonDisabled}
            fullWidth
            onClick={handleLogIn}
            style={{marginTop: "30px"}}
            variant="contained"
          >
            {LOG_IN}
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
      )}
      {status === HTTPStatus.ERROR && <Error message={errors.http} title={"Error"} />}
    </>
  )
}
