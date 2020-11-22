import {gql} from "@apollo/client"
import {client, history} from "../index"

const DELETE = gql`
  mutation Delete($id: ID!) {
    delete(id: $id)
  }
`

export async function deleteOne(id: string): Promise<void> {
  await client.mutate({mutation: DELETE, variables: {id: id}})

  history.push("/")
}
