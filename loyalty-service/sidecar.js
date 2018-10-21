const uuid4 = require('uuid/v4');
const request = require('request');

module.exports.register = function () {
    let register_uri = "http://localhost:8081/registry";
    let service = "gcic.loyalty";
    let headers = {'content-type': 'application/json'};
    let payload = {
        "name": service,
        "nodes": [{
            "id": service + "-" + uuid4(),
            "address": "127.0.0.1",
            "port": 50051,
        }],
    };
    request({
        url: register_uri,
        method: "POST",
        json: true,   // <--Very important!!!
        headers: headers,
        body: payload
    }, function (error, response, body){
        //console.log("response", response);
    });
};
