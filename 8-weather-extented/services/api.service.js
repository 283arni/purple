import { getData, TOKEN_DICTIONARY } from "./storage.service.js";
import axios from 'axios';

const getIcon = (icon) => {
	switch (icon.slice(0, -1)) {
		case '01':
			return '☀️';
		case '02':
			return '🌤️';
		case '03':
			return '☁️';
		case '04':
			return '☁️';
		case '09':
			return '🌧️';
		case '10':
			return '🌦️';
		case '11':
			return '🌩️';
		case '13':
			return '❄️';
		case '50':
			return '🌫️';
	}
};

const fetchData = async (town, token, lang) => {
		const {data} = await axios.get('https://api.openweathermap.org/data/2.5/weather',{
			params: {
				appid: token,
				q: town,
				lang: lang,
				units: 'metric',
			}
		}) 

		return data;
	}

const fetchWeather = async () => {
    const token = await getData(TOKEN_DICTIONARY.token);
    const townsStr = await getData(TOKEN_DICTIONARY.city);
	const lang = await getData(TOKEN_DICTIONARY.lang) || 'ru';
	
    if(!token) {
        throw new Error('Не указан ключ, указать можно с помощью -t [KEY]')
    }

	const towns = await townsStr.split(',')

	const data = await Promise.all(towns.map((town) =>  fetchData(town, token, lang)))


    return data;
}

export {fetchWeather, getIcon}