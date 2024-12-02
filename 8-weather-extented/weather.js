#!/usr/bin/env node
import { getArgs } from "./helpers/args.js"; 
import { fetchWeather } from "./services/api.service.js";
import { printHelp, printSuccess, printError, printWeather } from "./services/log.service.js";
import { saveData, TOKEN_DICTIONARY } from "./services/storage.service.js";

const saveToken = async (token) => {
    if(!token.length) {
        printError('Ключ не передан');
        return;
    }

    try {
        await saveData(TOKEN_DICTIONARY.token, token);
        printSuccess('Token saved')
        
    } catch (error) {
        console.log(printError(error))
    }
}

const saveLang = async (lang) => {
    if(lang !== 'ru' && lang !== 'eng') {
        printError('Не указан язык');
        return;
    }

    try {
        await saveData(TOKEN_DICTIONARY.lang, lang);
        printSuccess('Languege saved')
        
    } catch (error) {
        console.log(printError(error))
    }
}

const saveCity = async (city) => {
    if(!city.length) {
        printError('Город не передан');
        return;
    }

    try {
        await saveData(TOKEN_DICTIONARY.city, city);
        printSuccess('City saved')
        
    } catch (error) {
        console.log(printError(error))
    }
}

const weatherOutside = async () => {
    try {
        const weather = await fetchWeather(process.env.CITY);
        weather.forEach((weatherCity) => printWeather(weatherCity));
        
    } catch (error) {
        if(error?.responce?.code == 404) {
            printError('Неверно указан город')
        } else if(error?.responce?.code == 401) {
            printError('Неверно указан ключ')
        } else {
            printError(error.message)
        }
    }
}

const initCLI = () => {
    const argv = getArgs(process.argv);
    
    if(argv.h) {
        printHelp();
    }

    if(argv.s) {
        saveCity(argv.s);
    }

    if(argv.t) {
        saveToken(argv.t);
    }

    if(argv.l) {
        saveLang(argv.l);
    }

    weatherOutside();
}

initCLI();