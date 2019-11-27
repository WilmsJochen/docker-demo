const url = process.env.URL || "http://localhost:8080";
const axios = require('axios');

async function sendHttpCall(){
    axios.get(url).then(response => console.log(response.data))
}

console.log("Request send to URL: ", url);
sendHttpCall();


