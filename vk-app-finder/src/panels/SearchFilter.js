import {Div, ModalPage, ModalRoot, Select} from "@vkontakte/vkui";
import {ReactDadata} from "react-dadata";
import FormLayout from "@vkontakte/vkui/dist/components/FormLayout/FormLayout";
import FormLayoutGroup from "@vkontakte/vkui/dist/components/FormLayoutGroup/FormLayoutGroup";
import Radio from "@vkontakte/vkui/dist/components/Radio/Radio";
import Input from "@vkontakte/vkui/dist/components/Input/Input";
import React from "react";
import FilterHeader from "../components/filter/FilterHeader";
import {Search} from "@vkontakte/vkui/dist/es6";
import './SearchFilter.css'

const MODAL_PAGE_FILTERS = 'filters';

const getDaysCount = (date) => {
  return new Date(date.getFullYear(), date.getMonth() + 1, 0).getDate();
};

const daysOptions = (daysNumber) => {
  const days = [];
  for (let i = 1; i <= daysNumber; i++) {
    days.push(i);
  }

  return days.map((day) => (
    <option value={day}>{day}</option>
  ));
};

const months = [
  'Январь', 'Февраль', 'Март', 'Апрель', 'Май', 'Июнь', 'Июль',
  'Август', 'Сентябрь', 'Октябрь', 'Ноябрь', 'Декабрь'
];

const monthsOptions = () => {
  return months.map((month, index) => (
    <option value={index}>{month}</option>
  ));
};

const yearsOptions = () => {
  const years = [];
  const nowDate = new Date();

  for (let i = nowDate.getFullYear(); i > 1950; i--) {
    years.push(i);
  }

  return years.map(value => (
    <option value={value}>{value}</option>
  ));
};

const SearchFilter = (props) => {

  return (
    <ModalRoot
      activeModal={props.activeModal}
      onClose={props.onClose}
    >
      <ModalPage
        id={MODAL_PAGE_FILTERS}
        onClose={props.onClose}
        header={
          <FilterHeader onDone={props.onClose} onClose={props.onClose}/>
        }
      >
        <FormLayout>
          <Div>
            <ReactDadata token={'API_TOKEN'}
                         placeholder={'Место пропажи'}/>
          </Div>
          <FormLayoutGroup className={'filter__date'} top={'Дата пропажи'}>
            <Select className={'filter__date__day'} placeholder={'День'}>
              {daysOptions(31)}
            </Select>
            <Select className={'filter__date__month'} placeholder={'Месяц'}>
              {monthsOptions()}
            </Select>
            <Select className={'filter__date__year'} placeholder={'Год'}>
              {yearsOptions()}
            </Select>
          </FormLayoutGroup>

          <FormLayoutGroup top={'Порода'}>
            <Input placeholder={'Введите породу'}/>
          </FormLayoutGroup>

          <FormLayoutGroup top="Пол">
            <Radio name="sex" value={0} defaultChecked>Любой</Radio>
            <Radio name="sex" value={1}>Мужской</Radio>
            <Radio name="sex" value={2}>Женский</Radio>
          </FormLayoutGroup>
        </FormLayout>
      </ModalPage>
    </ModalRoot>
  );
};

export default SearchFilter;