import {ModalPage, ModalPageHeader, ModalRoot, PanelHeaderButton} from "@vkontakte/vkui";
import FormLayout from "@vkontakte/vkui/dist/components/FormLayout/FormLayout";
import FormLayoutGroup from "@vkontakte/vkui/dist/components/FormLayoutGroup/FormLayoutGroup";
import SelectMimicry from "@vkontakte/vkui/dist/components/SelectMimicry/SelectMimicry";
import Radio from "@vkontakte/vkui/dist/components/Radio/Radio";
import Checkbox from "@vkontakte/vkui/dist/components/Checkbox/Checkbox";
import Input from "@vkontakte/vkui/dist/components/Input/Input";
import Icon24Cancel from '@vkontakte/icons/dist/24/cancel';
import Icon24Done from '@vkontakte/icons/dist/24/done';
import React from "react";

const MODAL_PAGE_FILTERS = 'filters';

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
          <ModalPageHeader
            left={
              <PanelHeaderButton onClick={props.onClose}>
                <Icon24Cancel/>
              </PanelHeaderButton>
            }
            right={<PanelHeaderButton onClick={props.onClose}>
              <Icon24Done/>
            </PanelHeaderButton>}
          >
            Фильтры
          </ModalPageHeader>
        }
      >
        <FormLayout>
          <SelectMimicry top="Город" placeholder="Выбрать город" disabled/>

          <FormLayoutGroup top="Пол">
            <Radio name="sex" value={0} defaultChecked>Любой</Radio>
            <Radio name="sex" value={1}>Мужской</Radio>
            <Radio name="sex" value={2}>Женский</Radio>
          </FormLayoutGroup>

          <SelectMimicry top="Школа" placeholder="Выбрать школу" disabled/>
          <SelectMimicry top="Университет" placeholder="Выбрать университет" disabled/>

          <FormLayoutGroup top="Дополнительно">
            <Checkbox>С фотографией</Checkbox>
            <Checkbox>Сейчас на сайте</Checkbox>
          </FormLayoutGroup>

          <FormLayoutGroup top="Работа">
            <Input placeholder="Место работы"/>
            <Input placeholder="Должность"/>
          </FormLayoutGroup>

          <FormLayoutGroup top="Дата рождения">
            <SelectMimicry placeholder="День рождения" disabled/>
            <SelectMimicry placeholder="Месяц рождения" disabled/>
            <SelectMimicry placeholder="Год рождения" disabled/>
          </FormLayoutGroup>
        </FormLayout>
      </ModalPage>
    </ModalRoot>
  );
};

export default SearchFilter;