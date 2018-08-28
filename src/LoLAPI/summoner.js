

/*
	This file contains logic that handling summoner information such as:
		- Match History
		- Live Game 
		- Account information
*/

import {getSummonerStat, getMatches} from './api.js'
// import polyfill from 'babel-polyfill'

export let getSummoner = async (_summoner) => {
	let info = await getSummonerStat(_summoner);

	let match_history = await getMatches(info.accountId);

	let stat = {
		info: info,
		matches: match_history
	};

	return stat;
}