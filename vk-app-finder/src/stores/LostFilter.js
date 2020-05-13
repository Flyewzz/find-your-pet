import {decorate, observable} from "mobx";

class LostFilter {
  fields = {
    type: '0',
    sex: '0',
    breed: '',
  };
}

decorate(LostFilter, {
  fields: observable,
});

export default LostFilter;