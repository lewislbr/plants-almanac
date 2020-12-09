import React, {ChangeEvent, useEffect, useState} from "react"
import {Link as RouterLink} from "react-router-dom"
import {
  FormControl,
  Link,
  MenuItem,
  Select,
  Typography,
} from "@material-ui/core"
import {Skeleton} from "@material-ui/lab"
import {PlantCard} from "../components"
import {listAll} from "../services/plant"
import {Plants} from "../graphql"
import {FetchStatus, SORT_METHOD, SortMethods} from "../constants"

function getSortMethod(): string {
  const storedSortMethod = localStorage.getItem(SORT_METHOD)

  return storedSortMethod || SortMethods.Created
}

export function PlantList(): JSX.Element {
  const [data, setData] = useState({} as Plants)
  const [fetchStatus, setFetchStatus] = useState(FetchStatus.Idle)
  const [sortMethod, setSortMethod] = useState(getSortMethod())

  useEffect(() => {
    setFetchStatus(FetchStatus.Loading)
    ;(async (): Promise<void> => {
      try {
        const result = await listAll()

        switch (sortMethod) {
          case SortMethods.Created:
            setData({
              plants: result.data.plants
                ?.slice()
                .sort((a, b) => b?.created_at.localeCompare(a?.created_at)),
            })

            break
          case SortMethods.Edited:
            setData({
              plants: result.data.plants
                ?.slice()
                .sort((a, b) => b?.edited_at.localeCompare(a?.edited_at)),
            })

            break
          case SortMethods.Name:
            setData({
              plants: result.data.plants?.slice().sort((a, b) => {
                if (a?.name && b?.name) {
                  return a?.name.localeCompare(b?.name)
                }

                return 0
              }),
            })

            break
        }

        setFetchStatus(FetchStatus.Success)
      } catch (error) {
        setFetchStatus(FetchStatus.Error)

        console.error(error)
      }
    })()
  }, [sortMethod])

  function sortBy(event: ChangeEvent<{name?: string; value: unknown}>): void {
    switch (event.target.value) {
      case SortMethods.Created:
        setData({
          plants: data.plants
            ?.slice()
            .sort((a, b) => b?.created_at.localeCompare(a?.created_at)),
        })
        setSortMethod(SortMethods.Created)

        localStorage.setItem(SORT_METHOD, SortMethods.Created)

        break
      case SortMethods.Edited:
        setData({
          plants: data.plants
            ?.slice()
            .sort((a, b) => b?.edited_at.localeCompare(a?.edited_at)),
        })
        setSortMethod(SortMethods.Edited)

        localStorage.setItem(SORT_METHOD, SortMethods.Edited)

        break
      case SortMethods.Name:
        setData({
          plants: data.plants?.slice().sort((a, b) => {
            if (a?.name && b?.name) {
              return a?.name.localeCompare(b?.name)
            }

            return 0
          }),
        })
        setSortMethod(SortMethods.Name)

        localStorage.setItem(SORT_METHOD, SortMethods.Name)

        break
    }
  }

  return (
    <>
      <div
        style={{
          alignItems: "baseline",
          display: "flex",
          justifyContent: "space-between",
        }}
      >
        <Typography gutterBottom variant="h1">
          {"Plants"}
        </Typography>
        <FormControl
          style={{fontSize: "15px", padding: "0"}}
          variant="outlined"
        >
          <Select onChange={(event): void => sortBy(event)} value={sortMethod}>
            <MenuItem value={SortMethods.Created}>
              {"Sort by: Created"}
            </MenuItem>
            <MenuItem value={SortMethods.Edited}>{"Sort by: Edited"}</MenuItem>
            <MenuItem value={SortMethods.Name}>{"Sort by: Name"}</MenuItem>
          </Select>
        </FormControl>
      </div>
      <section style={{marginTop: "20px"}}>
        {fetchStatus === FetchStatus.Loading ? (
          <>
            <Skeleton animation="wave" height={80} />
            <Skeleton animation="wave" height={80} />
            <Skeleton animation="wave" height={80} />
            <Skeleton animation="wave" height={80} />
            <Skeleton animation="wave" height={80} />
          </>
        ) : fetchStatus === FetchStatus.Error ? (
          <Typography>{"ERROR"}</Typography>
        ) : (
          <div>
            {data.plants?.map((plant) => (
              <Link component={RouterLink} key={plant?.id} to={`/${plant?.id}`}>
                <PlantCard {...{name: plant?.name}} />
              </Link>
            ))}
          </div>
        )}
      </section>
    </>
  )
}
