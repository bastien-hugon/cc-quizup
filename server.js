var app = require('express')();
var express = require('express');
var server = require('http').Server(app);
var io = require('socket.io')(server);
var fs = require('fs');
var util = require('util');

// Launch Express Servee
server.listen(80, function () {
	console.log('Server running on port 80');
});

app.use('/', express.static(__dirname + '/html/'));

// Get API Content
eval(fs.readFileSync('assets/api/api.js')+'');