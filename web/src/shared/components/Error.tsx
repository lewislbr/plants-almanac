import React from "react"
import {Typography} from "@material-ui/core"

export function Error({message}: {message: string}): JSX.Element {
  return (
    <Typography
      style={{
        alignItems: "center",
        display: "flex",
        height: "100%",
        justifyContent: "center",
        paddingTop: "30%",
        width: "100%",
      }}
    >
      {message}
    </Typography>
  )
}
