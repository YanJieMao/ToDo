package controller

import (
	"errors"
	"fmt"
	"log"

	"github.com/YanJieMao/ToDo/model"
	"github.com/YanJieMao/ToDo/model/pojo"
	"github.com/YanJieMao/ToDo/model/reqo"
	"github.com/YanJieMao/ToDo/model/reso"
	"github.com/YanJieMao/ToDo/tool"
	"github.com/kataras/iris/v12"
)

// PostLogin user login
func PostLogin(ctx iris.Context) {
	req := reqo.PostLogin{}
	ctx.ReadJSON(&req)
	fmt.Println(req)
	// Query user by username
	user, err := userService.QueryByUsername(req.Username)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.ErrorQueryDatabase(err))
		return
	}

	log.Println(user, req)
	// If passwd are inconsistent
	if user.Passwd != req.Passwd {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.ErrorVerification(errors.New("用户名或密码错误")))
		return
	}

	// Login Ok
	// Get token
	token, err := tool.GetJWTString(user.Username, user.ID)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(model.ErrorBuildJWT(err))
	}

	res := reso.PostLogin{
		Username: user.Username,
		ID:       user.ID,
		Token:    token,
	}
	ctx.JSON(res)

}

// PostUser user register
func PostUser(ctx iris.Context) {
	req := reqo.PostUser{}
	ctx.ReadJSON(&req)
	fmt.Println(req)

	// Username and passwd can't be blank
	if req.Username == "" || req.Passwd == "" {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.ErrorIncompleteData(errors.New("用户名和密码不能为空")))
		return
	}

	// Query user with same username
	exist, _ := userService.QueryByUsername(req.Username)

	// Can't be same username
	if exist.Username != "" {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(model.ErrorVerification(errors.New("用户名已存在")))
		return

	}

	// New user and insert into database
	newUser := pojo.User{
		Username: req.Username,
		Passwd:   req.Passwd,
		Gender:   req.Gender,
		Age:      req.Age,
	}
	userID, err := userService.Insert(newUser)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(model.ErrorInsertDatabase(err))
		return
	}

	res := reso.PostUser{
		Username: newUser.Username,
		ID:       userID,
	}
	ctx.JSON(res)
}

// GetUser return user list
func GetUser(ctx iris.Context) {
	req := reqo.GetUser{}
	ctx.ReadQuery(&req)
	resList := []reso.GetUser{}

	userList, err := userService.Query(req.Username, req.ID)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(model.ErrorQueryDatabase(err))
		return
	}

	for _, user := range userList {
		res := reso.GetUser{
			ID:       user.ID,
			Username: user.Username,
			Passwd:   user.Passwd,
			Age:      user.Age,
			Gender:   user.Gender,
		}

		resList = append(resList, res)
	}
	ctx.JSON(resList)
}

// PutUser update user information
func PutUser(ctx iris.Context) {
	req := reqo.PutUser{}

	ctx.ReadJSON(&req)
	fmt.Println("putuser-req")
	fmt.Println(req)
	//logined := ctx.Values().Get("logined").(model.Logined)

	// // Query user by userID
	/* user, err := userService.QueryByID(req.ID)
	fmt.Println(user)
	if err != nil {
		ctx.JSON(err.Error)
		return
	} */

	user := pojo.User{
		ID:       req.ID,
		Username: req.Username,
		Passwd:   req.Passwd,
		Age:      req.Age,
		Gender:   req.Gender,
	}
	fmt.Println(user)
	// Replace if set
	if req.Gender != 0 {
		user.Gender = req.Gender
	}
	if req.Age != 0 {
		user.Age = req.Age
	}

	// Update user
	err := userService.UpdateByID(user.ID, user)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(model.ErrorQueryDatabase(err))
		return
	}

	// Get updated user
	updatedUser, err := userService.QueryByID(user.ID)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(model.ErrorQueryDatabase(err))
		return
	}

	res := reso.PutUser{

		Username: updatedUser.Username,
		Passwd:   updatedUser.Passwd,
		Gender:   updatedUser.Gender,
		Age:      updatedUser.Age,
	}
	ctx.JSON(res)
}
