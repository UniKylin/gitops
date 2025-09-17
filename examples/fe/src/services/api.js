import axios from 'axios'

// 创建axios实例
const api = axios.create({
  baseURL: '/api',
  timeout: 5000,
  headers: {
    'Content-Type': 'application/json'
  }
})

// Todo API服务
export const todoService = {
  // 获取所有todos
  async getTodos() {
    try {
      const response = await api.get('/todos')
      return response.data
    } catch (error) {
      console.error('获取todos失败:', error)
      throw error
    }
  },

  // 创建新的todo
  async createTodo(todoData) {
    try {
      const response = await api.post('/todos', todoData)
      return response.data
    } catch (error) {
      console.error('创建todo失败:', error)
      throw error
    }
  },

  // 更新todo状态
  async updateTodoStatus(id, completed) {
    try {
      const response = await api.put(`/todos/${id}`, { completed })
      return response.data
    } catch (error) {
      console.error('更新todo状态失败:', error)
      throw error
    }
  },

  // 删除todo
  async deleteTodo(id) {
    try {
      const response = await api.delete(`/todos/${id}`)
      return response.data
    } catch (error) {
      console.error('删除todo失败:', error)
      throw error
    }
  }
}

export default api
