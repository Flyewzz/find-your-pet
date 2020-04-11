import {Select, FormLayoutGroup, Div} from "@vkontakte/vkui";
import React from "react";
import FormLabel from "./FormLabel";

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

const DateInput = (props) => {
  const nowDate = new Date();

  return (
    <>
      <FormLabel text={'Дата пропажи'}/>
      <FormLayoutGroup className={'filter__date'} top="Дата пропажи">
        <Select defaultValue={nowDate.getDate()} className={'filter__date__day'}>
          {daysOptions(getDaysCount(nowDate))}
        </Select>
        <Select defaultValue={nowDate.getMonth()} className={'filter__date__month'}>
          {monthsOptions()}
        </Select>
        <Select defaultValue={nowDate.getFullYear()} className={'filter__date__year'}>
          {yearsOptions()}
        </Select>
      </FormLayoutGroup>
    </>
  );
};

export default DateInput;