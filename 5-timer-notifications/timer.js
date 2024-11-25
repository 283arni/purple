const EventEmitter = require('events');
const validateTime = require('./validate.js');
const notifier = require('node-notifier');
const path = require('path');

const events = new EventEmitter();

const time = process.argv.slice(2);



const initTimer = (time) => {
    const hours = parseFloat(time[0]) * 3600000 || 0;
    const minutes = parseFloat(time[1]) * 60000 || 0;
    const seconds = parseFloat(time[2]) * 1000 || 0;
    const fullTime = hours + minutes + seconds;

    const timer = setTimeout(() => {
        notifier.notify(
            {
              title: 'Уведомление',
              message: `Таймер сработал через ${time.join(' ')} !`,
              icon: path.join(`${__dirname}/img`, 'noty.png'), 
              sound: true,
              wait: true 
            }
        );
        clearTimeout(timer);
    }, fullTime);
}



events.on('initTimer', initTimer);

if(validateTime(time)) {
    events.emit('initTimer', time)
}

