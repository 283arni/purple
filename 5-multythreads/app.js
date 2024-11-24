const { performance, PerformanceObserver } = require('perf_hooks');
const os = require('node:os');
const { Worker } = require('worker_threads');
const { numsLength, numsArrays }  = require('./nums.js');

const performanceObserver = new PerformanceObserver((items) => {
    items.getEntries().forEach((item) => {
        console.log(`${item.name}: ${item.duration}`);
    });
})

performanceObserver.observe({entryTypes: ['measure']});

const numsLengthSimple = () =>  {
    performance.mark('simple start');


    const length = numsLength();

    performance.mark('simple end');
    performance.measure('simple', 'simple start', 'simple end');
    
    return length;
}

const compute = (arr) => {
    return new Promise((resolve, reject) => {

        const worker = new Worker('./worker', {
            workerData: { arr }
        });

        worker.on('message', (msg) => {
            resolve(msg);
        });

        worker.on('error', (e) => {
            reject(e);
        });
    })
}

const numsLengthMulty = async () => {
    performance.mark('multy start');

    const numCPUs = os.cpus().length;
    const numsArr = numsArrays(Math.ceil(300000 / numCPUs));

    const res = await Promise.all(numsArr.map((chunk) => compute(chunk)));

    performance.mark('multy end');
    performance.measure('multy', 'multy start', 'multy end');

    console.log(`Без остатка на 3 делятся (multy): ${res.reduce((acc, num) => acc + num, 0)}`);
};

const main = () => {
    numsLengthSimple();
    numsLengthMulty();
}

main();
