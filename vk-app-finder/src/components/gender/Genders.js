
import male from '../../img/gender/male.png';
import female from '../../img/gender/female.png';

function getGenderInfo(gender) {
    /* types:
        m - male
        f - female
        n/a - undefined
    */
   switch (gender) {
        case 'm':
           return {
               name: 'Мужской',
               picture: male,
           }
        case 'f':
            return {
                name: 'Женский',
                picture: female,
            }
        case 'n/a':
            return {
                name: 'Не определен',
                picture: null,
            };
   }
}

export default getGenderInfo;