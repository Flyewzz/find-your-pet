import React from "react";
import {CardGrid} from "@vkontakte/vkui";
import ProfileCard from "./ProfileCard";
import {decorate, observable, runInAction} from "mobx";
import {observer} from "mobx-react";
import ProfileService from '../../services/ProfileService';
import Icon56InfoOutline from '@vkontakte/icons/dist/56/info_outline';
import Placeholder from "@vkontakte/vkui/dist/components/Placeholder/Placeholder";
import Button from "@vkontakte/vkui/dist/components/Button/Button";
import GeocodingService from "../../services/GeocodingService";


class LostTab extends React.Component {
  constructor(props) {
    super(props);
    this.profileService = new ProfileService();
    this.geocodingService = new GeocodingService();
  }

  animals = null;
  addresses = [];

  componentDidMount() {
    this.props.userStore.getId().then(
      result => this.fetchLost(result.id)
    );
  }

  updateAddress = (index, result) => {
    const city = result.City === '' ? result.District : result.City;
    const address = result.MetroArea === '' ? result.Address : result.MetroArea;
    this.addresses[index] = city + (address === '' ? '' : ', ' + address);
  };

  fetchLost = (id) => {
    this.profileService.getLost(id).then(
      (result) => {
        runInAction(() => {
          this.animals = result.payload;
          this.addresses = result.payload === null ? [] : this.animals.map(() => '')
        });
        this.animals.forEach((value, index) => {
          const {longitude, latitude} = value;
          this.geocodingService.addressByCoords(longitude, latitude).then(
            result => this.updateAddress(index, result.address)
          );
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
        this.props.goBack();
      }
    );
  };

  animalsToCards = () => {
    return this.animals.map((animal, index) => (
      <ProfileCard
        onClick={() => this.props.toLost(animal.id)}
        cancel={() => {
          this.props.openDestructive(() => this.onRemove(animal.id))
        }}
        key={animal.id}
        animal={animal}
        address={this.addresses[index]}
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
  addresses: observable,
});

export default observer(LostTab);
