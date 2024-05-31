<template>
  <el-breadcrumb class="app-breadcrumb hidden-sm-and-down" separator="/">
    <transition-group appear name="breadcrumb">
      <el-breadcrumb-item v-for="(item, index) in levelList" :key="item.path">
        <span
          v-if="item.redirect === 'noRedirect' || index == levelList.length - 1"
          class="no-redirect"
        >{{  isBackMenu ? item.meta.title : $t(item.meta.title) }}</span>
        <a v-else @click.prevent="handleLink(item)">
          {{ isBackMenu ? item.meta.title : $t(item.meta.title) }}
        </a>
      </el-breadcrumb-item>
    </transition-group>
  </el-breadcrumb>
</template>

<script setup>
import { ref, defineComponent, watch } from "vue";
import { useRoute, useRouter } from "vue-router";

    const levelList = ref([]);
    const route = useRoute();
    const router = useRouter();
    const getBreadcrumb = () => {
      let matched = route.matched.filter(item => item.meta && item.meta.title);
      const first = matched[0];
      levelList.value = matched.filter(
        item => item.meta && item.meta.title && item.meta.breadcrumb !== false
      );
    };
    getBreadcrumb();

    const handleLink = (item) => {
      const { redirect, path } = item;
      if (redirect) {
        router.push(redirect.toString());
        return;
      }
      router.push(path);
    };

</script>

<style scoped >
.app-breadcrumb.el-breadcrumb {
  display: inline-block;
  font-size: 14px;
  line-height: 50px;
  .no-redirect {
    color: var(--system-header-breadcrumb-text-color);
    cursor: text;
  }
  a {
    color: var(--system-header-text-color);
  }
}
</style>