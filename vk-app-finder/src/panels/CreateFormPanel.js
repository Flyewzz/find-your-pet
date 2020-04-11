import React from "react";
import {
  Div, PanelHeader, Input, Separator, Header,
  Radio, FormLayoutGroup, FormLayout, Button, Select
} from "@vkontakte/vkui";
import {PanelHeaderBack} from "@vkontakte/vkui/dist/es6";
import DG from '2gis-maps';
import DateInput from "../components/forms/DateInput";
import AddressInput from "../components/forms/AddressInput";
import ImageLoader from "../components/forms/ImageLoader";
import './CreateFormPanel.css'
import Textarea from "@vkontakte/vkui/dist/components/Textarea/Textarea";
import FormLabel from "../components/forms/FormLabel";

class CreateFormPanel extends React.Component {
  constructor(props) {
    super(props);
  }

  componentDidMount() {
  }

  render() {
    return (
      <>
        <PanelHeader
          left={<PanelHeaderBack onClick={this.props.toMain}/>}>
          Я потерял
        </PanelHeader>
        <FormLayout separator="hide">
          <AddressInput title={'Местро пропажи'} placeholder={'Введите адрес'}/>
          <DateInput/>

          <div className={'selects-wrapper'}>
            <FormLayoutGroup className={'half-width'}>
              <FormLabel text={'Вид животного'}/>
              <Select>
                <option value={0} defaultChecked>Собака</option>
                <option value={1}>Кошка</option>
                <option value={2}>Другой</option>
              </Select>
            </FormLayoutGroup>
            <FormLayoutGroup className={'half-width'}>
              <FormLabel text={'Пол животного'}/>
              <Select>
                <option value={0} defaultChecked>Мужской</option>
                <option value={1}>Женской</option>
              </Select>
            </FormLayoutGroup>
          </div>

          <ImageLoader/>
          <FormLayoutGroup top={'Порода'}>
            <Input placeholder={'Введите породу'}/>
          </FormLayoutGroup>
          <Textarea top="Описание"
                    placeholder="Напишите все, что может помочь узнать вашего питомца"/>
          <Button className={'create-form__submit'} size={'l'}>Завершить</Button>
        </FormLayout>
      </>
    );
  }
}

export default CreateFormPanel;
