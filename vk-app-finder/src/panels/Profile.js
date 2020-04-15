import React from "react";
import {Group, PanelHeader} from "@vkontakte/vkui";
import {PanelHeaderBack} from "@vkontakte/vkui/dist/es6";
import ProfileService from "../services/ProfileService";
import {decorate, observable, runInAction} from "mobx";
import {observer} from "mobx-react";
import {Tabs} from '@vkontakte/vkui'
import {TabsItem} from '@vkontakte/vkui'
import LostTab from '../components/profile/LostTab';
import './Profile.css';

class ProfilePanel extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      mapView: false,
      places: [],
      activePanel: "main-panel",
      contextOpened: false,
      mode: "all",
      activeTab: "lost",
    };

    this.changeView = () => {
      const current = this.state.mapView;
      this.setState({mapView: !current});
    };

    this.select = this.select.bind(this);
  }

  select(e) {
    const mode = e.currentTarget.dataset.mode;
    this.setState({mode, contextOpened: false});
  }

  toLost = (id) => {
    this.setState({
      id: id,
      lostPanel: 'lost',
    });
  };

  render() {
    return (
      <>
        <PanelHeader left={<PanelHeaderBack/>}>Ваши объявления</PanelHeader>
        <Tabs>
          <TabsItem
            onClick={() => this.setState({activeTab: 'lost'})}
            selected={this.state.activeTab === 'lost'}>
            Пропавшие
          </TabsItem>
          <TabsItem
            onClick={() => this.setState({activeTab: 'found'})}
            selected={this.state.activeTab === 'found'}>
            Найденные
          </TabsItem>
        </Tabs>
        <Group separator="hide">
          {this.state.activeTab === 'lost'
          && <LostTab userStore={this.props.userStore}
                      openDestructive={this.props.openDestructive}
                      toLost={this.props.toLost}/>}
        </Group>
      </>
    );
  }
}

decorate(ProfilePanel, {});

export default observer(ProfilePanel);
