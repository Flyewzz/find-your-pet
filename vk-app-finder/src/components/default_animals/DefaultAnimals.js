
import vk_dog from '../../img/default_animals/vk_dog.jpg';
import vk_cat from '../../img/default_animals/vk_cat.jpg';
import vk_hamster from '../../img/default_animals/vk_hamster.jpg';

function getDefaultAnimal(type_id) {
    /* types:
        1 - dog
        2 - cat
        3 - others
    */
   switch (type_id) {
        case 1:
           return vk_dog;
        case 2:
            return vk_cat;
        case 3:
            return vk_hamster;
   }
}

export default getDefaultAnimal;