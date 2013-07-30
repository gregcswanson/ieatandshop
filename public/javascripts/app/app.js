function MessagesController($scope){
	$scope.errors = []; //[{"message": "test"}];
	$scope.messages = []; //[{"message": "test"}];
	$scope.infos = []; //[{"message": "test"}];
	$scope.warnings = []; //[{"message": "test"}];
	
	$scope.clearError = function(item) {
		var index=$scope.errors.indexOf(item)
  		$scope.errors.splice(index,1); 
	};
	
	$scope.clearMessage = function(item) {
		var index=$scope.messages.indexOf(item)
  		$scope.messages.splice(index,1); 
	};
	
	$scope.clearInfo = function(item) {
		var index=$scope.infos.indexOf(item)
  		$scope.infos.splice(index,1); 
	};
	
	$scope.clearWarning = function(item) {
		var index=$scope.warnings.indexOf(item)
  		$scope.warnings.splice(index,1); 
	};
	
	$scope.$on('error-message', function(event, data) {
    	$scope.errors.push({"message": data});
	});
	
	$scope.$on('info-message', function(event, data) {
    	$scope.infos.push({"message": data});
	});
	
	$scope.$on('warning-message', function(event, data) {
    	$scope.warnings.push({"message": data});
	});
	
	$scope.$on('message-message', function(event, data) {
    	$scope.messages.push({"message": data});
	});
}