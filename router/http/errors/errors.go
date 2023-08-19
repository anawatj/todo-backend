package router

import (
	"net/http"
	domains "todo-backend/domains/errors"

	"github.com/gin-gonic/gin"
)

func Handler(c *gin.Context) {
	// Execute request handlers and then handle any errors
	c.Next()

	errs := c.Errors

	if len(errs) > 0 {
		err, ok := errs[0].Err.(*domains.AppError)
		if ok {
			switch err.Type {
			case domains.NotFound:
				c.JSON(http.StatusNotFound, err.Error())
				return
			case domains.ValidationError:
				c.JSON(http.StatusBadRequest, err.Error())
				return
			case domains.ResourceAlreadyExists:
				c.JSON(http.StatusConflict, err.Error())
				return
			case domains.NotAuthenticated:
				c.JSON(http.StatusUnauthorized, err.Error())
				return
			case domains.NotAuthorized:
				c.JSON(http.StatusForbidden, err.Error())
				return
			case domains.RepositoryError:
			default:
				c.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		// Error is not AppError, return a generic internal server error
		c.JSON(http.StatusInternalServerError, "Internal Server Errror")
		return
	}
}
