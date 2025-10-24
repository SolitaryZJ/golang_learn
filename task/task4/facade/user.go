package facade

import (
	"github.com/gin-gonic/gin"
	"github.com/task/web/dao"
	"github.com/task/web/middleware"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// Register
// @Summary  Register
// @Produce  json
// @Success  200  {object}  map[string]string
// @Router   /user/register [post]
func Register(c *gin.Context) {
	var user dao.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var count int64
	dao.DB.Debug().Model(&dao.User{}).Where("username = ?", user.Username).Count(&count)
	if count != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	if err := dao.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login
// @Summary  login
// @Produce  json
// @Success  200  {object}  map[string]string
// @Router   /user/login [post]
func Login(c *gin.Context) {
	var user dao.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var storedUser dao.User
	if err := dao.DB.Where("username = ?", user.Username).First(&storedUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 生成 JWT
	token, err := middleware.GenerateToken(storedUser.ID, storedUser.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
