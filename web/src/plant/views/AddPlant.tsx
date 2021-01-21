import React, {ChangeEvent, useEffect, useLayoutEffect, useState} from "react"
import {useHistory, useLocation} from "react-router-dom"
import {Button, TextField, Typography} from "@material-ui/core"
import {Error, Loading} from "../../shared/components"
import * as plantService from "../services/plant"
import * as plantCopy from "../constants/copy"
import * as sharedCopy from "../../shared/constants/copy"
import {HTTPStatus} from "../../shared/constants/http"
import {EditVariables} from "../interfaces/Edit"

export function AddPlant(): JSX.Element {
  const [errors, setErrors] = useState({
    name: false,
    http: "",
  })
  const [buttonDisabled, setButtonDisabled] = useState(true)
  const [status, setStatus] = useState(HTTPStatus.IDLE)
  const [name, setName] = useState("")
  const [otherNames, setOtherNames] = useState("")
  const [description, setDescription] = useState("")
  const [plantSeason, setPlantSeason] = useState("")
  const [harvestSeason, setHarvestSeason] = useState("")
  const [pruneSeason, setPruneSeason] = useState("")
  const [tips, setTips] = useState("")
  const history = useHistory()
  const location = useLocation()
  const prevState = location.state as EditVariables
  const [isEditMode] = useState(Boolean(prevState))
  const missingFields = !name
  const activeErrors = Object.values(errors).includes(true)
  const noChanges =
    isEditMode &&
    name === prevState.name &&
    otherNames === prevState.other_names &&
    description === prevState.description &&
    plantSeason === prevState.plant_season &&
    harvestSeason === prevState.harvest_season &&
    pruneSeason === prevState.prune_season &&
    tips === prevState.tips

  useLayoutEffect(() => {
    window.scrollTo(0, 0)
  }, [location])

  useEffect(() => {
    if (missingFields || activeErrors || noChanges) {
      setButtonDisabled(true)
    } else {
      setButtonDisabled(false)
    }
  }, [missingFields, activeErrors, noChanges])

  useEffect(() => {
    if (isEditMode) {
      setStatus(HTTPStatus.LOADING)
      setName(prevState.name)
      setOtherNames(prevState.other_names || "")
      setDescription(prevState.description || "")
      setPlantSeason(prevState.plant_season || "")
      setHarvestSeason(prevState.harvest_season || "")
      setPruneSeason(prevState.prune_season || "")
      setTips(prevState.tips || "")
      setStatus(HTTPStatus.SUCCESS)
    }
  }, [isEditMode, prevState])

  function updateName(event: ChangeEvent<HTMLInputElement>): void {
    if (!event.target.value) {
      setErrors((errors) => ({...errors, name: true}))
    } else {
      setErrors((errors) => ({...errors, name: false}))
    }

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
    setStatus(HTTPStatus.LOADING)

    try {
      await plantService.addOne({
        name,
        otherNames,
        description,
        plantSeason,
        harvestSeason,
        pruneSeason,
        tips,
      })

      setStatus(HTTPStatus.SUCCESS)

      history.push("/plants")
    } catch (error) {
      setErrors((errors) => ({...errors, http: String(error)}))
      setStatus(HTTPStatus.ERROR)

      console.error(error)
    }
  }

  async function editPlant(): Promise<void> {
    setStatus(HTTPStatus.LOADING)

    try {
      await plantService.editOne(prevState.id, {
        name,
        otherNames,
        description,
        plantSeason,
        harvestSeason,
        pruneSeason,
        tips,
      })

      setStatus(HTTPStatus.SUCCESS)

      history.push("/plants/" + prevState.id)
    } catch (error) {
      setErrors((errors) => ({...errors, http: String(error)}))
      setStatus(HTTPStatus.ERROR)

      console.error(error)
    }
  }

  function cancel(): void {
    history.push("/plants")
  }

  return (
    <>
      {status === HTTPStatus.ERROR && (
        <Error message={errors.http} title={"Error"} />
      )}
      <Typography variant="h1">
        {isEditMode ? plantCopy.EDIT_PLANT : plantCopy.ADD_PLANT}
      </Typography>
      {status === HTTPStatus.LOADING ? (
        <Loading />
      ) : (
        <>
          <section style={{marginTop: "30px"}}>
            <TextField
              error={errors.name}
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
              disabled={buttonDisabled}
              fullWidth
              onClick={editPlant}
              style={{marginTop: "30px"}}
              variant="contained"
            >
              {plantCopy.SAVE_CHANGES}
            </Button>
          ) : (
            <Button
              color="primary"
              disabled={buttonDisabled}
              fullWidth
              onClick={addPlant}
              style={{marginTop: "30px"}}
              variant="contained"
            >
              {plantCopy.ADD_PLANT}
            </Button>
          )}
          <Button
            color="secondary"
            fullWidth
            onClick={cancel}
            style={{marginTop: "30px"}}
            variant="contained"
          >
            {sharedCopy.CANCEL}
          </Button>
        </>
      )}
    </>
  )
}
