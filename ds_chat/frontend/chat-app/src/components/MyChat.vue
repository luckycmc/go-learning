<template>
  <div>
    <h2>Chat</h2>
    <div>
      <input v-model="username" placeholder="Your name" />
    </div>
    <div>
      <input v-model="message" placeholder="Your message" @keyup.enter="sendMessage" />
      <button @click="sendMessage">Send</button>
    </div>
    <ul>
      <li v-for="(msg, index) in messages" :key="index">
        <strong>{{ msg.username }}:</strong> {{ msg.content }}
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: '',
      message: '',
      messages: [],
      ws: null
    }
  },
  created() {
    this.connectWebSocket()
  },
  methods: {
    connectWebSocket() {
      this.ws = new WebSocket('ws://localhost:8080/ws')

      this.ws.onmessage = (event) => {
        const msg = JSON.parse(event.data)
        this.messages.push(msg)
      }

      this.ws.onclose = () => {
        console.log('WebSocket connection closed')
      }
    },
    sendMessage() {
      if (this.message && this.username) {
        const msg = {
          username: this.username,
          content: this.message
        }
        this.ws.send(JSON.stringify(msg))
        this.message = ''
      }
    }
  }
}
</script>
