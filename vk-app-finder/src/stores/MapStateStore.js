import {decorate, observable} from "mobx";

class MapStateStore {
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