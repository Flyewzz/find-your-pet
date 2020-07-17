import React from "react";
import Card from "@vkontakte/vkui/dist/components/Card/Card";
import "./ProfileLostCard.css";
import Icon24DismissDark from '@vkontakte/icons/dist/24/dismiss_dark';
import config from "../../config";
import {Group, List} from "@vkontakte/vkui";
import getDefaultAnimal from "../default_animals/DefaultAnimals";
import LinesEllipsis from 'react-lines-ellipsis'

const ProfileLostCard = (props) => {
  const animal = props.animal;
  const breed = animal.breed === "" ? "Порода не указана" : animal.breed;
  const date = new Date(animal.date.replace(" ", "T"))
    .toLocaleDateString()
    .replace(/\//g, ".");

  return (
    <Card
      onClick={() => props.open(animal.id)}
      className={"profile__card"}
      size="l"
      mode="shadow">
      <div className={"profile__card__container"}>
        <div className={"profile__card__photo__container"}>
          <div className={"profile__card__photo__corner"}>
            <div className={"profile__card__photo__breed"}>{breed}</div>
          </div>
          <img className={"profile__card__photo"}
               src={
                 animal.picture_id
                   ? config.baseUrl + `lost/img?id=${animal.picture_id}`
                   : getDefaultAnimal(animal.type_id)
               }
               alt={""}/>
        </div>
        <Group className={"profile__card__details"}>
          <List className={"profile__attributes_list"}>
            <Icon24DismissDark
              onClick={(e) => {
                e.stopPropagation();
                props.cancel(animal.id)
              }}
              className={"profile__cancel-icon"}
            />
            <div className={"profile__card__cells"}>
              <div className={"profile__card__info_cell"}>
                <span style={{fontWeight: "bold"}}>Адрес: </span>
                {animal.address}
              </div>

              <div className={"profile__card__info_cell"}>
                <span style={{fontWeight: "bold"}}>Дата создания объявления: </span>
                {date}
              </div>

              <div className={"profile__card__info_cell"}>
                <span style={{fontWeight: "bold"}}>Описание: </span>
                <LinesEllipsis
                  text={animal.description}
                  maxLine='2'
                  ellipsis='...'
                  trimRight
                  basedOn='letters'
                  componenr={<span/>}
                />
              </div>
            </div>
          </List>
        </Group>
      </div>
    </Card>
  );
};

export default ProfileLostCard;