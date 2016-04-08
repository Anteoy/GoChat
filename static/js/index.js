var app = angular.module('gochat', [] ,function ($httpProvider) {
    // Use x-www-form-urlencoded Content-Type
    $httpProvider.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded;charset=utf-8';

    /**
    * The workhorse; converts an object to x-www-form-urlencoded serialization.
    * @param {Object} obj
    * @return {String}
    */
    var param = function (obj) {
        var query = '', name, value, fullSubName, subName, subValue, innerObj, i;

        for (name in obj) {
            value = obj[name];

            if (value instanceof Array) {
                for (i = 0; i < value.length; ++i) {
                    subValue = value[i];
                    fullSubName = name + '[' + i + ']';
                    innerObj = {};
                    innerObj[fullSubName] = subValue;
                    query += param(innerObj) + '&';
                }
            }
            else if (value instanceof Object) {
                for (subName in value) {
                    subValue = value[subName];
                    fullSubName = name + '[' + subName + ']';
                    innerObj = {};
                    innerObj[fullSubName] = subValue;
                    query += param(innerObj) + '&';
                }
            }
            else if (value !== undefined && value !== null)
                query += encodeURIComponent(name) + '=' + encodeURIComponent(value) + '&';
        }

        return query.length ? query.substr(0, query.length - 1) : query;
    };

    // Override $http service's default transformRequest
    $httpProvider.defaults.transformRequest = [function (data) {
        return angular.isObject(data) && String(data) !== '[object File]' ? param(data) : data;
    } ];
})

app.filter('myMsgInRight',function(){
	return function(msg){
		if(msg) return "'right"
		return "left"
	}
})
app.controller('chat',function($scope,$http){

	$scope.msg = ""
	$scope.chatMsg = []
	$http({url:"/myid"}).success(function(data){
		$scope.myid = data
	})
	
	var chat = new WebSocket("ws://localhost:3333/ws/chat");
	chat.onopen = function(){
		alert("连接成功")
	}
	chat.onmessage = function(e){
		var msg = JSON.parse(e.data)
		if(typeof msg.msg == 'string'){
			if(msg.id == $scope.myid){
				msg.rightOrLeft = {float:"right"}
			}else{
				msg.rightOrLeft = {float:"left"}
			}
			$scope.chatMsg.push(msg)
		}
		if(typeof msg[0] == 'number'){
			$scope.onlineusers = msg
		}
		$scope.$apply()
	}
	chat.onclose = function(){
		alert("连接中断了")
	}

	$scope.sendMsg = function(){
		$http({
			url:"/chat" ,
			method:"post" ,
			data:{msg : $scope.msg}
		}).success(function(data){
			if(data == "success"){
				$scope.msg=""
			}else{
				$scope.chatMsg.push({data:"This message send in error!"})
			}
		}).error(function(data){
			$scope.chatMsg.push({data:"This message send in error!"})
		})
	}

})

document.onkeydown=function(event){
            var e = event || window.event || arguments.callee.caller.arguments[0];
            if(e && e.keyCode==13){ // 按 Enter 
                 $("#send").click()
               }
 }; 


//var chatfriends = new WebSocket("ws://localhost:8080/ws/chatFriends");

