import React from "react";
import {CellButton, Div, Group, Header, PanelHeader} from "@vkontakte/vkui";
import {Cell, PanelHeaderBack, Separator} from "@vkontakte/vkui/dist/es6";
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
import Icon24Share from '@vkontakte/icons/dist/24/share';
import './LostAnimalPanel.css';

class LostAnimalPanel extends React.Component {
  constructor(props) {
    super(props);
    this.lostService = new LostService();
  }

  animal = {date: ''};
  // TODO add default avatar
  author = {first_name: '', last_name: '', photo_50: ''};

  componentDidMount() {
    this.lostService.getById(this.props.id).then(
      result => {
        runInAction(() => {
          this.animal = result;
        });
        this.props.userStore.getUserById(this.animal.vk_id).then(
          result => {
            runInAction(() => {
              this.author = result.response[0];
            })
          });
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

  writeHref = () => {
    if (this.author.can_write_private_message) {
      return {
        href: `https://vk.com/im?sel=${this.animal.vk_id}`,
        description: 'Напишите автору, если у вас есть информация о животном'
      };
    }
    return {
      href: `https://vk.com/id${this.animal.vk_id}`,
      description: 'ЛС автора закрыто, но вы можете отправить заявку на добавление в друзья'
    };
  };

  getShareText = () => {
    const {type_id} = this.animal;
    const breed = this.animal.breed === '' ? 'порода не указана' : this.animal.breed;
    // TODO getting address from coords
    return `Потерян питомец: ${config.types[type_id - 1]}, ${breed}. ` +
      `Адрес: ${'тут должен быть адрес'}. ${config.appUrl}`
  };

  share = () => {
    this.props.userStore.share(this.getShareText());
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
          <div style={{display: 'flex', alignItems: 'center'}}>
            <Header style={{paddingBottom: '10px'}}>{type + ', ' + breed}</Header>
            <Div onClick={this.share}
                 style={{display: 'flex', alignItems: 'center', marginLeft: 'auto'}}>
              {'Поделиться'}<Icon24Share style={{marginLeft: '5px'}}/>
            </Div>
          </div>
          <Group header={<Header mode={'secondary'}>Автор</Header>}>
            <Cell
              before={<Avatar src={this.author.photo_50}/>}>
              {this.author.first_name + ' ' + this.author.last_name}
            </Cell>
            <Cell multiline={true} description={this.writeHref().description}>
              <CellButton href={this.writeHref().href}
                          target={'_blank'}
                          before={<Icon24Write/>} className={'author__action-button'}>Написать</CellButton>
            </Cell>
            <Cell multiline={true} description={'Или дайте ему знать, мы пришлем ему уведомление'}>
              <CellButton before={<Icon24Done/>} className={'author__action-button'}>Я нашел!</CellButton>
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
                <InfoRow header="Пол животного">
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
  author: observable,
});

export default observer(LostAnimalPanel);
