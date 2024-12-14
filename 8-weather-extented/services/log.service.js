import chank from 'chalk';
import dedent from 'dedent-js';
import { getData, TOKEN_DICTIONARY } from './storage.service.js';

import { getIcon } from './api.service.js';

const printError = (err) => {
    console.log(`${chank.bgRedBright('ERROR')} ${err}`);
}

const printSuccess= (message) => {
    console.log(`${chank.bgGreenBright('SUCCESS')} ${message}`);
}

const printWeather = async (weather) => {
    const { weather: sky, main, wind, name } = weather;

    if( await getData(TOKEN_DICTIONARY.lang) == 'eng') {
        console.log(
            dedent`${chank.bgGreen(`Today's weather overview: ${name}`)}
            Today on the street ${chank.blueBright(sky[0].description)} ${getIcon(sky[0].icon)}
            Temperature ${chank.redBright(Math.floor(main.temp))} deg.
            Wind ${chank.yellowBright(Math.floor(wind.speed))} m/s
            `
        );
    } else {
        console.log(
            dedent`${chank.bgGreen(`Обзор погоды на сегодня: ${name}`)}
            Сегодня на улице ${chank.blueBright(sky[0].description)} ${getIcon(sky[0].icon)}
            Температура ${chank.redBright(Math.floor(main.temp))} гр.
            Ветер ${chank.yellowBright(Math.floor(wind.speed))} м/с
            `
        );
    }

}

const printHelp = () => {
    console.log(
        dedent`${chank.bgBlueBright('HELP')}
        -s [SITY] для установки города, несколько годов можно установить через запятую без пробелов. Пример: ${chank.bold('moscow,omsk')}
        -t [KEY] для установки токена
        -l [LANG] ru или eng
        -h список команд
        `
    );
}


export { printError, printSuccess, printHelp, printWeather}