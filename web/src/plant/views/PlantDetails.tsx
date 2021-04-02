import React, {useEffect, useState} from "react"
import {useHistory, useParams} from "react-router-dom"
import {Button, IconButton, Typography} from "@material-ui/core"
import CancelIcon from "@material-ui/icons/Cancel"
import {Dialog, Error, Loading} from "../../shared/components"
import {deleteOne, listOne} from "../services/plant"
import {DELETE_PLANT, EDIT_PLANT, NO_DATA} from "../constants/copy"
import {CANCEL} from "../../shared/constants/copy"
import {HTTPStatus} from "../../shared/constants/http"
import {Plant} from "../interfaces/plant"

export function PlantDetails(): JSX.Element {
  const [errors, setErrors] = useState({
    http: "",
  })
  const [data, setData] = useState({} as Plant)
  const [status, setStatus] = useState(HTTPStatus.IDLE)
  const [dialogOpen, setDialogOpen] = useState(false)
  const {id} = useParams<{id: Plant["id"]}>()
  const history = useHistory()

  useEffect(() => {
    setStatus(HTTPStatus.LOADING)
    ;(async (): Promise<void> => {
      try {
        const result = await listOne(id)

        setData(result)
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
    history.push({pathname: "/edit/" + id, state: data})
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
      await deleteOne(id)

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
            <Typography variant="h1">{data.name}</Typography>
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
                {data.other_names || NO_DATA}
              </Typography>
            </div>
            <div style={{marginBottom: "30px"}}>
              <Typography gutterBottom variant="h6">
                {"Description"}
              </Typography>
              <Typography gutterBottom variant="body1">
                {data.description || NO_DATA}
              </Typography>
            </div>
            <div style={{marginBottom: "30px"}}>
              <Typography gutterBottom variant="h6">
                {"Plant Season"}
              </Typography>
              <Typography gutterBottom variant="body1">
                {data.plant_season || NO_DATA}
              </Typography>
            </div>
            <div style={{marginBottom: "30px"}}>
              <Typography gutterBottom variant="h6">
                {"Harvest Season"}
              </Typography>
              <Typography gutterBottom variant="body1">
                {data.harvest_season || NO_DATA}
              </Typography>
            </div>
            <div style={{marginBottom: "30px"}}>
              <Typography gutterBottom variant="h6">
                {"Prune Season"}
              </Typography>
              <Typography gutterBottom variant="body1">
                {data.prune_season || NO_DATA}
              </Typography>
            </div>
            <div style={{marginBottom: "30px"}}>
              <Typography gutterBottom variant="h6">
                {"Tips"}
              </Typography>
              <Typography gutterBottom variant="body1">
                {data.tips || NO_DATA}
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
            {EDIT_PLANT}
          </Button>
          <Button
            color="secondary"
            fullWidth
            onClick={openDialog}
            style={{marginTop: "30px"}}
            variant="contained"
          >
            {DELETE_PLANT}
          </Button>
          {dialogOpen && (
            <Dialog
              action={deletePlant}
              actionText={DELETE_PLANT}
              cancel={closeDialog}
              cancelText={CANCEL}
              message={data.name + " will be deleted."}
              open={dialogOpen}
              title={DELETE_PLANT}
            />
          )}
          {status === HTTPStatus.ERROR && <Error message={errors.http} title={"Error"} />}
        </>
      )}
    </>
  )
}
