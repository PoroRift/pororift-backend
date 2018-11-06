// const express = require("express");
// const app = express();

import polyfill from 'babel-polyfill'
import express from "express";
import path from 'path';
import { getMatch } from "./LoLAPI/api.js"; 
import { getSummoner } from "./LoLAPI/summoner.js"
import fs from 'fs';

// console.log(process.cwd());

const app = express();

app.get("/", (req, res) => res.send("Hello World!"));

app.get("/url", (req, res) => {
  res.json({ "return" : "Test API"});
});


app.get("/summoner/:summoner", (req, res) => {
  getSummoner(req.params.summoner).then( (s) => {
    res.json(s);
  });
});

app.get("/match/:matchid", (req, res) => {
  getMatch(req.params.matchid).then( (m) => {
    res.json(m);
  });
});


/*** Get Champions ***/
import { getChampList, getChampionIDs, getChampByKey, getChampByName, getChampIcon, getImages, getRotation } from "./LoLAPI/champion.js"

app.get("/champ_by_key/:key", (req, res) => {
	getChampByKey(req.params.key).then( (champ) => {
		res.json(champ);
	});
});

app.get("/champ_by_name/:name", (req, res) => {
	getChampByName(req.params.name).then( (champ) => {
		res.json(champ);
	});
});

app.get("/championList", (req, res) => {
  getChampionIDs().then( (list) => {
    res.json(list);
  })
});

app.get("/champ_icon/:id", (req, res) => {
  getChampIcon(req.params.id).then( (img) => {
    res.writeHead(200, {'Content-Type': 'image/png'});
    res.end(img, 'binary');
  });
});

app.get("/images", (req, res) => {
  getImages().then( (j) => {
    res.json(j);
  })
});

app.get("/champ_rotation", (req, res) => {
  getRotation().then( (rot) => {
    return res.json(rot);
  })
});


/** API KEY **/

app.get("/update_key", (req, res) => {
  if(req.query.key != undefined){
    if(req.query.key.length != 42){
      res.json({message: "Invalid Key"});
    } else {
      let path = `${process.cwd()}/api_key`;
      fs.writeFileSync(path, req.query.key);
      res.json({message: "Update Key Successfully"});
    }
    
  } else {
    res.json({message: "Empty Key"});
  }
});


app.use('/champion', express.static(path.join(process.cwd(), 'data/img/champion')));



app.listen(3000, () => console.log("Listening to port 3000!"));


	
/*********** TESTING ASYNC & ES8 ************/
function bar() {
    return new Promise( (resolve, reject) => {
        setTimeout(() => {
            resolve("hello");
        }, 3000);
    });
}
app.get("/es8", (req, res)=>{
  
  let test = async () =>{
    console.log("Test");
    await bar()
    console.log("THIS")
  } 

  test();
  res.send("test");
});
