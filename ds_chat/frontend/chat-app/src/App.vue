<template>
  <div id="app">
    <h1>Chat App</h1>
    <MyChat />
  </div>
</template>

<script>
import MyChat from './components/MyChat.vue'

export default {
  name: 'App',
  components: {
    MyChat
  }
}
</script>
