import chank from 'chalk';
import dedent from 'dedent-js';

const printError = (err) => {
    console.log(`${chank.bgRedBright('ERROR')} ${err}`);
}

const printSuccess= (message) => {
    console.log(`${chank.bgGreenBright('SUCCESS')} ${message}`);
}

const printHelp = () => {
    console.log(
        dedent`${chank.bgBlueBright('HELP')}
        -s [SITY] для установки города
        -t [KEY] для установки токена
        -h вызов команд
        `
    );
}


export { printError, printSuccess, printHelp}