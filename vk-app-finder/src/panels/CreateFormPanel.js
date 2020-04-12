import React from "react";
import {
  PanelHeader, Input, FormLayoutGroup, FormLayout, Button, Select
} from "@vkontakte/vkui";
import {observer} from 'mobx-react';
import {PanelHeaderBack} from "@vkontakte/vkui/dist/es6";
import DateInput from "../components/forms/DateInput";
import AddressInput from "../components/forms/AddressInput";
import ImageLoader from "../components/forms/ImageLoader";
import './CreateFormPanel.css';
import Textarea from "@vkontakte/vkui/dist/components/Textarea/Textarea";
import FormLabel from "../components/forms/FormLabel";
import LostStore from "../stores/LostStore";
import config from '../config';

class CreateFormPanel extends React.Component {
  constructor(props) {
    super(props);
    this.lostStore = new LostStore();
  }

  onSubmit = () => {
    // this.lostStore.submit().then(
    //   (result) => {
    //     alert('res:' + result);
    //   },
    //   (error) => {
    //     alert('error:' + error);
    //   })
    const formData = new FormData();
    const fields = this.lostStore.form.fields;
    formData.append("type_id", fields.typeId.value);
    formData.append("author_id", fields.authorId.value);
    formData.append("picture", fields.picture.value);
    formData.append("sex", fields.sex.value);
    formData.append("breed", fields.breed.value);
    formData.append("description", fields.description.value);
    formData.append("latitude", fields.latitude.value);
    formData.append("longitude", fields.longitude.value);

    const request = new XMLHttpRequest();
    request.open("POST", config.baseUrl + 'lost');
    request.send(formData);
  };

  onAddressChange = (data) => {
    const longitude = 50; //data.data.geo_lat;
    const latitude = 35; //data.data.geo_lon;

    this.lostStore.form.fields.longitude.value = longitude;
    this.lostStore.form.fields.latitude.value = latitude;
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
                <option value={0} defaultChecked>Собака</option>
                <option value={1}>Кошка</option>
                <option value={2}>Другой</option>
              </Select>
            </FormLayoutGroup>
            <FormLayoutGroup className={'half-width'}>
              <FormLabel text={'Пол животного'}/>
              <Select onChange={this.onSexChange} value={fields.sex.value}>
                <option value={0}>Не определен</option>
                <option value={'m'}>Мужской</option>
                <option value={'f'}>Женской</option>
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
