import React, {ChangeEvent, useEffect, useState} from "react"
import {Link as RouterLink} from "react-router-dom"
import {
  FormControl,
  Link,
  MenuItem,
  Select,
  Typography,
} from "@material-ui/core"
import {PlantCard} from "../components"
import {Error, Loading, NavBar} from "../../shared/components"
import * as plantService from "../services/plant"
import * as sortService from "../services/sort"
import * as plantCopy from "../constants/copy"
import {HTTPStatus} from "../../shared/constants/http"
import * as sortConstant from "../constants/sort"
import {Plants} from "../interfaces/Plants"

export function PlantList(): JSX.Element {
  const [errors, setErrors] = useState({
    http: "",
  })
  const [data, setData] = useState({} as Plants)
  const [status, setStatus] = useState(HTTPStatus.IDLE)
  const [sortMethod, setSortMethod] = useState(
    localStorage.getItem(sortConstant.SORT_METHOD) ??
      sortConstant.Options.Created.KEY,
  )

  useEffect(() => {
    setStatus(HTTPStatus.LOADING)
    ;(async (): Promise<void> => {
      try {
        const result = await plantService.listAll()

        switch (sortMethod) {
          case sortConstant.Options.Created.KEY:
            setData({
              plants: result.data.plants
                ?.slice()
                .sort(sortService.desc("created_at")),
            })

            break
          case sortConstant.Options.Edited.KEY:
            setData({
              plants: result.data.plants
                ?.slice()
                .sort(sortService.desc("edited_at")),
            })

            break
          case sortConstant.Options.Name.KEY:
            setData({
              plants: result.data.plants?.slice().sort(sortService.asc("name")),
            })

            break
        }

        setStatus(HTTPStatus.SUCCESS)
      } catch (error) {
        setErrors((errors) => ({...errors, http: String(error)}))
        setStatus(HTTPStatus.ERROR)

        console.error(error)
      }
    })()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  function sortBy(event: ChangeEvent<{name?: string; value: unknown}>): void {
    switch (event.target.value) {
      case sortConstant.Options.Created.KEY:
        setData({
          plants: data.plants?.slice().sort(sortService.desc("created_at")),
        })
        setSortMethod(sortConstant.Options.Created.KEY)

        localStorage.setItem(
          sortConstant.SORT_METHOD,
          sortConstant.Options.Created.KEY,
        )

        break
      case sortConstant.Options.Edited.KEY:
        setData({
          plants: data.plants?.slice().sort(sortService.desc("edited_at")),
        })
        setSortMethod(sortConstant.Options.Edited.KEY)

        localStorage.setItem(
          sortConstant.SORT_METHOD,
          sortConstant.Options.Edited.KEY,
        )

        break
      case sortConstant.Options.Name.KEY:
        setData({
          plants: data.plants?.slice().sort(sortService.asc("name")),
        })
        setSortMethod(sortConstant.Options.Name.KEY)

        localStorage.setItem(
          sortConstant.SORT_METHOD,
          sortConstant.Options.Name.KEY,
        )

        break
    }
  }

  return (
    <>
      <div
        style={{
          alignItems: "flex-end",
          display: "flex",
          justifyContent: "space-between",
        }}
      >
        <Typography variant="h1">{plantCopy.PLANTS}</Typography>
        <FormControl
          style={{fontSize: "15px", padding: "0"}}
          variant="outlined"
        >
          <Select onChange={(event): void => sortBy(event)} value={sortMethod}>
            <MenuItem value={sortConstant.Options.Created.KEY}>
              {sortConstant.Options.Created.TEXT}
            </MenuItem>
            <MenuItem value={sortConstant.Options.Edited.KEY}>
              {sortConstant.Options.Edited.TEXT}
            </MenuItem>
            <MenuItem value={sortConstant.Options.Name.KEY}>
              {sortConstant.Options.Name.TEXT}
            </MenuItem>
          </Select>
        </FormControl>
      </div>
      <section style={{marginTop: "30px"}}>
        {status === HTTPStatus.LOADING ? (
          <Loading />
        ) : (
          <div>
            {data.plants?.map((plant) => (
              <Link
                component={RouterLink}
                key={plant?.id}
                to={`/plants/${plant?.id}`}
                underline="none"
              >
                <PlantCard {...{name: plant?.name}} />
              </Link>
            ))}
          </div>
        )}
      </section>
      {status === HTTPStatus.ERROR && (
        <Error message={errors.http} title={"Error"} />
      )}
      <NavBar />
    </>
  )
}
