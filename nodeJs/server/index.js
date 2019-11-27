const express = require('express');

const port = process.env.PORT || 8080;

const app = express();

app.use('/', message);

app.listen(port, () => {
    console.log(`Orchestrator listening on port ${port}`);
});

async function message (req, res) {
    console.log('Http request Received');
    res.send("Nodejs says: wE lOvE KUbErNetEs")
}

module.exports = app;