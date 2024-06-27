package handler

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lrayt/small-sparrow/ts_error"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

type HttpHandler struct {
	api *gin.RouterGroup
	srv *http.Server
}

func NewHttpHandler() *HttpHandler {
	r := gin.Default()

	return &HttpHandler{
		api: r.Group("/api/v1"),
		srv: &http.Server{
			Addr:    ":8080",
			Handler: r,
		},
	}
}

type bufferPool struct {
	pool *sync.Pool
}

func (b *bufferPool) Get() []byte {
	var data = b.pool.Get().([]byte)
	log.Printf("get buffer:%s\n", data)
	return data
}

func (b *bufferPool) Put(buf []byte) {
	var uuid = []byte(`,"uuid":"123"}`)
	if len(buf) > 0 {
		buf = buf[:len(buf)-2]
		buf = append(buf, uuid...)
	}
	log.Printf("put buffer:%s\n", buf)
	b.pool.Put(buf[:cap(buf)])
}

func newBufferPool(size int) *bufferPool {
	return &bufferPool{
		pool: &sync.Pool{
			New: func() interface{} {
				return make([]byte, size)
			},
		},
	}
}

func (h HttpHandler) proxyRouter() {
	h.api.Any("/proxy/baidu", func(c *gin.Context) {
		targetURL, err := url.Parse("http://127.0.0.1:8081/data?abc=123")
		if err != nil {
			c.JSON(http.StatusBadGateway, &ts_error.BaseResponse{
				Code: 1000,
				Msg:  err.Error(),
			})
		}

		proxy := httputil.NewSingleHostReverseProxy(targetURL)
		proxy.BufferPool = newBufferPool(32 * 1024)
		proxy.Director = func(req *http.Request) {
			req.URL.Scheme = targetURL.Scheme
			req.URL.Host = targetURL.Host
			req.URL.Path = targetURL.Path
			req.Host = targetURL.Host
			// 复制原始请求的头部
			for key, value := range c.Request.Header {
				req.Header[key] = value
			}
			log.Printf("Director:1--%s---%s---%s", targetURL.Path, targetURL.RawQuery, c.Request.URL.RawQuery)
		}
		proxy.ErrorHandler = func(writer http.ResponseWriter, request *http.Request, err error) {
			if errors.Is(err, context.Canceled) {
				log.Println("context.Canceled---->")
			}
			log.Printf("ErrorHandler:%v", err.Error())
		}
		proxy.ModifyResponse = func(res *http.Response) error {
			res.Header.Get("")
			log.Printf("ModifyResponse:1")
			//return errors.New("new err-->")
			return nil
		}
		log.Println("111")
		proxy.ServeHTTP(c.Writer, c.Request)
		log.Println("222")
	})
}

func (h *HttpHandler) Run() error {
	log.Println("------")
	//h.srv = &http.Server{
	//	Addr:    ":8080",
	//	Handler: h.router,
	//}
	h.proxyRouter()
	if err := h.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (h HttpHandler) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return h.srv.Shutdown(ctx)
}
