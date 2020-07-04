import React from "react";
import {CardGrid, Spinner} from "@vkontakte/vkui";
import ProfileFoundCard from "./ProfileFoundCard";
import {decorate, observable, runInAction} from "mobx";
import {observer} from "mobx-react";
import ProfileService from "../../services/ProfileService";
import Icon56InfoOutline from "@vkontakte/icons/dist/56/info_outline";
import Placeholder from "@vkontakte/vkui/dist/components/Placeholder/Placeholder";
import Button from "@vkontakte/vkui/dist/components/Button/Button";


class FoundTab extends React.Component {
  constructor(props) {
    super(props);
    this.profileService = new ProfileService();
  }

  animals = null;
  loading = true;

  componentDidMount() {
    this.props.userStore.getId().then((result) => this.fetchFound(result.id));
  }

  fetchFound = (id) => {
    this.profileService.getFound(id).then(
      (result) => {
        runInAction(() => {
          this.animals = result.payload;
        });

        this.loading = false;
      },
      (error) => {
        console.log(error);
      }
    );
  };

  onRemove = (id) => {
    const vkId = this.props.userStore.id;
    this.profileService.closeFound(id, vkId).then(() => {
      this.props.goBack();
    });
  };

  animalsToCards = () => {
    return this.animals.map((animal) => (
      <ProfileFoundCard
        onClick={() => {
          this.props.toFound(animal.id);
        }}
        cancel={() => {
          this.props.openDestructive(() => this.onRemove(animal.id));
        }}
        key={animal.id}
        animal={animal}
      />
    ));
  };

  render() {
    return (
      <CardGrid>
        {this.loading &&
        <Spinner size="large" style={{marginTop: 20, color: "rgb(83, 118, 164)"}}/>
        }
        {!this.loading && !this.animals && (
          <Placeholder
            icon={<Icon56InfoOutline/>}
            action={
              <Button onClick={this.props.toMainForm} size="l"
                      style={{cursor: 'pointer'}}>
                Cоздать объявление
              </Button>
            }
          >
            У Вас пока нет объявлений о находке
          </Placeholder>
        )}
        {this.animals && this.animalsToCards()}
      </CardGrid>
    );
  }
}

decorate(FoundTab, {
  animals: observable,
  loading: observable,
});

export default observer(FoundTab);
