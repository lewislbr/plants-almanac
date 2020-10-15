import * as React from "react"
import {Link} from "react-router-dom"
import {PlantCard} from "../components"
import {usePlantsQuery} from "../graphql/types"

export function Plants(): JSX.Element {
  const {data, loading, error, refetch} = usePlantsQuery()

  React.useEffect(() => {
    refetch()
  }, [refetch])

  return (
    <>
      <section>
        <h1 className="page-title">{"Plants"}</h1>
      </section>
      <section className="mt-8">
        {loading ? (
          <p>{"Loading..."}</p>
        ) : error ? (
          <p>{"ERROR"}</p>
        ) : (
          <div>
            {data?.plants?.map((plant) => (
              <Link to={`/${plant?._id}`} key={plant?._id || 0}>
                <PlantCard {...{name: plant?.name}} />
              </Link>
            ))}
          </div>
        )}
      </section>
    </>
  )
}
