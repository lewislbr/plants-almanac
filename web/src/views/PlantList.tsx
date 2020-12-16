import React, {ChangeEvent, useEffect, useState} from "react"
import {Link as RouterLink} from "react-router-dom"
import {
  FormControl,
  Link,
  MenuItem,
  Select,
  Typography,
} from "@material-ui/core"
import {Error, Loading, NavBar, PlantCard} from "../components"
import * as plantService from "../services/plant"
import * as sortService from "../services/sort"
import * as storageService from "../services/storage"
import * as copyConstant from "../constants/copy"
import * as errorConstant from "../constants/error"
import * as fetchConstant from "../constants/fetch"
import * as sortConstant from "../constants/sort"
import {Plants} from "../graphql/Plants"

export function PlantList(): JSX.Element {
  const [data, setData] = useState({} as Plants)
  const [fetchStatus, setFetchStatus] = useState(fetchConstant.Status.IDLE)
  const [sortMethod, setSortMethod] = useState(
    storageService.retrieve(sortConstant.SORT_METHOD) ??
      sortConstant.Options.Created.KEY,
  )

  useEffect(() => {
    setFetchStatus(fetchConstant.Status.LOADING)
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

        setFetchStatus(fetchConstant.Status.SUCCESS)
      } catch (error) {
        setFetchStatus(fetchConstant.Status.ERROR)

        console.error(error)
      }
    })()
  }, [sortMethod])

  function sortBy(event: ChangeEvent<{name?: string; value: unknown}>): void {
    switch (event.target.value) {
      case sortConstant.Options.Created.KEY:
        setData({
          plants: data.plants?.slice().sort(sortService.desc("created_at")),
        })
        setSortMethod(sortConstant.Options.Created.KEY)
        storageService.store(
          sortConstant.SORT_METHOD,
          sortConstant.Options.Created.KEY,
        )

        break
      case sortConstant.Options.Edited.KEY:
        setData({
          plants: data.plants?.slice().sort(sortService.desc("edited_at")),
        })
        setSortMethod(sortConstant.Options.Edited.KEY)
        storageService.store(
          sortConstant.SORT_METHOD,
          sortConstant.Options.Edited.KEY,
        )

        break
      case sortConstant.Options.Name.KEY:
        setData({
          plants: data.plants?.slice().sort(sortService.asc("name")),
        })
        setSortMethod(sortConstant.Options.Name.KEY)
        storageService.store(
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
        <Typography variant="h1">{copyConstant.PLANTS}</Typography>
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
        {fetchStatus === fetchConstant.Status.LOADING ? (
          <Loading />
        ) : fetchStatus === fetchConstant.Status.ERROR ? (
          <Error message={errorConstant.GENERIC_MESSAGE} />
        ) : (
          <div>
            {data.plants?.map((plant) => (
              <Link
                component={RouterLink}
                key={plant?.id}
                to={`/${plant?.id}`}
                underline="none"
              >
                <PlantCard {...{name: plant?.name}} />
              </Link>
            ))}
          </div>
        )}
      </section>
      <NavBar />
    </>
  )
}
