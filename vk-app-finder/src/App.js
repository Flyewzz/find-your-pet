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
import FoundPanel from "./panels/Found";
import FoundFilterStore from "./stores/FoundFilterStore";
import CreateFormFoundPanel from "./panels/CreateFormFoundPanel";
import FoundAnimalPanel from "./panels/FoundAnimalPanel";
import FoundMapStore from "./stores/FoundMapStore";

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

      popout: null,
    };
    this.onStoryChange = this.onStoryChange.bind(this);

    this.modalBack = (changed) => {
      if (changed) {
        this.lostFilterStore.animals = undefined;
      }
      this.setActiveModal(this.state.modalHistory[this.state.modalHistory.length - 2]);
    };
    this.userStore = new UserStore();
    this.mapStore = new LostMapStore(this.userStore);
    this.foundMapStore = new FoundMapStore(this.userStore);
    this.lostFilterStore = new LostFilterStore();
    this.foundFilterStore = new FoundFilterStore();
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
    console.log('Открыть popup кликнуто');
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

  toCreateLostForm = () => {
    this.setState({mainPanel: 'new_lost'});
  };
  toCreateFoundForm = () => {
    this.setState({mainPanel: 'new_found'});
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
  toFoundList = () => {
    this.setState({
      foundPanel: 'messages',
    });
  };
  toFound = (id) => {
    this.setState({
      id: id,
      foundPanel: 'found',
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
            onClick={(e) => {
              this.setState({mainPanel: 'main'});
              this.onStoryChange(e);
            }}
            selected={this.state.activeStory === 'main'}
            data-story="main"
            text="Главная"
          ><Icon28HomeOutline/></TabbarItem>
          <TabbarItem
            onClick={(e) => {
              this.setState({lostPanel: 'losts'});
              this.onStoryChange(e);
            }}
            selected={this.state.activeStory === 'lost'}
            data-story="lost"
            text="Потерялись"
          ><Icon28Menu/></TabbarItem>
          <TabbarItem
            onClick={(e) => {
              this.setState({foundPanel: 'messages'});
              this.onStoryChange(e);
            }}
            selected={this.state.activeStory === 'messages'}
            data-story="messages"
            text="Нашлись"
          ><Icon28ListCheckOutline/></TabbarItem>
          <TabbarItem
            onClick={(e) => {
              this.setState({profilePanel: 'more'});
              this.onStoryChange(e);
            }}
            selected={this.state.activeStory === 'more'}
            data-story="more"
            text="Профиль"
          ><Icon28User/></TabbarItem>
        </Tabbar>
      }>
        <View id="main" activePanel={this.state.mainPanel}>
          <Panel id="main">
            <MainPanel toCreateFoundForm={this.toCreateFoundForm}
                       toCreateLostForm={this.toCreateLostForm}/>
          </Panel>
          <Panel id="new_lost">
            <CreateFormPanel userStore={this.userStore}
                             toProfile={this.toProfileMain}
                             toMain={this.toMain}/>
          </Panel>
          <Panel id="new_found">
            <CreateFormFoundPanel userStore={this.userStore}
                                  toProfile={this.toProfileMain}
                                  toMain={this.toMain}/>
          </Panel>
        </View>
        <View popout={this.state.popout} id="lost" activePanel={this.state.lostPanel} modal={
          <SearchFilter activeModal={this.state.activeModal}
                        filterStore={this.lostFilterStore}
                        onClose={this.modalBack}/>
        }>
          <Panel id="losts">
            <LostPanel toLost={this.toLost}
                       lostFilterStore={this.lostFilterStore}
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
                        filterStore={this.lostFilterStore}
                        onClose={this.modalBack}/>
        }>
          <Panel id="messages">
            <FoundPanel toFound={this.toFound}
                        foundFilterStore={this.foundFilterStore}
                        mapStore={this.foundMapStore}
                        openFilters={this.openFilters}
                        />
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
                          openDestructive={this.openDestructive}
                          goBack={this.toProfile}
                          toMainForm={this.toMainForm}
                          toMainFoundForm={this.toMainFoundForm}
                          toLost={this.toProfileLost}
                          toFound={this.toProfileFound}
                          />
          </Panel>
          <Panel id="lost">
            <LostAnimalPanel userStore={this.userStore}
                             openDestructive={this.openDestructive}
                             goBack={this.toProfile}
                             id={this.state.profileId}/>
          </Panel>
          <Panel id="found">
            <FoundAnimalPanel userStore={this.userStore}
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

