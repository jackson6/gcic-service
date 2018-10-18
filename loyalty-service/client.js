var grpc = require('grpc');

var loyaltyProto = grpc.load('loyalty.proto');

var client = new loyaltyProto.LoyaltyService('gcic.loyalty',
    grpc.credentials.createInsecure());

console.log(client)

client.listPartners({}, function(error, partners) {
    if (error)
        console.log('Error: ', error);
    else
        console.log(partners);
});