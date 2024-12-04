const { parentPort, workerData} = require('worker_threads');
const { numsLengthWorker } = require('./nums.js');

parentPort.postMessage(numsLengthWorker(workerData))