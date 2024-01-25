package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testTask/models"
)

func (h *Handler) AddPerson(c *gin.Context) {
	var person models.Person

	if err := c.BindJSON(&person); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.AddPerson(person)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) GetPerson(c *gin.Context) {
	var params models.Person

	if err := c.BindJSON(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	persons, err := h.service.GetPerson(params)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, persons)

}

func (h *Handler) DeletePerson(c *gin.Context) {
	id := c.Query("id")

	err := h.service.DeletePerson(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) UpdatePerson(c *gin.Context) {
	var params models.Person
	id := c.Query("id")

	if err := c.BindJSON(&params); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.UpdatePerson(id, params)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
