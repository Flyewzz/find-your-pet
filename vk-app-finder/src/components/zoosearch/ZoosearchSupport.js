import React from 'react';
import Logo from './ZoosearchLogo';
import './ZoosearchSupport.css';
import {Div, Link} from "@vkontakte/vkui";

const ZoosearchSupport = () => {
  return (
    <Div className={'zoosearch'}>
      <div className={'zoosearch__line'}>
        <Link className={'zoosearch__link'}
              target={'_blank'}
              href={'https://zoosearch.ru'}>
          <div className={'zoosearch__text'}>При поддержке</div>
          <Logo className={'zoosearch__logo'}/>
        </Link>
      </div>
    </Div>
  );
};

export default ZoosearchSupport;