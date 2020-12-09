import React, {ChangeEvent, useEffect, useLayoutEffect, useState} from "react"
import {useHistory, useLocation} from "react-router-dom"
import {
  Button,
  CircularProgress,
  TextField,
  Typography,
} from "@material-ui/core"
import {addOne, editOne} from "../services/plant"
import {EditVariables} from "../graphql"
import {FetchStatus} from "../constants"

export function AddPlant(): JSX.Element {
  const location = useLocation()
  const prevState = location.state as EditVariables
  const [isEditMode] = useState(prevState)
  const [fetchStatus, setFetchStatus] = useState(FetchStatus.Idle)
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

  useLayoutEffect(() => {
    window.scrollTo(0, 0)
  }, [location])

  useEffect(() => {
    if (isEditMode) {
      setFetchStatus(FetchStatus.Loading)
      setName(prevState.name)
      setOtherNames(prevState.other_names || "")
      setDescription(prevState.description || "")
      setPlantSeason(prevState.plant_season || "")
      setHarvestSeason(prevState.harvest_season || "")
      setPruneSeason(prevState.prune_season || "")
      setTips(prevState.tips || "")
      setFetchStatus(FetchStatus.Success)
    }
  }, [isEditMode, prevState])

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

  async function editPlant(): Promise<void> {
    try {
      await editOne(prevState.id, plantState)

      history.push("/" + prevState.id)
    } catch (error) {
      console.error(error)
    }
  }

  function cancel(): void {
    history.push("/")
  }

  return (
    <>
      {fetchStatus === FetchStatus.Loading ? (
        <div
          style={{
            display: "flex",
            justifyContent: "center",
            marginTop: "100px",
          }}
        >
          <CircularProgress />
        </div>
      ) : fetchStatus === FetchStatus.Error ? (
        <Typography>{"ERROR"}</Typography>
      ) : (
        <>
          <Typography gutterBottom variant="h1">
            {(isEditMode ? "Edit" : "Add") + " Plant"}
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
          {isEditMode ? (
            <Button
              color="primary"
              fullWidth
              onClick={editPlant}
              style={{marginTop: "30px"}}
              variant="contained"
            >
              {"Save edits"}
            </Button>
          ) : (
            <Button
              color="primary"
              fullWidth
              onClick={addPlant}
              style={{marginTop: "30px"}}
              variant="contained"
            >
              {"Add plant"}
            </Button>
          )}
          <Button
            color="secondary"
            fullWidth
            onClick={cancel}
            style={{marginTop: "30px"}}
            variant="contained"
          >
            {"Cancel"}
          </Button>
        </>
      )}
    </>
  )
}
