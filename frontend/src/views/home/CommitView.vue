<template>
  <n-card :bordered="false" header-style="padding: 0;"
          content-style="padding: 0;height: 100%;" style="height: 100%">
    <div>
      <n-space>
        <n-card title="新增文件">
          <div v-for="item in updateFiles.addFiles" :key="item">
            <n-tag type="success">{{ item }}</n-tag>
          </div>
        </n-card>
        <n-card title="删除文件">
          <div v-for="item in updateFiles.deleteFiles" :key="item">
            <n-tag type="error">{{ item }}</n-tag>
          </div>
        </n-card>
        <n-card title="修改文件">
          <div v-for="item in updateFiles.modifyFiles" :key="item">
            <n-tag type="warning">{{ item }}</n-tag>
          </div>
        </n-card>
      </n-space>
      <div  style="padding: 20px 0">
        <n-input type="textarea" v-model:value="commitMessage" placeholder="请输入commit信息" />
      </div>
    </div>
    <template v-slot:footer>
      <n-flex justify="end">
        <n-button type="primary" @click="Commit">确 认</n-button>
      </n-flex>
    </template>
  </n-card>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { GreetService } from '#/changeme/index'
import { useMessage } from 'naive-ui'
import { useAppStore } from '@/stores/appStore'

const appStore = useAppStore()
const message = useMessage()
const commitMessage = ref('')
const updateFiles = ref({})

function Commit() {
  GreetService.CommitChanges(appStore.directory, commitMessage.value).then(() => {
    message.success('提交成功')
  }).catch(() => {
    message.error('提交失败')
  })
}


onMounted(() => {
  GetUpdateFiles()
})

// 获取变更文件列表
function GetUpdateFiles() {
  GreetService.GetUpdateFiles(appStore.directory).then((res) => {
    console.log(res)
    updateFiles.value = res
  })
}


</script>

<style scoped>

</style>