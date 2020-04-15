import React from "react";
import {CardGrid} from "@vkontakte/vkui";
import ProfileCard from "./ProfileCard";
import {decorate, observable, runInAction} from "mobx";
import {observer} from "mobx-react";
import ProfileService from '../../services/ProfileService';
import Icon56InfoOutline from '@vkontakte/icons/dist/56/info_outline';
import Placeholder from "@vkontakte/vkui/dist/components/Placeholder/Placeholder";
import Button from "@vkontakte/vkui/dist/components/Button/Button";


class LostTab extends React.Component {
  constructor(props) {
    super(props);
    this.profileService = new ProfileService();
  }

  animals = null;

  componentDidMount() {
    this.props.userStore.getId().then(
      result => this.fetchLost(result.id)
    );
  }

  fetchLost = (id) => {
    this.profileService.getLost(id).then(
      (result) => {
        runInAction(() => {
          this.animals = result.payload;
        });
      },
      (error) => {
        alert(error);
      }
    );
  };

  onRemove = (id) => {
    const vkId = this.props.userStore.id;
    this.profileService.close(id, vkId).then(
      () => {
        console.log('я тута');
        this.props.goBack();
      }
    );
  };

  animalsToCards = () => {
    return this.animals.map(animal => (
      <ProfileCard
        onClick={() => this.props.toLost(animal.id)}
        cancel={() => {
          this.props.openDestructive(() => this.onRemove(animal.id))
        }}
        key={animal.id}
        animal={animal}
      />
    ));
  };

  render() {
    return (
      <CardGrid>
        {!this.animals
        && <Placeholder stretched={true}
                        icon={<Icon56InfoOutline/>}
                        action={<Button onClick={this.toMainForm} size="l">
                          Cоздать новое объявление
                        </Button>}>
          У вас пока нет объявлений о пропаже
        </Placeholder>}
        {this.animals && this.animalsToCards()}
      </CardGrid>);
  }
}

decorate(LostTab, {
  animals: observable,
});

export default observer(LostTab);
