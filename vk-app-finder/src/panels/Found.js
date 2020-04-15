import React from "react";
import {Card, CardGrid, Div, Group, PanelHeader} from "@vkontakte/vkui";
import AnimalCard from "../components/cards/AnimalCard";
import FilterLine from "../components/cards/FilterLine";
import {decorate, observable} from "mobx";
import {observer} from "mobx-react";
import './Lost.css';
import {Map, Placemark, YMaps, ZoomControl} from 'react-yandex-maps';
import Placeholder from "@vkontakte/vkui/dist/components/Placeholder/Placeholder";
import Icon56InfoOutline from '@vkontakte/icons/dist/56/info_outline';
import GeocodingService from "../services/GeocodingService";

class FoundPanel extends React.Component {
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

    this.geocodingService = new GeocodingService();
    this.props.lostFilterStore.onFetch = this.computeAddresses;
  }

  addresses = [];

  componentDidMount() {
    this.props.lostFilterStore.fetch();
  }

  computeAddresses = () => {
    const store = this.props.lostFilterStore;
    this.addresses = store.animals === null ? [] : store.animals.map(() => '');

    if (store.animals !== null) {
      store.animals.forEach((value, index) => {
        const {longitude, latitude} = value;
        this.geocodingService.addressByCoords(longitude, latitude).then(
          result => this.updateAddress(index, result.address)
        );
      });
    }
  };

  updateAddress = (index, result) => {
    const city = result.City === '' ? result.District : result.City;
    const address = result.MetroArea === '' ? result.Address : result.MetroArea;
    this.addresses[index] = city + (address === '' ? '' : ', ' + address);
  };

  createMarkers = () => {
    return this.props.lostFilterStore.animals.map(value =>
      <Placemark onClick={() => this.props.toLost(value.id)}
                 geometry={[value.latitude, value.longitude]}/>
    );
  };

  animalsToCards = () => {
    return this.props.lostFilterStore.animals.map((animal, index) =>
      <React.Fragment key={1}>
        {!(index % 2) && <Card key={-animal.id} size="l" styles={{height: 0}}/>}
        <AnimalCard onClick={() => this.props.toLost(animal.id)}
                    address={this.addresses[index]}
                    key={animal.id} animal={animal}/>
      </React.Fragment>
    );
  };

  onBoundsChange = e => {
    this.props.mapStore.center = e.get('target').getCenter();
    this.props.mapStore.zoom = e.get('target').getZoom();
  };

  render() {
    const animals = this.props.lostFilterStore.animals;
    const mapStyle = {
      display: this.props.mapStore.isMapView ? undefined : 'none',
      height: '490px',
      width: '100%',
    };

    return (
      <>
        <PanelHeader>Потерялись</PanelHeader>
        <Group separator="hide">
          <FilterLine isMap={this.state.mapView}
                      filterStore={this.props.lostFilterStore}
                      changeView={this.changeView}
                      openFilters={this.props.openFilters}/>

          {!this.props.mapStore.isMapView && animals
          && <CardGrid>{this.animalsToCards()}</CardGrid>}
          {!this.props.mapStore.isMapView && !animals
          && <Placeholder icon={<Icon56InfoOutline/>}>
            Ничего не найдено<br/>Попробуйте позже или измените фильтры
          </Placeholder>}

          <Div><YMaps>
            <div>
              <Map style={mapStyle}
                   onBoundsChange={this.onBoundsChange}
                   state={{
                     center: this.props.mapStore.center,
                     zoom: this.props.mapStore.zoom,
                   }}>
                <ZoomControl/>
                {animals && this.createMarkers()}
              </Map>
            </div>
          </YMaps></Div>

        </Group>
      </>
    );
  }
}

decorate(FoundPanel, {
  addresses: observable,
});

export default observer(FoundPanel);
