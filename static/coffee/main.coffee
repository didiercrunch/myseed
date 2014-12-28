require.config
    baseUrl: "js"

    # alias libraries paths.  Must set 'angular'
    paths:
        "angular"      : "/bower_components/angular/angular.min"
        "angular-route": "/bower_components/angular-route/angular-route.min"
        "angularAMD"   : "/bower_components/angularAMD/angularAMD.min"


    # Add angular modules that does not support AMD out of the box, put it in a shim
    shim:
        angularAMD: ["angular"]
        "angular-route": ["angular"]


    # kick start application
    deps: ["app"]
