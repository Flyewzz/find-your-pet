import config from '../config';

class LostService {
  get = async () => {
    const url = config.baseUrl + 'losts';
    const options = {method: 'GET'};
    const request = new Request(url, options);
    const response = await fetch(request);
    return response.json();
  };

  getById = async (id) => {
    const url = config.baseUrl + `lost?id=${id}`;
    const options = {method: 'GET'};
    const request = new Request(url, options);
    const response = await fetch(request);
    return response.json();
  };

  create = (type, authorId, sex, breed, description,
            latitude, longitude, picture, callback) => {
    const formData = new FormData();

    formData.append("type_id", type);
    formData.append("vk_id", authorId);
    formData.append("picture", picture);
    formData.append("sex", sex);
    formData.append("breed", breed);
    formData.append("description", description);
    formData.append("latitude", latitude);
    formData.append("longitude", longitude);

    const request = new XMLHttpRequest();
    request.open("POST", config.baseUrl + 'lost');
    request.send(formData);

    request.onload = () => {
      callback(request.status, request.response);
    };
  };
}

export default LostService;