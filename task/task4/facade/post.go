package facade

import (
	"github.com/gin-gonic/gin"
	"github.com/task/web/dao"
	"net/http"
)

// CreatePost
// @Summary  CreatePost
// @Produce  json
// @Success  200  {object}  map[string]string
// @Router   /post/create [post]
func CreatePost(c *gin.Context) {
	userId, _ := c.Get("userID")
	var post dao.Post
	c.BindJSON(&post)
	post.UserID = userId.(uint)
	dao.DB.Create(&post)
	c.JSON(http.StatusOK, gin.H{
		"message": "Created successfully",
	})
}

// GetPost
// @Summary  GetPost
// @Produce  json
// @Success  200  {object}  map[string]string
// @Router   /post/read/:id [get]
func GetPost(c *gin.Context) {
	postId := c.Param("id")
	var post dao.Post
	result := dao.DB.First(&post, postId)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Post not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": post,
	})
}

type UpdatePostRequest struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// UpdatePost
// @Summary  UpdatePost
// @Produce  json
// @Success  200  {object}  map[string]string
// @Router   /post/update [post]
func UpdatePost(c *gin.Context) {
	var post UpdatePostRequest
	userId, _ := c.Get("userID")
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dbPost := dao.Post{
		Title:   post.Title,
		Content: post.Content,
	}
	dao.DB.Debug().Model(&dao.Post{}).Where("id = ? and user_id = ?", post.ID, userId).Updates(&dbPost)
	c.JSON(http.StatusOK, gin.H{
		"message": "Updated successfully",
	})
}

// DeletePost
// @Summary  DeletePost
// @Produce  json
// @Success  200  {object}  map[string]string
// @Router   /post/delete [delete]
func DeletePost(c *gin.Context) {
	var post dao.Post
	userId, _ := c.Get("userID")
	postId := c.Query("postId")
	dao.DB.Where("id = ? and user_id = ?", postId, userId).Delete(&post)
	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted successfully",
	})
}

// ListPost
// @Summary  ListPost
// @Produce  json
// @Success  200  {object}  map[string]string
// @Router   /post/list [get]
func ListPost(c *gin.Context) {
	userId := c.Query("userId")
	var posts []dao.Post
	dao.DB.Debug().Model(&dao.Post{}).Where("user_id = ?", userId).Find(&posts)
	c.JSON(http.StatusOK, gin.H{
		"data": posts,
	})
}
