import React from "react"
import {Paper, Typography} from "@material-ui/core"

export function PlantCard({
  name,
}: {
  name: string | null | undefined
}): JSX.Element {
  return (
    <Paper
      elevation={3}
      style={{
        marginBottom: "20px",
        padding: "12px 20px",
      }}
    >
      <Typography variant="h5">{name}</Typography>
    </Paper>
  )
}
