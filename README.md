# micro-project
* 预计实现一个基于微服务架构的招聘系统
## 1、时间安排
* 11.3 ~ 11.5 : 实现登录模块 web服务器，前端设计
* 11.6 ~ 11.10 ：实现etcd服务注册与发现
* 11.11~ 11.12 ： 实现招聘模块
## 2、web设计
### 技术选型
* gin,templates,mysql,etcd
### login
* 1.web服务器
* 2.html,css,javascript
* 3.mysql
* 4.部署位置：windows
* main.go
  ```
  r := gin.Default()
	r.LoadHTMLGlob("D:\\Develop\\gopath\\micro\\login\\html\\login.tmpl") //模板解析
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.tmpl", gin.H{"title": "登录页面"}) //模板渲染
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
  ```
* login.tmpl
  ```
  <!DOCTYPE html>
<html>
<head>
    <style>
        #character {
            width: 25px;
            height: 50px;
            background-color: black;
            position: absolute;
        }
    </style>
</head>
<body>
<div id="character"></div>
<script>
    var character = document.getElementById("character");

    document.addEventListener("click", function(event) {
        var targetX = event.pageX - (character.offsetWidth / 2);
        var targetY = event.pageY - (character.offsetHeight / 2);

        character.style.left = targetX + "px";
        character.style.top = targetY + "px";
    });
</script>
</body>
</html>
```
## 3、RPC设计
