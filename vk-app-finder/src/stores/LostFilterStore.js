import {decorate, observable} from "mobx";
import LostService from "../services/LostService";

class LostFilterStore {
  constructor() {
    this.lostService = new LostService();
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
    return this.lostService.get(type, sex, breed).then(result => {
      this.animals = (result.payload !== null && result.payload.length === 0)
        ? null : result.payload;
      this.onFetch();
    }).catch(err => { console.log(err); });
  };
}

decorate(LostFilterStore, {
  fields: observable,
  animals: observable,
});

export default LostFilterStore;