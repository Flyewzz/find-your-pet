import {action, toJS, decorate} from 'mobx'
import Validator from 'validatorjs';

class FormStore {
  getFlattenedValues = (valueKey = 'value') => {
    let data = {};
    let form = toJS(this.form).fields;
    Object.keys(form).forEach(key => {
      data[key] = form[key][valueKey]
    });
    return data
  };

  onFieldChange = (field, value) => {
    console.log(field, value);
    this.form.fields[field].value = value;
    const validation = new Validator(
      this.getFlattenedValues('value'),
      this.getFlattenedValues('rule'));
    this.form.meta.isValid = validation.passes(validation);
    this.form.fields[field].error = validation.errors.first(field)
  };

  setError = (errMsg) => {
    this.form.meta.error = errMsg
  };
}

decorate(FormStore, {
  onFieldChange: action,
  setError: action,
});

export default FormStore;