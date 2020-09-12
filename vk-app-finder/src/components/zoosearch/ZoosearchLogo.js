import React from 'react';
import logoImage from '../../img/zoosearch.png';

const Logo = (props) => {
  return(
    <img className={props.className} src={logoImage} alt={'zoosearch_logo'}/>
  );
};

export default Logo;