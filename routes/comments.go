package routes

import (
    "../models"
    "fmt"
    "net/http"
    "strconv"
    "github.com/go-martini/martini"
    "github.com/coopernurse/gorp"
)

func GetComments(enc Encoder, db gorp.SqlExecutor) (int, string) {
    var comments []models.Comment
    _, err := db.Select(&comments, "select * from comments order by id")
    if err != nil {
        checkErr(err, "select failed")
        return http.StatusInternalServerError, ""
    }
    return http.StatusOK, Must(enc.Encode(commentsToIface(comments)...))
}

func GetComment(enc Encoder, db gorp.SqlExecutor, parms martini.Params) (int, string) {
    id, err := strconv.Atoi(parms["id"])
    obj, _ := db.Get(models.Comment{}, id)
    if err != nil || obj == nil {
        checkErr(err, "get failed")
        // Invalid id, or does not exist
        return http.StatusNotFound, ""
    }
    entity := obj.(*models.Comment)
    return http.StatusOK, Must(enc.EncodeOne(entity))
}

func AddComment(entity models.Comment, w http.ResponseWriter, enc Encoder, db gorp.SqlExecutor) (int, string) {
    err := db.Insert(&entity)
    if err != nil {
        checkErr(err, "insert failed")
        return http.StatusConflict, ""
    }
    w.Header().Set("Location", fmt.Sprintf("/myapp/comments/%d", entity.Id))
    return http.StatusCreated, Must(enc.EncodeOne(entity))
}

func UpdateComment(entity models.Comment, enc Encoder, db gorp.SqlExecutor, parms martini.Params) (int, string) {
    id, err := strconv.Atoi(parms["id"])
    obj, _ := db.Get(models.Comment{}, id)
    if err != nil || obj == nil {
        checkErr(err, "get failed")
        // Invalid id, or does not exist
        return http.StatusNotFound, ""
    }
    oldEntity := obj.(*models.Comment)

    entity.Id = oldEntity.Id
    _, err = db.Update(&entity)
    if err != nil {
        checkErr(err, "update failed")
        return http.StatusConflict, ""
    }
    return http.StatusOK, Must(enc.EncodeOne(entity))
}

func DeleteComment(db gorp.SqlExecutor, parms martini.Params) (int, string) {
    id, err := strconv.Atoi(parms["id"])
    obj, _ := db.Get(models.Comment{}, id)
    if err != nil || obj == nil {
        checkErr(err, "get failed")
        // Invalid id, or does not exist
        return http.StatusNotFound, ""
    }
    entity := obj.(*models.Comment)
    _, err = db.Delete(entity)
    if err != nil {
        checkErr(err, "delete failed")
        return http.StatusConflict, ""
    }
    return http.StatusNoContent, ""
}

func commentsToIface(v []models.Comment) []interface{} {
    if len(v) == 0 {
        return nil
    }
    ifs := make([]interface{}, len(v))
    for i, v := range v {
        ifs[i] = v
    }
    return ifs
}
