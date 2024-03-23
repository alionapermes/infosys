<template>
  <h1>{{ isConnected ? 'Connected' : 'Disconnected' }}</h1>
  <button @click="changeConnState">Change Connection State</button>
</template>

<script lang="ts" setup>
import { ref } from 'vue'

interface Response {
  result: {
    is_connected: boolean
    error?: string | undefined
  }
}

const isConnected = ref<boolean>(false)

const changeConnState = () => {
  fetch('//api:8081/conn', {
    method: 'GET',
    headers: {
      //'Accept': 'application/json',
      'Content-Type': 'application/json',
      // 'Allow-Control-Request': '',
    },
  })
    .then(async (response) => {
      const data: Response = await response.json()
      console.log(data)

      if (typeof data.result.error !== 'undefined') {
        console.error(data.result.error)
        return
      }

      isConnected.value = data.result.is_connected
    })
}
</script>
