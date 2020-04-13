import React from "react";
import {Button, Card, CardGrid, CellButton, Div, Group, Header, PanelHeader} from "@vkontakte/vkui";
import {Cell, PanelHeaderBack, Separator} from "@vkontakte/vkui/dist/es6";
import AnimalCard from "../components/cards/AnimalCard";
import FilterLine from "../components/cards/FilterLine";
import DG from '2gis-maps';
import LostService from '../services/LostService';
import {decorate, observable, runInAction} from "mobx";
import {observer} from "mobx-react";
import './Lost.css';
import config from "../config";
import List from "@vkontakte/vkui/dist/components/List/List";
import InfoRow from "@vkontakte/vkui/dist/components/InfoRow/InfoRow";
import Avatar from "@vkontakte/vkui/dist/components/Avatar/Avatar";
import Icon24Write from '@vkontakte/icons/dist/24/write';
import Icon24Done from '@vkontakte/icons/dist/24/done';
import './LostAnimalPanel.css';

class LostAnimalPanel extends React.Component {
  constructor(props) {
    super(props);
    this.lostService = new LostService();
  }

  animal = {date: ''};

  componentDidMount() {
    this.lostService.getById(this.props.id).then(
      result => {
        runInAction(() => {
          this.animal = result;
        })
      },
      error => {
        alert(error);
      })
  }

  sexText = {
    'm': 'Мужской',
    'f': 'Женский',
    'n/a': 'Не указан',
  };

  render() {
    const {picture_id, type_id, description} = this.animal;
    const breed = this.animal.breed === '' ? 'порода не указана' : this.animal.breed;
    const date = new Date(this.animal.date.replace(' ', 'T')).toLocaleDateString();
    const type = config.types[type_id - 1];

    return (
      <>
        <PanelHeader left={<PanelHeaderBack onClick={this.props.goBack}/>}>
          Потерянный питомец
        </PanelHeader>
        <Group separator="hide">
          <Div>
            <img style={{width: '100%'}}
                 src={config.baseUrl + `lost/img?id=${picture_id}`} alt={''}/>
          </Div>
          <Header>{type + ', ' + breed}</Header>
          <Group header={<Header mode={'secondary'}>Автор</Header>}>
            <Cell
              before={<Avatar src={config.baseUrl + `lost/img?id=${picture_id}`}/>}
            >
              Артур Стамбульцян
            </Cell>
            <Cell multiline={true} description={'Напишите автору, если у вас есть информация о животном'}>
              <CellButton before={<Icon24Write/>} className={'author__action-button'}>Написать</CellButton>
            </Cell>
            <Cell  multiline={true} description={'Или дайте ему знать, мы пришлем ему уведомление'}>
              <CellButton  before={<Icon24Done/>} className={'author__action-button'}>Я нашел!</CellButton>
            </Cell>
          </Group>
          <Separator/>
          <Group>
            <List>
              <Cell>
                <InfoRow header="Место пропажи">
                  {'тут должен быть адрес'}
                </InfoRow>
              </Cell>
              <Cell>
                <InfoRow header="Пол">
                  {this.sexText[this.animal.sex]}
                </InfoRow>
              </Cell>
              {description !== '' && <Cell multiline={true}>
                <InfoRow header="Описание">
                  {description}
                </InfoRow>
              </Cell>}
              <Cell>
                <InfoRow header="Дата создания объявления">
                  {date}
                </InfoRow>
              </Cell>
            </List>
          </Group>
        </Group>
      </>
    );
  }
}

decorate(LostAnimalPanel, {
  animal: observable,
});

export default observer(LostAnimalPanel);
