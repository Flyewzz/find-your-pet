import {decorate, observable} from "mobx";

class MapStateStore {
  constructor(userStore) {
    userStore.getPosition().then(result => {
      if (result.available) {
        this.center = [result.lat, result.long];
      }
    })
  }

  isMapView = false;
  center = [55.75, 37.57];
  zoom = 9;
}

decorate(MapStateStore, {
  isMapView: observable,
  center: observable,
  zoom: observable,
});

export default MapStateStore;