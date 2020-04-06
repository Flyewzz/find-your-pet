import cat from "../../img/persik.png";
import {Cell} from "@vkontakte/vkui";
import React from "react";
import Card from "@vkontakte/vkui/dist/components/Card/Card";

const AnimalCard = (props) => {
  return (
    <Card key={props.key} style={{height: 250}} className="animal__card" size="m" mode="shadow">
      <img src={cat} height='180px'/>
      <Cell>{`Dog ${props.animal}`}</Cell>
    </Card>
  );
};

export default AnimalCard;