package router

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/vimsucks/blogsucks/router/response"
	"fmt"
	"github.com/vimsucks/blogsucks/model"
)

func GetPost(c echo.Context) error {
	guid := c.QueryParam("guid")
	if len(guid) == 0 {
		resp := &response.Error{ErrorCode: response.ErrorNotEnoughArguments, Message: "Please Enter An Guid"}
		return c.JSON(http.StatusBadRequest, resp)
	}
	post := model.GetPost(guid)

	// 如果文章不存在
	if len(post.Guid) == 0 {
		resp := &response.Error{ErrorCode: response.ErrorContentNotExists, Message: "Post not exists"}
		return c.JSON(http.StatusBadRequest, resp)
	}

	return c.String(http.StatusOK,
		fmt.Sprintf("Guid: %s, Title: %s, Content: %s", post.Guid, post.Title, post.Content))
}

func PostPost(c echo.Context) error {
	title := c.FormValue("title")
	content := c.FormValue("content")

	// 合法性检查
	if len(title) == 0 || len(content) == 0 {
		resp := &response.Error{ErrorCode: response.ErrorNotEnoughArguments, Message: "Not Enough Arguments"}
		return c.JSON(http.StatusBadRequest, resp)
	}

	guid := guidGenerator.String()
	c.Logger().Debug(fmt.Sprintf("New article uid \"%s\", title \"%s\", content \"%s\"", guid, title, content))
	model.NewPost(guid, title, content)
	return c.JSON(http.StatusOK, &response.Ok{Message: "发表成功"})
}

func PutPost(c echo.Context) error {
	guid := c.QueryParam("guid")
	title := c.FormValue("title")
	content := c.FormValue("content")

	if len(guid) == 0 {
		resp := &response.Error{ErrorCode: response.ErrorNotEnoughArguments, Message: "Guid can not be empty"}
		return c.JSON(http.StatusBadRequest, resp)
	}

	// 标题和内容至少有一项不为空
	if len(title) == 0 && len(content) == 0 {
		resp := &response.Error{ErrorCode: response.ErrorNotEnoughArguments, Message: "Either title or content can not be empty"}
		return c.JSON(http.StatusBadRequest, resp)
	}

	post := model.GetPost(guid)
	// 如果文章不存在
	if len(post.Guid) == 0 {
		resp := &response.Error{ErrorCode: response.ErrorContentNotExists, Message: "Post not exists"}
		return c.JSON(http.StatusBadRequest, resp)
	}

	model.UpdatePost(post, title, content)
	c.Logger().Debug(fmt.Sprintf("Post %s updated, title: %s, content: %s", guid, title, content))

	return c.String(http.StatusOK,
		fmt.Sprintf("Post guid: %s updated", post.Guid))
}

func DeletePost(c echo.Context) error {
	guid := c.QueryParam("guid")
	if len(guid) == 0 {
		resp := &response.Error{ErrorCode: response.ErrorNotEnoughArguments, Message: "Please Enter An Guid"}
		return c.JSON(http.StatusBadRequest, resp)
	}
	post := model.GetPost(guid)

	// 如果文章不存在
	if len(post.Guid) == 0 {
		resp := &response.Error{ErrorCode: response.ErrorContentNotExists, Message: "Post not exists"}
		return c.JSON(http.StatusBadRequest, resp)
	}

	model.DeletePost(post)
	c.Logger().Debug(fmt.Sprintf("Post %s deleted", guid))

	return c.String(http.StatusOK,
		fmt.Sprintf("Post guid: %s deleted", post.Guid))
}
