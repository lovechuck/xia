package blog

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/model"
	"server/mongo"
	"server/redis"
	"time"
)

func Search(ctx *gin.Context) {

	var feed model.Feed

	key := time.Now().Format("cnblogs:posts:20060102")
	err := redis.Get(key,&feed)

	if err != nil {

		const url = "http://wcf.open.cnblogs.com/blog/48HoursTopViewPosts/24"
		client := &http.Client{}
		request, err := http.NewRequest(http.MethodGet, url, nil)
		response, err := client.Do(request)
		defer response.Body.Close()
		if response.StatusCode == 200 {
			decoder := xml.NewDecoder(response.Body)
			err = decoder.Decode(&feed)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, model.Response{Code: 10001, Message: err.Error()})
				return
			}

			redis.Set(key, feed, 2*3600)
			ctx.JSON(http.StatusOK, model.Response{Code: 0, Data: feed})
			return
		}

		ctx.JSON(http.StatusInternalServerError, model.Response{Code: 10002, Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{Code: 0, Data: feed.Entry})
	return
}

func Favor(ctx *gin.Context) {
	var entry model.Entry

	err := ctx.ShouldBindJSON(&entry)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{Code: 10000, Message: err.Error()})
		return
	}

	err = mongodb.Insert("blogs", "book_mark",entry)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{Code: 10003, Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{Code: 0})
	return
}

func Bookmarks(ctx *gin.Context) {
	var results []model.Entry

	err := mongodb.GetAll("blogs", "book_mark", &results)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{Code: 10004, Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, model.Response{Code: 0, Data: results})
	return
}
