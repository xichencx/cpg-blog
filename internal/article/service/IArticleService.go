package service

import (
	"cpg-blog/internal/article/model"
	"github.com/gin-gonic/gin"
)

type IArticle interface {
	//Info 查询文章详情
	Info(ctx *gin.Context)

	// List 搜索文章
	List(ctx *gin.Context)

	// Add 新增文章
	Add(ctx *gin.Context)

	// Delete 删除文章
	Delete(ctx *gin.Context)

	// Update 更新文章
	Update(ctx *gin.Context)

	// UpdateArticleEx 服务间更新文章扩展信息
	UpdateArticleEx(ctx *gin.Context, sn int64, view bool, cmt bool, zan bool, add bool) error

	// FindArticles 服务间查询文章信息，支持list
	FindArticles(ctx *gin.Context, sn []int64) (articlesMap map[int64]model.Article)
}
