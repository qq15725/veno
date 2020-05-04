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
		"github.com/qq15725/go/database/model"
	)
    
	type User struct {
		model.Model
		table      string `model:"User"`
		primaryKey uint   `model:"UserId"`
	}
    ```
  
    main.go
  
    ```go
	package main
  
    import (
		"fmt"
		_ "github.com/go-sql-driver/mysql"
		"github.com/qq15725/go/app/models"
		"github.com/qq15725/go/database/dsn"
		"github.com/qq15725/go/database/model"
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
			model.Query((*models.User)(nil)).Where("UserId", ">", 2).Get(),
		)
	}
    ```
  
- router
    
    基于前缀树的简单路由
    
- veno 
    
    简单的 web 框架
    
    配置 config/database.json
    
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
    
    main.go
    
    ```go
	package main
  
	import "github.com/qq15725/go/veno"

	func main() {
		app := veno.New()
		app.GET("/", func (ctx *veno.HttpContext) {
			ctx.String(200, "Hello, World!")  
		})
		app.Run(":80")
	}
    ```
    
    访问 `http://localhost/`