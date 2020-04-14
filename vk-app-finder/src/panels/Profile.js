import React from "react";
import { Card, CardGrid, Div, Group, PanelHeader } from "@vkontakte/vkui";
import { PanelHeaderBack } from "@vkontakte/vkui/dist/es6";
import ProfileCard from "../components/cards/ProfileCard";
import FilterLine from "../components/cards/FilterLine";
import DG from "2gis-maps";
import ProfileService from "../services/ProfileService";
import { decorate, observable, runInAction } from "mobx";
import { observer } from "mobx-react";
import LostPanel from "./Lost";
import { Tabs } from '@vkontakte/vkui'
import { TabsItem } from '@vkontakte/vkui'
import { Separator } from '@vkontakte/vkui'
import { Panel } from '@vkontakte/vkui'
import { View } from '@vkontakte/vkui'
import LostTab from './ProfileTabs/LostTab';
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
      this.setState({ mapView: !current });
    };

    this.profileService = new ProfileService();
    this.select = this.select.bind(this);
  }

  animals = [];

  select(e) {
    const mode = e.currentTarget.dataset.mode;
    this.setState({ mode, contextOpened: false });
  }

  toLost = (id) => {
    this.setState({
      id: id,
      lostPanel: 'lost',
    });
  };

  componentDidMount() {
    this.profileService.get().then(
      (result) => {
        runInAction(() => {
          this.animals = result.payload;
        });
      },
      (error) => {
        alert(error);
      }
    );
  }

  animalsToCards = () => {
    return this.animals.map((animal, index) => (
      <>
        {/* <Card key={-animal.id} size="l" styles={{height: 0}}/> */}
        <ProfileCard
          onClick={() => this.props.toLost(animal.id)}
          key={animal.id}
          animal={animal}
        />
      </>
    ));
  };

  render() {
    return (
      <>
        <PanelHeader left={<PanelHeaderBack />}>Ваши объявления</PanelHeader>
        <View activePanel={this.state.activePanel} header={false}>
        <Panel id="main-panel" separator={false}>
            <Tabs>
              <TabsItem
                onClick={() => this.setState({ activeTab: 'lost' })}
                selected={this.state.activeTab === 'lost'}
              >
                Пропавшие
              </TabsItem>
              <TabsItem
                onClick={() => this.setState({ activeTab: 'found' })}
                selected={this.state.activeTab === 'found'}
              >
                Найденные
              </TabsItem>
            </Tabs>
            <Separator />
        <Group separator="hide">
          <FilterLine
            isMap={this.state.mapView}
            changeView={this.changeView}
            openFilters={this.props.openFilters}
          />

          {this.state.activeTab === 'lost' &&
           <LostTab toLost={this.toLost} />}
        </Group>
      </Panel>
      </View>
      </>
    );
  }
}

decorate(ProfilePanel, {
  animals: observable,
});

export default observer(ProfilePanel);
