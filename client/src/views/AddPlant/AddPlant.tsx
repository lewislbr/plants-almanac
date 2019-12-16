import React, { useRef } from 'react';
import { useMutation } from '@apollo/react-hooks';
import gql from 'graphql-tag';

import {
  Button,
  H1,
  Input,
  Label,
  Section,
  TextArea,
  TextTip,
} from '../../components';
import { ButtonDiv } from './AddPlantStyles';

const ADD_PLANT = gql`
  mutation AddPlant(
    $name: String!
    $otherNames: String
    $description: String
    $plantSeason: String
    $harvestSeason: String
    $pruneSeason: String
    $tips: String
  ) {
    createPlant(
      name: $name
      otherNames: $otherNames
      description: $description
      plantSeason: $plantSeason
      harvestSeason: $harvestSeason
      pruneSeason: $pruneSeason
      tips: $tips
    ) {
      _id
      name
      otherNames
      description
      plantSeason
      harvestSeason
      pruneSeason
      tips
    }
  }
`;

export function AddPlant(): JSX.Element {
  const [addPlant] = useMutation(ADD_PLANT);

  const nameElement = useRef<HTMLInputElement>(null!);
  const otherNamesElement = useRef<HTMLInputElement>(null!);
  const descriptionElement = useRef<HTMLTextAreaElement>(null!);
  const plantSeasonElement = useRef<HTMLInputElement>(null!);
  const harvestSeasonElement = useRef<HTMLInputElement>(null!);
  const pruneSeasonElement = useRef<HTMLInputElement>(null!);
  const tipsElement = useRef<HTMLTextAreaElement>(null!);

  function cancelAddPlant(): void {
    console.log('Canceled');
  }

  function confirmAddPlant(event: React.FormEvent<HTMLFormElement>): void {
    event.preventDefault();

    const name = nameElement.current.value;
    const otherNames =
      otherNamesElement.current.value == ''
        ? null
        : otherNamesElement.current.value;
    const description =
      descriptionElement.current.value == ''
        ? null
        : descriptionElement.current.value;
    const plantSeason =
      plantSeasonElement.current.value == ''
        ? null
        : plantSeasonElement.current.value;
    const harvestSeason =
      harvestSeasonElement.current.value == ''
        ? null
        : harvestSeasonElement.current.value;
    const pruneSeason =
      pruneSeasonElement.current.value == ''
        ? null
        : pruneSeasonElement.current.value;
    const tips =
      tipsElement.current.value == '' ? null : tipsElement.current.value;

    if (!name) return event.preventDefault();

    addPlant({
      variables: {
        name: name,
        otherNames: otherNames,
        description: description,
        plantSeason: plantSeason,
        harvestSeason: harvestSeason,
        pruneSeason: pruneSeason,
        tips: tips,
      },
    });
  }

  return (
    <>
      <Section>
        <div>
          <H1>Add Plant</H1>
        </div>
      </Section>
      <Section>
        <form onSubmit={confirmAddPlant}>
          <div>
            <Label>
              Name <TextTip>(Required)</TextTip>
            </Label>
            <Input type="text" ref={nameElement} />
          </div>
          <div>
            <Label>Other Names</Label>
            <Input type="text" ref={otherNamesElement} />
          </div>
          <div>
            <Label>Description</Label>
            <TextArea rows={4} ref={descriptionElement} />
          </div>
          <div>
            <Label>Plant Season</Label>
            <Input type="text" ref={plantSeasonElement} />
          </div>
          <div>
            <Label>Harvest Season</Label>
            <Input type="text" ref={harvestSeasonElement} />
          </div>
          <div>
            <Label>Prune Season</Label>
            <Input type="text" ref={pruneSeasonElement} />
          </div>
          <div>
            <Label>Tips</Label>
            <TextArea rows={4} ref={tipsElement} />
          </div>
          <ButtonDiv>
            <Button type="button" onClick={cancelAddPlant} secondary>
              Cancel
            </Button>
            <Button type="submit" primary>
              Confirm
            </Button>
          </ButtonDiv>
        </form>
      </Section>
    </>
  );
}
