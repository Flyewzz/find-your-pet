import {decorate, observable} from "mobx";
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
    query: '',
  };

  hasMore = true;
  animals = null;
  page = 1;

  onFieldChange = (key, value) => {
    this.fields[key] = value;
  };

  incPage = () => {
    this.page++
  };

  clearPage = () => {
    this.page = 1;
    this.animals = undefined;
  };

  fetch = async () => {
    const {type, sex, breed, query} = this.fields;
    return this.foundService.get(type, sex, breed, query, this.page).then(result => {
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
        // noinspection JSUnresolvedVariable
        this.hasMore = result.has_more;
        return result;
      },
      (error) => {
        console.log(error);
        this.animals = null;
      })
  };
}

decorate(FoundFilterStore, {
  fields: observable,
  animals: observable,
  hasMore: observable,
});

export default FoundFilterStore;