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
import {NavBar, PlantCard} from "../components"
import {retrieve, store} from "../services/storage"
import {listAll} from "../services/plant"
import {asc, desc} from "../services/sort"
import {Plants} from "../graphql"
import {FetchStatus, SORT_METHOD, SortMethods} from "../constants"

export function PlantList(): JSX.Element {
  const [data, setData] = useState({} as Plants)
  const [fetchStatus, setFetchStatus] = useState(FetchStatus.Idle)
  const [sortMethod, setSortMethod] = useState(
    retrieve(SORT_METHOD) ?? SortMethods.Created,
  )

  useEffect(() => {
    setFetchStatus(FetchStatus.Loading)
    ;(async (): Promise<void> => {
      try {
        const result = await listAll()

        switch (sortMethod) {
          case SortMethods.Created:
            setData({
              plants: result.data.plants?.slice().sort(desc("created_at")),
            })

            break
          case SortMethods.Edited:
            setData({
              plants: result.data.plants?.slice().sort(desc("edited_at")),
            })

            break
          case SortMethods.Name:
            setData({
              plants: result.data.plants?.slice().sort(asc("name")),
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
          plants: data.plants?.slice().sort(desc("created_at")),
        })
        setSortMethod(SortMethods.Created)
        store(SORT_METHOD, SortMethods.Created)

        break
      case SortMethods.Edited:
        setData({
          plants: data.plants?.slice().sort(desc("edited_at")),
        })
        setSortMethod(SortMethods.Edited)
        store(SORT_METHOD, SortMethods.Edited)

        break
      case SortMethods.Name:
        setData({
          plants: data.plants?.slice().sort(asc("name")),
        })
        setSortMethod(SortMethods.Name)
        store(SORT_METHOD, SortMethods.Name)

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
        <Typography variant="h1">{"Plants"}</Typography>
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
      <section style={{marginTop: "30px"}}>
        {fetchStatus === FetchStatus.Loading ? (
          <>
            <Skeleton
              animation="wave"
              height={150}
              style={{marginTop: "-30px"}}
            />
            <Skeleton
              animation="wave"
              height={150}
              style={{marginTop: "-30px"}}
            />
            <Skeleton
              animation="wave"
              height={150}
              style={{marginTop: "-30px"}}
            />
            <Skeleton
              animation="wave"
              height={150}
              style={{marginTop: "-30px"}}
            />
            <Skeleton
              animation="wave"
              height={150}
              style={{marginTop: "-30px"}}
            />
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
      <NavBar />
    </>
  )
}
