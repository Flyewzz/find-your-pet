import config from '../config';

class LostService {
  get = async () => {
    const url = config + 'losts';
    const options = {method: 'GET'};
    const request = new Request(url, options);
    const response = await fetch(request);
    return response.json();
  };

  create = async (type, authorId, sex, breed, description,
                  latitude, longitude, picture) => {
    const url = config.baseUrl +
      `lost?type_id=${type}&author_id=${authorId}&sex=${sex}` +
      `&breed=${breed}&description=${description}&latitude=${latitude}` +
      `&longitude=${longitude}&picture=${picture}`;
    const options = {method: 'POST'};
    const request = new Request(url, options);
    const response = await fetch(request);
    return response.json();
  };
}

export default LostService;