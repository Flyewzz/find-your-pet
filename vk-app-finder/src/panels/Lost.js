import React from "react";
import {
  Card,
  CardGrid,
  Group,
  PanelHeader,
  } from "@vkontakte/vkui";
import {PanelHeaderBack} from "@vkontakte/vkui/dist/es6";
import AnimalCard from "../components/card/AnimalCard";

const getDogs = () => {
  let dogs = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

  return dogs.map((animal, index) =>
    <>
      {!(index % 2) && <Card size="l" styles={{height: 0}}/>}
      <AnimalCard key={index} animval={animal}/>
    </>
  );
};

const LostPanel = () => {
  return (
    <>
      <PanelHeader left={<PanelHeaderBack/>}>Потерялись</PanelHeader>
      <Group separator="hide">
        <CardGrid>{getDogs()}</CardGrid>
      </Group>
    </>
  );
};

export default LostPanel;
