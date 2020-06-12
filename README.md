- config

    读取 json 配置

- container

    IOC 容器

- context

    ...

- database

    简单的 ORM 实现
    
    models/user.go
    
    ```go
	package models
    
	import (
		"github.com/qq15725/veno/database/model"
	)
    
	type User struct {
		model.Model
		table      string `model:"users"`
		primaryKey uint   `model:"id"`
	}
    ```
  
    main.go
  
    ```go
	package main
  
    import (
		"fmt"
		_ "github.com/go-sql-driver/mysql"
		"github.com/qq15725/veno/database/dsn"
		"github.com/qq15725/veno/database/model"
		"models"
    )
  
	func main() {
		model.SetConnections(map[string]dsn.Connector{
			"mysql": &dsn.Mysql{
				Host:     "127.0.0.1",
				Port:     3306,
				Database: "database",
				Username: "root",
				Password: "",
				Charset:  "utf8mb4",
			},
		})
  
		defer model.CloseConnections()
  
		fmt.Println(
			model.Query((*models.User)(nil)).Find(1),
		)
	}
    ```
  
- router
    
    基于前缀树的简单路由
    
- veno 
    
    简单的 web 框架
    
    config/database.json
    
    ```json
    {
      "connections": {
        "mysql": {
          "driver": "mysql",
          "host": "127.0.0.1",
          "port": 3306,
          "database": "database",
          "username": "root",
          "password": "",
          "charset": "utf8mb4"
        }
      }
    }

    ```
  
    models/user.go
    
    ```go
    package models
    
    import (
    	"github.com/qq15725/veno/database/model"
    )
    
    type User struct {
    	model.Model
    	table      string `model:"users"`
    	primaryKey uint   `model:"id"`
    }
    ```
  
    controller/user.go
    
    ```go
  
    package controller
    
    import (
    	"github.com/qq15725/veno/database/model"
    	"github.com/qq15725/veno/veno"
    	"net/http"
    	"models"
    )
    
    type UserController struct {
    }
    
    func (uc *UserController) Index(ctx *veno.HttpContext) {
    	ctx.JSON(
    		http.StatusOK,
    		model.Query((*models.User)(nil)).Get(),
    	)
    }

    ```
    
    main.go
    
    ```go
    package main
    
    import (
        _ "github.com/go-sql-driver/mysql"
        "github.com/qq15725/veno/veno"
        "controller"
        "log"
        "path"
        "runtime"
    )
    
    var (
        ROOT string
    )
    
    func init() {
        if _, file, _, ok := runtime.Caller(0); ok {
            ROOT = path.Dir(file)
            log.Println("ROOT:", ROOT)
        }
    }
    
    func main() {
        app := veno.New(ROOT)
    
        app.Router.GET("/users", (&controller.UserController{}).Index)
    
        app.Run(":80")
    }

    ```
    
    访问 `http://localhost/`