const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');
// here we actually load the proto content
const packageDefinition = protoLoader.loadSync('./proto/loyalty/loyalty.proto');
const loyaltyProto = grpc.loadPackageDefinition(packageDefinition);
const service = require('./sidecar.js');

const server = new grpc.Server();

server.addService(loyaltyProto.LoyaltyService.service, {
    listPartners: (_, callback) => {
        let partners = [{partnerId: "1234", address: "New Kingston"}];
        callback(null, partners)
    },
    getPartner: (call, callback) => {
        let partner = {partnerId: "1234", address: "New Kingston"};
        callback(null, partner)
    },
    getMember: (call, callback) => {
        let member = {memberId: "1234", firstName: "O'Dane", lastName: "Jackson", points: 5000.00}
        callback(null, member)
    },
    getPartnerTransaction: (call, callback) => {
        let partner = {partnerId: "1234", address: "New Kingston"};
        let member = {memberId: "1234", firstName: "O'Dane", lastName: "Jackson", points: 5000.00}
        let transactions = [{member: member, partner: partner, points: 400.00, transactionId: "dfcvfdfsderefwd3232", timestamp: new Date()}];
        callback(null, transactions)
    },
    getMemberTransaction: (call, callback) => {
        let partner = {partnerId: "1234", address: "New Kingston"};
        let member = {memberId: "1234", firstName: "O'Dane", lastName: "Jackson", points: 5000.00}
        let transactions = [{member: member, partner: partner, points: 400.00, transactionId: "dfcvfdfsderefwd3232", timestamp: new Date()}];
        callback(null, transactions)
    },
    usePoints: (call, callback) => {
        let partner = {partnerId: "1234", address: "New Kingston"};
        let member = {memberId: "1234", firstName: "O'Dane", lastName: "Jackson", points: 5000.00}
        let transaction = {member: member, partner: partner, points: 400.00, transactionId: "dfcvfdfsderefwd3232", timestamp: new Date()};
        callback(null, transaction)
    },
    earnPoints: (call, callback) => {
        let member = {memberId: "1234", firstName: "O'Dane", lastName: "Jackson", points: 5000.00}
        let transaction = [{member: member, points: 900.00, transactionId: "dfcvfdfsderefwd3232", timestamp: new Date()}];
        callback(null, transaction)
    },
});

service.register();

server.bind('127.0.0.1:50051',
    grpc.ServerCredentials.createInsecure());
console.log('Server running at http://0.0.0.0:50051');
server.start();