import React, {lazy, StrictMode, Suspense} from "react"
import {Route, Switch} from "react-router-dom"
import {Header} from "./components"

const AddPlant = lazy(() =>
  import("./views/AddPlant" /* webpackChunkName: 'AddPlant' */).then(
    ({AddPlant}) => ({
      default: AddPlant,
    }),
  ),
)
const PlantDetails = lazy(() =>
  import("./views/PlantDetails" /* webpackChunkName: 'PlantDetails' */).then(
    ({PlantDetails}) => ({
      default: PlantDetails,
    }),
  ),
)
const PlantList = lazy(() =>
  import("./views/PlantList" /* webpackChunkName: 'PlantList' */).then(
    ({PlantList}) => ({
      default: PlantList,
    }),
  ),
)

export function App(): JSX.Element {
  return (
    <StrictMode>
      <Header />
      <main className="max-w-5xl mx-auto p-5pc">
        <Suspense fallback={<div>{"Loading..."}</div>}>
          <Switch>
            <Route exact path="/" component={PlantList} />
            <Route path="/add-plant" component={AddPlant} />
            <Route path="/:id" component={PlantDetails} />
          </Switch>
        </Suspense>
      </main>
    </StrictMode>
  )
}
