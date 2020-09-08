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
import {RouteNode} from "react-router5";

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

      wasEverFetched: true,
      needLostFetch: true,
      needFoundFetch: true,

      createLostStage: 0,
      createFoundStage: 0,

      popout: null,
      formPopout: null,
    };
    this.onStoryChange = this.onStoryChange.bind(this);

    this.modalBack = () => {
      window.history.back();
      this.setActiveModal(this.state.modalHistory[this.state.modalHistory.length - 2]);
    };
    this.userStore = new UserStore(props.search);
    this.mapStore = new LostMapStore(this.userStore);
    this.foundMapStore = new FoundMapStore(this.userStore);
    this.lostFilterStore = new LostFilterStore();
    this.lostFilter = new LostFilter();
    this.foundFilterStore = new FoundFilterStore();
    this.foundFilter = new FoundFilter();
    this.openFilters = () => {
      this.props.router.navigate('filters');
      this.setActiveModal('filters');
    };

    props.router.subscribe(this.routeChange);
  }

  routeChange = (value) => {
    this.setState({popout: null});
    this.setActiveModal(null);

    let name = value.route.name;
    if (name === 'lost' || name === 'found' || name === 'my_lost' || name === 'my_found') {
      name += value.route.params.id;
    }

    this.chooseRoute(name, value.route.params);
  };

  componentDidMount() {
    const hash = this.props.hash.slice(1);
    if (hash.startsWith('lost') && hash !== 'losts') {
      this.setState({wasEverFetched: false});
      const id = hash.slice(4);
      this.props.router.navigate('lost', {id});
    } else if (hash.startsWith('found') && hash !== 'founds') {
      this.setState({wasEverFetched: false});
      const id = hash.slice(5);
      this.props.router.navigate('found', {id});
    }
  }

  chooseRoute = (location, params) => {
    if (location === 'home') {
      this.setState({activeStory: 'main', mainPanel: 'main'});
    } else if (location === 'create_lost') {
      this.setState({
        activeStory: 'main',
        mainPanel: 'new_lost',
        createLostStage: params.stage
      });
    } else if (location === 'create_found') {
      this.setState({
        activeStory: 'main',
        mainPanel: 'new_found',
        createFoundStage: params.stage,
      });
    } else if (location === 'losts') {
      this.setState({activeStory: 'lost', lostPanel: 'losts'});
    } else if (location === 'founds') {
      this.setState({activeStory: 'messages', foundPanel: 'messages'});
    } else if (location === 'profile') {
      this.setState({activeStory: 'more', profilePanel: 'more'});
    } else if (location === 'profile_lost') {
      this.setState(
        {activeStory: 'more', profilePanel: 'more', profileTab: 'lost'}
      );
    } else if (location === 'profile_found') {
      this.setState(
        {activeStory: 'more', profilePanel: 'more', profileTab: 'found'}
      );
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
    } else if (location.slice(0, 7) === 'my_lost') {
      const id = location.slice(7);
      this.setState({
        activeStory: 'more',
        profilePanel: 'lost',
        profileId: id,
      });
    } else if (location.slice(0, 8) === 'my_found') {
      const id = location.slice(8);
      this.setState({
        activeStory: 'more',
        profilePanel: 'found',
        profileId: id,
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
    this.props.router.navigate('destructive');

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
    window.history.back();
    this.setState({popout: null});
  };

  openScreenSpinner = () => {
    this.setState({formPopout: <ScreenSpinner/>})
  };

  closeScreenSpinner = () => {
    this.setState({formPopout: null})
  };

  toCreateLostForm = () => {
    this.props.router.navigate('create_lost',
      {stage: this.state.createLostStage}
    );
  };

  toCreateFoundForm = () => {
    this.props.router.navigate('create_found',
      {stage: this.state.createFoundStage}
    );
  };

  toMain = () => {
    this.props.router.navigate('home');
  };

  toLostList = (needFetch = true) => {
    this.setState({
      needLostFetch: needFetch,
      wasEverFetched: true,
    });
    this.props.router.navigate('losts');
  };

  toLost = (id) => {
    this.props.router.navigate('lost', {id});
  };

  toFoundList = (needFetch = true) => {
    this.setState({
      needFoundFetch: needFetch,
      wasEverFetched: true,
    });
    this.props.router.navigate('founds');
  };

  toFound = (id) => {
    this.props.router.navigate('found', {id});
  };

  toProfileLostTab = () => {
    this.props.router.navigate('profile_lost')
  };

  toProfileFoundTab = () => {
    this.props.router.navigate('profile_found');
  };

  toProfileLost = (id) => {
    this.props.router.navigate('my_lost', {id});
  };

  toProfileFound = (id) => {
    this.props.router.navigate('my_found', {id});
  };

  toMainForm = () => {
    this.props.router.navigate('create_lost', {stage: 0});
  };

  toMainFoundForm = () => {
    this.props.router.navigate('create_found', {stage: 0});
  };

  toProfileMain = () => {
    this.props.router.navigate('profile')
  };

  render() {
    let {router} = this.props;

    return (
      <Epic activeStory={this.state.activeStory} tabbar={
        <Tabbar>
          <TabbarItem
            onClick={() => this.props.router.navigate('home')}
            selected={this.state.activeStory === 'main'}
            data-story="main"
            text="Главная"
          ><Icon28HomeOutline/></TabbarItem>
          <TabbarItem
            onClick={() => this.props.router.navigate('losts')}
            selected={this.state.activeStory === 'lost'}
            data-story="lost"
            text="Потерялись"
          ><Icon28Menu/></TabbarItem>
          <TabbarItem
            onClick={() => this.props.router.navigate('founds')}
            selected={this.state.activeStory === 'messages'}
            data-story="messages"
            text="Нашлись"
          ><Icon28ListCheckOutline/></TabbarItem>
          <TabbarItem
            onClick={() => this.props.router.navigate('profile')}
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
                             stage={this.state.createLostStage}
                             router={router}
                             toMain={this.toMain}/>
          </Panel>
          <Panel id="new_found">
            <CreateFormFoundPanel userStore={this.userStore}
                                  toProfile={this.toProfileFoundTab}
                                  openPopout={this.openScreenSpinner}
                                  closePopout={this.closeScreenSpinner}
                                  stage={this.state.createFoundStage}
                                  router={router}
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
                             wasEverFetched={this.state.wasEverFetched}
                             goBack={(needFetch) => this.toLostList(needFetch)}
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
                              wasEverFetched={this.state.wasEverFetched}
                              goBack={(needFetch) => this.toFoundList(needFetch)}
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

export default observer((props) => (
  <RouteNode>
    {({route}) => <App route={route} {...props}/>}
  </RouteNode>
));
