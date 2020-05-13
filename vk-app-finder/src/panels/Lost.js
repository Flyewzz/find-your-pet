import React from "react";
import {Card, CardGrid, Div, Group, PanelHeader, Spinner, Footer} from "@vkontakte/vkui";
import AnimalCard from "../components/cards/AnimalCard";
import FilterLine from "../components/cards/FilterLine";
import {observer} from "mobx-react";
import './Lost.css';
import {Map, Placemark, YMaps, ZoomControl} from 'react-yandex-maps';
import Placeholder from "@vkontakte/vkui/dist/components/Placeholder/Placeholder";
import Icon56InfoOutline from '@vkontakte/icons/dist/56/info_outline';
import Icon28CancelOutline from "@vkontakte/icons/dist/28/cancel_outline";
import InfiniteScroll from 'react-infinite-scroll-component';


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
  }

  lostFilterStore = this.props.lostFilterStore;

  componentDidMount() {
    this.props.lostFilterStore.animals = undefined;
    this.filterChanged = true;
    this.props.lostFilterStore.fetch();
  }

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
                    key={animal.id} animal={animal} type={'lost'}/>
      </React.Fragment>
    );
  };

  onBoundsChange = e => {
    this.props.mapStore.center = e.get('target').getCenter();
    this.props.mapStore.zoom = e.get('target').getZoom();
  };

  onSearchInput = (e) => {
    this.props.lostFilterStore.fields.query = e.target.value;
    this.props.lostFilterStore.fetch();
  };

  fetchNext = () => {
    this.lostFilterStore.incPage();
    this.props.lostFilterStore.fetch();
  };

  render() {
    const animals = this.props.lostFilterStore.animals;
    const mapStyle = {
      display: this.props.mapStore.isMapView && animals ? undefined : 'none',
    };

    return (
      <>
        <PanelHeader>Потерялись</PanelHeader>
        <Group separator="hide">

          <FilterLine isMap={this.state.mapView}
                      onChange={this.onSearchInput}
                      query={this.props.lostFilterStore.fields.query}
                      filterStore={this.props.lostFilterStore}
                      changeView={this.changeView}
                      openFilters={this.props.openFilters}/>

          {animals === undefined && this.filterChanged &&
          <Spinner size="large" style={{marginTop: 20, color: "rgb(83, 118, 164)"}}/>
          }

          {!this.props.mapStore.isMapView && animals
          && <CardGrid>
            <InfiniteScroll dataLength={animals.length}
                            hasMore={this.lostFilterStore.hasMore}
                            next={this.fetchNext}
                            style={{overflow: 'hidden'}}
                            loader={<Spinner size="large"
                                             style={{
                                               marginTop: 20,
                                               color: "rgb(83, 118, 164)",
                                             }}/>}>
              {this.animalsToCards()}
            </InfiniteScroll>
          </CardGrid>}
          {!this.props.mapStore.isMapView && animals && !this.lostFilterStore.hasMore &&
          <Footer>Больше ничего не найдено</Footer>}
          {!this.props.mapStore.isMapView && (animals === null)
          && <Placeholder icon={<Icon56InfoOutline/>}>
            Ничего не найдено<br/>Попробуйте позже или измените фильтры
          </Placeholder>}

          <Div><YMaps>
            <div style={mapStyle} className={'map__container'}>
              <Icon28CancelOutline
                style={mapStyle}
                onClick={() => {
                  this.props.mapStore.isMapView = false;
                }}
                className={"cancel-icon"}
              />
              <Map style={{
                height: '490px',
                width: '100%',
              }}
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

export default observer(LostPanel);
