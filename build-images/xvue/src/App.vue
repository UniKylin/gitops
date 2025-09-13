<template>
  <div class="app">
    <header class="header">
      <h1>🎉 Welcome to Xvue!</h1>
      <p>A Vue 3 Application</p>
    </header>
    
    <main class="main">
      <div class="card">
        <h2>Hello App 👋</h2>
        
        <div class="counter">
          <p>Counter: {{ count }}</p>
          <button @click="increment" class="btn">Click me!</button>
        </div>
        
        <div class="message">
          <input 
            v-model="message" 
            placeholder="Type a message..."
            class="input"
          />
          <p v-if="message">You typed: {{ message }}</p>
        </div>
      </div>
    </main>
  </div>
</template>

<script>
import { ref } from 'vue'

export default {
  name: 'App',
  setup() {
    const count = ref(0)
    const message = ref('')
    
    const increment = () => {
      count.value++
    }
    
    return {
      count,
      message,
      increment
    }
  }
}
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  min-height: 100vh;
}

.app {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 20px;
}

.header {
  text-align: center;
  margin-bottom: 40px;
  color: white;
}

.header h1 {
  font-size: 3rem;
  margin-bottom: 10px;
  text-shadow: 2px 2px 4px rgba(0,0,0,0.3);
}

.header p {
  font-size: 1.2rem;
  opacity: 0.9;
}

.main {
  width: 100%;
  max-width: 500px;
}

.card {
  background: white;
  border-radius: 20px;
  padding: 40px;
  box-shadow: 0 20px 40px rgba(0,0,0,0.1);
  text-align: center;
}

.card h2 {
  color: #333;
  margin-bottom: 20px;
  font-size: 2rem;
}

.card p {
  color: #666;
  margin-bottom: 30px;
  line-height: 1.6;
}

.counter {
  margin-bottom: 30px;
  padding: 20px;
  background: #f8f9fa;
  border-radius: 10px;
}

.counter p {
  font-size: 1.5rem;
  font-weight: bold;
  color: #667eea;
  margin-bottom: 15px;
}

.btn {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  padding: 12px 24px;
  border-radius: 25px;
  font-size: 1rem;
  cursor: pointer;
  transition: transform 0.2s, box-shadow 0.2s;
}

.btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 10px 20px rgba(102, 126, 234, 0.3);
}

.message {
  padding: 20px;
  background: #f8f9fa;
  border-radius: 10px;
}

.input {
  width: 100%;
  padding: 12px;
  border: 2px solid #e9ecef;
  border-radius: 8px;
  font-size: 1rem;
  margin-bottom: 15px;
  transition: border-color 0.2s;
}

.input:focus {
  outline: none;
  border-color: #667eea;
}

.message p {
  color: #28a745;
  font-weight: bold;
  margin: 0;
}
</style>
