<script setup lang="ts">
const props = defineProps({
  model: String,
  label: String,
})
</script>

<template>
  <v-text-field
    v-model="props.model"
    :label="props.label"
    hide-details
    required
  ></v-text-field>
</template>
