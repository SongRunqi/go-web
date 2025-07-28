<template>
  <div class="chat-container">
    <div class="messages-container" ref="messagesContainer">
      <div
        v-for="(message, index) in messages"
        :key="index"
        :class="['message', message.role === 'user' ? 'user-message' : 'ai-message']"
      >
        <div class="message-content">
          <strong>{{ message.role === 'user' ? 'You' : 'AI' }}:</strong>
          {{ message.content }}
        </div>
      </div>
      <div v-if="isLoading" class="message ai-message">
        <div class="message-content">
          <strong>AI:</strong>
          <span class="loading">Thinking...</span>
        </div>
      </div>
    </div>
    
    <div class="input-container">
      <div class="model-selector">
        <label for="model">Model:</label>
        <select id="model" v-model="selectedModel">
          <option value="gpt-4o-mini">GPT-4o Mini</option>
          <option value="gpt-3.5-turbo">GPT-3.5 Turbo</option>
          <option value="gpt-4">GPT-4</option>
        </select>
      </div>
      
      <div class="message-input">
        <textarea
          v-model="inputMessage"
          @keydown.enter.prevent="sendMessage"
          placeholder="Type your message here... (Press Enter to send)"
          :disabled="isLoading"
        ></textarea>
        <button @click="sendMessage" :disabled="isLoading || !inputMessage.trim()">
          Send
        </button>
      </div>
      
      <button @click="clearChat" class="clear-button">Clear Chat</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, nextTick } from 'vue'
import { chatService } from '../services/chatService'

interface Message {
  role: 'user' | 'assistant'
  content: string
}

const messages = ref<Message[]>([])
const inputMessage = ref('')
const selectedModel = ref('gpt-4o-mini')
const isLoading = ref(false)
const messagesContainer = ref<HTMLElement>()

const scrollToBottom = async () => {
  await nextTick()
  if (messagesContainer.value) {
    messagesContainer.value.scrollTop = messagesContainer.value.scrollHeight
  }
}

const sendMessage = async () => {
  if (!inputMessage.value.trim() || isLoading.value) return

  const userMessage: Message = {
    role: 'user',
    content: inputMessage.value.trim()
  }

  messages.value.push(userMessage)
  const currentInput = inputMessage.value.trim()
  inputMessage.value = ''
  isLoading.value = true
  
  await scrollToBottom()

  try {
    const response = await chatService.sendMessage({
      model: selectedModel.value,
      messages: messages.value
    })

    const aiMessage: Message = {
      role: 'assistant',
      content: response
    }

    messages.value.push(aiMessage)
  } catch (error) {
    console.error('Chat error:', error)
    const errorMessage: Message = {
      role: 'assistant',
      content: 'Sorry, I encountered an error. Please try again.'
    }
    messages.value.push(errorMessage)
  } finally {
    isLoading.value = false
    await scrollToBottom()
  }
}

const clearChat = () => {
  messages.value = []
}
</script>

<style scoped>
.chat-container {
  display: flex;
  flex-direction: column;
  height: 100%;
  max-height: calc(100vh - 120px);
}

.messages-container {
  flex: 1;
  overflow-y: auto;
  padding: 1rem;
  background-color: #f5f5f5;
}

.message {
  margin-bottom: 1rem;
  padding: 0.75rem;
  border-radius: 8px;
  max-width: 80%;
}

.user-message {
  background-color: #007bff;
  color: white;
  margin-left: auto;
  text-align: right;
}

.ai-message {
  background-color: white;
  border: 1px solid #ddd;
  margin-right: auto;
}

.message-content {
  word-wrap: break-word;
}

.loading {
  font-style: italic;
  opacity: 0.7;
}

.input-container {
  padding: 1rem;
  background-color: white;
  border-top: 1px solid #ddd;
}

.model-selector {
  margin-bottom: 1rem;
}

.model-selector label {
  display: inline-block;
  margin-right: 0.5rem;
  font-weight: bold;
}

.model-selector select {
  padding: 0.25rem;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.message-input {
  display: flex;
  gap: 0.5rem;
  margin-bottom: 1rem;
}

.message-input textarea {
  flex: 1;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  resize: none;
  min-height: 60px;
  font-family: inherit;
}

.message-input button {
  padding: 0.75rem 1.5rem;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-weight: bold;
}

.message-input button:hover:not(:disabled) {
  background-color: #0056b3;
}

.message-input button:disabled {
  background-color: #ccc;
  cursor: not-allowed;
}

.clear-button {
  padding: 0.5rem 1rem;
  background-color: #dc3545;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.clear-button:hover {
  background-color: #c82333;
}
</style>