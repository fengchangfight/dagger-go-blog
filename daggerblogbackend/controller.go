package main

import (
	"encoding/json"
	"daggerblogbackend/config"
	"daggerblogbackend/entity"
	"daggerblogbackend/model"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func randStringWithCharset(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func getFilenameExtension(filename string) string {
	parts := strings.Split(filename, ".")
	if len(parts) > 1 {
		return parts[len(parts)-1]
	} else {
		return ""
	}
}

func UploadImage(ctx *gin.Context) {
	env := getEnv()
	var baseImgDir string
	if env == "prod" {
		baseImgDir = config.IMG_DIR_PROD
	} else if env == "qa" {
		baseImgDir = config.IMG_DIR_QA
	} else {
		baseImgDir = config.IMG_DIR_DEV
	}
	// do the real stuff uploading image
	file, header, err := ctx.Request.FormFile("files")
	filename := header.Filename

	extension := getFilenameExtension(filename)

	var newfname string
	newfname = randStringWithCharset(16) + "." + extension
	existingFileMap := listDir(baseImgDir)
	var exist bool
	_, ok := existingFileMap[newfname]
	exist = ok
	for exist {
		newfname = randStringWithCharset(16) + "." + extension
		_, ok := existingFileMap[newfname]
		exist = ok
	}

	out, err := os.Create(baseImgDir + "/" + newfname)
	if err != nil {
		fmt.Println("Creating file failed...")
		ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": "上传失败", "data": "fail..."})
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "已成功上传图片", "data": newfname})
}

func listDir(dirname string) map[string]int {
	result := make(map[string]int)
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		fmt.Println(err)
	}
	for _, f := range files {
		result[f.Name()] = 1
	}
	return result
}

func ClearUnreadComment(ctx *gin.Context) {
	pid := ctx.Param("id")

	result := config.RDB_CONN.Table("bgo_post_comment").Where("post_id = ?", pid).Updates(map[string]interface{}{"read": 1})

	if result.Error != nil {
		ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": "设置已读评论错误", "data": false})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功设置已读评论", "data": result})
	}
}

func EditBlog(ctx *gin.Context) {
	pid := ctx.Param("id")

	x, _ := ioutil.ReadAll(ctx.Request.Body)
	var mapResult map[string]interface{}
	if err := json.Unmarshal([]byte(x), &mapResult); err != nil {
		fmt.Println(err)
	}
	// get 5 parameters: longitude, latitude, title, content, icon
	title := mapResult["title"].(string)
	content := mapResult["content"].(string)
	category_idf := mapResult["category_id"].(float64)

	category_id := int64(category_idf)

	var post entity.Post
	i, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": "不合法的ID", "data": err})
	}
	post.ID = i
	config.RDB_CONN.Model(&post).UpdateColumns(entity.Post{Title: title, Content: content, CategoryID: category_id})

	ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "已成功更新博文", "data": i})
}

func TrashBlog(ctx *gin.Context) {
	pid := ctx.Param("id")

	var post entity.Post
	i, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": "不合法的ID", "data": err})
	}
	post.ID = i
	config.RDB_CONN.Model(&post).UpdateColumns(entity.Post{IsDeleted: true})

	ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "已成功将博文移到垃圾箱", "data": i})
}

func CommentStuff(ctx *gin.Context) {
	x, _ := ioutil.ReadAll(ctx.Request.Body)
	var mapResult map[string]interface{}
	if err := json.Unmarshal([]byte(x), &mapResult); err != nil {
		fmt.Println(err)
	}

	sessionId := mapResult["session_id"].(string)
	commenter := mapResult["commenter"].(string)
	commentContent := mapResult["comment_content"].(string)
	blogIdstr := mapResult["blog_id"].(string)

	blogId, _ := strconv.ParseInt(blogIdstr, 10, 64)

	blogIdInCache := getCacheByKey(config.REDIS_CONN, sessionId)

	if blogIdInCache == "" || blogIdInCache != blogIdstr {
		ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": "无效的验证", "data": sessionId})
		return
	} else {
		// 验证通过， 存入评论
		current_time := time.Now()

		comment := entity.Comment{
			Commenter:      commenter,
			CommentContent: commentContent,
			PostID:         blogId,
			CommentTime:    current_time,
			Read:           false,
		}

		createResult := config.RDB_CONN.Create(&comment)
		if createResult.Error != nil {
			// error
			ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": createResult.Error.(*mysql.MySQLError).Message, "data": comment.ID})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功评论", "data": comment.ID})
		}

	}

}

func OpenCommentWindow4Blog(ctx *gin.Context) {
	x, _ := ioutil.ReadAll(ctx.Request.Body)
	var mapResult map[string]interface{}
	if err := json.Unmarshal([]byte(x), &mapResult); err != nil {
		fmt.Println(err)
	}

	sessionId := mapResult["sessionId"].(string)
	blogId := mapResult["blogId"].(string)

	setCacheItemVal(config.REDIS_CONN, sessionId, blogId, 120)
}

func CreateNewBlog(ctx *gin.Context) {
	x, _ := ioutil.ReadAll(ctx.Request.Body)
	var mapResult map[string]interface{}
	if err := json.Unmarshal([]byte(x), &mapResult); err != nil {
		fmt.Println(err)
	}
	// get 5 parameters: longitude, latitude, title, content, icon
	title := mapResult["title"].(string)
	content := mapResult["content"].(string)
	category_idf := mapResult["category_id"].(float64)

	category_id := int64(category_idf)

	// create 2 parameters: current timestamp, and login user
	current_time := time.Now()
	session := sessions.Default(ctx)
	current_uid := session.Get("uid").(int64)

	// use the 7 parameter all together to create a new record or memory point
	post := entity.Post{
		Title:      title,
		Content:    content,
		UserID:     current_uid,
		CategoryID: category_id,
		CreatedAt:  current_time,
		IsDeleted:  false,
	}

	createResult := config.RDB_CONN.Create(&post)

	if createResult.Error != nil {
		// error
		ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": createResult.Error.(*mysql.MySQLError).Message, "data": post.ID})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功创建博文", "data": post.ID})
	}
}

func GetAllCategories(ctx *gin.Context) {
	var result []model.CategoryVO

	config.RDB_CONN.Table("bgo_category").Scan(&result)

	ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功获取分类列表", "data": result})
}

func DeleteComment(ctx *gin.Context) {
	cidstr := ctx.Param("id")

	cid, err := strconv.ParseInt(cidstr, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": "无效的评论id", "data": cidstr})
		return
	} else {
		config.RDB_CONN.Where("id = ?", cid).Delete(entity.Comment{})
		ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功删除评论", "data": cidstr})
	}
}

func GetCommentList(ctx *gin.Context) {
	params := ctx.Request.URL.Query()
	pageStr := params["page"]
	var page int
	var err error
	if pageStr != nil && len(pageStr) > 0 {
		page, err = strconv.Atoi(pageStr[0])
		if err != nil {
			page = 1
		}
	} else {
		page = 1
	}

	var result []model.CommentVO
	qResult := config.RDB_CONN.Table("bgo_post_comment").Offset((page - 1) * config.ITEMS_PER_PAGE).Limit(config.ITEMS_PER_PAGE).Order("comment_time desc").Select("bgo_post_comment.id, commenter, comment_content, comment_time, c.title as post_title, c.id as post_id").Joins("inner join bgo_post c on c.id = bgo_post_comment.post_id").Scan(&result)

	var count int

	config.RDB_CONN.Table("bgo_post_comment").Count(&count)

	pageObj := model.PaginationVO{
		TotalCount:  count,
		PageSize:    config.ITEMS_PER_PAGE,
		CurrentPage: page,
		Data:        result,
	}

	fmt.Println(qResult)
	if qResult.Error != nil {
		ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": "载入未读评论错误", "data": false})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功载入未读评论列表", "data": pageObj})
	}
}

func CountUnreadComment(ctx *gin.Context) {
	var count int
	qResult := config.RDB_CONN.Table("bgo_post_comment").Where(map[string]interface{}{"read": false}).Count(&count)

	if qResult.Error != nil {
		ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": "查看未读评论错误", "data": false})
	} else {
		ctx.String(http.StatusOK, strconv.Itoa(count))
	}

}

func GetComment4Blog(ctx *gin.Context) {
	cid := ctx.Param("id")
	var result []model.CommentVO
	qResult := config.RDB_CONN.Table("bgo_post_comment").Order("comment_time desc").Select("id, commenter, comment_content, comment_time").Where("post_id = ?", cid).Scan(&result)

	fmt.Println(qResult)
	if qResult.Error != nil {
		ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": "载入评论错误", "data": false})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功评论列表", "data": result})
	}
}

func GetPostDetail(ctx *gin.Context) {
	pid := ctx.Param("id")

	var result model.PostVO
	qResult := config.RDB_CONN.Table("bgo_post").Select("bgo_post.id, title, content, c.name as category, category_id, created_at").Joins("inner join bgo_category c on c.id = bgo_post.category_id").Where("bgo_post.id = ?", pid).Scan(&result)

	fmt.Println(qResult)
	if qResult.Error != nil {
		ctx.JSON(http.StatusOK, gin.H{"ok": false, "message": "博客不存在", "data": false})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功获取博客详情", "data": result})
	}
}

func GetAllBlogsByPage(ctx *gin.Context) {
	params := ctx.Request.URL.Query()
	pageStr := params["page"]
	categoryIdStr := params["category_id"]

	var page int
	var category_id int64
	var err error
	if pageStr != nil && len(pageStr) > 0 {
		page, err = strconv.Atoi(pageStr[0])
		if err != nil {
			page = 1
		}
	} else {
		page = 1
	}

	if categoryIdStr != nil && len(categoryIdStr) > 0 {
		category_id, err = strconv.ParseInt(categoryIdStr[0], 10, 64)
	} else {
		category_id = -1
	}

	var result []model.PostListVO
	var count int

	if category_id == -1 {
		config.RDB_CONN.Table("bgo_post").Where("is_deleted = ? ", false).Count(&count)
		config.RDB_CONN.Table("bgo_post").Offset((page-1)*config.ITEMS_PER_PAGE).Limit(config.ITEMS_PER_PAGE).Order("created_at desc").Select("bgo_post.id, title, c.name as category, created_at").Joins("inner join bgo_category c on c.id = bgo_post.category_id").Where("is_deleted = ? ", false).Scan(&result)
	} else {
		config.RDB_CONN.Table("bgo_post").Where("is_deleted = ? and category_id=?", false, category_id).Count(&count)
		config.RDB_CONN.Table("bgo_post").Offset((page-1)*config.ITEMS_PER_PAGE).Limit(config.ITEMS_PER_PAGE).Order("created_at desc").Select("bgo_post.id, title, c.name as category, created_at").Joins("inner join bgo_category c on c.id = bgo_post.category_id").Where("is_deleted = ? and category_id=?", false, category_id).Scan(&result)
	}

	// // fmt.Println(result)
	pageObj := model.PaginationVO{
		TotalCount:  count,
		PageSize:    config.ITEMS_PER_PAGE,
		CurrentPage: page,
		Data:        result,
	}

	ctx.JSON(http.StatusOK, gin.H{"ok": true, "message": "成功获取博文列表", "data": pageObj})
}
