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
  const nameElement = React.useRef<HTMLInputElement>(null)
  const otherNamesElement = React.useRef<HTMLInputElement>(null)
  const descriptionElement = React.useRef<HTMLTextAreaElement>(null)
  const plantSeasonElement = React.useRef<HTMLInputElement>(null)
  const harvestSeasonElement = React.useRef<HTMLInputElement>(null)
  const pruneSeasonElement = React.useRef<HTMLInputElement>(null)
  const tipsElement = React.useRef<HTMLTextAreaElement>(null)

  async function submitAddPlant(
    event: React.FormEvent<HTMLFormElement>,
  ): Promise<void> {
    event.preventDefault()

    const newPlant = {
      name: nameElement.current?.value || "",
      other_names: otherNamesElement.current?.value || null,
      description: descriptionElement.current?.value || null,
      plant_season: plantSeasonElement.current?.value || null,
      harvest_season: harvestSeasonElement.current?.value || null,
      prune_season: pruneSeasonElement.current?.value || null,
      tips: tipsElement.current?.value || null,
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
            <input className="input" type="text" ref={nameElement} required />
          </div>
          <div>
            <label className="label">{"Other Names"}</label>
            <input className="input" type="text" ref={otherNamesElement} />
          </div>
          <div>
            <label className="label">{"Description"}</label>
            <textarea className="input" rows={4} ref={descriptionElement} />
          </div>
          <div>
            <label className="label">{"Plant Season"}</label>
            <input className="input" type="text" ref={plantSeasonElement} />
          </div>
          <div>
            <label className="label">{"Harvest Season"}</label>
            <input className="input" type="text" ref={harvestSeasonElement} />
          </div>
          <div>
            <label className="label">{"Prune Season"}</label>
            <input className="input" type="text" ref={pruneSeasonElement} />
          </div>
          <div>
            <label className="label">{"Tips"}</label>
            <textarea className="input" rows={4} ref={tipsElement} />
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
