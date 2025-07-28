<script setup lang="ts">
import ChatInterface from './components/ChatInterface.vue'
</script>

<template>
  <div id="app">
    <header>
      <h1>AI Chat</h1>
    </header>
    <main>
      <ChatInterface />
    </main>
  </div>
</template>

<style scoped>
#app {
  min-height: 100vh;
  display: flex;
  flex-direction: column;
}

header {
  background-color: #2c3e50;
  color: white;
  padding: 1rem;
  text-align: center;
}

header h1 {
  margin: 0;
  font-size: 1.5rem;
}

main {
  flex: 1;
  display: flex;
  flex-direction: column;
}
</style>
