import React from 'react';
import Card from '@vkontakte/vkui/dist/components/Card/Card';
import './ProfileCard.css';
import Icon28CancelOutline from '@vkontakte/icons/dist/28/cancel_outline';
import config from '../../config';
import {Group, InfoRow, List} from "@vkontakte/vkui";
import {Cell} from "@vkontakte/vkui/dist/es6";

const ProfileCard = (props) => {
  const animal = props.animal;
  const breed = animal.breed === '' ? 'порода не указана' : animal.breed;
  const date = new Date(animal.date.replace(' ', 'T')).toLocaleDateString();
  const place = 'Адрес, улица, дом';

  return (
    <Card
      onClick={props.onClick}
      style={{height: 181}}
      className={'profile__card__container'}
      size="l"
      mode="shadow"
    >
      <div style={{display: 'flex', height: '100%', width: '100%'}}>
        <img
          className={'profile__card__photo'}
          src={config.baseUrl + `lost/img?id=${animal.picture_id}`}
          alt={""}
        />
        <Group className={'profile__card__details'}>
          <Icon28CancelOutline onClick={() => props.cancel(animal.id)} className={'profile__cancel-icon'}/>
          <Cell>
            <InfoRow className={'profile__card__info'}>
              <span style={{fontWeight: 'bold'}}>Потерялся: </span>
              {`${config.types[animal.type_id - 1]}, ${breed}`}
            </InfoRow>
          </Cell>
          <Cell>
            <InfoRow className={'profile__card__address'}>
              <span style={{fontWeight: 'bold'}}>Адрес: </span>{place}
            </InfoRow>
          </Cell>
          <Cell>
            <InfoRow className={'profile__card__date'}>
              <span style={{fontWeight: 'bold'}}>Дата пропажи: </span>{date}
            </InfoRow>
          </Cell>
          <Cell>
            <InfoRow>
              <span style={{fontWeight: 'bold'}}>Описание: </span>
              {animal.description}
            </InfoRow>
          </Cell>
        </Group>
      </div>
    </Card>
  );
};

export default ProfileCard;
