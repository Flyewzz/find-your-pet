import config from '../config';

class FoundService {
  get = async (type, sex, breed) => {
    let url = config.baseUrl + 'founds?';
    if (type !== '0') {
      url += '&type_id=' + type;
    }
    if (sex !== '0') {
      url += '&sex=' + sex;
    }
    if (breed !== '') {
      url += '&breed=' + breed;
    }
    const options = {method: 'GET'};
    const request = new Request(url, options);
    const response = await fetch(request);
    return response.json();
  };

  getById = async (id) => {
    const url = config.baseUrl + `found?id=${id}`;
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
    request.open("POST", config.baseUrl + 'found');
    request.send(formData);

    request.onload = () => {
      callback(request.status, request.response);
    };
  };
}

export default FoundService;