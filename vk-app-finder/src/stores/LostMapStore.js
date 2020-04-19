import {decorate, observable} from "mobx";

class LostMapStore {
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

decorate(LostMapStore, {
  isMapView: observable,
  center: observable,
  zoom: observable,
});

export default LostMapStore;