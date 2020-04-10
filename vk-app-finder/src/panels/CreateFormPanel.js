import React from "react";
import {
  Div, PanelHeader, Input, Separator, Header,
  Radio, FormLayoutGroup, FormLayout, Button
} from "@vkontakte/vkui";
import {PanelHeaderBack} from "@vkontakte/vkui/dist/es6";
import DG from '2gis-maps';
import DateInput from "../components/filter/DateInput";
import AddressInput from "../components/filter/AddressInput";
import ImageLoader from "../components/filter/ImageLoader";
import './CreateFormPanel.css'
import Textarea from "@vkontakte/vkui/dist/components/Textarea/Textarea";

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

          <FormLayoutGroup top={'Вид животного'}>
            <Radio name="type" value={0} defaultChecked>Собака</Radio>
            <Radio name="type" value={1}>Кошка</Radio>
            <Radio name="type" value={2}>Другой</Radio>
          </FormLayoutGroup>
          <FormLayoutGroup top={'Пол животного'}>
            <Radio name="sex" value={0} defaultChecked>Любой</Radio>
            <Radio name="sex" value={1}>Мужской</Radio>
            <Radio name="sex" value={2}>Женский</Radio>
          </FormLayoutGroup>
          <ImageLoader/>
          <FormLayoutGroup top={'Порода'}>
            <Input placeholder={'Введите породу'}/>
          </FormLayoutGroup>
          <Textarea top="Описание"
                    placeholder="Напишите все, что может помочь узнать вашего питомца" />
          <Button size={'l'}>Завершить</Button>
        </FormLayout>
      </>
    );
  }
}

export default CreateFormPanel;
