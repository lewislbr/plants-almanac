import React, {useEffect, useState} from "react"
import {useParams} from "react-router-dom"
import {
  Button,
  CircularProgress,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
  Typography,
} from "@material-ui/core"
import {deleteOne, listOne} from "../services/plant"
import {Plant} from "../graphql"
import {DataStatus} from "../constants"

export function PlantDetails(): JSX.Element {
  const [data, setData] = useState({} as Plant)
  const [dataStatus, setDataStatus] = useState(DataStatus.Idle)
  const [alertOpen, setAlertOpen] = useState(false)
  const {id} = useParams<{id: string}>()

  useEffect(() => {
    setDataStatus(DataStatus.Loading)
    ;(async (): Promise<void> => {
      try {
        const result = await listOne(id)

        setData(result.data as Plant)
        setDataStatus(DataStatus.Success)
      } catch (error) {
        setDataStatus(DataStatus.Error)

        console.error(error)
      }
    })()
  }, [id])

  function openAlert(): void {
    setAlertOpen(true)
  }

  function closeAlert(): void {
    setAlertOpen(false)
  }

  async function deletePlant(): Promise<void> {
    await deleteOne(id)
  }

  return (
    <>
      {dataStatus === DataStatus.Loading ? (
        <div
          style={{
            display: "flex",
            justifyContent: "center",
            marginTop: "100px",
          }}
        >
          <CircularProgress />
        </div>
      ) : dataStatus === DataStatus.Error ? (
        <Typography>{"ERROR"}</Typography>
      ) : (
        <>
          <Typography gutterBottom variant="h1">
            {data.plant?.name}
          </Typography>
          <section>
            <div style={{marginBottom: "30px"}}>
              <Typography gutterBottom variant="h6">
                {"Other Names"}
              </Typography>
              <Typography gutterBottom variant="body1">
                {data.plant?.other_names || "No data yet"}
              </Typography>
            </div>
            <div style={{marginBottom: "30px"}}>
              <Typography gutterBottom variant="h6">
                {"Description"}
              </Typography>
              <Typography gutterBottom variant="body1">
                {data.plant?.description || "No data yet"}
              </Typography>
            </div>
            <div style={{marginBottom: "30px"}}>
              <Typography gutterBottom variant="h6">
                {"Plant Season"}
              </Typography>
              <Typography gutterBottom variant="body1">
                {data.plant?.plant_season || "No data yet"}
              </Typography>
            </div>
            <div style={{marginBottom: "30px"}}>
              <Typography gutterBottom variant="h6">
                {"Harvest Season"}
              </Typography>
              <Typography gutterBottom variant="body1">
                {data.plant?.harvest_season || "No data yet"}
              </Typography>
            </div>
            <div style={{marginBottom: "30px"}}>
              <Typography gutterBottom variant="h6">
                {"Prune Season"}
              </Typography>
              <Typography gutterBottom variant="body1">
                {data.plant?.prune_season || "No data yet"}
              </Typography>
            </div>
            <div style={{marginBottom: "30px"}}>
              <Typography gutterBottom variant="h6">
                {"Tips"}
              </Typography>
              <Typography gutterBottom variant="body1">
                {data.plant?.tips || "No data yet"}
              </Typography>
            </div>
          </section>
          <Button
            color="secondary"
            fullWidth
            onClick={openAlert}
            style={{marginTop: "30px"}}
            variant="contained"
          >
            {"Delete plant"}
          </Button>
          <Dialog onClose={closeAlert} open={alertOpen}>
            <DialogTitle>{"Delete plant"}</DialogTitle>
            <DialogContent>
              <DialogContentText>
                {"The plant will be deleted."}
              </DialogContentText>
            </DialogContent>
            <DialogActions>
              <Button color="primary" onClick={closeAlert}>
                {"Cancel"}
              </Button>
              <Button
                autoFocus
                color="primary"
                onClick={deletePlant}
                variant="contained"
              >
                {"Delete"}
              </Button>
            </DialogActions>
          </Dialog>
        </>
      )}
    </>
  )
}
