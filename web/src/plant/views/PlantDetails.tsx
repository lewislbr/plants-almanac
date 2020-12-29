import React, {useEffect, useState} from "react"
import {useHistory, useParams} from "react-router-dom"
import {Button, IconButton, Typography} from "@material-ui/core"
import CancelIcon from "@material-ui/icons/Cancel"
import {Alert, Error, Loading} from "../../shared/components"
import * as plantService from "../services/plant"
import * as plantCopy from "../constants/copy"
import * as sharedCopy from "../../shared/constants/copy"
import * as errorConstant from "../../shared/constants/error"
import * as fetchConstant from "../../shared/constants/fetch"
import {Plant} from "../interfaces/Plant"

export function PlantDetails(): JSX.Element {
  const [data, setData] = useState({} as Plant)
  const [fetchStatus, setFetchStatus] = useState(fetchConstant.Status.IDLE)
  const [alertOpen, setAlertOpen] = useState(false)
  const {id} = useParams<{id: string}>()
  const history = useHistory()

  useEffect(() => {
    setFetchStatus(fetchConstant.Status.LOADING)
    ;(async (): Promise<void> => {
      try {
        const result = await plantService.listOne(id)

        setData(result.data as Plant)
        setFetchStatus(fetchConstant.Status.SUCCESS)
      } catch (error) {
        setFetchStatus(fetchConstant.Status.ERROR)

        console.error(error)
      }
    })()
  }, [id])

  function close(): void {
    history.push("/plants")
  }

  function editPlant(): void {
    history.push({pathname: "/edit/" + id, state: data.plant})
  }

  function openAlert(): void {
    setAlertOpen(true)
  }

  function closeAlert(): void {
    setAlertOpen(false)
  }

  async function deletePlant(): Promise<void> {
    setFetchStatus(fetchConstant.Status.LOADING)

    try {
      await plantService.deleteOne(id)

      setFetchStatus(fetchConstant.Status.SUCCESS)

      history.push("/plants")
    } catch (error) {
      setFetchStatus(fetchConstant.Status.ERROR)

      console.error(error)
    }
  }

  return (
    <>
      {fetchStatus === fetchConstant.Status.LOADING ? (
        <Loading />
      ) : fetchStatus === fetchConstant.Status.ERROR ? (
        <Error message={errorConstant.GENERIC_MESSAGE} />
      ) : (
        <>
          <div
            style={{
              alignItems: "center",
              display: "flex",
              justifyContent: "space-between",
            }}
          >
            <Typography variant="h1">{data.plant?.name}</Typography>
            <IconButton onClick={close}>
              <CancelIcon />
            </IconButton>
          </div>
          <section style={{marginTop: "30px"}}>
            <div style={{marginBottom: "30px"}}>
              <Typography gutterBottom variant="h6">
                {"Other Names"}
              </Typography>
              <Typography gutterBottom variant="body1">
                {data.plant?.other_names || plantCopy.NO_DATA}
              </Typography>
            </div>
            <div style={{marginBottom: "30px"}}>
              <Typography gutterBottom variant="h6">
                {"Description"}
              </Typography>
              <Typography gutterBottom variant="body1">
                {data.plant?.description || plantCopy.NO_DATA}
              </Typography>
            </div>
            <div style={{marginBottom: "30px"}}>
              <Typography gutterBottom variant="h6">
                {"Plant Season"}
              </Typography>
              <Typography gutterBottom variant="body1">
                {data.plant?.plant_season || plantCopy.NO_DATA}
              </Typography>
            </div>
            <div style={{marginBottom: "30px"}}>
              <Typography gutterBottom variant="h6">
                {"Harvest Season"}
              </Typography>
              <Typography gutterBottom variant="body1">
                {data.plant?.harvest_season || plantCopy.NO_DATA}
              </Typography>
            </div>
            <div style={{marginBottom: "30px"}}>
              <Typography gutterBottom variant="h6">
                {"Prune Season"}
              </Typography>
              <Typography gutterBottom variant="body1">
                {data.plant?.prune_season || plantCopy.NO_DATA}
              </Typography>
            </div>
            <div style={{marginBottom: "30px"}}>
              <Typography gutterBottom variant="h6">
                {"Tips"}
              </Typography>
              <Typography gutterBottom variant="body1">
                {data.plant?.tips || plantCopy.NO_DATA}
              </Typography>
            </div>
          </section>
          <Button
            color="secondary"
            fullWidth
            onClick={editPlant}
            style={{marginTop: "30px"}}
            variant="contained"
          >
            {plantCopy.EDIT_PLANT}
          </Button>
          <Button
            color="secondary"
            fullWidth
            onClick={openAlert}
            style={{marginTop: "30px"}}
            variant="contained"
          >
            {plantCopy.DELETE_PLANT}
          </Button>
          <Alert
            action={deletePlant}
            actionText={plantCopy.DELETE_PLANT}
            cancel={closeAlert}
            cancelText={sharedCopy.CANCEL}
            message={data.plant?.name + " will be deleted."}
            open={alertOpen}
            title={plantCopy.DELETE_PLANT}
          />
        </>
      )}
    </>
  )
}
