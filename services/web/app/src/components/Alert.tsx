import React, {Dispatch, SetStateAction} from "react"

export function Alert({
  deletePlant,
  id,
  setAlertOpen,
}: {
  deletePlant: (id: string) => Promise<void>
  id: string
  setAlertOpen: Dispatch<SetStateAction<boolean>>
}): JSX.Element {
  function closeAlert(): void {
    setAlertOpen(false)
  }

  return (
    <div className="items-center bg-backdrop fixed flex h-full justify-center left-0 right-0 top-0 w-full z-20">
      <div className="content-center bg-gray-100 fixed flex flex-col m-auto p-5 rounded-lg z-30">
        <p className="block mb-6">
          {"The plant will be deleted from the database."}
        </p>
        <button
          className="button button-danger mb-4"
          onClick={(): Promise<void> => deletePlant(id)}
          type="button"
        >
          {"Delete"}
        </button>
        <button
          className="button button-secondary"
          onClick={closeAlert}
          type="button"
        >
          {"Cancel"}
        </button>
      </div>
    </div>
  )
}
