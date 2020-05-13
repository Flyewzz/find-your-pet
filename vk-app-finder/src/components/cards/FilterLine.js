import React from 'react';
import {Search, Div, Button} from "@vkontakte/vkui";
import Icon28SlidersOutline from '@vkontakte/icons/dist/28/sliders_outline';
import Icon28PlaceOutline from '@vkontakte/icons/dist/28/place_outline';
import Icon28CancelOutline from "@vkontakte/icons/dist/28/cancel_outline";
import Icon24Cancel from '@vkontakte/icons/dist/24/cancel';
import config from '../../config';
import './FilterLine.css'
import {observer} from "mobx-react";

const ActiveFilter = (props) => {
  return (
    <Button className={'active-filter'}
            onClick={props.onClick}
            mode={'outline'}
            after={<Icon24Cancel/>}>
      {props.child}
    </Button>
  );
};

const getActiveFilters = (store) => {
  const activeFilters = [];

  if (store.fields.type !== '0') {
    const onClick = () => {
      store.fields.type = '0';
      store.clearPage();
      store.fetch();
    };
    activeFilters.push(
      <ActiveFilter child={config.types[+store.fields.type - 1]}
                    onClick={onClick}/>
    );
  }
  if (store.fields.sex !== '0') {
    const onClick = () => {
      store.fields.sex = '0';
      store.clearPage();
      store.fetch();
    };
    const sex = store.fields.sex;
    const text = sex === 'm' ? 'Мужской' : sex === 'f' ? 'Женский' : 'Пол не определен';
    activeFilters.push(<ActiveFilter child={text} onClick={onClick}/>);
  }
  if (store.fields.breed !== '') {
    const onClick = () => {
      store.fields.breed = '';
      store.clearPage();
      store.fetch();
    };
    activeFilters.push(
      <ActiveFilter child={store.fields.breed}
                    onClick={onClick}/>
    );
  }

  return activeFilters;
};

const FilterLine = (props) => {
  const className = props.isMap ? 'checked' : '';

  return (
    <>
      <Div className={'filter__wrapper'}>
        {props.isMap &&
        <Icon28CancelOutline
          onClick={() => {
          }}
          className={"cancel-icon"}
        />}
        <Search onChange={props.onChange}
                value={props.query}
                type="text"
                placeholder="Найти..."/>
        <Icon28PlaceOutline onClick={props.changeView} className={'filter__map ' + className}/>
        <Icon28SlidersOutline onClick={props.openFilters} className={'filter__button'}/>
      </Div>
      <Div style={{marginTop: 0, paddingTop: 0}}>
        {getActiveFilters(props.filterStore)}
      </Div>
    </>
  );
};

export default observer(FilterLine);