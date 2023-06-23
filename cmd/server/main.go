package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AbdulwahabNour/movies/api/handlers"
	"github.com/AbdulwahabNour/movies/api/routes"
	"github.com/AbdulwahabNour/movies/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

 
var conf  config.Config 
var app  config.App

func main(){
    

    flag.IntVar(&conf.Port, "port", 4000, "API server port")
    flag.StringVar(&conf.Env, "env", "development", "Environment (development|staging|production)")
    flag.Parse()

    
   
 
    app.InfoLog = log.New(os.Stdout, "INFO:\t", log.Ldate | log.Ltime)
 
    app.ErrorLog = log.New(os.Stderr, "ERROR:\t", log.Ldate | log.Ltime)
    conf.Version = "1.0.0"
    
    app.Config   = &conf
   

    r := gin.Default()
    binding.EnableDecoderDisallowUnknownFields = true
 
    handler := handlers.NewApiHandlers(&app)
    routers := routes.NewApiHandlers(handler)

    routers.SetRoutes(r)


    r.MaxMultipartMemory = 1 << 20 // 1MB
 
    srv := &http.Server{
          Addr: fmt.Sprintf(":%d", app.Config.Port),
          Handler: r,
          IdleTimeout: time.Minute,
          WriteTimeout: 30*time.Second,
          ReadTimeout: 10*time.Second,
    }

    app.InfoLog.Printf("stating %s sever on %s", app.Config.Env, srv.Addr)
    err := srv.ListenAndServe()
    app.ErrorLog.Println(err)

}


 