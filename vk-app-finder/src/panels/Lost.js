import React from "react";
import {
  Card,
  CardGrid,
  Group,
  Button,
  Div,
  PanelHeader,
  List,
  Cell,
} from "@vkontakte/vkui";
import { FixedLayout, PanelHeaderBack } from "@vkontakte/vkui/dist/es6";
import "./Main.css";
import "../elements/cards/Card.css";
import cat from '../img/persik.png';

const getDogs = () => {
  let dogs = [];
  for (let i = 0; i < 2; ++i) {
    dogs.push(
      <Card className="animal__card" size="m" mode="shadow">
        <div style={{ height: 250 }}>
          <List>
            <Cell>
              <img src={cat} height='180px'></img>
            </Cell>
            <Cell>{`Dog ${i}`}</Cell>
          </List>
        </div>
      </Card>
    );
  }
  return dogs;
};
const LostPanel = () => {
  return (
    <>
      <PanelHeader left={<PanelHeaderBack />}>Потерялись</PanelHeader>
      <Group separator="hide">
        <CardGrid>{getDogs()}</CardGrid>
      </Group>
      <Group separator="hide">
        <CardGrid>{getDogs()}</CardGrid>
      </Group>
      <Group separator="hide">
        <CardGrid>{getDogs()}</CardGrid>
      </Group>
      <Group separator="hide">
        <CardGrid>{getDogs()}</CardGrid>
      </Group>
      <Group separator="hide">
        <CardGrid>{getDogs()}</CardGrid>
      </Group>
      <Group separator="hide">
        <CardGrid>{getDogs()}</CardGrid>
      </Group>
    </>
  );
};

export default LostPanel;
