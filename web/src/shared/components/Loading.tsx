import React from "react"
import {CircularProgress} from "@material-ui/core"

export function Loading(): JSX.Element {
  return (
    <div
      style={{
        alignItems: "center",
        display: "flex",
        height: "100%",
        justifyContent: "center",
        paddingTop: "30%",
        width: "100%",
      }}
    >
      <CircularProgress />
    </div>
  )
}
