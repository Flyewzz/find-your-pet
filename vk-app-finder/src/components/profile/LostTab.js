import React from "react";
import {CardGrid} from "@vkontakte/vkui";
import ProfileCard from "./ProfileCard";
import {decorate, observable, runInAction} from "mobx";
import {observer} from "mobx-react";
import ProfileService from '../../services/ProfileService';


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

  animalsToCards = () => {
    return this.animals.map(animal => (
      <ProfileCard
        onClick={() => this.props.toLost(animal.id)}
        cancel={this.props.openDestructive}
        key={animal.id}
        animal={animal}
      />
    ));
  };

  render() {
    return (
      <CardGrid>
        {!this.animals && 'пуста'}
        {this.animals && this.animalsToCards()}
      </CardGrid>);
  }
}

decorate(LostTab, {
  animals: observable,
});

export default observer(LostTab);
