
import fs from 'fs';

// let DIR = "./data/content/data/en_US/"
let DIR = `${process.cwd()}/data/content/data/en_US`

let champFileList = fs.readdirSync(`${DIR}/champion`);

let championInfo = readJsonFile(`${DIR}/championFull.json`);

let champImg = `${process.cwd()}/data/content/img/champion`


export let getChampByKey = async (key) => {
  let name = championInfo.keys[key];
  return  championInfo.data[name];
};

export let getChampByName = async(name) => {
  for (let k in championInfo.keys){
    if (k.toLowerCase() == k.toLowerCase()) {
      let key = championInfo.keys[k];
      return championInfo.data[key]
    }
  }
  return { "Msg" : `Cannot find champion by the name ${name}`};
}

export let getChampionIDs = async() => {
  return championInfo.keys;
}

export let getChampIcon = async(id) => {
  let name = championInfo.keys[id];
  
  return fs.readFileSync(`${champImg}/${name}.png`)
};


export let getImages = async() => {
  let list = {
    "loading": fs.readdirSync(`${process.cwd()}/data/img/champion/loading`),
    "splash" : fs.readdirSync(`${process.cwd()}/data/img/champion/splash`),
    "tiles"  : fs.readdirSync(`${process.cwd()}/data/img/champion/tiles`)
  }

  return list;
}



function readJsonFile(_path){

  let readFile = fs.readFileSync(_path);
  let jsonData = JSON.parse(readFile.toString());
  return jsonData;
}