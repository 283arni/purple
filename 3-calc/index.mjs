import { divide, multiply, add, subtract } from "./utils.mjs";
import { validateDivide, validateСommand, validateNum } from "./validate.mjs";

const command = process.argv[2];
const firstNum =  +process.argv[3];
const secondNum = +process.argv[4];

if(validateСommand(command) && validateNum(firstNum) && validateNum(secondNum) && validateDivide(command, secondNum)) {
    switch (command) {
        case 'add':
            add(firstNum, secondNum);
            break;
        case 'divide':
            divide(firstNum, secondNum);
            break;
    
        case 'multiply':
            multiply(firstNum, secondNum);
            break;
    
        case 'subtract':
            subtract(firstNum, secondNum);
            break;
    }
}


