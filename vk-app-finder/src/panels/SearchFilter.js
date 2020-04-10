import {ModalPage, ModalRoot, FormLayout, FormLayoutGroup, Radio, Input} from "@vkontakte/vkui";
import React from "react";
import FilterHeader from "../components/filter/FilterHeader";
import DateInput from "../components/filter/DateInput";
import AddressInput from "../components/filter/AddressInput";
import './SearchFilter.css'

const MODAL_PAGE_FILTERS = 'filters';

const SearchFilter = (props) => {
  return (
    <ModalRoot
      activeModal={props.activeModal}
      onClose={props.onClose}>
      <ModalPage
        id={MODAL_PAGE_FILTERS}
        onClose={props.onClose}
        header={<FilterHeader onDone={props.onClose} onClose={props.onClose}/>}>
        <FormLayout>
          <AddressInput title={'Местро пропажи'} placeholder={'Место пропажи'}/>
          <DateInput/>
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