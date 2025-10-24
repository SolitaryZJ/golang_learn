package facade

import (
	"github.com/gin-gonic/gin"
	"github.com/task/web/dao"
	"net/http"
)

type CommentRequest struct {
	PostID  uint   `json:"post_id"`
	Content string `json:"content"`
}

// CreateComment
// @Summary  CreateComment
// @Produce  json
// @Success  200  {object}  map[string]string
// @Router   /comment/create [post]
func CreateComment(c *gin.Context) {
	userId, _ := c.Get("userID")
	var commentRequest CommentRequest
	if err := c.ShouldBindJSON(&commentRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate the post  exists
	var post dao.Post
	if dao.DB.Debug().Where("id = ?", commentRequest.PostID).First(&post).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Post not found"})
		return
	}
	comment := dao.Comment{
		Content: commentRequest.Content,
		UserID:  userId.(uint),
		PostID:  commentRequest.PostID,
	}
	dao.DB.Create(&comment)
	c.JSON(http.StatusOK, gin.H{
		"data": comment,
	})
}

// ListComments
// @Summary  ListComments
// @Produce  json
// @Success  200  {object}  map[string]string
// @Router   /comment/list [get]
func ListComments(c *gin.Context) {
	postId := c.Query("postId")
	var comments []dao.Comment
	dao.DB.Where("post_id = ?", postId).Find(&comments)
	c.JSON(http.StatusOK, gin.H{
		"data": comments,
	})
}
