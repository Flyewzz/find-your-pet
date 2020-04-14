import {decorate, observable} from 'mobx';
import bridge from '@vkontakte/vk-bridge';
import config from '../config';

class UserStore {
  constructor() {
    bridge.send("VKWebAppGetUserInfo", {}).then(
      result => {
        this.id = result.id;
      }
    );
  }

  id = -1;

  share(text) {
    console.log(text);
    bridge.send("VKWebAppShowWallPostBox", {message: text}).then(
      result => console.log(result)
    );
  }

  getUserById = async (id) => {
    return await bridge.send('VKWebAppCallAPIMethod', {
      method: 'users.get',
      request_id: 'user.get' + id,
      params: {
        v: '5.103',
        user_ids: id,
        fields: 'can_write_private_message,photo_50',
        access_token: config.serviceKey,
      },
    })
  };
}

decorate(UserStore, {
  id: observable,
});

export default UserStore;