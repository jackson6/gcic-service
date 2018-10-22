const express = require('express');

const logger = require('morgan');

const bodyParser = require('body-parser');

const cors = require('cors')

const app = express();

app.use(logger('dev'));

app.use(cors())

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));

app.get('*', (req, res) => res.status(200).send({
    message: 'Welcome to the beginning of nothingness.',
}));

app.get('/partner', (req, res) => res.status(200).send({
    message: 'Welcome to the beginning of partner.',
}));

app.get('/api/member/:id', (req, res) => res.status(200).send({
    message: 'Welcome to the beginning of member.',
}));

app.get('/api/partner/transaction/:id', (req, res) => res.status(200).send({
    message: 'Welcome to the beginning of partner transaction.',
}));

app.get('/api/member/transaction/:id', (req, res) => res.status(200).send({
    message: 'Welcome to the beginning of member transactio.',
}));

app.post('/api/earn/points', (req, res) => res.status(200).send({
    message: 'Welcome to the beginning of earn points.',
}));

app.post('/api/use/points', (req, res) => res.status(200).send({
    message: 'Welcome to the beginning of use points.',
}));
module.exports = app;