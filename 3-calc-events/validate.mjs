const commandsAll = ['add', 'divide', 'multiply', 'subtract']


const validateСommand = (command) => {
    return commandsAll.find((str) => str === command)
}

const validateDivide = (command, num) => {
    if(command === 'divide' && num <= 0) {
        return false;
    }

    return true;
}

const validateNum = (num) => {
    return  Number.isFinite(num);
}

export { validateDivide, validateСommand, validateNum }