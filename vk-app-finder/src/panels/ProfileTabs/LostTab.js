import React from "react";
import { CardGrid } from "@vkontakte/vkui";
import ProfileCard from "../../components/cards/ProfileCard";
import {decorate, observable, runInAction} from "mobx";
import {observer} from "mobx-react";
import DG from '2gis-maps';
import ProfileService from '../../services/LostService';


class LostTab extends React.Component {
  constructor(props) {
    super(props);
    this.profileService = new ProfileService();
    this.animals = [];
  }

  componentDidMount() {
    this.profileService.get().then(
      (result) => {
        runInAction(() => {
          this.animals = result.payload;
        });
      },
      (error) => {
        alert(error);
      }
    );
  }

  animalsToCards = () => {
    return this.animals.map((animal, index) => (
      <>
        <ProfileCard
          onClick={() => this.props.toLost(animal.id)}
          key={animal.id}
          animal={animal}
        />
      </>
    ));
  };

  render() {
    return <CardGrid>{this.animalsToCards()}</CardGrid>;
  }
}

decorate(LostTab, {
  animals: observable,
});

export default observer(LostTab);
