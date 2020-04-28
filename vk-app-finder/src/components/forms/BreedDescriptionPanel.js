import {Button, FormLayout, FormLayoutGroup, FormStatus, InfoRow, Input} from "@vkontakte/vkui";
import React from "react";
import Textarea from "@vkontakte/vkui/dist/components/Textarea/Textarea";
import Div from "@vkontakte/vkui/dist/components/Div/Div";
import {observer} from "mobx-react";
import {decorate, observable} from "mobx";
import {Spinner} from "@vkontakte/vkui/dist/es6";

class BreedDescriptionPanel extends React.Component {
  breeds = this.props.breeds;

  breedsToButtons = () => {
    const currentBreed = this.props.fields.breed.value;
    return this.breeds.map(breed => {
      const mode = breed === currentBreed ? 'primary' : 'outline';
      const onClick = () => {
        this.props.fields.breed.value = breed;
      };

      return (
        <Button key={breed} style={{ cursor: 'pointer' }} onClick={onClick} mode={mode}>
          {breed}
        </Button>
      );
    });
  };

  render() {
    const props = this.props;
    console.log(this.breeds);
    return (
      <FormLayout>
        {this.breeds !== undefined && this.breeds !== 'error' && this.breeds !== null &&
        <FormLayoutGroup top={'Возможно, это порода животного?'}>
          {this.breedsToButtons()}
        </FormLayoutGroup>}

        <FormLayoutGroup top={'Порода'}>
          <Input value={props.fields.breed.value === '' ? undefined : props.fields.breed.value}
                 onChange={props.onBreedChange}
                 placeholder={'Введите породу'}/>
        </FormLayoutGroup>
        <Textarea top="Описание"
                  onChange={props.onDescriptionChange}
                  defaultValue={props.fields.description.value === '' ?
                    undefined : props.fields.description.value}
                  placeholder="Напишите все, что может помочь узнать вашего питомца"/>
        <Div style={{paddingLeft: 0}}>
          <Button onClick={props.onSubmit} style={{ cursor: 'pointer' }} className={'create-form__submit'} size={'l'}>
            Завершить
          </Button>
        </Div>
      </FormLayout>
    );
  }
}

decorate(BreedDescriptionPanel, {
  breeds: observable,
});

export default observer(BreedDescriptionPanel);