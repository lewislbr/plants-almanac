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
const PlantDetails = React.lazy(() =>
  import("./views/PlantDetails" /* webpackChunkName: 'PlantDetails' */).then(
    ({PlantDetails}) => ({
      default: PlantDetails,
    }),
  ),
)
const Plants = React.lazy(() =>
  import("./views/Plants" /* webpackChunkName: 'Plants' */).then(
    ({Plants}) => ({
      default: Plants,
    }),
  ),
)

export function App(): JSX.Element {
  return (
    <BrowserRouter>
      <Header />
      <main className="max-w-5xl mx-auto p-5pc">
        <React.Suspense fallback={<div>{"Loading..."}</div>}>
          <Switch>
            <Route exact path="/" component={Plants} />
            <Route path="/add-plant" component={AddPlant} />
            <Route path="/:id" component={PlantDetails} />
          </Switch>
        </React.Suspense>
      </main>
    </BrowserRouter>
  )
}
