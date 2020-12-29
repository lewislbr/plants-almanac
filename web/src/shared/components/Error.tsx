import React, {useState} from "react"
import {
  Dialog,
  DialogContent,
  DialogContentText,
  DialogTitle,
  IconButton,
} from "@material-ui/core"
import CancelIcon from "@material-ui/icons/Cancel"
import {makeStyles} from "@material-ui/core/styles"
import {GENERIC_MESSAGE} from "../constants/error"

const useDialogStyles = makeStyles({
  root: {
    background: "linear-gradient(45deg, #FE6B8B 30%, #ef3b53  90%)",
  },
})

export function Error({
  message = GENERIC_MESSAGE,
  title,
}: {
  message?: string
  title: string
}): JSX.Element {
  const [isOpen, setIsOpen] = useState(true)
  const dialogStyles = useDialogStyles()

  function closeError(): void {
    setIsOpen(false)
  }

  return (
    <Dialog
      classes={{
        paper: dialogStyles.root,
      }}
      fullWidth
      maxWidth="xl"
      open={isOpen}
    >
      <div
        style={{
          alignItems: "center",
          display: "flex",
          justifyContent: "space-between",
        }}
      >
        <DialogTitle>{title}</DialogTitle>
        <IconButton onClick={closeError}>
          <CancelIcon />
        </IconButton>
      </div>
      <DialogContent>
        <DialogContentText style={{color: "#000000"}}>
          {message}
        </DialogContentText>
      </DialogContent>
    </Dialog>
  )
}
