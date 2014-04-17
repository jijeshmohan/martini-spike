'use strict';

angular.module('myapp')
  .controller('CommentController', ['$scope', '$modal', 'resolvedComment', 'Comment',
    function ($scope, $modal, resolvedComment, Comment) {

      $scope.comments = resolvedComment;

      $scope.create = function () {
        $scope.clear();
        $scope.open();
      };

      $scope.update = function (id) {
        $scope.comment = Comment.get({id: id});
        $scope.open(id);
      };

      $scope.delete = function (id) {
        Comment.delete({id: id},
          function () {
            $scope.comments = Comment.query();
          });
      };

      $scope.save = function (id) {
        if (id) {
          Comment.update({id: id}, $scope.comment,
            function () {
              $scope.comments = Comment.query();
              $scope.clear();
            });
        } else {
          Comment.save($scope.comment,
            function () {
              $scope.comments = Comment.query();
              $scope.clear();
            });
        }
      };

      $scope.clear = function () {
        $scope.comment = {
          
          "text": "",
          
          "id": ""
        };
      };

      $scope.open = function (id) {
        var commentSave = $modal.open({
          templateUrl: 'comment-save.html',
          controller: CommentSaveController,
          resolve: {
            comment: function () {
              return $scope.comment;
            }
          }
        });

        commentSave.result.then(function (entity) {
          $scope.comment = entity;
          $scope.save(id);
        });
      };
    }]);

var CommentSaveController =
  function ($scope, $modalInstance, comment) {
    $scope.comment = comment;

    

    $scope.ok = function () {
      $modalInstance.close($scope.comment);
    };

    $scope.cancel = function () {
      $modalInstance.dismiss('cancel');
    };
  };
