'use strict';

angular.module('myapp')
  .controller('PostController', ['$scope', '$modal', 'resolvedPost', 'Post','$sce',
    function ($scope, $modal, resolvedPost, Post,$sce) {

      $scope.posts = resolvedPost;

      $scope.create = function () {
        $scope.clear();
        $scope.open();
      };

      $scope.update = function (id) {
        $scope.post = Post.get({id: id});
        $scope.open(id);
      };

      $scope.delete = function (id) {
        Post.delete({id: id},
          function () {
            $scope.posts = Post.query();
          });
      };

      $scope.save = function (id) {
        if (id) {
          Post.update({id: id}, $scope.post,
            function () {
              $scope.posts = Post.query();
              $scope.clear();
            });
        } else {
          Post.save($scope.post,
            function () {
              $scope.posts = Post.query();
              $scope.clear();
            });
        }
      };

      $scope.clear = function () {
        $scope.post = {
          
          "title": "",
          
          "body": "",
          
          "id": ""
        };
      };
      $scope.to_trusted = function(html_code) {
         var converter = new Showdown.converter();
          return $sce.trustAsHtml(converter.makeHtml(html_code || ''));
      };
      $scope.open = function (id) {
        var postSave = $modal.open({
          templateUrl: 'post-save.html',
          controller: PostSaveController,
          resolve: {
            post: function () {
              return $scope.post;
            }
          }
        });

        postSave.result.then(function (entity) {
          $scope.post = entity;
          $scope.save(id);
        });
      };
    }]);

var PostSaveController =
  function ($scope, $modalInstance, post) {
    $scope.post = post;

    

    $scope.ok = function () {
      $modalInstance.close($scope.post);
    };

    $scope.cancel = function () {
      $modalInstance.dismiss('cancel');
    };
  };
