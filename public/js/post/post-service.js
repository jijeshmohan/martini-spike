'use strict';

angular.module('myapp')
  .factory('Post', ['$resource', function ($resource) {
    return $resource('myapp/posts/:id', {}, {
      'query': { method: 'GET', isArray: true},
      'get': { method: 'GET'},
      'update': { method: 'PUT'}
    });
  }]);
