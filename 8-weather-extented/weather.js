#!/usr/bin/env node
import { getArgs } from "./helpers/args.js"; 
import { printHelp } from "./services/log.service.js";

const initCLI = () => {
    const argv = getArgs(process.argv);
    
    if(argv.h) {
        printHelp()
    }

    if(argv.s) {

    }

    if(argv.t) {

    }
}

initCLI();