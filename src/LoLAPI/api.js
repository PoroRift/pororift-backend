
/*
  This file contain ONLY LoL API and does not contain any logic for gathering data

  Every function here is a call to LoL API only and should not do any complicated logi
*/
import fs from 'fs';
import request from 'request'
import polyfill from 'babel-polyfill'

let URL_GETSTAT = 'https://na1.api.riotgames.com/lol/summoner/v3/summoners/by-name/'; // This need to be move out ot config file
let URL_MATCHES = 'https://na1.api.riotgames.com/lol/match/v3/matchlists/by-account/'
let URL_MATCHE = 'https://na1.api.riotgames.com/lol/match/v3/timelines/by-match/'
let URL_CHAMP = 'https://na1.api.riotgames.com//lol/platform/v3/champion-rotations/'


let key = fs.readFileSync('./api_key'); // Config file need to plan and implemented later

key = '?api_key=' + key;


export let getSummonerStat = async (_summonerName) => {

  try{
    let stat = await doRequest(`${URL_GETSTAT}${_summonerName}/${key}`, 'GET');
    return stat;
  } catch(err) {
    console.log(err.stack);
    return err.stack;
  }
};

export let getMatches = async (_accountID) => {
  try{
    let matches = await doRequest(`${URL_MATCHES}${_accountID}/${key}`, 'GET');
    return matches;
  } catch(err){
    console.log(err.stack);
    return err.stack;
  }
};


export let getMatch = async (_matchID) => {
  try{
    let match = await doRequest(`${URL_MATCHE}${_matchID}/${key}`, 'GET');
    return match;
  } catch(err){
    console.log(err.stack);
    return err.stack;
  }
};

export let getChamp = async () => {
  try{
    let champs = await doRequest(`${URL_CHAMP}${key}`, 'GET');
    return champs;
  } catch (err) {
    console.log(err.stack);
    return err.stack;
  }
};


function doRequest(_url, _method){
  console.log(_url);
  return new Promise( (resolve, reject) => {
    try{
      var options = {
        url: _url,
        method: _method,
        headers: {
          'Content-Type': 'application/json'
        },
        json: true
      };

      request(options, (err, res, body)=>{
        if (res && (res.statusCode === 200 || res.statusCode === 201)) {
          // console.log(" response ", res.body.rates.INR);
          resolve(res.body);
        } else {
          console.log('error body', body);
          reject(new Error(`Error body: ${JSON.stringify(body)}`));
        }
      });
    } catch (err) {
      console.log('error request', err);
      process.exit();
    }
  });
}

// Just in case i need to calculate circumference...
const pi = 3.14159265359;
export var calcCircumference = (radius) => {
    return 2 * radius * pi;
};

  