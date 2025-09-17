package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// Todo 结构体定义
type Todo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   int       `json:"completed"` // 0=未完成, 1=已完成
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// 数据库连接
var db *sql.DB

func main() {
	// 初始化数据库连接
	initDB()
	defer db.Close()

	// 创建Gin路由
	r := gin.Default()

	// 配置CORS中间件
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// API路由
	api := r.Group("/api")
	{
		api.GET("/todos", getTodos)
		api.POST("/todos", createTodo)
		api.PUT("/todos/:id", updateTodoStatus)
		api.DELETE("/todos/:id", deleteTodo)
	}

	// 启动服务器
	log.Println("服务器启动在端口 :9000")
	r.Run(":9000")
}

// 初始化数据库连接
func initDB() {
	var err error

	// 从环境变量获取数据库配置，如果没有则使用默认值
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "password")
	dbName := getEnv("DB_NAME", "todo_db")

	// PostgreSQL连接字符串
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s",
		dbUser, dbPassword, dbName, dbHost, dbPort)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// 测试数据库连接
	if err = db.Ping(); err != nil {
		log.Fatal("数据库ping失败:", err)
	}

	// 创建todos表
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description TEXT,
		completed INTEGER DEFAULT 0,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("创建表失败:", err)
	}

	// 插入一些示例数据
	insertSampleData()

	log.Println("数据库初始化完成")
}

// 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// 插入示例数据
func insertSampleData() {
	// 检查是否已有数据
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM todos").Scan(&count)
	if err != nil {
		log.Println("检查数据失败:", err)
		return
	}

	if count == 0 {
		sampleTodos := []string{
			"学习Go语言",
			"完成API项目",
			"阅读技术文档",
			"准备面试",
		}

		for _, title := range sampleTodos {
			_, err := db.Exec("INSERT INTO todos (title, description) VALUES ($1, $2)",
				title, "这是一个示例todo项目")
			if err != nil {
				log.Println("插入示例数据失败:", err)
			}
		}
		log.Println("示例数据插入完成")
	}
}

// GET /api/todos - 获取所有todo列表
func getTodos(c *gin.Context) {
	rows, err := db.Query("SELECT id, title, description, completed, created_at, updated_at FROM todos ORDER BY created_at DESC")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  500,
			"error": "查询数据库失败",
		})
		return
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":  500,
				"error": "扫描数据失败",
			})
			return
		}
		todos = append(todos, todo)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": todos,
	})
}

// POST /api/todos - 创建新的todo
func createTodo(c *gin.Context) {
	var request struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  400,
			"error": "请求数据格式错误",
		})
		return
	}

	// 验证标题不能为空
	if request.Title == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":  400,
			"error": "标题不能为空",
		})
		return
	}

	// 插入新todo
	var todo Todo
	err := db.QueryRow("INSERT INTO todos (title, description, completed) VALUES ($1, $2, $3) RETURNING id, title, description, completed, created_at, updated_at",
		request.Title, request.Description, 0).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  500,
			"error": "创建失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"data":    todo,
		"message": "创建成功",
	})
}

// PUT /api/todos/:id - 更新todo完成状态
func updateTodoStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  400,
			"error": "无效的ID",
		})
		return
	}

	var request struct {
		Completed int `json:"completed"` // 0=未完成, 1=已完成
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  400,
			"error": "请求数据格式错误",
		})
		return
	}

	// 验证completed字段值
	if request.Completed != 0 && request.Completed != 1 {
		c.JSON(http.StatusOK, gin.H{
			"code":  400,
			"error": "completed字段值必须为0或1",
		})
		return
	}

	// 更新数据库
	_, err = db.Exec("UPDATE todos SET completed = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2",
		request.Completed, id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  500,
			"error": "更新失败",
		})
		return
	}

	// 查询更新后的数据
	var todo Todo
	err = db.QueryRow("SELECT id, title, description, completed, created_at, updated_at FROM todos WHERE id = $1",
		id).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  500,
			"error": "查询更新后数据失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"data":    todo,
		"message": "更新成功",
	})
}

// DELETE /api/todos/:id - 删除todo
func deleteTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  400,
			"error": "无效的ID",
		})
		return
	}

	// 检查todo是否存在
	var exists bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM todos WHERE id = $1)", id).Scan(&exists)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  500,
			"error": "检查todo存在性失败",
		})
		return
	}

	if !exists {
		c.JSON(http.StatusOK, gin.H{
			"code":  404,
			"error": "Todo不存在",
		})
		return
	}

	// 删除todo
	_, err = db.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  500,
			"error": "删除失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "删除成功",
	})
}
