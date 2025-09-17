<template>
  <div id="app">
    <TodoList />
  </div>
</template>

<script>
import TodoList from './components/TodoList.vue'

export default {
  name: 'App',
  components: {
    TodoList
  }
}
</script>

<style>
#app {
  min-height: 100vh;
  background: #f8f9fa;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
}

/* 全局样式优化 */
* {
  box-sizing: border-box;
}

body {
  margin: 0;
  padding: 0;
  background: #f8f9fa;
  color: #2c3e50;
  line-height: 1.4;
  font-size: 14px;
}

/* 滚动条样式 */
::-webkit-scrollbar {
  width: 8px;
  height: 8px;
}

::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 4px;
}

::-webkit-scrollbar-thumb {
  background: #c1c1c1;
  border-radius: 4px;
}

::-webkit-scrollbar-thumb:hover {
  background: #a8a8a8;
}

/* 按钮通用样式 */
button {
  font-family: inherit;
  cursor: pointer;
}

button:disabled {
  cursor: not-allowed;
}

/* 输入框通用样式 */
input, textarea {
  font-family: inherit;
}

input:focus, textarea:focus {
  outline: none;
}

/* 链接样式 */
a {
  color: #007bff;
  text-decoration: none;
}

a:hover {
  text-decoration: underline;
}

/* 响应式设计 */
@media (max-width: 1200px) {
  #app {
    padding: 0;
  }
}

@media (max-width: 768px) {
  #app {
    background: #f8f9fa;
    padding: 0;
  }
  
  body {
    font-size: 13px;
  }
}
</style>
