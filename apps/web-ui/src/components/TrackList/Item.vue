<script setup lang="ts">
import { Track } from '.'

const props = defineProps<{
  instance: Track
}>()
</script>

<template>
  <div>
    <div>{{ props.instance.title }}</div>
    <div>{{ props.instance.album }}</div>
    <div>{{ props.instance.band }}</div>
  </div>
</template>
