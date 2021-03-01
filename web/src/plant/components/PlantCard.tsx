import React from "react"
import {Paper, Typography} from "@material-ui/core"

export function PlantCard({name}: {name: string | null | undefined}): JSX.Element {
  return (
    <Paper
      elevation={3}
      style={{
        boxShadow: "0px 2px 10px 1px #100d0d0d",
        display: "flex",
        justifyContent: "flex-start",
        height: "100px",
        marginBottom: "20px",
        padding: "12px",
      }}
    >
      <div
        style={{
          backgroundColor: "lightgray",
          borderRadius: 30,
          height: "100%",
          width: "76px",
        }}
      ></div>
      <Typography style={{marginLeft: "10px"}} variant="h5">
        {name}
      </Typography>
    </Paper>
  )
}
