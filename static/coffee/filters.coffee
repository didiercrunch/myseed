define ["angular", "services"], (angular, services) ->

  # Filters
  angular.module("myApp.filters", ["myApp.services"]).filter "interpolate", [
    "version"
    (version) ->
      return (text) ->
        String(text).replace /\%VERSION\%/g, version
  ]
  return
