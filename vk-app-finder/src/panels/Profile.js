import React from 'react';
import {Group, PanelHeader} from '@vkontakte/vkui';
import {decorate} from 'mobx';
import {observer} from 'mobx-react';
import {Tabs} from '@vkontakte/vkui'
import {TabsItem} from '@vkontakte/vkui'
import LostTab from '../components/profile/LostTab';
import './Profile.css';
import FoundTab from "../components/profile/FoundTab";

class ProfilePanel extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      mapView: false,
      places: [],
      contextOpened: false,
      mode: "all",
      activeTab: props.activeTab,
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

  render() {
    return (
      <>
        <PanelHeader>Ваши объявления</PanelHeader>
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
                      toMainForm={this.props.toMainForm}
                      goBack={this.props.goBackLost}
                      toLost={this.props.toLost}/>}
          {this.state.activeTab === 'found'
          && <FoundTab userStore={this.props.userStore}
                       openDestructive={this.props.openDestructive}
                       toMainForm={this.props.toMainFoundForm}
                       goBack={this.props.goBackFound}
                       toFound={this.props.toFound}/>}
        </Group>
      </>
    );
  }
}

decorate(ProfilePanel, {});

export default observer(ProfilePanel);
