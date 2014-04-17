'use strict';

angular.module('myapp')
  .config(['$routeProvider', function ($routeProvider) {
    $routeProvider
      .when('/comments', {
        templateUrl: 'views/comment/comments.html',
        controller: 'CommentController',
        resolve:{
          resolvedComment: ['Comment', function (Comment) {
            return Comment.query();
          }]
        }
      })
    }]);
