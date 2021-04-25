import React from "react"
import {Link} from "react-router-dom"
import {BottomNavigation, BottomNavigationAction} from "@material-ui/core"

export function NavBar(): JSX.Element {
  return (
    <BottomNavigation
      showLabels
      style={{
        bottom: 0,
        left: 0,
        position: "fixed",
        right: 0,
        width: "100%",
        zIndex: 10,
      }}
    >
      <BottomNavigationAction component={Link} icon={"🌱"} label="Plants" to="/" />
      <BottomNavigationAction component={Link} icon={"➕"} label="Add plant" to="/add-plant" />
      <BottomNavigationAction component={Link} icon={"👤"} label="Account" to="/account" />
    </BottomNavigation>
  )
}
