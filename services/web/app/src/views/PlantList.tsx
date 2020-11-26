import React, {useEffect, useState} from "react"
import {Link} from "react-router-dom"
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
      <section>
        <h1 className="page-title">{"Plants"}</h1>
      </section>
      <section className="mt-8">
        {dataStatus === DataStatus.Loading ? (
          <p>{"Loading..."}</p>
        ) : dataStatus === DataStatus.Error ? (
          <p>{"ERROR"}</p>
        ) : (
          <div>
            {data?.plants?.map((plant) => (
              <Link to={`/${plant?.id}`} key={plant?.id}>
                <PlantCard {...{name: plant?.name}} />
              </Link>
            ))}
          </div>
        )}
      </section>
    </>
  )
}
