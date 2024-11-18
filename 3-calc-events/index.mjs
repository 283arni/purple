import add from "./add.mjs";
import divide from "./divide.mjs";
import multiply from "./multiply.mjs";
import subtract from "./subtract.mjs";
import EventEmitter from "events";

const events = new EventEmitter();

const nums = process.argv.slice(2, 4);
const name = process.argv.slice(4)[0];

events.on('add', add);
events.on('divide', divide);
events.on('multiply', multiply);
events.on('subtract', subtract);

switch (name) {
    case 'add':
        events.emit('add', nums);
        break;
    case 'divide':
        events.emit('divide', nums);
        break;

    case 'multiply':
        events.emit('multiply', nums);
        break;

    case 'subtract':
        events.emit('subtract', nums);
        break;
}
