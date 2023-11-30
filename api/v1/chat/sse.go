package chat

import (
	"gateway/util"
	"github.com/gin-gonic/gin"
	"time"
)

func SSEHandler(c *gin.Context) {
	sseWriter := util.NewSSEWriter(c.Writer)

	messageChan := make(chan string)
	defer close(messageChan)

	go func() {
		for {
			time.Sleep(5 * time.Second)
			messageChan <- "New message!"
		}
	}()

	for {
		select {
		case <-c.Request.Context().Done():
			return
		case msg := <-messageChan:
			if err := sseWriter.SendEvent(msg); err != nil {
				return
			}
		}
	}
}
