import * as React from "react"
import {gql, useMutation} from "@apollo/client"
import {Add} from "../graphql"

const ADD = gql`
  mutation Add(
    $name: String!
    $other_names: String
    $description: String
    $plant_season: String
    $harvest_season: String
    $prune_season: String
    $tips: String
  ) {
    add(
      name: $name
      other_names: $other_names
      description: $description
      plant_season: $plant_season
      harvest_season: $harvest_season
      prune_season: $prune_season
      tips: $tips
    )
  }
`

export function AddPlant({history}: {history: any}): JSX.Element {
  const [addPlant] = useMutation<Add>(ADD)
  const [name, setName] = React.useState("")
  const [otherNames, setOtherNames] = React.useState("")
  const [description, setDescription] = React.useState("")
  const [plantSeason, setPlantSeason] = React.useState("")
  const [harvestSeason, setHarvestSeason] = React.useState("")
  const [pruneSeason, setPruneSeason] = React.useState("")
  const [tips, setTips] = React.useState("")

  function updateName(event: React.ChangeEvent<HTMLInputElement>) {
    setName(event.target.value)
  }

  function updateOtherNames(event: React.ChangeEvent<HTMLInputElement>) {
    setOtherNames(event.target.value)
  }

  function updateDescription(event: React.ChangeEvent<HTMLTextAreaElement>) {
    setDescription(event.target.value)
  }

  function updatePlantSeason(event: React.ChangeEvent<HTMLInputElement>) {
    setPlantSeason(event.target.value)
  }

  function updateHarvestSeason(event: React.ChangeEvent<HTMLInputElement>) {
    setHarvestSeason(event.target.value)
  }

  function updatePruneSeason(event: React.ChangeEvent<HTMLInputElement>) {
    setPruneSeason(event.target.value)
  }

  function updateTips(event: React.ChangeEvent<HTMLTextAreaElement>) {
    setTips(event.target.value)
  }

  async function submitAddPlant(
    event: React.FormEvent<HTMLFormElement>,
  ): Promise<void> {
    event.preventDefault()

    const newPlant = {
      name: name,
      other_names: otherNames,
      description: description,
      plant_season: plantSeason,
      harvest_season: harvestSeason,
      prune_season: pruneSeason,
      tips: tips,
    }

    await addPlant({variables: newPlant})

    history.push("/")
  }

  return (
    <>
      <section>
        <div>
          <h1 className="page-title">{"Add Plant"}</h1>
        </div>
      </section>
      <section>
        <form onSubmit={submitAddPlant}>
          <div>
            <label className="label">
              {"Name"}{" "}
              <span className="text-gray-500 text-xs">{"(Required)"}</span>
            </label>
            <input
              className="input"
              onChange={updateName}
              required
              type="text"
              value={name}
            />
          </div>
          <div>
            <label className="label">{"Other Names"}</label>
            <input
              className="input"
              onChange={updateOtherNames}
              type="text"
              value={otherNames}
            />
          </div>
          <div>
            <label className="label">{"Description"}</label>
            <textarea
              className="input"
              onChange={updateDescription}
              rows={4}
              value={description}
            />
          </div>
          <div>
            <label className="label">{"Plant Season"}</label>
            <input
              className="input"
              onChange={updatePlantSeason}
              type="text"
              value={plantSeason}
            />
          </div>
          <div>
            <label className="label">{"Harvest Season"}</label>
            <input
              className="input"
              onChange={updateHarvestSeason}
              type="text"
              value={harvestSeason}
            />
          </div>
          <div>
            <label className="label">{"Prune Season"}</label>
            <input
              className="input"
              onChange={updatePruneSeason}
              type="text"
              value={pruneSeason}
            />
          </div>
          <div>
            <label className="label">{"Tips"}</label>
            <textarea
              className="input"
              onChange={updateTips}
              rows={4}
              value={tips}
            />
          </div>
          <div className="flex justify-center pt-6">
            <button className="button button-primary" type="submit">
              {"Add plant"}
            </button>
          </div>
        </form>
      </section>
    </>
  )
}
