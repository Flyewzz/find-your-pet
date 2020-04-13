import {observable, runInAction, decorate} from 'mobx'
import GenericFormStore from './GenericFormStore'
import LostService from "../services/LostService";

class LostStore extends GenericFormStore {
  constructor() {
    super();
    this.userService = new LostService();
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

  submit = (callback) => {
    try {
      this.userService.create(
        this.form.fields.typeId.value,
        this.form.fields.authorId.value,
        this.form.fields.sex.value,
        this.form.fields.breed.value,
        this.form.fields.description.value,
        this.form.fields.latitude.value,
        this.form.fields.longitude.value,
        this.form.fields.picture.value,
        callback,
      );
    } catch (error) {
      runInAction(() => {
        console.log(error);
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