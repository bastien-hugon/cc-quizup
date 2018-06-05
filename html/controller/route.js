var routageModule = angular.module('routageModule', ['ngRoute']);

routageModule.config(['$routeProvider', function ($routeProvider) {
    $routeProvider
    .when('/', {
        templateUrl: "/view/index.html",
        controller: "login"
    })
    .when('/ready', {
        templateUrl: "/view/ready.html",
        controller: "ready"
    })
    .otherwise({
        redirectTo: '/'
    })
}]);