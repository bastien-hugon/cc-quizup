quizzUpController.controller('login', ['$scope', '$location', '$rootScope', '$route', '$http', function($scope, $location, $rootScope, $route, $http) {
    $scope.get_ready = function(){
/*         $http.post('/someUrl', data, config).then(successCallback, errorCallback);
 */        $location.path('/ready')
    }
}]);