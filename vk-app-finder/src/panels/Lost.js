import React from "react";
import {
  Card,
  CardGrid,
  Group, Panel, View,
  PanelHeader,
  ModalPage,
  ModalRoot,
  ModalPageHeader,
  PanelHeaderButton,
} from "@vkontakte/vkui";
import {PanelHeaderBack} from "@vkontakte/vkui/dist/es6";
import AnimalCard from "../components/cards/AnimalCard";
import FilterLine from "../components/cards/FilterLine";

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

const MODAL_PAGE_FILTERS = 'filters';

class LostPanel extends React.Component {
  render() {
    return (
      <>
        <PanelHeader left={<PanelHeaderBack/>}>Потерялись</PanelHeader>
        <Group separator="hide">
          <FilterLine openFilters={this.props.openFilters}/>
          <CardGrid>{getDogs()}</CardGrid>
        </Group>
      </>
    );
  }
}

export default LostPanel;
