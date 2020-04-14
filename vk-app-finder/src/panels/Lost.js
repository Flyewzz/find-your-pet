import React from "react";
import {Card, CardGrid, Div, Group, PanelHeader} from "@vkontakte/vkui";
import {PanelHeaderBack} from "@vkontakte/vkui/dist/es6";
import AnimalCard from "../components/cards/AnimalCard";
import FilterLine from "../components/cards/FilterLine";
import LostService from '../services/LostService';
import {decorate, observable, runInAction} from "mobx";
import {observer} from "mobx-react";
import './Lost.css';
import {YMaps, Map, Placemark, ZoomControl} from 'react-yandex-maps';

class LostPanel extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      mapView: props.mapStore.isMapView,
      places: []
    };
    this.changeView = () => {
      const current = props.mapStore.isMapView;
      props.mapStore.isMapView = !current;
    };

    this.lostService = new LostService();
  }

  animals = [];

  componentDidMount() {
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
    return this.animals.map(value =>
      <Placemark onClick={() => this.props.toLost(value.id)}
                 geometry={[value.longitude, value.latitude]}/>
    )
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

  onBoundsChange = e => {
    this.props.mapStore.center = e.get('target').getCenter();
    this.props.mapStore.zoom = e.get('target').getZoom();
  };

  render() {
    const mapStyle = {
      display: this.props.mapStore.isMapView ? undefined : 'none',
      height: '490px',
      width: '100%',
    };

    return (
      <>
        <PanelHeader left={<PanelHeaderBack/>}>Потерялись</PanelHeader>
        <Group separator="hide">
          <FilterLine isMap={this.state.mapView}
                      changeView={this.changeView}
                      openFilters={this.props.openFilters}/>

          {!this.props.mapStore.isMapView
          && <CardGrid>{this.animalsToCards()}</CardGrid>}

          <Div><YMaps>
            <div>
              <Map style={mapStyle}
                   onBoundsChange={this.onBoundsChange}
                   state={{
                     center: this.props.mapStore.center,
                     zoom: this.props.mapStore.zoom,
                   }}>
                <ZoomControl/>
                {this.createMarkers()}
              </Map>
            </div>
          </YMaps></Div>

        </Group>
      </>
    );
  }
}

decorate(LostPanel, {
  animals: observable,
});

export default observer(LostPanel);
