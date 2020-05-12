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
    query: '',
  };

  animals = null;

  onFieldChange = (key, value) => {
    this.fields[key] = value;
  };

  fetch = async (page) => {
    const {type, sex, breed, query} = this.fields;
    return this.lostService.get(type, sex, breed, query, page).then(result => {
      const animals = (result.payload !== null && result.payload.length === 0)
        ? null : result.payload;
      if (this.animals === null && animals !== null) {
        this.animals = []
      }
      this.animals.push(...animals);
      return result;
    });
  };
}

decorate(LostFilterStore, {
  fields: observable,
  animals: observable,
});

export default LostFilterStore;