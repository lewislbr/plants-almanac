import React, {useEffect, useState} from "react"
import {useHistory, useParams} from "react-router-dom"
import {Button, IconButton, Typography} from "@material-ui/core"
import CancelIcon from "@material-ui/icons/Cancel"
import {Dialog, Error, Loading} from "../../shared/components"
import * as plantService from "../services/plant"
import * as plantCopy from "../constants/copy"
import * as sharedCopy from "../../shared/constants/copy"
import {HTTPStatus} from "../../shared/constants/http"
import {Plant} from "../interfaces/Plant"

export function PlantDetails(): JSX.Element {
  const [errors, setErrors] = useState({
    http: "",
  })
  const [data, setData] = useState({} as Plant)
  const [status, setStatus] = useState(HTTPStatus.IDLE)
  const [dialogOpen, setDialogOpen] = useState(false)
  const {id} = useParams<{id: string}>()
  const history = useHistory()

  useEffect(() => {
    setStatus(HTTPStatus.LOADING)
    ;(async (): Promise<void> => {
      try {
        const result = await plantService.listOne(id)

        setData(result.data as Plant)
        setStatus(HTTPStatus.SUCCESS)
      } catch (error) {
        setErrors((errors) => ({...errors, http: String(error)}))
        setStatus(HTTPStatus.ERROR)

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

  function openDialog(): void {
    setDialogOpen(true)
  }

  function closeDialog(): void {
    setDialogOpen(false)
  }

  async function deletePlant(): Promise<void> {
    setStatus(HTTPStatus.LOADING)

    try {
      await plantService.deleteOne(id)

      setStatus(HTTPStatus.SUCCESS)

      history.push("/plants")
    } catch (error) {
      setErrors((errors) => ({...errors, http: String(error)}))
      setStatus(HTTPStatus.ERROR)

      console.error(error)
    }
  }

  return (
    <>
      {status === HTTPStatus.LOADING ? (
        <Loading />
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
            onClick={openDialog}
            style={{marginTop: "30px"}}
            variant="contained"
          >
            {plantCopy.DELETE_PLANT}
          </Button>
          {dialogOpen && (
            <Dialog
              action={deletePlant}
              actionText={plantCopy.DELETE_PLANT}
              cancel={closeDialog}
              cancelText={sharedCopy.CANCEL}
              message={data.plant?.name + " will be deleted."}
              open={dialogOpen}
              title={plantCopy.DELETE_PLANT}
            />
          )}
          {status === HTTPStatus.ERROR && (
            <Error message={errors.http} title={"Error"} />
          )}
        </>
      )}
    </>
  )
}
