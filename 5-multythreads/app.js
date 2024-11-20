const { performance, PerformanceObserver } = require('perf_hooks');
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
    return new Promise((resolve, rejects) => {

        const worker = new Worker('./worker', {
            workerData: { arr }
        });

        worker.on('message', (msg) => {
            resolve(msg);
        });

        worker.on('error', (e) => {
            rejects(e);
        });
    })
}

const numsLengthMulty = async () =>  {
    performance.mark('multy start');

    const numsArr = numsArrays(50000);

    const res = await Promise.all([
        compute(numsArr[0]),
        compute(numsArr[1]),
        compute(numsArr[2]),
        compute(numsArr[3]),
        compute(numsArr[4]),
        compute(numsArr[5]),
    ]);

    performance.mark('multy end');
    performance.measure('multy', 'multy start', 'multy end');

    console.log(`Без остатка на 3 делятся: ${res.reduce(( acc, num) => acc + num, 0)}`);
}

const main = () => {
    numsLengthSimple();
    numsLengthMulty();
}

main();
