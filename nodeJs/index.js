const express = require('express');

const port = process.env.PORT || 8080;
const url = process.env.url || "http://localhost:8080";


const app = express();
const axios = require('axios');

async function message (req, res) {
    res.send("Nodejs says: wE lOvE KUbErNetEs")
}

async function sendHttpCall(){
    axios.get(url).then(response => console.log(response.data))
}

app.use('/', message);

app.listen(port, () => {
    console.log(`Orchestrator listening on port ${port}`);
});

sendHttpCall();



module.exports = app;