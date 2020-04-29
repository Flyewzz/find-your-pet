import {decorate, observable} from "mobx";
import LostService from "../services/LostService";
import FoundService from "../services/FoundService";

class FoundFilterStore {
  constructor() {
    this.foundService = new FoundService();
  }

  fields = {
    address: '',
    type: '0',
    sex: '0',
    breed: '',
  };

  animals = null;

  onFieldChange = (key, value) => {
    this.fields[key] = value;
  };

  fetch = async () => {
    const {type, sex, breed} = this.fields;
    return this.foundService.get(type, sex, breed).then(result => {
      this.animals = (result.payload !== null && result.payload.length === 0)
        ? null : result.payload;
      this.onFetch();
    })
  };
}

decorate(FoundFilterStore, {
  fields: observable,
  animals: observable,
});

export default FoundFilterStore;