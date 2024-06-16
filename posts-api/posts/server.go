package posts

import (
	"net/http"
	"posts/posts/models"
	"posts/uidgen"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type PostsServer struct {
	UidGenNode *uidgen.UniqueIdGenerator
}

// POST /posts
func (p PostsServer) CreatePost(context *gin.Context) {
	var input models.CreatePostRequest
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	new_post_id := p.UidGenNode.GeanerateId()
	result, err := WriteNewPostRecord(
		strconv.FormatUint(uint64(new_post_id), 10),
		&input,
	)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"post": *result})
}

// GET /posts (individual posts only)
func (p PostsServer) GetPost(context *gin.Context) {
	var input models.IndividualPostRequest
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := GetPostRecord(input.ID)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusInternalServerError, gin.H{})
		return
	}
	context.JSON(http.StatusOK, gin.H{"post": *result})
}
