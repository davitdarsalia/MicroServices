package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

var handler = Handler{services: nil}

func TestHandler(t *testing.T) {
	nullishEngine := gin.Default()

	for i := 0; i < 100; i++ {
		res := handler.InitRoutes()
		var (
			notEqual, notNil = assert.NotEqual(t, nullishEngine, res), assert.NotNil(t, res)
		)
		if !notEqual || !notNil {
			t.Error("Incorrect Core Engine")
			t.Fail()
		} else {
			t.Logf("Engine N-%d - Created Successfully", i)
		}
	}
}
