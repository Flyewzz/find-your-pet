import React from "react";
import {Card, CardGrid, Div, Group, PanelHeader} from "@vkontakte/vkui";
import {PanelHeaderBack} from "@vkontakte/vkui/dist/es6";
import AnimalCard from "../components/cards/AnimalCard";
import FilterLine from "../components/cards/FilterLine";
import DG from '2gis-maps';
import LostService from '../services/LostService';
import {decorate, observable, runInAction} from "mobx";
import {observer} from "mobx-react";
import './Lost.css';

class LostPanel extends React.Component {
  constructor(props) {
    super(props);

    this.state = {mapView: false, places: []};
    this.changeView = () => {
      const current = this.state.mapView;
      this.setState({mapView: !current});
    };

    this.lostService = new LostService();
  }

  animals = [];

  componentDidMount() {
    // TODO get position of user and filter by it
    this.map = DG.map('map', {
      center: [54.98, 82.89],
      zoom: 13,
    });

    this.lostService.get().then(
      result => {
        runInAction(() => {
          this.animals = result.payload;
        })
      },
      error => {
        alert(error);
      })
  }

  createMarkers = () => {
    this.animals.forEach(value => {
      DG.marker([value.longitude, value.latitude]).addTo(this.map);
    });
  };

  animalsToCards = () => {
    return this.animals.map((animal, index) =>
      <>
        {!(index % 2) && <Card key={-animal.id} size="l" styles={{height: 0}}/>}
        <AnimalCard onClick={() => this.props.toLost(animal.id)}
                    key={animal.id} animal={animal}/>
      </>
    );
  };

  render() {
    const mapStyle = {
      display: this.state.mapView ? undefined : 'none',
      height: '500x'
    };
    this.createMarkers();

    return (
      <>
        <PanelHeader left={<PanelHeaderBack/>}>Потерялись</PanelHeader>
        <Group separator="hide">
          <FilterLine isMap={this.state.mapView}
                      changeView={this.changeView}
                      openFilters={this.props.openFilters}/>

          {!this.state.mapView && <CardGrid>{this.animalsToCards()}</CardGrid>}

          <Div>
            <div style={mapStyle} id={'map'}/>
          </Div>
        </Group>
      </>
    );
  }
}

decorate(LostPanel, {
  animals: observable,
});

export default observer(LostPanel);
