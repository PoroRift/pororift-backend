
import fs from 'fs';
import { champRot } from './api';

// let DIR = "./data/content/data/en_US/"
let DIR = `${process.cwd()}/data/content/data/en_US`

let champFileList = fs.readdirSync(`${DIR}/champion`);

let championInfo = readJsonFile(`${DIR}/championFull.json`);

let champImg = `${process.cwd()}/data/content/img/champion`

let championMap = init();



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
  return championMap.list;
}

export let getChampIcon = async(id) => {
  let name = championInfo.keys[id];
  
  return fs.readFileSync(`${champImg}/${name}.png`)
};

export let getRotation = async() => {
  let champ_rotation = await champRot();

  let freeChamp = champ_rotation["freeChampionIds"].map( x => championMap.map[x]);
  let freeChamp_new = champ_rotation["freeChampionIdsForNewPlayers"].map( x => championMap.map[x]);

  // Deep clone of JSON object (Following immutable methodology)
  let newChamp_rotation = JSON.parse(JSON.stringify(champ_rotation));
  newChamp_rotation["freeChampionIds"] = freeChamp;
  newChamp_rotation["freeChampionIdsForNewPlayers"] = freeChamp_new;

  return newChamp_rotation;
};


export let getImages = async() => {
  let list = {
    "loading": fs.readdirSync(`${process.cwd()}/data/img/champion/loading`),
    "splash" : fs.readdirSync(`${process.cwd()}/data/img/champion/splash`),
    "tiles"  : fs.readdirSync(`${process.cwd()}/data/img/champion/tiles`)
  }

  return list;
}



function init(){

  let champList = [];
  let champMap = {};
  for (var key in championInfo.keys){
    let id = key;
    let value = championInfo.keys[id];
    let champ_json = {id: key, name: value};
    champList.push(champ_json);
    champMap[key] = champ_json;
  }
  return {list: champList, map: champMap };
}


function readJsonFile(_path){

  let readFile = fs.readFileSync(_path);
  let jsonData = JSON.parse(readFile.toString());
  return jsonData;
}