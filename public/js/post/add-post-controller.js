'use strict';

angular.module('myapp')
  .controller('AddPostController', ['$scope','Post', '$location',
    function ($scope,  Post, $location) {
        $scope.post = {
          "title": "",
          "body": "",
          "id": null
        }

      $scope.save = function () {
          Post.save($scope.post,
            function () {
             $location.path( "/posts" );
            });
      };

 }]);