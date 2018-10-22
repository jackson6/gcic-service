const net = require('net');
const client = new net.Socket();
const port = 50051;
const host = '127.0.0.1';

client.connect(port, host, function() {
    console.log('Connected');
    client.write("Hello From Client " + client.address().address);
});