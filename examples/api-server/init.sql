-- PostgreSQL数据库初始化脚本
-- 创建数据库
CREATE DATABASE todo_db;

-- 切换到todo_db数据库
\c todo_db;

-- 创建todos表
CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 插入示例数据
INSERT INTO todos (title, description) VALUES 
('学习Go语言', '深入学习Go语言的基础语法和高级特性'),
('完成API项目', '使用Go Gin框架完成todo list API项目'),
('阅读技术文档', '阅读相关技术文档，提升开发技能'),
('准备面试', '准备技术面试，复习相关知识点');

-- 查看数据
SELECT * FROM todos;
