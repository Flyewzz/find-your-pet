import React, {useState, useEffect} from 'react';
import View from '@vkontakte/vkui/dist/components/View/View';
import '@vkontakte/vkui/dist/vkui.css';
import Tabbar from "@vkontakte/vkui/dist/components/Tabbar/Tabbar";
import TabbarItem from "@vkontakte/vkui/dist/components/TabbarItem/TabbarItem";
import Icon28Menu from '@vkontakte/icons/dist/28/menu';
import {Epic, Panel, PanelHeader} from "@vkontakte/vkui";
import MainPanel from "./panels/Main";

class App extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      activeStory: 'main'
    };
    this.onStoryChange = this.onStoryChange.bind(this);
  }

  onStoryChange(e) {
    this.setState({activeStory: e.currentTarget.dataset.story})
  }

  render() {
    return (
      <Epic activeStory={this.state.activeStory} tabbar={
        <Tabbar>
          <TabbarItem
            onClick={this.onStoryChange}
            selected={this.state.activeStory === 'main'}
            data-story="main"
            text="Главная"
          ><Icon28Menu/></TabbarItem>
          <TabbarItem
            onClick={this.onStoryChange}
            selected={this.state.activeStory === 'discover'}
            data-story="discover"
            text="Нашлись"
          ><Icon28Menu/></TabbarItem>
          <TabbarItem
            onClick={this.onStoryChange}
            selected={this.state.activeStory === 'messages'}
            data-story="messages"
            text="Потерялись"
          ><Icon28Menu/></TabbarItem>
          <TabbarItem
            onClick={this.onStoryChange}
            selected={this.state.activeStory === 'more'}
            data-story="more"
            text="Мои объявления"
          ><Icon28Menu/></TabbarItem>
        </Tabbar>
      }>
        <View id="main" activePanel="main">
          <Panel id="main">
            <MainPanel/>
          </Panel>
        </View>
        <View id="discover" activePanel="discover">
          <Panel id="discover">
            <PanelHeader>Поиск</PanelHeader>
          </Panel>
        </View>
        <View id="messages" activePanel="messages">
          <Panel id="messages">
            <PanelHeader>Сообщения</PanelHeader>
          </Panel>
        </View>
        <View id="more" activePanel="more">
          <Panel id="more">
            <PanelHeader>Ещё</PanelHeader>
          </Panel>
        </View>
      </Epic>
    );
  }
}

export default App;

