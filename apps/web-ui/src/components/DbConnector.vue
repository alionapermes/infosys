<script lang="ts" setup>
import { ref } from 'vue'

import { post, Response } from '../util'

import TextInput from './TextInput.vue'

const isConnected = ref<boolean>(false)

const changeConnState = () => {
  post('/conn', {
    dbname: dbName,
    dbuser: dbUser,
    dbpass: dbPass,
    dbport: parseInt(dbPort),
  }).then(async (response) => {
    const data: Response = await response.json()
    console.log(data)

    if (typeof data.result.error !== 'undefined') {
      console.error(data.result.error)
      return
    }

    isConnected.value = data.result.is_connected
  })
}

let dbName = 'postgres'
let dbUser = 'postgres'
let dbPass = 'postgres'
let dbPort = '5432'
</script>

<template>
  <v-sheet class="mx-auto" width="300" color="black">
    <TextInput :model="dbName" label="db name"/>
    <TextInput :model="dbUser" label="db user"/>
    <TextInput :model="dbPass" label="db pass"/>
    <TextInput :model="dbPort" label="db port"/>
    <v-btn
      @click="changeConnState"
      :color="isConnected ? 'red' : 'light-green'"
      class="my-2"
    >{{ isConnected ? 'Закрыть соединение' : 'Открыть соединение' }}
    </v-btn>
  </v-sheet>
</template>
