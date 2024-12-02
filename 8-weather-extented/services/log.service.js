import chank from 'chalk';
import dedent from 'dedent-js';

import { getIcon } from './api.service.js';

const printError = (err) => {
    console.log(`${chank.bgRedBright('ERROR')} ${err}`);
}

const printSuccess= (message) => {
    console.log(`${chank.bgGreenBright('SUCCESS')} ${message}`);
}

const printWeather = (weather) => {
    const { weather: sky, main, wind, name } = weather;
    console.log(
        dedent`${chank.bgGreen(`Обзор погоды на сегодня: ${name}`)}
        Сегодня на улице ${chank.blueBright(sky[0].description)} ${getIcon(sky[0].icon)}
        Температура ${chank.redBright(Math.floor(main.temp))} гр.
        Ветер ${chank.yellowBright(Math.floor(wind.speed))} м/с
        `
    );
}

const printHelp = () => {
    console.log(
        dedent`${chank.bgBlueBright('HELP')}
        -s [SITY] для установки города
        -t [KEY] для установки токена
        -l [LANG] ru или eng
        -h список команд
        `
    );
}


export { printError, printSuccess, printHelp, printWeather}