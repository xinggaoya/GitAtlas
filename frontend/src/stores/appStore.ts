import { ref } from 'vue'
import { defineStore } from 'pinia'
import { GreetService } from '#/changeme'

export const useAppStore = defineStore('app-config', () => {
  const directory = ref<string>('')
  const remoteUrlList = ref<string[]>([])

  // 获取远程仓库
  async function GetRemoteURL(): Promise<void> {
    await GreetService.GetRemoteURL(directory.value).then((urls: string[]) => {
      remoteUrlList.value = urls
    })
    return
  }

  return { directory, remoteUrlList, GetRemoteURL }
})
