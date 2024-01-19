package users

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"test-backend/internal/appresult"
	"test-backend/internal/handlers"
	"test-backend/pkg/logging"
)

const (
	getAllURL = "/get-all"
	getOneURL = "/get"
	addURL    = "/add"
	deleteURL = "/delete"
)

type handler struct {
	repository Repository
	logger     *logging.Logger
}

func NewHandler(repository Repository, logger *logging.Logger) handlers.Handler {
	return &handler{
		repository: repository,
		logger:     logger,
	}
}

func (h *handler) Register(router *gin.RouterGroup) {
	router.POST(getAllURL, h.GetAll)
	router.POST(addURL, h.Create)
	router.POST(deleteURL, h.Delete)

}

// GetAll godoc
// @Description all data users
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       users  body PaginationDTO true  "Get All JSON"
// @Success     200  {array}  DataDTO
// @Router      /users/get-all [post]
func (h *handler) GetAll(c *gin.Context) {
	var reqDTO PaginationDTO
	err := c.ShouldBindJSON(&reqDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, appresult.ErrMissingParam)
		return
	}

	result, err := h.repository.GetAllData(context.Background(), reqDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, appresult.ErrInternalServer)
		return
	}

	successResult := appresult.Success
	successResult.Data = result
	c.JSON(http.StatusOK, successResult)

	return
}

// Create godoc
// @Description create and update data  user
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       users  body ReqUser  true  "Create and Update JSON"
// @Success     200  {string}  string
// @Router      /users/add [post]
func (h *handler) Create(c *gin.Context) {
	var (
		err    error
		reqDTO ReqUser
	)

	err = c.ShouldBindJSON(&reqDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, appresult.ErrMissingParam)
		return
	}

	fmt.Println("uuid :::", reqDTO.Id)
	if reqDTO.Id != 0 {
		err = h.repository.UpdateData(context.Background(), reqDTO)
	} else {
		id, _ := h.repository.AddData(context.Background(), reqDTO)
		go h.CustomAgeEnRich(reqDTO.Name, id)
		go h.CustomGenderEnRich(reqDTO.Name, id)
		go h.CustomNationalizeEnRich(reqDTO.Name, id)
	}

	if err != nil {
		return
	}

	successResult := appresult.Success
	successResult.Data = ""
	c.JSON(http.StatusOK, successResult)
	return
}

// Delete godoc
// @Description delete data  user
// @Tags        users
// @Accept      json
// @Produce     json
// @Param       users  body ReqIdDTO  true  "Delete JSON"
// @Success     200  {string}  string
// @Router      /users/delete [post]
func (h *handler) Delete(c *gin.Context) {
	var reqDTO ReqIdDTO
	err := c.ShouldBindJSON(&reqDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, appresult.ErrMissingParam)
		return
	}

	err = h.repository.DeleteData(context.Background(), reqDTO.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, appresult.ErrMissingParam)
		return
	}

	successResult := appresult.Success
	successResult.Data = ""
	c.JSON(http.StatusOK, successResult)
	return
}

func (h *handler) CustomAgeEnRich(name string, id int) {
	var u ResUserAge

	payload := url.Values{}
	payload.Add("name", name)

	req, _ := http.NewRequest("GET", "https://api.agify.io/?"+payload.Encode(), nil)
	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		h.logger.Error("age request give err", err)
	}

	body, err := io.ReadAll(res.Body)
	err = res.Body.Close()

	err = json.Unmarshal(body, &u)
	if err != nil {
		h.logger.Error("body error", err)
	}
	u.Id = id
	err = h.repository.UpdateAgeData(context.Background(), u)
}

func (h *handler) CustomGenderEnRich(name string, id int) {
	var u ResUserGender

	payload := url.Values{}
	payload.Add("name", name)

	req, _ := http.NewRequest("GET", "https://api.genderize.io/?"+payload.Encode(), nil)
	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		h.logger.Error("gender request give err", err)
	}

	body, err := io.ReadAll(res.Body)
	err = res.Body.Close()

	err = json.Unmarshal(body, &u)
	if err != nil {
		h.logger.Error("body error", err)
	}

	u.Id = id
	err = h.repository.UpdateGenderData(context.Background(), u)

}

func (h *handler) CustomNationalizeEnRich(name string, id int) {
	var u ResUserCountry

	payload := url.Values{}
	payload.Add("name", name)

	req, _ := http.NewRequest("GET", "https://api.nationalize.io/?"+payload.Encode(), nil)
	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		h.logger.Error("age request give err", err)
	}

	body, err := io.ReadAll(res.Body)
	err = res.Body.Close()

	err = json.Unmarshal(body, &u)
	if err != nil {
		h.logger.Error("body error", err)
	}

	u.Id = id
	err = h.repository.UpdateNationalizeData(context.Background(), u)

}
