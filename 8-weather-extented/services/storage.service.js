import { homedir } from 'os';
import { join } from 'path';
import { promises } from 'fs';


const TOKEN_DICTIONARY = {
    token: 'token',
    city: 'city',
    lang: 'lang'
}
const pathData = join(homedir(), 'w-data.json');

const saveData = async ( key, value) => {
    let data = {}

    if( await isExist(pathData)) {
        const file = await promises.readFile(pathData);
        data = JSON.parse(file)
    }

    data[key] = value;

    await promises.writeFile(pathData, JSON.stringify(data))
}

const getData = async (key) => {
    if( await isExist(pathData)) {
        const file = await promises.readFile(pathData);
        const data = JSON.parse(file);

        return data[key]
    }

    return undefined;
}

const isExist = async (path) => {
    try {
        await promises.stat(path);
        return true
    } catch (error) {
        return false;
    }
}

export { saveData, getData, TOKEN_DICTIONARY }