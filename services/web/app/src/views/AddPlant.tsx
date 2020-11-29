import React, {ChangeEvent, useState} from "react"
import {useHistory} from "react-router-dom"
import {Button, TextField, Typography} from "@material-ui/core"
import {addOne} from "../services/plant"

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
  const history = useHistory()

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

  async function addPlant(): Promise<void> {
    try {
      await addOne(plantState)

      history.push("/")
    } catch (error) {
      console.error(error)
    }
  }

  return (
    <>
      <Typography gutterBottom variant="h1">
        {"Add Plant"}
      </Typography>
      <TextField
        label="Name"
        fullWidth
        margin={"dense"}
        onChange={updateName}
        required
        value={name}
        variant="outlined"
      />
      <TextField
        label="Other names"
        fullWidth
        margin={"dense"}
        onChange={updateOtherNames}
        value={otherNames}
        variant="outlined"
      />
      <TextField
        label="Description"
        fullWidth
        margin={"dense"}
        multiline
        onChange={updateDescription}
        rows={4}
        value={description}
        variant="outlined"
      />
      <TextField
        label="Plant season"
        fullWidth
        margin={"dense"}
        onChange={updatePlantSeason}
        value={plantSeason}
        variant="outlined"
      />
      <TextField
        label="Harvest season"
        fullWidth
        margin={"dense"}
        onChange={updateHarvestSeason}
        value={harvestSeason}
        variant="outlined"
      />
      <TextField
        label="Prune season"
        fullWidth
        margin={"dense"}
        onChange={updatePruneSeason}
        value={pruneSeason}
        variant="outlined"
      />
      <TextField
        label="Tips"
        fullWidth
        margin={"dense"}
        multiline
        onChange={updateTips}
        rows={4}
        value={tips}
        variant="outlined"
      />
      <Button
        color="primary"
        fullWidth
        onClick={addPlant}
        style={{marginTop: "30px"}}
        variant="contained"
      >
        {"Add plant"}
      </Button>
    </>
  )
}
