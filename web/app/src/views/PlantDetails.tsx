import * as React from "react"
import {gql, useMutation, useQuery} from "@apollo/client"
import {Alert} from "../components"
import {Delete, Plant} from "../graphql"

const PLANT = gql`
  query Plant($id: ID!) {
    plant(id: $id) {
      id
      name
      other_names
      description
      plant_season
      harvest_season
      prune_season
      tips
    }
  }
`

const DELETE = gql`
  mutation Delete($id: ID!) {
    delete(id: $id)
  }
`

export function PlantDetails({
  history,
  match,
}: {
  history: any
  match: any
}): JSX.Element {
  const {data, loading, error} = useQuery<Plant>(PLANT, {
    variables: {id: match.params.id},
  })
  const [deletePlant] = useMutation<Delete>(DELETE)
  const [alertOpen, setAlertOpen] = React.useState(false)

  function openAlert(): void {
    setAlertOpen(true)
  }

  async function submitDeletePlant(event: React.SyntheticEvent): Promise<void> {
    event.preventDefault()

    await deletePlant({
      variables: {id: data?.plant?.id as string},
    })

    history.push("/")
  }

  return (
    <>
      {loading ? (
        <p>{"Loading..."}</p>
      ) : error ? (
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
            <button
              className="button button-danger"
              type="button"
              onClick={openAlert}
            >
              {"Delete plant"}
            </button>
          </div>
          {alertOpen ? (
            <Alert {...{deletePlant: submitDeletePlant, setAlertOpen}} />
          ) : null}
        </>
      )}
    </>
  )
}
