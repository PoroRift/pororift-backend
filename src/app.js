import express from 'express';
import { key } from '../config';

const app = express();

app.get('/', (req, res) => res.send('Hello World!' + key)); //how to access personal lol api key.

app.listen(3000, () => console.log('Listening to port 3000!'));