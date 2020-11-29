import React from "react"
import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
} from "@material-ui/core"

export function Alert({
  action,
  cancel,
  message,
  open,
  title,
}: {
  action: () => void
  cancel: () => void
  message: string
  open: boolean
  title: string
}): JSX.Element {
  return (
    <Dialog fullWidth maxWidth="xl" onClose={cancel} open={open}>
      <DialogTitle>{title}</DialogTitle>
      <DialogContent>
        <DialogContentText>{message}</DialogContentText>
      </DialogContent>
      <DialogActions>
        <Button color="primary" onClick={cancel}>
          {"Cancel"}
        </Button>
        <Button autoFocus color="primary" onClick={action} variant="contained">
          {"Delete"}
        </Button>
      </DialogActions>
    </Dialog>
  )
}
