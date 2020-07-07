import React from 'react';
import View from '@vkontakte/vkui/dist/components/View/View';
import '@vkontakte/vkui/dist/vkui.css';
import Tabbar from "@vkontakte/vkui/dist/components/Tabbar/Tabbar";
import TabbarItem from "@vkontakte/vkui/dist/components/TabbarItem/TabbarItem";
import Icon28Menu from '@vkontakte/icons/dist/28/menu';
import Icon28User from '@vkontakte/icons/dist/28/user';
import Icon28HomeOutline from '@vkontakte/icons/dist/28/home_outline';
import Icon28ListCheckOutline from '@vkontakte/icons/dist/28/list_check_outline';
import {Epic, Panel} from "@vkontakte/vkui";
import MainPanel from "./panels/Main";
import LostPanel from "./panels/Lost";
import SearchFilter from "./panels/SearchFilter";
import CreateFormPanel from "./panels/CreateFormPanel";
import LostAnimalPanel from "./panels/LostAnimalPanel";
import LostMapStore from "./stores/LostMapStore";
import UserStore from "./stores/UserStore";
import ProfilePanel from "./panels/Profile";
import {Alert} from "@vkontakte/vkui";
import LostFilterStore from "./stores/LostFilterStore";
import LostFilter from "./stores/LostFilter"
import FoundPanel from "./panels/Found";
import FoundFilterStore from "./stores/FoundFilterStore";
import FoundFilter from "./stores/FoundFilter";
import CreateFormFoundPanel from "./panels/CreateFormFoundPanel";
import FoundAnimalPanel from "./panels/FoundAnimalPanel";
import FoundMapStore from "./stores/FoundMapStore";
import {observer} from "mobx-react";
import ScreenSpinner from "@vkontakte/vkui/dist/components/ScreenSpinner/ScreenSpinner";
// import bridge from "@vkontakte/vk-bridge";

class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      activeModal: null,
      modalHistory: [],
      activeStory: 'main',
      mainPanel: 'main',
      lostPanel: 'losts',
      foundPanel: 'messages',
      profilePanel: 'more',
      profileTab: 'lost',

      needLostFetch: true,
      needFoundFetch: true,

      popout: null,
      formPopout: null,
    };
    this.onStoryChange = this.onStoryChange.bind(this);

    this.modalBack = () => {
      this.setActiveModal(this.state.modalHistory[this.state.modalHistory.length - 2]);
    };
    this.userStore = new UserStore();
    this.mapStore = new LostMapStore(this.userStore);
    this.foundMapStore = new FoundMapStore(this.userStore);
    this.lostFilterStore = new LostFilterStore();
    this.lostFilter = new LostFilter();
    this.foundFilterStore = new FoundFilterStore();
    this.foundFilter = new FoundFilter();
    this.openFilters = () => {
      this.setActiveModal('filters');
    };
  }

  componentDidMount() {
    this.chooseRoute(window.location.hash.slice(1));
  }

  chooseRoute = (location) => {
    if (location === 'home') {
      this.setState({activeStory: 'main', mainPanel: 'main'});
    } else if (location === 'losts') {
      this.setState({activeStory: 'lost', lostPanel: 'losts'});
    } else if (location === 'founds') {
      this.setState({activeStory: 'messages', foundPanel: 'messages'});
    } else if (location === 'profile') {
      this.setState({activeStory: 'more', profilePanel: 'more'});
    } else if (location.slice(0, 5) === 'found') {
      const id = location.slice(5);
      this.setState({
        activeStory: 'messages',
        id: id,
        foundPanel: 'found',
      });
    } else if (location.slice(0, 4) === 'lost') {
      const id = location.slice(4);
      this.setState({
        activeStory: 'lost',
        id: id,
        lostPanel: 'lost',
      });
    }
  };

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
    this.setState({
      popout:
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
    this.setState({popout: null});
  };

  openScreenSpinner = () => {
    this.setState({formPopout: <ScreenSpinner/>})
  };
  closeScreenSpinner = () => {
    this.setState({formPopout: null})
  };

  toCreateLostForm = () => {
    this.setState({mainPanel: 'new_lost'});
  };
  toCreateFoundForm = () => {
    this.setState({mainPanel: 'new_found'});
  };
  toMain = () => {
    this.setState({mainPanel: 'main'});
  };
  toLostList = (needFetch=true) => {
    this.setState({
      needFoundFetch: needFetch,
    });
    this.userStore.changeLocation(`losts`);
    this.chooseRoute('losts');
  };
  toLost = (id) => {
    this.userStore.changeLocation(`lost${id}`);
    this.chooseRoute(`lost${id}`);
  };
  toFoundList = (needFetch=true) => {
    this.userStore.changeLocation(`founds`);
    this.setState({
      needFoundFetch: needFetch,
    });
    this.chooseRoute('founds');
  };
  toFound = (id) => {
    this.userStore.changeLocation(`found${id}`);
    this.chooseRoute(`found${id}`)
  };
  toProfileLostTab = () => {
    this.setState({
      profilePanel: 'more',
      profileTab: 'lost',
    });
  };
  toProfileFoundTab = () => {
    this.setState({
      profilePanel: 'more',
      profileTab: 'found',
    });
  };
  toProfileLost = (id) => {
    this.setState({
      profilePanel: 'lost',
      profileId: id,
    });
  };
  toProfileFound = (id) => {
    this.setState({
      profilePanel: 'found',
      profileId: id,
    });
  };
  toMainForm = () => {
    this.setState({
      activeStory: 'main',
      mainPanel: 'new_lost',
    })
  };
  toMainFoundForm = () => {
    this.setState({
      activeStory: 'main',
      mainPanel: 'new_found',
    })
  };
  toProfileMain = () => {
    this.setState({
      activeStory: 'more',
      profilePanel: 'more',
    });
  };

  render() {
    return (
      <Epic activeStory={this.state.activeStory} tabbar={
        <Tabbar>
          <TabbarItem
            onClick={() => {
              this.userStore.changeLocation('home');
              this.chooseRoute('home');
            }}
            selected={this.state.activeStory === 'main'}
            data-story="main"
            text="Главная"
          ><Icon28HomeOutline/></TabbarItem>
          <TabbarItem
            onClick={() => {
              this.userStore.changeLocation('losts');
              this.chooseRoute('losts');
            }}
            selected={this.state.activeStory === 'lost'}
            data-story="lost"
            text="Потерялись"
          ><Icon28Menu/></TabbarItem>
          <TabbarItem
            onClick={() => {
              this.userStore.changeLocation('founds');
              this.chooseRoute('founds');
            }}
            selected={this.state.activeStory === 'messages'}
            data-story="messages"
            text="Нашлись"
          ><Icon28ListCheckOutline/></TabbarItem>
          <TabbarItem
            onClick={() => {
              this.userStore.changeLocation('profile');
              this.chooseRoute('profile');
            }}
            selected={this.state.activeStory === 'more'}
            data-story="more"
            text="Профиль"
          ><Icon28User/></TabbarItem>
        </Tabbar>
      }>
        <View popout={this.state.formPopout} id="main" activePanel={this.state.mainPanel}>
          <Panel id="main">
            <MainPanel toCreateFoundForm={this.toCreateFoundForm}
                       toCreateLostForm={this.toCreateLostForm}/>
          </Panel>
          <Panel id="new_lost">
            <CreateFormPanel userStore={this.userStore}
                             toProfile={this.toProfileMain}
                             openPopout={this.openScreenSpinner}
                             closePopout={this.closeScreenSpinner}
                             toMain={this.toMain}/>
          </Panel>
          <Panel id="new_found">
            <CreateFormFoundPanel userStore={this.userStore}
                                  toProfile={() => {
                                    this.toProfileFoundTab();
                                    this.toProfileMain();
                                  }}
                                  openPopout={this.openScreenSpinner}
                                  closePopout={this.closeScreenSpinner}
                                  toMain={this.toMain}/>
          </Panel>
        </View>
        <View popout={this.state.popout} id="lost" activePanel={this.state.lostPanel} modal={
          <SearchFilter activeModal={this.state.activeModal}
                        filterStore={this.lostFilterStore}
                        filter={this.lostFilter}
                        onClose={this.modalBack}/>
        }>
          <Panel id="losts">
            <LostPanel toLost={this.toLost}
                       lostFilterStore={this.lostFilterStore}
                       filter={this.lostFilter}
                       needFetch={this.state.needLostFetch}
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
        <View popout={this.state.popout} id="messages" activePanel={this.state.foundPanel} modal={
          <SearchFilter activeModal={this.state.activeModal}
                        filterStore={this.foundFilterStore}
                        filter={this.foundFilter}
                        onClose={this.modalBack}/>
        }>
          <Panel id="messages">
            <FoundPanel toFound={this.toFound}
                        foundFilterStore={this.foundFilterStore}
                        filter={this.foundFilter}
                        needFetch={this.state.needFoundFetch}
                        mapStore={this.foundMapStore}
                        openFilters={this.openFilters}/>
          </Panel>
          <Panel id="found">
            <FoundAnimalPanel userStore={this.userStore}
                              openDestructive={this.openDestructive}
                              goBack={this.toFoundList}
                              id={this.state.id}/>
          </Panel>
        </View>
        <View popout={this.state.popout} id="more" activePanel={this.state.profilePanel}>
          <Panel id="more">
            <ProfilePanel userStore={this.userStore}
                          activeTab={this.state.profileTab}
                          openDestructive={this.openDestructive}
                          goBackLost={this.toProfileLostTab}
                          goBackFound={this.toProfileFoundTab}
                          toMainForm={this.toMainForm}
                          toMainFoundForm={this.toMainFoundForm}
                          toLost={this.toProfileLost}
                          toFound={this.toProfileFound}
            />
          </Panel>
          <Panel id="lost">
            <LostAnimalPanel userStore={this.userStore}
                             openDestructive={this.openDestructive}
                             goBack={this.toProfileLostTab}
                             id={this.state.profileId}/>
          </Panel>
          <Panel id="found">
            <FoundAnimalPanel userStore={this.userStore}
                              openDestructive={this.openDestructive}
                              goBack={this.toProfileFoundTab}
                              id={this.state.profileId}/>
          </Panel>
        </View>
      </Epic>
    );
  }
}

export default observer(App);
