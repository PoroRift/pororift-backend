

import fs from 'fs';
import request from 'request'
import polyfill from 'babel-polyfill'

let url_getStat = "https://na1.api.riotgames.com/lol/summoner/v3/summoners/by-name/"; // This need to be move out ot config file

let key = fs.readFileSync('./api_key');

key = '?api_key=' + key;
console.log(key.toString());


export let getSummonerStat = async (_summonerName) => {
	// let stat = null;
	// await request(`${url_getStat}rubberice/${key}`, (err, response, body) => {
	// 	console.log(response);
	// 	stat = response
	// });


	// try{
	// 	let stat =  await 
	// 	console.log(stat.response);
	// 	return stat;
	// } catch(err) {
	// 	return err;
	// }

	try{
		let stat = await doRequest(`${url_getStat}${_summonerName}/${key}`, 'GET');
		// console.log(stat);
		return stat;
	} catch(err) {
		return err;
	}
}


function doRequest(_url, _method){
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

				console.log("Do Request");
				if (res && (res.statusCode === 200 || res.statusCode === 201)) {
					// console.log(' response ', res.body.rates.INR);
					resolve(res.body);
				} else {
					console.log('error body', body);
					reject(new Error(`Error body: ${JSON.stringify(body)}`));
				}
			});
		} catch (err) {
			console.log(" error request ", err);
			process.exit();
		}
	})
}

const pi = 3.14159265359;
export var calcCircumference = (radius) => {
    return 2 * radius * pi;
};

