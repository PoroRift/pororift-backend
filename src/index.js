// const express = require('express');
// const app = express();

import express from 'express';
import {getSummonerStat} from './LoLAPI/api.js'; // This going to be taken out
import {getSummoner} from './LoLAPI/summoner.js'



console.log(process.cwd());
const app = express();

app.get('/', (req, res) => res.send('Hello World!'));

app.get('/url', (req, res) => {
	res.json({ 'return' : "Test API"});
})

app.get('/stat/:summoner', (req,res) => {

	getSummonerStat(req.params.summoner).then((s) => {
		res.json(s);
	});

});


app.get('/summoner/:summoner', (req, res) => {
		getSummoner(req.params.summoner).then( (s) => {
			res.json(s);
		})
})

app.listen(3000, () => console.log('Listening to port 3000!'));



/*********** TESTING ASYNC & ES8 ************/
function bar() {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            resolve('hello');
        }, 3000);
    });
}
app.get('/es8', (req, res)=>{
	
	let test = async () =>{
		console.log("Test");
		await bar()
		console.log("THIS")
	} 

	test();
	res.send("test");
});
