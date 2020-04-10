import React from 'react';
import {ReactDadata} from 'react-dadata';
import {Div} from '@vkontakte/vkui';
import FormLabel from "./FormLabel";

const AddressInput = (props) => {
  return (
    <>
      <FormLabel text={props.title}/>
      <Div>
        <ReactDadata token={'API_TOKEN'}
                     placeholder={props.placeholder}/>
      </Div>
    </>
  );
};

export default AddressInput;