function EchoController($scope,$http, $rootScope){

	$scope.value = []; //{"Title": "Hello", "Lines": [{"Id": 1, "Name": "One"},{"Id": 2, "Name": "Two"}]};
	$scope.valueRaw = JSON.stringify($scope.value);

	$scope.post = function(){
		var postData = JSON.stringify($scope.value);
		console.log(postData);
		$http.post('/api/echo/post', $scope.value).success(function(data){
			console.log(data.Data);
			$rootScope.$broadcast('message-message', "Saved");
			$scope.value = data.Data;
			$scope.valueRaw = JSON.stringify($scope.value);
		}).error(function(data){
			console.log(data);
			$rootScope.$broadcast('error-message', data);
			$scope.valueRaw = data;
		});;
	}
	
	$scope.get = function(){
		$http.get('/api/echo/get').success(function(data){
			console.log(data);
			$scope.value = data.Data;
			$scope.valueRaw = JSON.stringify($scope.value);
		}).error(function(data){
			console.log(data);
			$rootScope.$broadcast('error-message', data);
			$scope.valueRaw = data;
		});
	}
	
	$scope.add = function(){
		$scope.value.Lines.push({"Id": "", "Name": "new"});
		$scope.valueRaw = JSON.stringify($scope.value);
	}
	
	$scope.remove = function(index)
	{
		var deletedLine = $scope.value.Lines[index];
        $scope.value.Lines.splice(index, 1);
        $scope.valueRaw = JSON.stringify($scope.value);
	}
	
	$scope.get();

}

