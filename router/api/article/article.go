package article

import (
	"net/http"
	"github.com/holehole5566/goproject/pkg/app"
	C "github.com/holehole5566/goproject/pkg/constant"
	"github.com/gin-gonic/gin"
	"github.com/holehole5566/goproject/service"
)

type Article struct {
	title string  `json:"title"`
	subject    string `json:"content"`
	date       string    `json:"date"`
}

func GetArticle(c *gin.Context) {

	appG := app.Gin{C: c}

	switch tour, err := service.article.GetArticle(c.Param("id")); err {

	case C.ErrArticleIDNotNumber:
		appG.Response(http.StatusBadRequest, C.INVALID_PARAMS, err.Error(), nil)

	case C.ErrArticleNotFound:
		appG.Response(http.StatusBadRequest, C.ERROR_GET_ARTICLE_NO_RECORD, err.Error(), nil)

	case C.ErrDatabase:
		appG.Response(http.StatusInternalServerError, C.ERROR_GET_ARTICLE_FAIL, err.Error(), nil)

	case nil:
		appG.Response(http.StatusOK, C.SUCCESS, C.SuccessMsg, tour)

	default:
		appG.Response(http.StatusInternalServerError, C.SERVER_ERROR, err.Error(), nil)
	}
}

func GetAllArticle(c *gin.Context) {
	appG := app.Gin{C: c}

	switch total, err := service.article.GetAllArticle(); err {

	case C.ErrDatabase:
		appG.Response(http.StatusInternalServerError, C.SERVER_ERROR, err.Error(), nil)

	case nil:
		appG.Response(http.StatusOK, C.SUCCESS, C.SuccessMsg, total)

	}
}


func AddArticle(c *gin.Context) {

	appG := app.Gin{C: c}

	var t article

	if err := c.BindJSON(&t); err != nil {
		appG.Response(http.StatusBadRequest, C.INVALID_PARAMS, err.Error(), nil)
		return
	}

	switch tourID, err := service.Article.AddArticle(t.Collects, t.Title); err {

	case C.ErrArticleAddFormatIncorrect:
		appG.Response(http.StatusBadRequest, C.ERROR_ADD_ARTICLE_FORMAT_INCORRECT, err.Error(), nil)

	case C.ErrArticleAddCollectsRecordNotFound:
		appG.Response(http.StatusBadRequest, C.ERROR_ADD_ARTICLE_NO_COLLECTS_RECORD, err.Error(), nil)

	case C.ErrDatabase:
		appG.Response(http.StatusInternalServerError, C.SERVER_ERROR, err.Error(), nil)

	case nil:
		appG.Response(http.StatusOK, C.SUCCESS, C.SuccessMsg, tourID)

	}
}

func DelArticle(c *gin.Context) {
	appG := app.Gin{C: c}

	switch err := service.article.DelArticle(c.Param("id")); err {

	case C.ErrArticleDelIDIncorrect:
		appG.Response(http.StatusBadRequest, C.ERROR_DEL_ARTICLE_ID_INCORRECT, err.Error(), nil)

	case C.ErrArticleDelDeleted:
		appG.Response(http.StatusGone, C.ERROR_DEL_ARTICLE_DELETED, err.Error(), nil)

	case C.ErrDatabase:
		appG.Response(http.StatusInternalServerError, C.SERVER_ERROR, err.Error(), nil)

	case nil:
		appG.Response(http.StatusOK, C.SUCCESS, C.SuccessMsg, nil)

	}
}