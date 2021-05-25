import React, {ChangeEvent, useEffect, useState} from "react"
import {Link as RouterLink} from "react-router-dom"
import {FormControl, Link, MenuItem, Select, Typography} from "@material-ui/core"
import {PlantCard} from "../components"
import {Error, Loading, NavBar} from "../../shared/components"
import {listAll} from "../services/plant"
import {asc, desc} from "../utils/sort"
import {PLANTS} from "../constants/copy"
import {HTTPStatus} from "../../shared/constants/http"
import {Options, SORT_METHOD} from "../constants/sort"
import {Plant} from "../interfaces/plant"

export function PlantList(): JSX.Element {
  const [errors, setErrors] = useState({
    http: "",
  })
  const [data, setData] = useState([] as Plant[])
  const [status, setStatus] = useState(HTTPStatus.IDLE)
  const [sortMethod, setSortMethod] = useState(
    localStorage.getItem(SORT_METHOD) ?? Options.Created.KEY,
  )

  useEffect(() => {
    setStatus(HTTPStatus.LOADING)
    ;(async (): Promise<void> => {
      try {
        const result = (await listAll()) || []

        switch (sortMethod) {
          case Options.Created.KEY:
            setData(result.slice().sort(desc("created_at")))

            break
          case Options.Edited.KEY:
            setData(result.slice().sort(desc("edited_at")))

            break
          case Options.Name.KEY:
            setData(result.slice().sort(asc("name")))

            break
        }

        setStatus(HTTPStatus.SUCCESS)
      } catch (error) {
        setErrors((errors) => ({...errors, http: error.message}))
        setStatus(HTTPStatus.ERROR)

        console.error(error)
      }
    })()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  function sortBy(event: ChangeEvent<{name?: string; value: unknown}>): void {
    switch (event.target.value) {
      case Options.Created.KEY:
        setData(data.slice().sort(desc("created_at")))
        setSortMethod(Options.Created.KEY)

        localStorage.setItem(SORT_METHOD, Options.Created.KEY)

        break
      case Options.Edited.KEY:
        setData(data.slice().sort(desc("edited_at")))
        setSortMethod(Options.Edited.KEY)

        localStorage.setItem(SORT_METHOD, Options.Edited.KEY)

        break
      case Options.Name.KEY:
        setData(data.slice().sort(asc("name")))
        setSortMethod(Options.Name.KEY)

        localStorage.setItem(SORT_METHOD, Options.Name.KEY)

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
        <Typography variant="h1">{PLANTS}</Typography>
        <FormControl style={{fontSize: "15px", padding: "0"}} variant="outlined">
          <Select onChange={(event): void => sortBy(event)} value={sortMethod}>
            <MenuItem value={Options.Created.KEY}>{Options.Created.TEXT}</MenuItem>
            <MenuItem value={Options.Edited.KEY}>{Options.Edited.TEXT}</MenuItem>
            <MenuItem value={Options.Name.KEY}>{Options.Name.TEXT}</MenuItem>
          </Select>
        </FormControl>
      </div>
      <section style={{marginTop: "30px"}}>
        {status === HTTPStatus.LOADING ? (
          <Loading />
        ) : (
          <div>
            {data.map((plant) => (
              <Link
                component={RouterLink}
                key={plant.id}
                to={`/plants/${plant.id}`}
                underline="none"
              >
                <PlantCard {...{name: plant.name}} />
              </Link>
            ))}
          </div>
        )}
      </section>
      {status === HTTPStatus.ERROR && <Error message={errors.http} title={"Error"} />}
      <NavBar />
    </>
  )
}
