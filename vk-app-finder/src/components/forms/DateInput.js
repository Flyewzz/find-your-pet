import {Select, FormLayoutGroup} from "@vkontakte/vkui";
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

const DateInput = () => {
  const nowDate = new Date();

  return (
    <>
      <FormLabel text={'Дата пропажи'}/>
      <FormLayoutGroup className={'filter__date'} top="Дата пропажи">
        <Select defaultValue={-1} className={'filter__date__day'}>
          <option value={-1}>День</option>
          {daysOptions(getDaysCount(nowDate))}
        </Select>
        <Select defaultValue={-1} className={'filter__date__month'}>
          <option value={-1}>Месяц</option>
          {monthsOptions()}
        </Select>
        <Select defaultValue={-1} className={'filter__date__year'}>
          <option value={-1}>Год</option>
          {yearsOptions()}
        </Select>
      </FormLayoutGroup>
    </>
  );
};

export default DateInput;