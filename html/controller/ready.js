quizzUpController.controller('ready', ['$scope', '$location', '$rootScope', '$route', '$http', function($scope, $location, $rootScope, $route, $http) {
    $scope.get_all_team_connected = function(){
        
    }
    $scope.getReady = function () {
        var button = $("#readyButton")

        if (button.text() == "Ready") {
            button.attr("class", "uk-button uk-button-danger uk-align-center")
            button.text("Not Ready") 
        } else {
            button.attr("class", "uk-button uk-label-success uk-align-center")
            button.text("Ready")
        }
    }
}]);