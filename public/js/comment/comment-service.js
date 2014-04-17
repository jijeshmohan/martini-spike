'use strict';

angular.module('myapp')
  .factory('Comment', ['$resource', function ($resource) {
    return $resource('myapp/comments/:id', {}, {
      'query': { method: 'GET', isArray: true},
      'get': { method: 'GET'},
      'update': { method: 'PUT'}
    });
  }]);
