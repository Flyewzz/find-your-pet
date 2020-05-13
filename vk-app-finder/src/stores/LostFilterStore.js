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
  hasMore = true;

  onFieldChange = (key, value) => {
    this.fields[key] = value;
  };

  fetch = async (page) => {
    const {type, sex, breed, query} = this.fields;
    return this.lostService.get(type, sex, breed, query, page).then(result => {
      const newAnimals = (result.payload !== null && result.payload.length === 0)
        ? null : result.payload;
      if ((this.animals === null || this.animals === undefined) && newAnimals !== null) {
        this.animals = []
      }
      if (this.animals !== null && this.animals !== undefined) {
        const animals = this.animals;
        animals.push(...newAnimals);
        this.animals = animals;
      }
      this.hasMore = result.has_more;
      return result;
    });
  };
}

decorate(LostFilterStore, {
  fields: observable,
  animals: observable,
  hasMore: observable,
});

export default LostFilterStore;