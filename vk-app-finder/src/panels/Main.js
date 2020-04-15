import React from 'react';
import {Button, Div, PanelHeader} from "@vkontakte/vkui";
import {FixedLayout, PanelHeaderBack} from "@vkontakte/vkui/dist/es6";
import './Main.css';

const MainPanel = (props) => {
  return (
    <>
      <PanelHeader>Главная</PanelHeader>
      <FixedLayout className={'main__wrapper'}>
        <Div className={'main__welcome-text'}>
          Здесь вам помогут найти<br/>
          Вашего любимца<br/>
          Или Вы можете помочь другим
        </Div>
        <Div className={'main__buttons-group'}>
          <Button onClick={props.toCreateLostForm} className={'main__button'}>Я потерял</Button>
          <Button className={'main__button'}>Я нашел</Button>
        </Div>
      </FixedLayout>
    </>
  )
};

export default MainPanel;