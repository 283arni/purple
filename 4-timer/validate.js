
const validateTime = (time) => {
    return  time.every((paiceTime) => Number.isFinite(parseFloat(paiceTime)));
}

module.exports =  validateTime;