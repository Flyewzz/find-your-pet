import React from "react";
import {
  PanelHeader, Input, FormLayoutGroup, FormLayout, Button, Select
} from "@vkontakte/vkui";
import {observer} from 'mobx-react';
import {PanelHeaderBack} from "@vkontakte/vkui/dist/es6";
import AddressInput from "../components/forms/AddressInput";
import ImageLoader from "../components/forms/ImageLoader";
import './CreateFormPanel.css';
import Textarea from "@vkontakte/vkui/dist/components/Textarea/Textarea";
import FormLabel from "../components/forms/FormLabel";
import LostStore from "../stores/LostStore";
import GeocodingService from "../services/GeocodingService";

class CreateFormPanel extends React.Component {
  constructor(props) {
    super(props);
    this.lostStore = new LostStore(props.userStore);
    this.geocodingService = new GeocodingService();
  }

  onSubmit = () => {
    this.lostStore.submit(
      (result) => {
        console.log('res:' + result);
      })
  };

  onAddressChange = (data) => {
    const address = data.value;
    this.geocodingService.getCoords(address).then(result => {
      const firstCandidate = result.candidates[0];
      // add some error here if candidate is undefined
      const location = firstCandidate.location;
      const longitude = location.x;
      const latitude = location.y;

      this.lostStore.form.fields.longitude.value = longitude;
      this.lostStore.form.fields.latitude.value = latitude;
    });
  };

  onPictureSet = (picture) => {
    this.lostStore.form.fields.picture.value = picture;
  };

  onTypeChange = (type) => {
    this.lostStore.form.fields.typeId.value = type.target.value;
  };

  onSexChange = (sex) => {
    this.lostStore.form.fields.sex.value = sex.target.value;
  };

  onBreedChange = (breed) => {
    this.lostStore.form.fields.breed.value = breed.target.value;
  };

  onDescriptionChange = (description) => {
    this.lostStore.form.fields.description.value = description.target.value;
  };

  render() {
    const fields = this.lostStore.form.fields;

    return (
      <>
        <PanelHeader left={<PanelHeaderBack onClick={this.props.toMain}/>}>
          Я потерял
        </PanelHeader>
        <FormLayout separator="hide">
          <AddressInput title={'Местро пропажи'}
                        onAddressChange={this.onAddressChange}
                        placeholder={'Введите адрес'}/>

          <div className={'selects-wrapper'}>
            <FormLayoutGroup className={'half-width'}>
              <FormLabel text={'Вид животного'}/>
              <Select onChange={this.onTypeChange} value={fields.typeId.value}>
                <option value={1} defaultChecked>Собака</option>
                <option value={2}>Кошка</option>
                <option value={3}>Другой</option>
              </Select>
            </FormLayoutGroup>
            <FormLayoutGroup className={'half-width'}>
              <FormLabel text={'Пол животного'}/>
              <Select onChange={this.onSexChange} value={fields.sex.value}>
                <option value={'n/a'}>Не определен</option>
                <option value={'m'}>Мужской</option>
                <option value={'f'}>Женский</option>
              </Select>
            </FormLayoutGroup>
          </div>

          <ImageLoader onPictureSet={this.onPictureSet}/>
          <FormLayoutGroup top={'Порода'}>
            <Input defaultValue={fields.breed.value === '' ? undefined : fields.breed.value}
                   onChange={this.onBreedChange}
                   placeholder={'Введите породу'}/>
          </FormLayoutGroup>
          <Textarea top="Описание"
                    onChange={this.onDescriptionChange}
                    defaultValue={fields.description.value === '' ? undefined : fields.description.value}
                    placeholder="Напишите все, что может помочь узнать вашего питомца"/>
          <Button onClick={this.onSubmit} className={'create-form__submit'} size={'l'}>Завершить</Button>
        </FormLayout>
      </>
    );
  }
}

export default observer(CreateFormPanel);
