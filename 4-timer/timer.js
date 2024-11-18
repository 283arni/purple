const EventEmitter = require('events');

const events = new EventEmitter();

const initTimer = (time) => {
    const hours = parseFloat(time[0]) * 3600000 || 0;
    const minutes = parseFloat(time[1]) * 60000 || 0;
    const seconds = parseFloat(time[2]) * 1000 || 0;
    const fullTime = hours + minutes + seconds;
    
    const timer = setTimeout(() => {
        console.log(`Таймер сработал через ${time.join(' ')} !`);
        clearTimeout(timer); 
    }, fullTime);
}

events.on('initTimer', initTimer);
events.emit('initTimer', process.argv.slice(2))
