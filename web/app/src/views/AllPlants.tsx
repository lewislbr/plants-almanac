import * as React from "react"
import {Link} from "react-router-dom"
import {gql, useQuery} from "@apollo/client"
import {PlantCard} from "../components"
import {Plants} from "../graphql"

const PLANTS = gql`
  query Plants {
    plants {
      id
      name
    }
  }
`

export function AllPlants(): JSX.Element {
  const {data, loading, error, refetch} = useQuery<Plants>(PLANTS)

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
              <Link to={`/${plant?.id}`} key={plant?.id || 0}>
                <PlantCard {...{name: plant?.name}} />
              </Link>
            ))}
          </div>
        )}
      </section>
    </>
  )
}
