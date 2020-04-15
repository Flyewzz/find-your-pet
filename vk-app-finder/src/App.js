import React from 'react';
import View from '@vkontakte/vkui/dist/components/View/View';
import '@vkontakte/vkui/dist/vkui.css';
import Tabbar from "@vkontakte/vkui/dist/components/Tabbar/Tabbar";
import TabbarItem from "@vkontakte/vkui/dist/components/TabbarItem/TabbarItem";
import Icon28Menu from '@vkontakte/icons/dist/28/menu';
import Icon28User from '@vkontakte/icons/dist/28/user';
import Icon28HomeOutline from '@vkontakte/icons/dist/28/home_outline';
import Icon28ListCheckOutline from '@vkontakte/icons/dist/28/list_check_outline';
import {Epic, Panel, PanelHeader} from "@vkontakte/vkui";
import MainPanel from "./panels/Main";
import LostPanel from "./panels/Lost";
import SearchFilter from "./panels/SearchFilter";
import CreateFormPanel from "./panels/CreateFormPanel";
import LostAnimalPanel from "./panels/LostAnimalPanel";
import MapStateStore from "./stores/MapStateStore";
import UserStore from "./stores/UserStore";
import ProfilePanel from "./panels/Profile";
import {Alert} from "@vkontakte/vkui";

class App extends React.Component {
  constructor(props) {
    super(props);

    this.state = {
      activeModal: null,
      modalHistory: [],
      activeStory: 'more',
      mainPanel: 'new_lost',
      lostPanel: 'losts',
      profilePanel: 'more',

      popout: null,
    };
    this.onStoryChange = this.onStoryChange.bind(this);

    this.modalBack = () => {
      this.setActiveModal(this.state.modalHistory[this.state.modalHistory.length - 2]);
    };
    this.mapStore = new MapStateStore();
    this.userStore = new UserStore();
    this.openFilters = () => {
      this.setActiveModal('filters');
    };
  }

  setActiveModal(activeModal) {
    activeModal = activeModal || null;
    let modalHistory = this.state.modalHistory ? [...this.state.modalHistory] : [];

    if (activeModal === null) {
      modalHistory = [];
    } else if (modalHistory.indexOf(activeModal) !== -1) {
      modalHistory = modalHistory.splice(0, modalHistory.indexOf(activeModal) + 1);
    } else {
      modalHistory.push(activeModal);
    }

    this.setState({
      activeModal,
      modalHistory,
    });
  };

  onStoryChange(e) {
    this.setState({activeStory: e.currentTarget.dataset.story})
  }

  openDestructive = (onAccept) => {
    console.log('открыть попап кликнуто');
    this.setState({ popout:
        <Alert
          actionsLayout="vertical"
          actions={[{
            title: 'Закрыть объявление',
            autoclose: true,
            mode: 'destructive',
            action: () => onAccept(),
          }, {
            title: 'Отмена',
            autoclose: true,
            mode: 'cancel'
          }]}
          onClose={this.closePopout}>
          <h2>Подтвердите действие</h2>
          <p>Вы уверены, что хотите закрыть объявление?</p>
        </Alert>
    });
  };

  closePopout = () => {
    this.setState({ popout: null });
  };

  toCreateLostForm = () => {
    this.setState({mainPanel: 'new_lost'});
  };
  toMain = () => {
    this.setState({mainPanel: 'main'});
  };
  toLostList = () => {
    this.setState({
      lostPanel: 'losts',
    });
  };
  toLost = (id) => {
    this.setState({
      id: id,
      lostPanel: 'lost',
    });
  };
  toProfile = () => {
    this.setState({profilePanel: 'more'});
  };
  toProfileLost = (id) => {
    this.setState({
      profilePanel: 'lost',
      profileId: id,
    });
  };

  render() {
    return (
      <Epic activeStory={this.state.activeStory} tabbar={
        <Tabbar>
          <TabbarItem
            onClick={this.onStoryChange}
            selected={this.state.activeStory === 'main'}
            data-story="main"
            text="Главная"
          ><Icon28HomeOutline/></TabbarItem>
          <TabbarItem
            onClick={this.onStoryChange}
            selected={this.state.activeStory === 'lost'}
            data-story="lost"
            text="Потерялись"
          ><Icon28Menu/></TabbarItem>
          <TabbarItem
            onClick={this.onStoryChange}
            selected={this.state.activeStory === 'messages'}
            data-story="messages"
            text="Нашлись"
          ><Icon28ListCheckOutline/></TabbarItem>
          <TabbarItem
            onClick={this.onStoryChange}
            selected={this.state.activeStory === 'more'}
            data-story="more"
            text="Профиль"
          ><Icon28User/></TabbarItem>
        </Tabbar>
      }>
        <View id="main" activePanel={this.state.mainPanel}>
          <Panel id="main">
            <MainPanel toCreateLostForm={this.toCreateLostForm}/>
          </Panel>
          <Panel id="new_lost">
            <CreateFormPanel userStore={this.userStore} toMain={this.toMain}/>
          </Panel>
        </View>
        <View popup={this.state.popout} id="lost" activePanel={this.state.lostPanel} modal={
          <SearchFilter activeModal={this.state.activeModal}
                        onClose={this.modalBack}/>
        }>
          <Panel id="losts">
            <LostPanel toLost={this.toLost}
                       mapStore={this.mapStore}
                       openFilters={this.openFilters}/>
          </Panel>
          <Panel id="lost">
            <LostAnimalPanel userStore={this.userStore}
                             openDestructive={this.openDestructive}
                             goBack={this.toLostList}
                             id={this.state.id}/>
          </Panel>
        </View>
        <View id="messages" activePanel="messages">
          <Panel id="messages">
            <PanelHeader>Сообщения</PanelHeader>
          </Panel>
        </View>
        <View popout={this.state.popout} id="more" activePanel={this.state.profilePanel}>
          <Panel id="more">
            <ProfilePanel userStore={this.userStore}
                          openDestructive={this.openDestructive}
                          goBack={this.toProfile}
                          toLost={this.toProfileLost}/>
          </Panel>
          <Panel id="lost">
            <LostAnimalPanel userStore={this.userStore}
                             openDestructive={this.openDestructive}
                             goBack={this.toProfile}
                             id={this.state.profileId}/>
          </Panel>
        </View>
      </Epic>
    );
  }
}

export default App;

