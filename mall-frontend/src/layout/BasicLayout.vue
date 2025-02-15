<template>
  <div id="BasicLayout">
    <a-layout>
      <a-layout-header class="header">
        <GlobalHeader> </GlobalHeader>
      </a-layout-header>
      <a-layout-content class="content">
        <router-view></router-view>
      </a-layout-content>
      <a-layout-footer class="footer">
        <a
          href="https://github.com/YuanxiLeeL?tab=repositories"
          target="_blank"
        >
          Yuanxi的个人仓库主页
        </a>
      </a-layout-footer>
    </a-layout>
  </div>
</template>
<script setup lang="ts">
import GlobalHeader from "@/components/GlobalHeader.vue";
</script>
<style scoped>
#BasicLayout .footer {
  background: rgb(235, 207, 207);
  text-align: center;
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  padding: 16px;
}

#BasicLayout .content {
  padding: 20px;
  margin-bottom: 20px;
  background: white;
}

#BasicLayout .header {
  background: white;
  margin-bottom: 16px;
  color: unset;
  padding-inline: 20px;
}
</style>
