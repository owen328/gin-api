package controller

import (
	"github.com/gin-gonic/gin"
	"go_learn/app/form"
	"go_learn/app/model"
	"go_learn/common"
	"go_learn/utils"
	"gorm.io/gorm"
	"net/http"
	"os"
	"path"
)

type IAdmin interface {
	Index(c *gin.Context)
	Register(c *gin.Context)
	Login(c *gin.Context)
	Home(c *gin.Context)
	Upload(c *gin.Context)
}

type Admin struct {
	db *gorm.DB
}

func NewAdmin() *Admin {
	return &Admin{
		db: common.DB,
	}
}

func (a Admin) Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func (a Admin) Register(c *gin.Context) {
	var createForm form.RegisterAdmin
	if err := c.ShouldBindJSON(&createForm); err != nil {
		//fmt.Println(err.Error())
		utils.Fail(c, err.Error())
		return
	}
	if createForm.Password != createForm.ConfirmPassword {
		utils.Fail(c, "确认密码不一致")
		return
	}
	password, err := utils.MakePassword(createForm.Password)
	if err != nil {
		utils.Fail(c, "创建账号失败")
	}

	if result := a.db.Create(&model.Admin{
		Username: createForm.Username,
		Password: password,
		Mobile:   createForm.Mobile,
	}).Error; result != nil {
		utils.Fail(c, "创建账号失败:"+result.Error())
		return
	}
	utils.Success(c, "创建成功")
	return
}

func (a Admin) Login(c *gin.Context) {
	var loginForm form.LoginAdmin
	if err := c.ShouldBindJSON(&loginForm); err != nil {
		utils.Fail(c, err.Error())
		return
	}
	var adminModel model.Admin
	if result := a.db.Where("username = ?", loginForm.Username).Select([]string{"username", "password", "mobile", "id"}).First(&adminModel); result.Error != nil {
		utils.Fail(c, "账号或密码错误")
		return
	}
	ok := utils.CheckPassword(loginForm.Password, adminModel.Password)
	if !ok {
		utils.Fail(c, "账号或密码错误")
		return
	}

	token, err := utils.GenerateToken(adminModel)
	if err != nil {
		utils.Fail(c, "账号或密码错误")
		return
	}
	utils.Success(c, gin.H{
		"id":       adminModel.Id,
		"username": adminModel.Username,
		"token":    token,
	})
	return
}

func (a Admin) Home(c *gin.Context) {
	username := c.MustGet("username")
	utils.Success(c, map[string]any{
		"username": username,
	})
}

func (a Admin) Upload(c *gin.Context) {
	file, err := c.FormFile("file_path")
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}

	fs, err := file.Open()
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}
	defer fs.Close()
	md5, err := common.FileMd5(fs)
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}
	//指针归零
	fs.Seek(0, 0)

	cur, err := os.Getwd()
	if err != nil {
		utils.Fail(c, err.Error())
		return
	}
	uploadDir := path.Join(cur, "uploads")
	os.Mkdir(uploadDir, 0666)
	uploadFilePath := path.Join(uploadDir, md5+path.Ext(file.Filename))

	_, err = common.SaveFile(fs, uploadFilePath)

	if err != nil {
		utils.Fail(c, err.Error())
		return
	}

	utils.Success(c, map[string]string{
		"file_path": uploadFilePath,
	})
	return
}
