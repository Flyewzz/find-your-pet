import {Div} from '@vkontakte/vkui';
import React from 'react';
import './FormLabel.css'

const FormLabel = (props) => {
  return (
    <Div className={'form__label'}>
      {props.text}
    </Div>
  );
};

export default FormLabel;