import {ModalPage, ModalRoot, FormLayout, FormLayoutGroup, Radio, Input, Select} from "@vkontakte/vkui";
import React from "react";
import FilterHeader from "../components/forms/FilterHeader";
import AddressInput from "../components/forms/AddressInput";
import './SearchFilter.css'
import FormLabel from "../components/forms/FormLabel";
import {observer} from "mobx-react";
import {decorate, observable} from "mobx";

const MODAL_PAGE_FILTERS = 'filters';

class SearchFilter extends React.Component {
  constructor(props) {
    super(props);
    this.sex = props.filterStore.fields.sex;
    this.type = props.filterStore.fields.type;
    this.breed = props.filterStore.fields.breed;
  }

  onSexChange = (e) => {
    this.sex = e.target.value;
  };

  onTypeChange = (e) => {
    this.type = e.target.value;
  };

  onBreedChange = (e) => {
    this.breed = e.target.value;
  };

  onAccept = () => {
    const onFieldChange = this.props.filterStore.onFieldChange;
    onFieldChange('sex', this.sex);
    onFieldChange('type', this.type);
    onFieldChange('breed', this.breed);
    this.props.filterStore.fetch();
    this.props.onClose();
  };

  render() {
    const props = this.props;
    return (
      <ModalRoot
        activeModal={props.activeModal}
        onClose={props.onClose}>
        <ModalPage
          id={MODAL_PAGE_FILTERS}
          onClose={props.onClose}
          header={<FilterHeader onDone={this.onAccept} onClose={props.onClose}/>}>
          <FormLayout>
            <AddressInput title={'Местро пропажи'} placeholder={'Место пропажи'}/>
            <div className={'selects-wrapper'}>
              <FormLayoutGroup className={'half-width'}>
                <FormLabel text={'Вид животного'}/>
                <Select onChange={this.onTypeChange}
                        value={this.type}>
                  <option value={0}>Любой</option>
                  <option value={1}>Собака</option>
                  <option value={2}>Кошка</option>
                  <option value={3}>Другой</option>
                </Select>
              </FormLayoutGroup>
              <FormLayoutGroup className={'half-width'}>
                <FormLabel text={'Пол животного'}/>
                <Select onChange={this.onSexChange}
                        value={this.sex}>
                  <option value={0}>Любой</option>
                  <option value={'n/a'}>Не определен</option>
                  <option value={'m'}>Мужской</option>
                  <option value={'f'}>Женский</option>
                </Select>
              </FormLayoutGroup>
            </div>
            <FormLayoutGroup top={'Порода'}>
              <Input onChange={this.onBreedChange}
                     value={this.breed}
                     placeholder={'Введите породу'}/>
            </FormLayoutGroup>
          </FormLayout>
        </ModalPage>
      </ModalRoot>
    );
  }
}

decorate(SearchFilter, {
  sex: observable,
  type: observable,
  breed: observable,
});

export default observer(SearchFilter);