import config from '../config';

class GeocodingService {
  get = async () => {
    const url = 'https://fezzo.ru/';
    const options = {method: 'GET'};
    const request = new Request(url, options);
    const response = await fetch(request);
    return response.json();
  };

  auth = async () => {
    const clientId = config.clientID;
    const clientSecret = config.clientSecret;
    const url = `https://www.arcgis.com/sharing/oauth2/token?client_id=${clientId}&grant_type=client_credentials&client_secret=${clientSecret}&f=json`;
    const options = {method: 'GET'};
    const request = new Request(url, options);
    const response = await fetch(request);
    return response.json();
  };

  _coordsRequest = async (address, token) => {
    const url = `https://geocode.arcgis.com/arcgis/rest/services/World/` +
      `GeocodeServer/findAddressCandidates?f=json&SingleLine=${address}&` +
      `token=${token}&forStorage=true`;
    const options = {method: 'GET'};
    const request = new Request(url, options);
    const response = await fetch(request);
    return response.json();
  };

  getCoords = async (address) => {
    return this.auth().then((result) => {
      const token = result.access_token;
      return this._coordsRequest(address, token);
    });
  };
}

export default GeocodingService;