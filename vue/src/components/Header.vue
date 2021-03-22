<template>
  <div class="header">
    <router-link to="/">
      <img
        v-if="info.logo"
        :src="info.logo"
        alt="BaiYue云盘"
        style="height: 56px; width: auto"
        id="logo"
      />
      <a-spin v-else />
    </router-link>
    <a-space>

      <a-space v-if="type === 'file'">
        <a-button
          type="primary"
          shape="circle"
          size="large"
          @click="copyFileLink"
        >
          <template #icon><copy /></template>
        </a-button>
        <a target="_blank" :href="downloadUrl">
          <a-button type="primary" shape="circle" size="large">
            <template #icon><download /></template>
          </a-button>
        </a>
      </a-space>

    </a-space>
  </div>
  <a-divider class="header-content" />
</template>

<script lang="ts">
import { GlobalDataProps } from "@/store"
import { computed, defineComponent, ref } from "vue"
import { useStore } from "vuex"
import { useRoute, useRouter } from "vue-router"
import useDownloadUrl from "../hooks/useDownloadUrl"
export default defineComponent({
  name: "Header",
  setup() {
    const store = useStore<GlobalDataProps>()
    const route = useRoute()
    const router = useRouter()
    const info = computed(() => store.state.info)
    const url = ref<string>(window.location.href)

    const type = computed(() => store.state.type)
    const { downloadUrl,copyFileLink } = useDownloadUrl()
    const keyword = ref<string>('');

    const onSearch = (searchValue: string) => {
      router.push(route.path+'?q='+searchValue)
    };
    return {
      info,
      url,
      type,
      copyFileLink,
      downloadUrl,
      keyword,
      onSearch
    };
  },
});
</script>

<style scoped>
.header {
  padding-top: 3px;
  height: 56px;
  width: 100%;
  display: flex;
  display: -webkit-flex; /* Safari */
  justify-content: space-between;
  align-items: center;
}
.header-content {
  margin: 10px 0 5px 0;
}
@media screen and (max-width: 600px) {
  .qrcode {
    display: none;
  }
}
</style>