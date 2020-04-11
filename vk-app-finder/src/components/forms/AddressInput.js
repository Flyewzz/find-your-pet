import React from 'react';
import {ReactDadata} from 'react-dadata';
import {Div} from '@vkontakte/vkui';
import FormLabel from "./FormLabel";

const AddressInput = (props) => {
  return (
    <>
      <FormLabel text={props.title}/>
      <Div>
        <ReactDadata token={'c2dbff3b19ebd89e5b88b6b70350deae900ea3d3'}
                     placeholder={props.placeholder}/>
      </Div>
    </>
  );
};

export default AddressInput;