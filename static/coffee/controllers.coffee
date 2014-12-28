define ["angular", "services"], (angular) ->

  # Controllers

  # Sample controller where service is being used

  # More involved example where controller is required from an external file
  controllers = angular.module("myApp.controllers", ["myApp.services"])

  controllers.controller("MyCtrl1", [
    "$scope"
    "$injector"
    ($scope, $injector) ->
      require ["controllers/myctrl1"], (myctrl1) ->

        # injector method takes an array of modules as the first argument
        # if you want your controller to be able to use components from
        # any of your other modules, make sure you include it together with 'ng'
        # Furthermore we need to pass on the $scope as it's unique to this controller
        $injector.invoke(myctrl1, this, {$scope: $scope})

        return
  ])

  controllers.controller( "MyCtrl2", [
    "$scope"
    "$injector"
    ($scope, $injector) ->
      require ["controllers/myctrl2"], (myctrl2) ->

        # injector method takes an array of modules as the first argument
        # if you want your controller to be able to use components from
        # any of your other modules, make sure you include it together with 'ng'
        # Furthermore we need to pass on the $scope as it's unique to this controller
        $injector.invoke(myctrl2, this, {$scope: $scope})

        return
  ])
