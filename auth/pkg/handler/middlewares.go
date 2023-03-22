package handler

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter"
	ginLimiter "github.com/ulule/limiter/drivers/middleware/gin"
	"github.com/ulule/limiter/drivers/store/memory"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

import (
	"fmt"
)

const (
	brightGreen  = "\033[38;5;10m"
	brightYellow = "\033[38;5;11m"
	brightBlue   = "\033[38;5;12m"
	brightOrange = "[38;5;208m"
	brightPink   = "[38;5;207m"
	reset        = "[0m"

	reqIdLength = 5

	separator = "-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------"
)

var counter uintptr = 1
var logPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 0, 256)
	},
}

type logMessage struct {
	message []byte
}

func requestQuantityLimiter() gin.HandlerFunc {
	rate, err := limiter.NewRateFromFormatted("1000000-H")
	if err != nil {
		fmt.Printf("Request Quantity Limiter: %s", err.Error())
	}

	store := memory.NewStore()

	limiter := limiter.New(store, rate)

	return ginLimiter.NewMiddleware(limiter)

}

func sizeLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		maxBodySize := int64(0)
		if c.Request.Body != nil {
			c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxBodySize)
			if _, err := io.Copy(ioutil.Discard, c.Request.Body); err != nil {
				c.AbortWithStatusJSON(http.StatusRequestEntityTooLarge, gin.H{
					"error": "request body size limit exceeded",
				})
				return
			}
		}

		c.Next()

		if c.Writer.Status() == http.StatusRequestEntityTooLarge {
			c.Abort()
		}
	}
}

func customLogger() gin.HandlerFunc {
	logChan := make(chan *logMessage, 100)
	var wg sync.WaitGroup
	wg.Add(1)

	go loggerWorker(logChan, &wg)
	shutdown := make(chan struct{})
	var once sync.Once

	go func() {
		<-shutdown
		close(logChan)
		wg.Wait()
	}()

	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)

		select {
		case <-c.Done():
			once.Do(func() {
				close(shutdown)
			})
		default:
			start := time.Now()

			c.Next()

			buf := logPool.Get().([]byte)
			buf = append(buf, separator...)
			buf = append(buf, fmt.Sprintf("\n%s \n[Request %sN%d%s - %sID %s%s] [%sUnauthorized User%s] -> Path: [%s%s%s] -> Date [%s%s%s] -> Latency: [%s%d%s Nano Seconds] -> StatusCode: [%s%d%s] -> [%s%f KB%s]\n",
				brightYellow, brightOrange, counter, reset,
				brightYellow,
				generateShortUUID(), brightYellow, brightPink, reset, brightGreen, c.Request.URL.Path, reset,
				brightBlue, time.Now().Format("2 January 2006 : 15:04:05:00"), reset,
				brightPink, time.Since(start).Milliseconds(), reset, brightOrange, c.Writer.Status(), reset,
				brightPink, float64(c.Writer.Size())/1024, reset,
			)...)

			logChan <- &logMessage{message: buf}
			counter++
		}
	}
}

func loggerWorker(logChan <-chan *logMessage, wg *sync.WaitGroup) {
	var buf bytes.Buffer
	defer wg.Done()

	for msg := range logChan {
		buf.Reset()
		buf.Write(msg.message)
		fmt.Println(buf.String())
		logPool.Put(msg.message[:0])
	}
}

func generateShortUUID() string {
	randBytes := make([]byte, reqIdLength)
	_, err := rand.Read(randBytes)
	if err != nil {
		return ""
	}
	return hex.EncodeToString(randBytes)
}
