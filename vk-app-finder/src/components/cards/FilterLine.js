import React from 'react';
import {Search, Div} from "@vkontakte/vkui";
import Icon28SlidersOutline from '@vkontakte/icons/dist/28/sliders_outline';
import Icon28PlaceOutline from '@vkontakte/icons/dist/28/place_outline';
import './FilterLine.css'

const FilterLine = (props) => {
  const className = props.isMap? 'checked': '';

  return (
    <Div className={'filter__wrapper'}>
      <Search type="text" placeholder="Найти..." />
      <Icon28PlaceOutline onClick={props.changeView} className={'filter__map ' + className}/>
      <Icon28SlidersOutline onClick={props.openFilters}/>
    </Div>
  );
};

export default FilterLine;