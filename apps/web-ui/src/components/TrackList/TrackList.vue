<script lang="ts" setup>
import { ref } from 'vue'

import { get } from '../../util'

import { Track, Response } from '.'
import Item from './Item.vue'

const tracks = ref<Track[]>([])

const getTracks = () => {
  get('/tracks').then(async (response) => {
    const data: Response = await response.json()
    console.log(data)

    if (typeof data === 'undefined') {
      console.error(data)
      return
    }

    tracks.value = data.tracks
  })
}
</script>

<template>
  <v-sheet class="mx-auto" width="600" color="black">
    <v-btn
      @click="getTracks"
      class="my-2"
    >
    Получить список треков
    </v-btn>
    <div v-for="(track, index) in tracks" :key="index">
      <Item :instance="track"/>
    </div>
  </v-sheet>
</template>
