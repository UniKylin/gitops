<template>
  <div class="todo-app">
    <!-- 顶部导航栏 -->
    <header class="app-header">
      <div class="header-content">
        <h1 class="app-title">
          <i class="icon">📝</i>
          Todo Manager
        </h1>
        <div class="header-actions">
          <div class="stats">
            <span class="stat-item">
              <strong>{{ todos.length }}</strong> 总计
            </span>
            <span class="stat-item">
              <strong>{{ completedCount }}</strong> 已完成
            </span>
            <span class="stat-item">
              <strong>{{ pendingCount }}</strong> 待完成
            </span>
          </div>
        </div>
      </div>
    </header>

    <div class="app-body">
      <!-- 左侧表单区域 -->
      <aside class="sidebar">
        <div class="sidebar-content">
          <h2 class="sidebar-title">添加新任务</h2>
          <TodoForm @todo-created="handleTodoCreated" />
        </div>
      </aside>

      <!-- 右侧主内容区域 -->
      <main class="main-content">
        <div class="content-header">
          <div class="content-title">
            <h2>任务列表</h2>
            <div class="loading" v-if="loading">
              <i class="spinner"></i>
              加载中...
            </div>
          </div>
          <div class="content-actions">
            <button 
              @click="refreshTodos" 
              class="btn-refresh"
              :disabled="loading"
            >
              <i class="icon">🔄</i>
              刷新
            </button>
          </div>
        </div>

        <div class="todo-table" v-if="!loading">
          <div class="table-header">
            <div class="col-title">任务</div>
            <div class="col-status">状态</div>
            <div class="col-time">创建时间</div>
            <div class="col-actions">操作</div>
          </div>
          
          <div class="table-body">
            <div 
              v-for="todo in todos" 
              :key="todo.id" 
              class="table-row"
              :class="{ completed: todo.completed === 1 }"
            >
              <div class="col-title">
                <div class="todo-title">{{ todo.title }}</div>
                <div class="todo-description" v-if="todo.description">
                  {{ todo.description }}
                </div>
              </div>
              
              <div class="col-status">
                <span 
                  class="status-badge"
                  :class="{ active: todo.completed === 1 }"
                >
                  {{ todo.completed === 1 ? '已完成' : '进行中' }}
                </span>
              </div>
              
              <div class="col-time">
                <div class="time-text">{{ formatTime(todo.created_at) }}</div>
                <div class="time-text" v-if="todo.updated_at !== todo.created_at">
                  更新: {{ formatTime(todo.updated_at) }}
                </div>
              </div>
              
              <div class="col-actions">
                <button 
                  @click="toggleStatus(todo)"
                  class="btn-action btn-toggle"
                  :class="{ active: todo.completed === 1 }"
                >
                  {{ todo.completed === 1 ? '标记未完成' : '标记完成' }}
                </button>
                
                <button 
                  @click="deleteTodo(todo.id)"
                  class="btn-action btn-delete"
                >
                  删除
                </button>
              </div>
            </div>
          </div>
          
          <div v-if="todos.length === 0" class="empty-state">
            <div class="empty-icon">📋</div>
            <div class="empty-text">暂无任务</div>
            <div class="empty-hint">在左侧添加您的第一个任务</div>
          </div>
        </div>

        <div v-if="error" class="error-message">
          <i class="icon">⚠️</i>
          {{ error }}
        </div>
      </main>
    </div>
  </div>
</template>

<script>
import { ref, onMounted, computed } from 'vue'
import { todoService } from '../services/api.js'
import TodoForm from './TodoForm.vue'

export default {
  name: 'TodoList',
  components: {
    TodoForm
  },
  setup() {
    const todos = ref([])
    const loading = ref(false)
    const error = ref('')

    // 计算属性
    const completedCount = computed(() => {
      return todos.value.filter(todo => todo.completed === 1).length
    })

    const pendingCount = computed(() => {
      return todos.value.filter(todo => todo.completed === 0).length
    })

    // 获取todos列表
    const fetchTodos = async () => {
      loading.value = true
      error.value = ''
      try {
        const response = await todoService.getTodos()
        if (response.code === 0) {
          todos.value = response.data
        } else {
          error.value = response.error || '获取数据失败'
        }
      } catch (err) {
        error.value = '网络请求失败'
        console.error(err)
      } finally {
        loading.value = false
      }
    }

    // 刷新todos
    const refreshTodos = () => {
      fetchTodos()
    }

    // 处理新创建的Todo
    const handleTodoCreated = (newTodo) => {
      // 将新Todo添加到列表顶部
      todos.value.unshift(newTodo)
    }

    // 切换todo状态
    const toggleStatus = async (todo) => {
      const newStatus = todo.completed === 1 ? 0 : 1
      try {
        const response = await todoService.updateTodoStatus(todo.id, newStatus)
        if (response.code === 0) {
          todo.completed = newStatus
          todo.updated_at = response.data.updated_at
        } else {
          error.value = response.error || '更新失败'
        }
      } catch (err) {
        error.value = '更新状态失败'
        console.error(err)
      }
    }

    // 删除todo
    const deleteTodo = async (id) => {
      if (!confirm('确定要删除这个Todo吗？')) {
        return
      }
      
      try {
        const response = await todoService.deleteTodo(id)
        if (response.code === 0) {
          todos.value = todos.value.filter(todo => todo.id !== id)
        } else {
          error.value = response.error || '删除失败'
        }
      } catch (err) {
        error.value = '删除失败'
        console.error(err)
      }
    }

    // 格式化时间
    const formatTime = (timeString) => {
      return new Date(timeString).toLocaleString('zh-CN')
    }

    // 组件挂载时获取数据
    onMounted(() => {
      fetchTodos()
    })

    return {
      todos,
      loading,
      error,
      completedCount,
      pendingCount,
      handleTodoCreated,
      refreshTodos,
      toggleStatus,
      deleteTodo,
      formatTime
    }
  }
}
</script>

<style scoped>
.todo-app {
  min-height: 100vh;
  background: #f8f9fa;
  display: flex;
  flex-direction: column;
}

/* 顶部导航栏 */
.app-header {
  background: #fff;
  border-bottom: 1px solid #e9ecef;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  position: sticky;
  top: 0;
  z-index: 100;
}

.header-content {
  max-width: 1400px;
  margin: 0 auto;
  padding: 0 24px;
  height: 64px;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.app-title {
  font-size: 18px;
  font-weight: 600;
  color: #2c3e50;
  margin: 0;
  display: flex;
  align-items: center;
  gap: 8px;
}

.stats {
  display: flex;
  gap: 24px;
}

.stat-item {
  font-size: 12px;
  color: #6c757d;
}

.stat-item strong {
  color: #495057;
  font-weight: 600;
}

/* 主体布局 */
.app-body {
  flex: 1;
  max-width: 1400px;
  margin: 0 auto;
  padding: 24px;
  display: grid;
  grid-template-columns: 320px 1fr;
  gap: 24px;
  width: 100%;
}

/* 侧边栏 */
.sidebar {
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  height: fit-content;
  position: sticky;
  top: 88px;
}

.sidebar-content {
  padding: 24px;
}

.sidebar-title {
  font-size: 16px;
  font-weight: 600;
  color: #2c3e50;
  margin: 0 0 16px 0;
}

/* 主内容区域 */
.main-content {
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  overflow: hidden;
}

.content-header {
  padding: 24px;
  border-bottom: 1px solid #e9ecef;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.content-title h2 {
  font-size: 16px;
  font-weight: 600;
  color: #2c3e50;
  margin: 0;
}

.loading {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #6c757d;
  font-size: 12px;
  margin-top: 6px;
}

.spinner {
  width: 16px;
  height: 16px;
  border: 2px solid #e9ecef;
  border-top: 2px solid #007bff;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

.btn-refresh {
  padding: 6px 12px;
  background: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  font-size: 12px;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 4px;
  transition: background 0.2s;
}

.btn-refresh:hover:not(:disabled) {
  background: #0056b3;
}

.btn-refresh:disabled {
  background: #6c757d;
  cursor: not-allowed;
}

/* 表格样式 */
.todo-table {
  overflow-x: auto;
}

.table-header {
  display: grid;
  grid-template-columns: 2fr 120px 180px 200px;
  background: #f8f9fa;
  border-bottom: 1px solid #e9ecef;
  font-weight: 600;
  color: #495057;
  font-size: 12px;
}

.table-header > div {
  padding: 12px 16px;
  border-right: 1px solid #e9ecef;
}

.table-header > div:last-child {
  border-right: none;
}

.table-body {
  min-height: 200px;
}

.table-row {
  display: grid;
  grid-template-columns: 2fr 120px 180px 200px;
  border-bottom: 1px solid #e9ecef;
  transition: background 0.2s;
}

.table-row:hover {
  background: #f8f9fa;
}

.table-row.completed {
  opacity: 0.7;
  background: #f8f9fa;
}

.table-row > div {
  padding: 12px 16px;
  border-right: 1px solid #e9ecef;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.table-row > div:last-child {
  border-right: none;
}

.col-title {
  min-width: 0;
}

.todo-title {
  font-weight: 500;
  color: #2c3e50;
  margin-bottom: 3px;
  word-break: break-word;
  font-size: 14px;
}

.todo-description {
  font-size: 12px;
  color: #6c757d;
  line-height: 1.4;
  word-break: break-word;
}

.status-badge {
  display: inline-block;
  padding: 3px 6px;
  border-radius: 10px;
  font-size: 11px;
  font-weight: 500;
  background: #e9ecef;
  color: #495057;
}

.status-badge.active {
  background: #d4edda;
  color: #155724;
}

.time-text {
  font-size: 12px;
  color: #6c757d;
  line-height: 1.3;
}

.col-actions {
  display: flex;
  gap: 8px;
  flex-direction: column;
}

.btn-action {
  padding: 4px 8px;
  border: none;
  border-radius: 3px;
  font-size: 11px;
  cursor: pointer;
  transition: all 0.2s;
  font-weight: 500;
}

.btn-toggle {
  background: #e9ecef;
  color: #495057;
}

.btn-toggle.active {
  background: #d4edda;
  color: #155724;
}

.btn-toggle:hover {
  background: #dee2e6;
}

.btn-toggle.active:hover {
  background: #c3e6cb;
}

.btn-delete {
  background: #f8d7da;
  color: #721c24;
}

.btn-delete:hover {
  background: #f5c6cb;
}

/* 空状态 */
.empty-state {
  text-align: center;
  padding: 60px 20px;
  color: #6c757d;
}

.empty-icon {
  font-size: 3rem;
  margin-bottom: 16px;
}

.empty-text {
  font-size: 1.1rem;
  font-weight: 500;
  margin-bottom: 8px;
}

.empty-hint {
  font-size: 0.9rem;
  color: #adb5bd;
}

/* 错误消息 */
.error-message {
  background: #f8d7da;
  color: #721c24;
  padding: 16px 24px;
  border-left: 4px solid #dc3545;
  display: flex;
  align-items: center;
  gap: 8px;
  font-weight: 500;
}

/* 响应式设计 */
@media (max-width: 1200px) {
  .app-body {
    grid-template-columns: 280px 1fr;
    gap: 20px;
  }
  
  .table-header,
  .table-row {
    grid-template-columns: 2fr 100px 160px 180px;
  }
}

@media (max-width: 768px) {
  .app-body {
    grid-template-columns: 1fr;
    padding: 16px;
    gap: 16px;
  }
  
  .sidebar {
    position: static;
    order: 2;
  }
  
  .main-content {
    order: 1;
  }
  
  .header-content {
    padding: 0 16px;
  }
  
  .stats {
    gap: 16px;
  }
  
  .table-header,
  .table-row {
    grid-template-columns: 1fr;
    gap: 8px;
  }
  
  .table-header > div,
  .table-row > div {
    border-right: none;
    border-bottom: 1px solid #e9ecef;
    padding: 12px 16px;
  }
  
  .col-actions {
    flex-direction: row;
    justify-content: flex-end;
  }
}
</style>
