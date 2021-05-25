import React, {useContext, useEffect, useState} from "react"
import {useHistory} from "react-router-dom"
import {Button, Typography} from "@material-ui/core"
import {getInfo, logOut} from "../services/user"
import {UserInfo} from "../interfaces/user"
import {AuthContext} from "../contexts/auth"
import {ACCOUNT, LOG_OUT} from "../constants/copy"
import {Error, Loading} from "../../shared/components"
import {transformDate} from "../../shared/utils/date"
import {CANCEL} from "../../shared/constants/copy"
import {HTTPStatus} from "../../shared/constants/http"

export function Account(): JSX.Element {
  const [errors, setErrors] = useState({
    http: "",
  })
  const [data, setData] = useState({} as UserInfo)
  const [status, setStatus] = useState(HTTPStatus.IDLE)
  const {setAuthenticatedUser} = useContext(AuthContext)
  const history = useHistory()

  useEffect(() => {
    setStatus(HTTPStatus.LOADING)
    ;(async (): Promise<void> => {
      try {
        const result = await getInfo()

        setData(result)
        setStatus(HTTPStatus.SUCCESS)
      } catch (error) {
        setErrors((errors) => ({...errors, http: error.message}))
        setStatus(HTTPStatus.ERROR)

        console.error(error)
      }
    })()
  }, [])

  async function handleLogOut(): Promise<void> {
    setStatus(HTTPStatus.LOADING)

    try {
      await logOut()

      setAuthenticatedUser(false)
      setStatus(HTTPStatus.SUCCESS)

      history.push("/welcome")
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
      {status === HTTPStatus.LOADING ? (
        <Loading />
      ) : (
        <>
          <Typography variant="h1">{ACCOUNT}</Typography>
          <section style={{marginTop: "30px"}}>
            <div style={{marginBottom: "30px"}}>
              <Typography gutterBottom variant="body1">
                {`Hello ${data.name}, your account is registered with the email ${
                  data.email
                } and was created on ${transformDate(data.created_at)}.`}
              </Typography>
            </div>
          </section>
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
          {status === HTTPStatus.ERROR && <Error message={errors.http} title={"Error"} />}
        </>
      )}
    </>
  )
}
