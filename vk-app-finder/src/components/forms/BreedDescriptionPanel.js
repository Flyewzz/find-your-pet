import {Button, FormLayout, FormLayoutGroup, FormStatus, InfoRow, Input} from "@vkontakte/vkui";
import React from "react";
import Textarea from "@vkontakte/vkui/dist/components/Textarea/Textarea";
import Div from "@vkontakte/vkui/dist/components/Div/Div";
import {observer} from "mobx-react";

class BreedDescriptionPanel extends React.Component {
  breedsToButtons = () => {
    const currentBreed = this.props.fields.breed.value;
    const breeds = this.props.breeds;

    return breeds.map(breed => {
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
    const breeds = this.props.breeds;

    return (
      <FormLayout>
        {breeds !== undefined && breeds !== 'error' && breeds !== null &&
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

export default observer(BreedDescriptionPanel);