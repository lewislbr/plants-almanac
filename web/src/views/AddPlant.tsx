import React, {ChangeEvent, useEffect, useLayoutEffect, useState} from "react"
import {useHistory, useLocation} from "react-router-dom"
import {Button, TextField, Typography} from "@material-ui/core"
import {Error, Loading} from "../components"
import * as plantService from "../services/plant"
import * as copyConstant from "../constants/copy"
import * as errorConstant from "../constants/error"
import * as fetchConstant from "../constants/fetch"
import {EditVariables} from "../graphql/Edit"

export function AddPlant(): JSX.Element {
  const location = useLocation()
  const prevState = location.state as EditVariables
  const [isEditMode] = useState(prevState)
  const [fetchStatus, setFetchStatus] = useState(fetchConstant.Status.IDLE)
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
      setFetchStatus(fetchConstant.Status.LOADING)
      setName(prevState.name)
      setOtherNames(prevState.other_names || "")
      setDescription(prevState.description || "")
      setPlantSeason(prevState.plant_season || "")
      setHarvestSeason(prevState.harvest_season || "")
      setPruneSeason(prevState.prune_season || "")
      setTips(prevState.tips || "")
      setFetchStatus(fetchConstant.Status.SUCCESS)
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
    setFetchStatus(fetchConstant.Status.LOADING)

    try {
      await plantService.addOne(plantState)

      setFetchStatus(fetchConstant.Status.SUCCESS)

      history.push("/plants")
    } catch (error) {
      setFetchStatus(fetchConstant.Status.ERROR)

      console.error(error)
    }
  }

  async function editPlant(): Promise<void> {
    setFetchStatus(fetchConstant.Status.LOADING)

    try {
      await plantService.editOne(prevState.id, plantState)

      setFetchStatus(fetchConstant.Status.SUCCESS)

      history.push("/plants/" + prevState.id)
    } catch (error) {
      setFetchStatus(fetchConstant.Status.ERROR)

      console.error(error)
    }
  }

  function cancel(): void {
    history.push("/plants")
  }

  return (
    <>
      {fetchStatus === fetchConstant.Status.LOADING ? (
        <Loading />
      ) : fetchStatus === fetchConstant.Status.ERROR ? (
        <Error message={errorConstant.GENERIC_MESSAGE} />
      ) : (
        <>
          <Typography variant="h1">
            {isEditMode ? copyConstant.EDIT_PLANT : copyConstant.ADD_PLANT}
          </Typography>
          <section style={{marginTop: "30px"}}>
            <TextField
              fullWidth
              label="Name"
              onChange={updateName}
              required
              value={name}
              variant="outlined"
            />
            <TextField
              fullWidth
              label="Other names"
              onChange={updateOtherNames}
              value={otherNames}
              variant="outlined"
            />
            <TextField
              fullWidth
              label="Description"
              multiline
              onChange={updateDescription}
              rows={6}
              value={description}
              variant="outlined"
            />
            <TextField
              fullWidth
              label="Plant season"
              onChange={updatePlantSeason}
              value={plantSeason}
              variant="outlined"
            />
            <TextField
              fullWidth
              label="Harvest season"
              onChange={updateHarvestSeason}
              value={harvestSeason}
              variant="outlined"
            />
            <TextField
              fullWidth
              label="Prune season"
              onChange={updatePruneSeason}
              value={pruneSeason}
              variant="outlined"
            />
            <TextField
              fullWidth
              label="Tips"
              multiline
              onChange={updateTips}
              rows={6}
              value={tips}
              variant="outlined"
            />
          </section>
          {isEditMode ? (
            <Button
              color="primary"
              fullWidth
              onClick={editPlant}
              style={{marginTop: "30px"}}
              variant="contained"
            >
              {copyConstant.SAVE_CHANGES}
            </Button>
          ) : (
            <Button
              color="primary"
              fullWidth
              onClick={addPlant}
              style={{marginTop: "30px"}}
              variant="contained"
            >
              {copyConstant.ADD_PLANT}
            </Button>
          )}
          <Button
            color="secondary"
            fullWidth
            onClick={cancel}
            style={{marginTop: "30px"}}
            variant="contained"
          >
            {copyConstant.CANCEL}
          </Button>
        </>
      )}
    </>
  )
}
