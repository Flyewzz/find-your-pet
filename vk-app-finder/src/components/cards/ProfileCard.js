import React from "react";
import Card from "@vkontakte/vkui/dist/components/Card/Card";
import "./ProfileCard.css";
import Div from "@vkontakte/vkui/dist/components/Div/Div";
import config from "../../config";

const ProfileCard = (props) => {
  const animal = props.animal;
  const breed = animal.breed === "" ? "порода не указана" : animal.breed;
  const date = new Date(animal.date.replace(" ", "T")).toLocaleDateString();

  return (
    <Card
      onClick={props.onClick}
      style={{ height: 181, display: "flex" }}
      className={"animal-card__container"}
      size="l"
      mode="shadow"
    >
      {/* <Div style={{display: "flex", flexDirection: "column"}}> */}
        <img
          className={"animal-card__photo"}
          src={config.baseUrl + `lost/img?id=${animal.picture_id}`}
          style={{ height: "180px", flexGrow: 2 }}
          alt={""}
        />
        <Div className={"animal-card__details"} style={{ flexGrow: 2 }}>
          <Div className={"animal-card__info"}>
            {`${config.types[animal.type_id - 1]}, ${breed}`}
          </Div>
          <Div className={"animal-card__address"}>{props.animal.place}</Div>
          <Div className={"animal-card__date"}>{date}</Div>
        </Div>
      {/* </Div> */}
    </Card>
  );
};

export default ProfileCard;
