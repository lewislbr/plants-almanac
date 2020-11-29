import React, {useEffect, useState} from "react"
import {Link as RouterLink} from "react-router-dom"
import {Link, Typography} from "@material-ui/core"
import {Skeleton} from "@material-ui/lab"
import {PlantCard} from "../components"
import {listAll} from "../services/plant"
import {Plants} from "../graphql"
import {FetchStatus} from "../constants"

export function PlantList(): JSX.Element {
  const [data, setData] = useState({} as Plants)
  const [fetchStatus, setFetchStatus] = useState(FetchStatus.Idle)

  useEffect(() => {
    setFetchStatus(FetchStatus.Loading)
    ;(async (): Promise<void> => {
      try {
        const result = await listAll()

        setData(result.data as Plants)
        setFetchStatus(FetchStatus.Success)
      } catch (error) {
        setFetchStatus(FetchStatus.Error)

        console.error(error)
      }
    })()
  }, [])

  return (
    <>
      <Typography gutterBottom variant="h1">
        {"Plants"}
      </Typography>
      <section>
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
