import React, {useEffect, useState} from "react"
import {useParams} from "react-router-dom"
import {Alert} from "../components"
import {deleteOne, listOne} from "../services"
import {Plant} from "../graphql"
import {DataStatus} from "../constants"

export function PlantDetails(): JSX.Element {
  const [data, setData] = useState({} as Plant)
  const [dataStatus, setDataStatus] = useState(DataStatus.Idle)
  const [alertOpen, setAlertOpen] = useState(false)
  const {id} = useParams<{id: string}>()

  useEffect(() => {
    setDataStatus(DataStatus.Loading)
    ;(async (): Promise<void> => {
      try {
        const result = await listOne(id)

        setData(result.data as Plant)
        setDataStatus(DataStatus.Success)
      } catch (error) {
        setDataStatus(DataStatus.Error)

        console.error(error)
      }
    })()
  }, [id])

  function openAlert(): void {
    setAlertOpen(true)
  }

  return (
    <>
      {dataStatus === DataStatus.Loading ? (
        <p>{"Loading..."}</p>
      ) : dataStatus === DataStatus.Error ? (
        <p>{"ERROR"}</p>
      ) : (
        <>
          <section>
            <h1 className="page-title">{data?.plant?.name}</h1>
          </section>
          <section className="mb-12">
            <h5 className="data-title">{"Other Names:"}</h5>
            <p className="data-body">
              {data?.plant?.other_names || "No data yet"}
            </p>
            <h5 className="data-title">{"Description:"}</h5>
            <p className="data-body">
              {data?.plant?.description || "No data yet"}
            </p>
            <h5 className="data-title">{"Plant Season:"}</h5>
            <p className="data-body">
              {data?.plant?.plant_season || "No data yet"}
            </p>
            <h5 className="data-title">{"Harvest Season:"}</h5>
            <p className="data-body">
              {data?.plant?.harvest_season || "No data yet"}
            </p>
            <h5 className="data-title">{"Prune Season:"}</h5>
            <p className="data-body">
              {data?.plant?.prune_season || "No data yet"}
            </p>
            <h5 className="data-title">{"Tips:"}</h5>
            <p className="data-body">{data?.plant?.tips || "No data yet"}</p>
          </section>
          <div className="flex justify-center">
            <button className="button button-danger" onClick={openAlert}>
              {"Delete plant"}
            </button>
          </div>
          {alertOpen ? (
            <Alert {...{deletePlant: deleteOne, id, setAlertOpen}} />
          ) : null}
        </>
      )}
    </>
  )
}
