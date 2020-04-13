import cat from "../../img/cat_example.jpg";
import React from "react";
import Card from "@vkontakte/vkui/dist/components/Card/Card";
import "./AnimalCard.css"
import Div from "@vkontakte/vkui/dist/components/Div/Div";
import config from '../../config';

const AnimalCard = (props) => {
  const animal = props.animal;
  const breed = animal.breed === ''? 'порода не указана': animal.breed;
  const date = new Date(animal.date.replace(' ', 'T')).toLocaleDateString();

  return (
    <Card key={props.key} style={{height: 275}}
          className="animal__card" size="m" mode="shadow">
      <img className={'animal-card__photo'}
           src={config.baseUrl + `lost/img?id=${animal.picture_id}`}
           height='180px' alt={''}/>
      <Div className={'animal-card__info'}>
        {`${config.types[animal.type_id - 1]}, ${breed}`}
      </Div>
      <Div className={'animal-card__address'}>{props.animal.place}</Div>
      <Div className={'animal-card__date'}>{date}</Div>
    </Card>
  );
};

export default AnimalCard;