const numsLength = () => {
    const nums = new Array(300000).fill(1).map((n, i) => n + i);
    return nums.filter((num) => num % 3 === 0).length;
}

const numsArrays = (chunkSize) => {
    const nums = new Array(300000).fill(1).map((n, i) => n + i);
    const res = [];

    for (let i = 0; i < nums.length; i += chunkSize) {
        const chunk = nums.slice(i, i + chunkSize);
        res.push(chunk);
    }

    return res;
}

const numsLengthWorker = ({ arr }) => {
    return arr.filter((num) => num % 3 === 0).length;
}

module.exports = { numsLength, numsArrays, numsLengthWorker } 