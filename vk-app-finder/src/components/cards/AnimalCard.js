import React from "react";
import Card from "@vkontakte/vkui/dist/components/Card/Card";
import "./AnimalCard.css";
import config from "../../config";
import {Div} from "@vkontakte/vkui";
import getDefaultAnimal from '../default_animals/DefaultAnimals';
import getGenderInfo from '../gender/Genders';

const AnimalCard = (props) => {
  const animal = props.animal;
  const breed = animal.breed === "" ? "Порода не указана" : animal.breed;
  const date = new Date(animal.date.replace(" ", "T"))
    .toLocaleDateString()
    .replace(/\//g, ".");
  const type = props.type;
  const cardSize = (window.innerWidth > 550) ? 'm' : 'l'; // 'l' for mobiles
  const gender = getGenderInfo(animal.sex);

  return (
    <Card
      onClick={props.onClick}
      style={{height: "max(315px, auto)"}}
      className="animal__card"
      size={cardSize}
      mode="shadow"
    >
      <div className={"animal-card__photo__container"}>
        <div className={"animal-card__photo__corner"}>
          <div className={"animal-card__photo__breed"}>{breed}</div>
        </div>
        <img
          className={"animal-card__photo"}
          src={animal.picture_id ? (config.baseUrl + `${type}/img?id=${animal.picture_id}`)
            : getDefaultAnimal(animal.type_id)}
          alt={""}
        />
      </div>
      <Div className={"animal-card__data"}>
        <p className={"animal-card__address"}>{props.address}</p>
        <p className={"animal-card__date"}>Дата {type === 'lost' ? 'пропажи' : 'находки'}: {date}</p>
        <p className={"animal-card__gender"}>Пол: {gender.name}
          <img className={"animal-card__gender__picture"}
               src={gender.picture} width='32' height='32' alt={''}/>
        </p>
      </Div>
    </Card>
  );
};

export default AnimalCard;