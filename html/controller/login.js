quizzUpController.controller('login', ['$scope', '$location', '$rootScope', '$route', '$http', function($scope, $location, $rootScope, $route, $http) {
    $scope.get_ready = function(){
        var teamName = angular.element("#teamName")[0].value

        if (teamName == '' || teamName == undefined) {
            alert('Empty team name')
            return;
        }
/*      $http.post('/someUrl', data, config).then(successCallback, errorCallback);
*/      $location.path('/ready')
    }
}]);