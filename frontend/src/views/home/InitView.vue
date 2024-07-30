<template>
  <n-card title="远程仓库" :bordered="false" header-style="padding: 0;"
          content-style="padding: 0;height: 100%;" style="height: 100%">
    <n-input v-model:value="remoteURL" placeholder="请输入远程仓库地址" />
    <template v-slot:footer>
      <n-flex justify="end">
        <n-button type="primary" @click="Commit" :loading>确 认</n-button>
      </n-flex>
    </template>
  </n-card>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { GreetService } from '#/changeme'
import { useMessage } from 'naive-ui'
import { useAppStore } from '@/stores/appStore'

const remoteURL = ref('')
const message = useMessage()
const loading = ref(false)
const appStore = useAppStore()

function Commit() {
  if (appStore.directory === '') {
    message.error('请先选择项目目录')
    return
  }
  loading.value = true
  GreetService.InitGitRepository(appStore.directory, remoteURL.value).then(() => {
    message.success('初始化成功')
    appStore.GetRemoteURL()
  }).catch(() => {
    message.error('初始化失败')
  }).finally(() => {
    loading.value = false
  })
}
</script>

<style scoped>

</style>