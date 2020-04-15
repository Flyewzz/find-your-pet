import config from '../config';

class ProfileService {
  getLost = async (vkId) => {
    const url = `${config.baseUrl}profile/lost?vk_id=${vkId}`;
    const options = {method: 'GET'};
    const request = new Request(url, options);
    const response = await fetch(request);
    return response.json();
  };

  close = async (id, vkId) => {
    const url = `${config.baseUrl}lost?vk_id=${vkId}&lost_id=${id}&opened=${0}`;
    const options = {method: 'PUT'};
    const request = new Request(url, options);
    return await fetch(request);
  };
}

export default ProfileService;