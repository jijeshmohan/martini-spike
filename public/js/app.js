// Declare app level module which depends on filters, and services
angular.module('myapp', ['ngResource', 'ngRoute', 'ui.bootstrap', 'ui.date'])
  .config(['$routeProvider', function ($routeProvider) {
    $routeProvider
      .when('/', {
        templateUrl: 'views/home/home.html', 
        controller: 'HomeController'})
      .otherwise({redirectTo: '/'});
  }])
  .directive('blogActiveLink', ['$location', function(location) {
    return {
      restrict: 'A',
      link: function(scope, element, attrs, controller) {
        var clazz = attrs.blogActiveLink;
        var path = attrs.href;
        path = path.substring(2); 
        scope.location = location;
        scope.$watch('location.path()', function(newPath) {
          if (path === newPath) {
            element.parent().addClass(clazz);
          } else {
            element.parent().removeClass(clazz);
          }
        });
      }
    };
  }]) 
  .directive('markdown', function ($http) {
    var converter = new Showdown.converter();
      
    return {
      restrict: 'E',
      require: 'model',
      scope: {
         post: '=model'
      },
      template:  '<div>{{post.body | markdown}}</div>'
    };
  }).filter('markdown', function () {
    var converter = new Showdown.converter();
    return function (value) {
        return converter.makeHtml(value || '');
    };
});