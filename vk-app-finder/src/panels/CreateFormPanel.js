import React from "react";
import {
  PanelHeader, FormLayoutGroup, FormLayout, Button, Select, FormStatus
} from "@vkontakte/vkui";
import {observer} from 'mobx-react';
import {PanelHeaderBack, Spinner} from "@vkontakte/vkui/dist/es6";
import AddressInput from "../components/forms/AddressInput";
import ImageLoader from "../components/forms/ImageLoader";
import './CreateFormPanel.css';
import FormLabel from "../components/forms/FormLabel";
import LostStore from "../stores/LostStore";
import GeocodingService from "../services/GeocodingService";
import Div from "@vkontakte/vkui/dist/components/Div/Div";
import BreedDescriptionPanel from "../components/forms/BreedDescriptionPanel";
import {decorate, observable, runInAction} from "mobx";
import BreedService from "../services/BreedService";

const FirstFormPanel = (props) => {
  return (
    <>
      <AddressInput title={'Место пропажи*'}
                    defaultValue={props.address}
                    onAddressChange={props.onAddressChange}
                    placeholder={'Введите адрес'}/>

      <div className={'selects-wrapper'}>
        <FormLayoutGroup className={'half-width'}>
          <FormLabel text={'Вид животного*'}/>
          <Select onChange={props.onTypeChange} value={props.fields.typeId.value}>
            <option value={1} defaultChecked>Собака</option>
            <option value={2}>Кошка</option>
            <option value={3}>Другой</option>
          </Select>
        </FormLayoutGroup>
        <FormLayoutGroup className={'half-width'}>
          <FormLabel text={'Пол животного*'}/>
          <Select onChange={props.onSexChange} value={props.fields.sex.value}>
            <option value={'n/a'}>Не определен</option>
            <option value={'m'}>Мужской</option>
            <option value={'f'}>Женский</option>
          </Select>
        </FormLayoutGroup>
      </div>
      <Div style={{paddingLeft: 0}}>
        <Button className={'create-form__submit'}
                disabled={props.addressLoading || props.fields.longitude.value === ''}
                onClick={props.toNext}
                size={'l'}>
          Далее
        </Button>
      </Div>
    </>
  );
};

class CreateFormPanel extends React.Component {
  constructor(props) {
    super(props);
    this.lostStore = new LostStore(props.userStore);
    this.geocodingService = new GeocodingService();
    this.breedService = new BreedService();
    this.state = {stage: 0}
  }

  breeds = undefined;
  addressLoading = false;

  onSubmit = () => {
    if (!this.lostStore.check()) {
      return;
    }

    this.props.openPopout();
    this.lostStore.submit(
      (result) => {
        console.log('res:' + result);
        this.props.toProfile();
        this.props.closePopout();
      })
  };

  onAddressChange = (data) => {
    this.lostStore.form.fields.address.value = data.value;
    this.addressLoading = true;
    this.geocodingService.getCoords(data.value).then(result => {
      const firstCandidate = result.candidates[0];
      // add some error here if candidate is undefined
      const location = firstCandidate.location;
      const longitude = location.x;
      const latitude = location.y;
      this.lostStore.form.fields.longitude.value = longitude;
      this.lostStore.form.fields.latitude.value = latitude;
      this.addressLoading = false;
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
    runInAction(() => {
      this.lostStore.form.fields.breed.value = breed.target.value;
    });
  };

  onBreedChoose = (breed) => {
    runInAction(() => {
      this.lostStore.form.fields.breed.value = breed;
    });
  };

  onDescriptionChange = (description) => {
    this.lostStore.form.fields.description.value = description.target.value;
  };

  back = () => {
    if (this.state.stage > 0) {
      this.setState({stage: this.state.stage - 1});
    } else {
      this.props.toMain();
    }
  };

  submitFirst = () => {
    const form = this.lostStore.form;
    if (form.fields.latitude.value === '') {
      form.meta.error = 'Необходимо ввести корректный адрес пропажи.';
      form.meta.isValid = false;
    } else {
      form.meta.error = '';
      form.meta.isValid = true;
      this.setState({stage: 1});
    }
  };

  submitSecond = () => {
    const image = this.lostStore.form.fields.picture.value;
    const animalType = this.lostStore.form.fields.typeId.value;
    if (image !== null && animalType === 1)  {
      this.breeds = undefined;
      this.breedService.getBreeds(image).then(
        (result) => {
          this.breeds = result;
        }, (error) => {
          console.log(error);
          this.breeds = 'error';
        }
      )
    } else {
      this.breeds = null;
    }
    this.setState({stage: 2});
  };

  render() {
    const fields = this.lostStore.form.fields;

    return (
      <>
        <PanelHeader left={<PanelHeaderBack onClick={this.back}/>}>
          Я потерял
        </PanelHeader>
        <FormLayout separator="hide">
          {this.state.stage === 0 && !this.lostStore.form.meta.isValid &&
          <FormStatus mode={'error'}>
            {this.lostStore.form.meta.error}
          </FormStatus>}
          {this.state.stage === 0 && this.addressLoading &&
          <FormStatus>
            <Spinner/>
            <div style={{width: '100%', textAlign: 'center'}}>
              Не уходите, мы проверяем адрес
            </div>
          </FormStatus>}
          {this.state.stage === 0
          && <FirstFormPanel onTypeChange={this.onTypeChange}
                             fields={fields}
                             address={fields.address.value}
                             addressLoading={this.addressLoading}
                             toNext={this.submitFirst}
                             onSexChange={this.onSexChange}
                             onAddressChange={this.onAddressChange}/>}

          {this.state.stage === 1 && <>
            <ImageLoader onPictureSet={this.onPictureSet}/>
            <Div style={{paddingLeft: 0}}>
              <Button onClick={this.submitSecond} size={'l'}>
                Далее
              </Button>
            </Div>
          </>}

          {this.state.stage === 2 && this.breeds === 'error' &&
          <FormStatus>
            {'Нам не удалось распознать породу животного'}
          </FormStatus>}
          {this.state.stage === 2 && this.breeds !== 'error' && !this.lostStore.form.meta.isValid &&
          <FormStatus mode={'error'}>
            {this.lostStore.form.meta.error}
          </FormStatus>}

          {this.state.stage === 2 && this.breeds === undefined &&
          <FormStatus>
            <Spinner/>
            <div style={{textAlign: 'center', width: '100%'}}>
              {'Пара секунд, мы пытаемся распознать породу животного'}
            </div>
          </FormStatus>}

          {this.state.stage === 2 && <>
            <BreedDescriptionPanel onBreedChange={this.onBreedChange}
                                   onBreedChoose={this.onBreedChoose}
                                   onDescriptionChange={this.onDescriptionChange}
                                   onSubmit={this.onSubmit}
                                   breeds={this.breeds}
                                   fields={fields}/>
          </>}

        </FormLayout>
      </>
    );
  }
}

decorate(CreateFormPanel, {
  breeds: observable,
  addressLoading: observable,
});

export default observer(CreateFormPanel);
