import { divide, multiply, add, subtract } from "./utils.mjs";
import { validateDivide, validateСommand, validateNum } from "./validate.mjs";
import EventEmitter from "events";

const events = new EventEmitter();

const command = process.argv[2];
const firstNum =  +process.argv[3];
const secondNum = +process.argv[4];

events.on('add', add);
events.on('divide', divide);
events.on('multiply', multiply);
events.on('subtract', subtract);

if(validateСommand(command) && validateNum(firstNum) && validateNum(secondNum) && validateDivide(command, secondNum)) {
    switch (command) {
        case 'add':
            events.emit('add', firstNum, secondNum);
            break;
        case 'divide':
            events.emit('divide', firstNum, secondNum);
            break;
    
        case 'multiply':
            events.emit('multiply', firstNum, secondNum);
            break;
    
        case 'subtract':
            events.emit('subtract', firstNum, secondNum);
            break;
    }
}
