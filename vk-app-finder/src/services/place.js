class PlaceService {
  get = async () => {
    const url = 'https://fezzo.ru/';
    const options = {method: 'GET'};
    const request = new Request(url, options);
    const response = await fetch(request);
    return response.json();
  };

  getCoords = async (address) => {
    const url = `http://search.maps.sputnik.ru/search/addr?q=${address}`;
    const options = {method: 'GET'};
    const request = new Request(url, options);
    const response = await fetch(request);
    return response.json();
  };
}

export default PlaceService;