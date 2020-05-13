import {decorate, observable} from "mobx";

class FoundFilter {
  fields = {
    type: '0',
    sex: '0',
    breed: '',
  };
}

decorate(FoundFilter, {
  fields: observable,
});

export default FoundFilter;