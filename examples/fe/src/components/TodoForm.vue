<template>
  <div class="todo-form">
    <h2>添加新的Todo</h2>
    <form @submit.prevent="handleSubmit">
      <div class="form-group">
        <label for="title">标题 *</label>
        <input
          type="text"
          id="title"
          v-model="form.title"
          placeholder="请输入Todo标题"
          required
          :disabled="loading"
        />
        <span v-if="errors.title" class="error">{{ errors.title }}</span>
      </div>
      
      <div class="form-group">
        <label for="description">描述</label>
        <textarea
          id="description"
          v-model="form.description"
          placeholder="请输入Todo描述（可选）"
          rows="3"
          :disabled="loading"
        ></textarea>
      </div>
      
      <div class="form-actions">
        <button 
          type="submit" 
          class="btn-submit"
          :disabled="loading || !form.title.trim()"
        >
          {{ loading ? '创建中...' : '创建Todo' }}
        </button>
        <button 
          type="button" 
          class="btn-cancel"
          @click="resetForm"
          :disabled="loading"
        >
          重置
        </button>
      </div>
    </form>
  </div>
</template>

<script>
import { ref, reactive } from 'vue'
import { todoService } from '../services/api.js'

export default {
  name: 'TodoForm',
  emits: ['todo-created'],
  setup(props, { emit }) {
    const loading = ref(false)
    const form = reactive({
      title: '',
      description: ''
    })
    const errors = reactive({
      title: ''
    })

    // 表单验证
    const validateForm = () => {
      errors.title = ''
      
      if (!form.title.trim()) {
        errors.title = '标题不能为空'
        return false
      }
      
      if (form.title.trim().length > 255) {
        errors.title = '标题不能超过255个字符'
        return false
      }
      
      return true
    }

    // 提交表单
    const handleSubmit = async () => {
      if (!validateForm()) {
        return
      }

      loading.value = true
      
      try {
        const response = await todoService.createTodo({
          title: form.title.trim(),
          description: form.description.trim()
        })
        
        if (response.code === 0) {
          // 触发父组件事件
          emit('todo-created', response.data)
          resetForm()
        } else {
          errors.title = response.error || '创建失败'
        }
      } catch (error) {
        errors.title = '网络请求失败'
        console.error(error)
      } finally {
        loading.value = false
      }
    }

    // 重置表单
    const resetForm = () => {
      form.title = ''
      form.description = ''
      errors.title = ''
    }

    return {
      loading,
      form,
      errors,
      handleSubmit,
      resetForm
    }
  }
}
</script>

<style scoped>
.todo-form {
  /* 移除背景和边框，因为现在在侧边栏中 */
  background: transparent;
  border: none;
  padding: 0;
  margin: 0;
  box-shadow: none;
}

.todo-form h2 {
  display: none; /* 隐藏标题，因为侧边栏已有标题 */
}

.form-group {
  margin-bottom: 20px;
}

.form-group label {
  display: block;
  margin-bottom: 4px;
  color: #495057;
  font-weight: 500;
  font-size: 12px;
}

.form-group input,
.form-group textarea {
  width: 100%;
  padding: 8px 10px;
  border: 1px solid #ced4da;
  border-radius: 4px;
  font-size: 12px;
  transition: border-color 0.2s ease;
  font-family: inherit;
  background: #fff;
}

.form-group input:focus,
.form-group textarea:focus {
  outline: none;
  border-color: #007bff;
  box-shadow: 0 0 0 2px rgba(0, 123, 255, 0.25);
}

.form-group input:disabled,
.form-group textarea:disabled {
  background-color: #f8f9fa;
  cursor: not-allowed;
  opacity: 0.7;
}

.form-group textarea {
  resize: vertical;
  min-height: 60px;
}

.error {
  color: #dc3545;
  font-size: 11px;
  margin-top: 4px;
  display: block;
  font-weight: 500;
}

.form-actions {
  display: flex;
  gap: 8px;
  flex-direction: column;
  margin-top: 20px;
}

.btn-submit,
.btn-cancel {
  padding: 8px 12px;
  border: none;
  border-radius: 4px;
  font-size: 12px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
  width: 100%;
}

.btn-submit {
  background: #007bff;
  color: white;
}

.btn-submit:hover:not(:disabled) {
  background: #0056b3;
  transform: translateY(-1px);
}

.btn-submit:disabled {
  background: #6c757d;
  cursor: not-allowed;
  transform: none;
}

.btn-cancel {
  background: #6c757d;
  color: white;
}

.btn-cancel:hover:not(:disabled) {
  background: #545b62;
  transform: translateY(-1px);
}

.btn-cancel:disabled {
  background: #adb5bd;
  cursor: not-allowed;
  transform: none;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .form-group {
    margin-bottom: 16px;
  }
  
  .form-group input,
  .form-group textarea {
    padding: 10px;
  }
  
  .form-actions {
    margin-top: 20px;
  }
  
  .btn-submit,
  .btn-cancel {
    padding: 10px 14px;
  }
}
</style>
