define ["app"], (app) ->
    app.controller "MyCtrl2", ($scope) ->
        $scope.message = "Message from MyCtrl2"
        $scope.oneAtATime = true
        $scope.groups = [
            {
                title: "Dynamic Group Header - 1"
                content: "Dynamic Group Body - 1"
            }
            {
                title: "Dynamic Group Header - 2"
                content: "Dynamic Group Body - 2"
            }
        ]

        $scope.message = "Message from MyCtrl2"

        $scope.items = [
            "Item 1"
            "Item 2"
            "Item 3"
        ]

        $scope.addItem = ->
            newItemNo = $scope.items.length + 1
            $scope.items.push("Item " + newItemNo)
            return


        return
