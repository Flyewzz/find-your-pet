import config from '../config';

class ProfileService {
  getLost = async (vkId) => {
    const url = `${config.baseUrl}profile/lost?vk_id=${vkId}`;
    const options = {method: 'GET'};
    const request = new Request(url, options);
    const response = await fetch(request);
    return response.json();
  };
}

export default ProfileService;