import config from '../config';

class ProfileService {
  getLost = async (vkId) => {
    const url = `${config.baseUrl}profile/lost?vk_id=${vkId}`;
    const options = {method: 'GET'};
    const request = new Request(url, options);
    const response = await fetch(request);
    return response.json();
  };

  getFound = async (vkId) => {
    const url = `${config.baseUrl}profile/found?vk_id=${vkId}`;
    const options = {method: 'GET'};
    const request = new Request(url, options);
    const response = await fetch(request);
    return response.json();
  };

  closeLost = async (id, vkId) => {
    const url = `${config.baseUrl}lost?vk_id=${vkId}&id=${id}`;
    const options = {method: 'DELETE'};
    const request = new Request(url, options);
    return await fetch(request);
  };

  closeFound = async (id, vkId) => {
    const url = `${config.baseUrl}found?vk_id=${vkId}&id=${id}`;
    const options = {method: 'DELETE'};
    const request = new Request(url, options);
    return await fetch(request);
  };
}

export default ProfileService;