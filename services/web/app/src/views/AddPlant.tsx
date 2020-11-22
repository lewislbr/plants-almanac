import React, {ChangeEvent, useState} from "react"
import {addOne} from "../services"

export function AddPlant(): JSX.Element {
  const [name, setName] = useState("")
  const [otherNames, setOtherNames] = useState("")
  const [description, setDescription] = useState("")
  const [plantSeason, setPlantSeason] = useState("")
  const [harvestSeason, setHarvestSeason] = useState("")
  const [pruneSeason, setPruneSeason] = useState("")
  const [tips, setTips] = useState("")
  const plantState = {
    name,
    otherNames,
    description,
    plantSeason,
    harvestSeason,
    pruneSeason,
    tips,
  }

  function updateName(event: ChangeEvent<HTMLInputElement>): void {
    setName(event.target.value)
  }

  function updateOtherNames(event: ChangeEvent<HTMLInputElement>): void {
    setOtherNames(event.target.value)
  }

  function updateDescription(event: ChangeEvent<HTMLTextAreaElement>): void {
    setDescription(event.target.value)
  }

  function updatePlantSeason(event: ChangeEvent<HTMLInputElement>): void {
    setPlantSeason(event.target.value)
  }

  function updateHarvestSeason(event: ChangeEvent<HTMLInputElement>): void {
    setHarvestSeason(event.target.value)
  }

  function updatePruneSeason(event: ChangeEvent<HTMLInputElement>): void {
    setPruneSeason(event.target.value)
  }

  function updateTips(event: ChangeEvent<HTMLTextAreaElement>): void {
    setTips(event.target.value)
  }

  return (
    <>
      <section>
        <div>
          <h1 className="page-title">{"Add Plant"}</h1>
        </div>
      </section>
      <section>
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
          <button
            className="button button-primary"
            onClick={(): Promise<void> => addOne(plantState)}
          >
            {"Add plant"}
          </button>
        </div>
      </section>
    </>
  )
}
