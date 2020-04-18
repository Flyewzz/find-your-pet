require('dotenv').load();

const config = {
  baseUrl: process.env.BASE_URL , // '/' at the end is strictly necessary (!)
  clientID: process.env.CLIENT_ID,
  clientSecret: process.env.CLIENT_SECRET,
  serviceKey: process.env.SERVICE_KEY,
  types: ["Собака", "Кошка", "Животное"],
  appUrl: process.env.APP_URL,
  appId: process.env.APPLICATION_ID,
};

export default config;
