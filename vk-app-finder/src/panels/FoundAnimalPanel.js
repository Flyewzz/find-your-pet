import React from "react";
import {CellButton, Div, Group, Header, PanelHeader} from "@vkontakte/vkui";
import {Cell, PanelHeaderBack, Separator} from "@vkontakte/vkui/dist/es6";
import {decorate, observable, runInAction} from "mobx";
import {observer} from "mobx-react";
import "./Lost.css"; // The same styles
import config from "../config";
import List from "@vkontakte/vkui/dist/components/List/List";
import InfoRow from "@vkontakte/vkui/dist/components/InfoRow/InfoRow";
import Avatar from "@vkontakte/vkui/dist/components/Avatar/Avatar";
import Icon24Write from "@vkontakte/icons/dist/24/write";
import Icon24Done from "@vkontakte/icons/dist/24/done";
import Icon24Share from "@vkontakte/icons/dist/24/share";
import Icon24Cancel from "@vkontakte/icons/dist/24/cancel";
import "./FoundAnimalPanel.css";
import FoundService from "../services/FoundService";
import ProfileService from "../services/ProfileService";
import getDefaultAnimal from '../components/default_animals/DefaultAnimals';


class FoundAnimalPanel extends React.Component {
  constructor(props) {
    super(props);
    this.foundService = new FoundService();
    this.profileService = new ProfileService();
  }

  animal = {date: ""};
  // TODO add default avatar
  author = {first_name: "", last_name: "", photo_50: ""};
  isMine = false;

  componentDidMount() {
    this.props.userStore.getId().then((resultId) =>
      this.foundService.getById(this.props.id).then(
        (result) => {
          runInAction(() => {
            this.animal = result;
            this.isMine = this.animal.vk_id === resultId.id;
          });
          this.props.userStore.getUserById(this.animal.vk_id).then((result) => {
            runInAction(() => {
              this.author = result.response[0];
            });
          });
        },
        (error) => {
          alert(error);
        }
      )
    );
  }

  sexText = {
    m: "Мужской",
    f: "Женский",
    "n/a": "Не указан",
  };

  writeHref = () => {
    if (this.author.can_write_private_message) {
      return {
        href: `https://vk.com/im?sel=${this.animal.vk_id}`,
        description: "Напишите автору, если у Вас есть информация о животном",
      };
    }
    return {
      href: `https://vk.com/id${this.animal.vk_id}`,
      description: `Отправить заявку на добавление в друзья`,
    };
  };

  getProfileLink() {
    return `https://vk.com/id${this.animal.vk_id}`;
  }

  getShareText = () => {
    const {type_id} = this.animal;
    const breed =
      this.animal.breed === "" ? "Порода не указана" : this.animal.breed;
    return (
      `Найден питомец: ${config.types[type_id - 1]}, ${breed}. ` +
      `Адрес: ${this.animal.address}. ${config.appUrl}`
    );
  };

  share = () => {
    this.props.userStore.share(this.getShareText());
  };

  onClose = () => {
    const vkId = this.props.userStore.id;
    this.profileService
      .closeFound(this.props.id, vkId)
      .then(() => this.props.goBack());
  };

  render() {
    const {picture_id, type_id, description} = this.animal;
    const breed =
      this.animal.breed === "" ? "Порода не указана" : this.animal.breed;
    const date = new Date(
      this.animal.date.replace(" ", "T")
    ).toLocaleDateString();
    const type = config.types[type_id - 1];

    return (
      <>
        <PanelHeader left={<PanelHeaderBack onClick={this.props.goBack}/>}>
          Найденный питомец
        </PanelHeader>
        <Group separator="hide">
          <Div>
            <img
              style={{width: "100%"}}
              src={
                picture_id
                  ? config.baseUrl + `found/img?id=${picture_id}`
                  : getDefaultAnimal(type_id)
              }
              alt={""}
            />
          </Div>
          <div style={{display: "flex", alignItems: "center"}}>
            <Header style={{paddingBottom: "10px"}}>
              {type + ", " + breed}
            </Header>
            <Div
              onClick={this.share}
              style={{
                cursor: "pointer",
                display: "flex",
                alignItems: "center",
                marginLeft: "auto",
              }}
            >
              {"Поделиться"}
              <Icon24Share style={{marginLeft: "5px"}}/>
            </Div>
          </div>
          <Group
            header={!this.isMine && <Header mode={"secondary"}>Автор</Header>}
          >
            {!this.isMine ? [
              <a href={this.getProfileLink()} style={{textDecoration: "none"}}>
                <Cell className={"author__info"} before={<Avatar src={this.author.photo_50}/>}>
                  {this.author.first_name + " " + this.author.last_name}
                </Cell>
              </a>,
              <Cell multiline={true}
                    description={this.writeHref().description}>
                <CellButton href={this.writeHref().href}
                            target={"_blank"}
                            before={<Icon24Write/>}
                            className={"author__action-button"}>
                  Написать
                </CellButton>
              </Cell>
            ] : [
              <CellButton
                before={<Icon24Cancel/>}
                onClick={() => this.props.openDestructive(this.onClose)}
                mode={"danger"}
              >
                Закрыть объявление
              </CellButton>,
            ]}
          </Group>
          <Separator/>
          <Group>
            <List>
              <Cell>
                <InfoRow header="Место находки">{this.animal.address}</InfoRow>
              </Cell>
              <Cell>
                <InfoRow header="Пол животного">
                  {this.sexText[this.animal.sex]}
                </InfoRow>
              </Cell>
              {description !== "" && (
                <Cell multiline={true}>
                  <InfoRow header="Описание">{description}</InfoRow>
                </Cell>
              )}
              <Cell>
                <InfoRow header="Дата создания объявления">{date}</InfoRow>
              </Cell>
            </List>
          </Group>
        </Group>
      </>
    );
  }
}

decorate(FoundAnimalPanel, {
  animal: observable,
  author: observable,
  isMine: observable,
});

export default observer(FoundAnimalPanel);
