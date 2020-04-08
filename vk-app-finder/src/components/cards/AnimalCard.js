import cat from "../../img/cat_example.jpg";
import React from "react";
import Card from "@vkontakte/vkui/dist/components/Card/Card";
import "./AnimalCard.css"
import Div from "@vkontakte/vkui/dist/components/Div/Div";

const AnimalCard = (props) => {
  return (
    <Card key={props.key} style={{height: 275}} className="animal__card" size="m" mode="shadow">
      <img className={'animal-card__photo'} src={cat} height='180px'/>
      <Div className={'animal-card__info'}>
        {`${props.animal.type}, ${props.animal.breed}`}
      </Div>
      <Div className={'animal-card__address'}>{props.animal.place}</Div>
      <Div className={'animal-card__date'}>
        {`Пропал: ${props.animal.date}`}
      </Div>
    </Card>
  );
};

export default AnimalCard;