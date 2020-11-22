import * as React from "react"
import {BrowserRouter, Route, Switch} from "react-router-dom"
import {Header} from "./components"

const AddPlant = React.lazy(() =>
  import("./views/AddPlant" /* webpackChunkName: 'AddPlant' */).then(
    ({AddPlant}) => ({
      default: AddPlant,
    }),
  ),
)
const AllPlants = React.lazy(() =>
  import("./views/AllPlants" /* webpackChunkName: 'AllPlants' */).then(
    ({AllPlants}) => ({
      default: AllPlants,
    }),
  ),
)
const PlantDetails = React.lazy(() =>
  import("./views/PlantDetails" /* webpackChunkName: 'PlantDetails' */).then(
    ({PlantDetails}) => ({
      default: PlantDetails,
    }),
  ),
)

export function App(): JSX.Element {
  return (
    <React.StrictMode>
      <BrowserRouter>
        <Header />
        <main className="max-w-5xl mx-auto p-5pc">
          <React.Suspense fallback={<div>{"Loading..."}</div>}>
            <Switch>
              <Route exact path="/" component={AllPlants} />
              <Route path="/add-plant" component={AddPlant} />
              <Route path="/:id" component={PlantDetails} />
            </Switch>
          </React.Suspense>
        </main>
      </BrowserRouter>
    </React.StrictMode>
  )
}
