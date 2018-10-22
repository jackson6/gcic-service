const service        = require('./sidecar.js');
const app             = require('./app');
const http            = require('http');
const port = parseInt(process.env.PORT, 10) || 50051;


service.register();
app.set('port', port);

const server = http.createServer(app);
server.listen(port);