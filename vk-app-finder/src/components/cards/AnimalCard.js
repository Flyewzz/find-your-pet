import React from "react";
import Card from "@vkontakte/vkui/dist/components/Card/Card";
import "./AnimalCard.css"
import config from '../../config';
import {Cell} from "@vkontakte/vkui";

const AnimalCard = (props) => {
  const animal = props.animal;
  const breed = animal.breed === ''? 'порода не указана': animal.breed;
  const date = new Date(animal.date.replace(' ', 'T')).toLocaleDateString();
  console.log(props.address);

  return (
    <Card onClick={props.onClick} style={{height: 315}}
          className="animal__card" size="m" mode="shadow">
      <img className={'animal-card__photo'}
           src={config.baseUrl + `lost/img?id=${animal.picture_id}`}
           height='180px' alt={''}/>
      <Cell className={'animal-card__address'}>{props.address}</Cell>
      <Cell className={'animal-card__info'}>
        {`${config.types[animal.type_id - 1]}, ${breed}`}
      </Cell>
      <Cell className={'animal-card__date'}>{date}</Cell>
    </Card>
  );
};

export default AnimalCard;