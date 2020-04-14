import config from '../config';

class ProfileService {
  get = async () => {
    const url = config.baseUrl + 'losts';
    const options = {method: 'GET'};
    const request = new Request(url, options);
    const response = await fetch(request);
    return response.json();
  };
}

export default ProfileService;