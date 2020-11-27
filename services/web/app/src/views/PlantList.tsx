import React, {useEffect, useState} from "react"
import {Link as RouterLink} from "react-router-dom"
import {Link, Typography} from "@material-ui/core"
import {Skeleton} from "@material-ui/lab"
import {PlantCard} from "../components"
import {listAll} from "../services/plant"
import {Plants} from "../graphql"
import {DataStatus} from "../constants"

export function PlantList(): JSX.Element {
  const [data, setData] = useState({} as Plants)
  const [dataStatus, setDataStatus] = useState(DataStatus.Idle)

  useEffect(() => {
    setDataStatus(DataStatus.Loading)
    ;(async (): Promise<void> => {
      try {
        const result = await listAll()

        setData(result.data as Plants)
        setDataStatus(DataStatus.Success)
      } catch (error) {
        setDataStatus(DataStatus.Error)

        console.error(error)
      }
    })()
  }, [])

  return (
    <>
      <Typography gutterBottom variant="h1">
        {"Plants"}
      </Typography>
      <section className="mt-8">
        {dataStatus === DataStatus.Loading ? (
          <>
            <Skeleton animation="wave" height={50} />
            <Skeleton animation="wave" height={50} />
            <Skeleton animation="wave" height={50} />
            <Skeleton animation="wave" height={50} />
            <Skeleton animation="wave" height={50} />
          </>
        ) : dataStatus === DataStatus.Error ? (
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
