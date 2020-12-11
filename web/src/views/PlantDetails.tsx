import React, {useEffect, useState} from "react"
import {useHistory, useParams} from "react-router-dom"
import {Button, CircularProgress, Typography} from "@material-ui/core"
import {Alert} from "../components"
import {deleteOne, listOne} from "../services/plant"
import {Plant} from "../graphql"
import {FetchStatus} from "../constants"

export function PlantDetails(): JSX.Element {
  const [data, setData] = useState({} as Plant)
  const [fetchStatus, setFetchStatus] = useState(FetchStatus.Idle)
  const [alertOpen, setAlertOpen] = useState(false)
  const {id} = useParams<{id: string}>()
  const history = useHistory()

  useEffect(() => {
    setFetchStatus(FetchStatus.Loading)
    ;(async (): Promise<void> => {
      try {
        const result = await listOne(id)

        setData(result.data as Plant)
        setFetchStatus(FetchStatus.Success)
      } catch (error) {
        setFetchStatus(FetchStatus.Error)

        console.error(error)
      }
    })()
  }, [id])

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
    try {
      await deleteOne(id)

      history.push("/")
    } catch (error) {
      console.error(error)
    }
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
          <Typography variant="h1">{data.plant?.name}</Typography>
          <section style={{marginTop: "30px"}}>
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
            onClick={editPlant}
            style={{marginTop: "30px"}}
            variant="contained"
          >
            {"Edit plant"}
          </Button>
          <Button
            color="secondary"
            fullWidth
            onClick={openAlert}
            style={{marginTop: "30px"}}
            variant="contained"
          >
            {"Delete plant"}
          </Button>
          <Alert
            action={deletePlant}
            cancel={closeAlert}
            message={data.plant?.name + " will be deleted."}
            open={alertOpen}
            title={"Delete plant"}
          />
        </>
      )}
    </>
  )
}
