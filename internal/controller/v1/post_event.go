package v1

import (
	"net/http"
	"time"

	"github.com/Maxson-dev/place-api/internal/pkg/cerror"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type PostEventRequest struct {
	EventType string    `json:"event_type" binding:"required"`
	Datetime  time.Time `json:"datetime" binding:"required"`
	Payload   string    `json:"payload" binding:"required"`
}

type PostEventResponse struct {
	ID string `json:"id"`
}

// PostEvent
//
// @Summary     Create scheduled event
// @Description Method for adding a scheduled event
// @ID          PostEvent
// @Tags  	    event_api
// @Accept      json
// @Produce     json
// @Param       request body PostEventRequest true "Scheduled event"
// @Success     200 {object} PostEventResponse
// @Failure     400 {object} Error
// @Failure     500 {object} Error
// @Router      /api/v1/event [post]
func (c *controller) PostEvent(ctx *gin.Context) {
	var req PostEventRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.Error(ctx, err, "bad request", http.StatusBadRequest)
		return
	}

	evt, err := c.eventUC.Add(ctx, req.EventType, req.Datetime, []byte(req.Payload))
	if err != nil {
		switch {
		case errors.Is(err, cerror.ErrBadInput):
			c.Error(ctx, err, "invalid event", http.StatusBadRequest)
		default:
			c.Error(ctx, err, "internal error", http.StatusInternalServerError)
		}
		return
	}

	ctx.JSON(http.StatusOK, PostEventResponse{ID: evt.ID.String()})
}
