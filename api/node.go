package api

import (
	"net/http"

	"github.com/Dataman-Cloud/rolex/util/dockerclient"
	"github.com/gin-gonic/gin"
)

func (api *Api) InspectNode(ctx *gin.Context) {
	nodes, err := api.GetDockerClient().Nodelist(dockerclient.NodeListOptions{})
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": nodes})
}

func (api *Api) ListNodes(ctx *gin.Context)  {}
func (api *Api) CreateNode(ctx *gin.Context) {}
func (api *Api) UpdateNode(ctx *gin.Context) {}
func (api *Api) RemoveNode(ctx *gin.Context) {}
