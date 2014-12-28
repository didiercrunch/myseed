define [
    "angularAMD"
    "angular-route"
    "foundation"
], (angularAMD) ->
    app = angular.module("webapp", ["ngRoute", "mm.foundation"]) # "['mm.foundation']"
    app.config ($routeProvider) ->
        $routeProvider.when("/view1", angularAMD.route(
            templateUrl: "partials/partial1.html"
            controller: "MyCtrl1"
            controllerUrl: "controllers/myctrl1"
        )).when("/view2", angularAMD.route(
            templateUrl: "partials/partial2.html"
            controller: "MyCtrl2"
            controllerUrl: "controllers/myctrl2"
        )).otherwise redirectTo: "/view1"
        return

    angularAMD.bootstrap app
