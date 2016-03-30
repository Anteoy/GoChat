var app = angular.module('gochat', []);

app.controller('chat',function($scope,$http){

})


var chat = new WebSocket("ws://localhost:8080/chat");
var chatfriends = new WebSocket("ws://localhost:8080/chatFriends");

