import add from "./add.mjs";
import divide from "./divide.mjs";
import multiply from "./multiply.mjs";
import subtract from "./subtract.mjs";

switch (process.argv[4]) {
    case 'add':
        add(+process.argv[2], +process.argv[3]);
        break;
    case 'divide':
        divide(+process.argv[2], +process.argv[3]);
        break;

    case 'multiply':
        multiply(+process.argv[2], +process.argv[3]);
        break;

    case 'subtract':
        subtract(+process.argv[2], +process.argv[3]);
        break;
}
