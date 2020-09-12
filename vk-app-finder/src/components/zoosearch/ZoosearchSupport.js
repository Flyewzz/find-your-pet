import React from 'react';
import Logo from './ZoosearchLogo';
import './ZoosearchSupport.css';
import {Div} from "@vkontakte/vkui";

const ZoosearchSupport = () => {
  return(
    <Div className={'zoosearch'}>
      <div className={'zoosearch__line'}>
        <div className={'zoosearch__text'}>При поддержке</div>
        <Logo className={'zoosearch__logo'}/>
      </div>
    </Div>
  );
};

export default ZoosearchSupport;