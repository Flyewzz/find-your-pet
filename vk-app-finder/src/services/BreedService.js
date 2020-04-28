import config from '../config';

class BreedService {
  __toBase64 = file => new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.readAsDataURL(file);
    reader.onload = () => resolve(reader.result);
    reader.onerror = error => reject(error);
  });

  getBreeds = async (image) => {
    const base64 = await this.__toBase64(image);
    const url = config.baseUrl + 'breed';
    const headers = {'Content-type': 'application/json'};

    const options = {
      method: 'POST',
      headers: headers,
      body: JSON.stringify({picture: base64}),
    };

    const request = new Request(url, options);
    
    const response = await fetch(request);
    return response.json();
  };
}

export default BreedService;