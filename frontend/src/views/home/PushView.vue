<template>
  <n-card title="推送仓库" :bordered="false" header-style="padding: 0;"
          content-style="padding: 0;height: 100%;" style="height: 100%">
    <n-flex>
      <n-radio-group v-model:value="remoteURL">
        <n-radio v-for="(item,index) in appStore.remoteUrlList" :key="index" :value="item">{{ item }}
        </n-radio>
      </n-radio-group>
    </n-flex>
    <template v-slot:footer>
      <n-flex justify="end">
        <n-button type="primary" @click="Commit" :loading>确 认</n-button>
      </n-flex>
    </template>
  </n-card>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { GreetService } from '#/changeme/index'
import { useMessage } from 'naive-ui'
import { useAppStore } from '@/stores/appStore'

const appStore = useAppStore()
const message= useMessage()
const remoteURL = ref('')
const loading = ref(false)


function Commit() {
  loading.value = true
  GreetService.PushChanges(appStore.directory, remoteURL.value, 'master').then(() => {
    message.success('推送成功')
  }).catch(() => {
    message.error('推送失败')
  }).finally(() => {
    loading.value = false
  })
}

</script>

<style scoped>

</style>