<script setup lang="ts">
import { RouterView } from 'vue-router'
import config from '@/assets/naive-ui-theme-overrides.json'

</script>

<template>
  <n-config-provider :theme-overrides="config">
    <n-message-provider>
      <n-dialog-provider>
        <RouterView />
      </n-dialog-provider>
    </n-message-provider>
  </n-config-provider>
</template>

<style scoped>
</style>
