<template>
  <n-flex vertical class="home">
    <n-input v-model:value="appStore.directory" @click="Init" placeholder="请设置工作目录" />
    <n-card title="远程仓库" :bordered="false" header-style="padding: 0;" content-style="padding: 0;">
      <n-flex vertical>
        <n-tag type="warning" v-if="appStore.remoteUrlList.length === 0">请初始化仓库</n-tag>
        <n-tag type="primary" v-for="item in appStore.remoteUrlList"
               :key="item"
        >{{ item }}
        </n-tag>
      </n-flex>
    </n-card>
    <div>
      <n-tabs v-model:value="activeTab" animated type="card" placement="left">
        <n-tab-pane name="0" tab="初始化" :disabled="appStore.remoteUrlList.length > 0">
          <InitView />
        </n-tab-pane>
        <n-tab-pane name="1" tab="拉取" :disabled="appStore.remoteUrlList.length == 0">
          <PullView />
        </n-tab-pane>
        <n-tab-pane name="2" tab="提交" :disabled="appStore.remoteUrlList.length == 0">
          <CommitView />
        </n-tab-pane>
        <n-tab-pane name="3" tab="推送" :disabled="appStore.remoteUrlList.length == 0">
          <PushView />
        </n-tab-pane>
      </n-tabs>
    </div>
  </n-flex>
</template>

<script setup lang="ts">
import { GreetService } from '#/changeme/index'
import { ref } from 'vue'
import { useAppStore } from '@/stores/appStore'
import InitView from '@/views/home/InitView.vue'
import PullView from '@/views/home/PullView.vue'
import CommitView from '@/views/home/CommitView.vue'
import PushView from '@/views/home/PushView.vue'

const appStore = useAppStore()
const activeTab = ref('0')
function Init() {
  GreetService.InitWordRepo().then((path: string) => {
    appStore.directory = path
    appStore.GetRemoteURL().then(() => {
      if (appStore.remoteUrlList.length > 0) {
        activeTab.value = '1'
      }
    })
  })
}

</script>

<style scoped>
.home {
  padding: 10px;
}
</style>
