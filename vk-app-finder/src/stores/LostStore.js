import {observable, runInAction, decorate} from 'mobx'
import GenericFormStore from './GenericFormStore'
import LostService from '../services/LostService';

class LostStore extends GenericFormStore {
  constructor(userStore) {
    super();
    this.lostService = new LostService();
    this.userStore = userStore;
  }

  form = {
    fields: {
      typeId: {
        value: 1,
        error: null,
        rule: 'required'
      },
      authorId: {
        value: 1,
        error: null,
        rule: 'required'
      },
      sex: {
        value: 'm',
        error: null,
        rule: 'required'
      },
      breed: {
        value: '',
        error: null,
        rule: []
      },
      description: {
        value: '',
        error: null,
        rule: []
      },
      latitude: {
        value: '',
        error: null,
        rule: 'required'
      },
      longitude: {
        value: '',
        error: null,
        rule: 'required'
      },
      address: {
        value: '',
        error: null,
        rule: 'required',
      },
      picture: {
        value: null,
        error: null,
        rule: 'required'
      },
    },
    meta: {
      isValid: true,
      error: null,
    },
  };

  check = () => {
    const {description, picture} = this.form.fields;
    this.form.meta.isValid = !(description.value === '' && picture.value === null);
    this.form.meta.error = 'Вы не загрузили фото и не добавили описание.' +
      ' Как другим узнать Вашего питомца?';
    return this.form.meta.isValid;
  };

  submit = (callback) => {
    try {
        this.lostService.create(
          this.form.fields.typeId.value,
          this.userStore.id,
          this.form.fields.sex.value,
          this.form.fields.breed.value,
          this.form.fields.description.value,
          this.form.fields.latitude.value,
          this.form.fields.longitude.value,
          this.form.fields.picture.value,
          this.form.fields.address.value,
          callback
        );
    } catch (error) {
      runInAction(() => {
        this.form.meta.isValid = false;
        this.form.meta.error = error;
      })
    }
  };
}

decorate(LostStore, {
  form: observable,
});

export default LostStore;