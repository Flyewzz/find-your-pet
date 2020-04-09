import React from "react";
import {Card, CardGrid, Div, Group, PanelHeader} from "@vkontakte/vkui";
import {PanelHeaderBack} from "@vkontakte/vkui/dist/es6";
import AnimalCard from "../components/cards/AnimalCard";
import FilterLine from "../components/cards/FilterLine";
import DG from '2gis-maps';

const getDogs = () => {
  const dogs = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
  const animal_test = {
    type: 'Кот',
    breed: 'Сфинкс',
    place: 'Раменское, ул. Свободы',
    date: '22 ноября, 14:51',
  };
  return dogs.map((animal, index) =>
    <>
      {!(index % 2) && <Card size="l" styles={{height: 0}}/>}
      <AnimalCard key={index} animal={animal_test}/>
    </>
  );
};

class LostPanel extends React.Component {
  constructor(props) {
    super(props);
    this.state = {mapView: false};
    this.changeView = () => {
      const current = this.state.mapView;
      this.setState({mapView: !current});
    };
  }

  componentDidMount() {
    DG.map('map', {
      center: [54.98, 82.89],
      zoom: 13,
    })
  }

  render() {
    const mapStyle = {
      display: this.state.mapView? undefined: 'none',
      height: '500px'
    };
    console.log(mapStyle.display);

    return (
      <>
        <PanelHeader left={<PanelHeaderBack/>}>Потерялись</PanelHeader>
        <Group separator="hide">
          <FilterLine changeView={this.changeView}
                      openFilters={this.props.openFilters}/>
          {!this.state.mapView && <CardGrid>{getDogs()}</CardGrid>}
          <Div>
            <div style={mapStyle} id={'map'}/>
          </Div>
        </Group>
      </>
    );
  }
}

export default LostPanel;
