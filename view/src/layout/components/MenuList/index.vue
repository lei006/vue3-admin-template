<template>
  <el-scrollbar>

    <div class="menu-box">
        <div class="btn-item" :class="activeMenu == '/report/edit'?'btn-item-active':'' " @click="jumpPath('/report/edit')">
            <div class="icon"><el-icon><User /></el-icon></div>
            <div class="text">撰写报告</div>
        </div>
        <div class="btn-item" :class="activeMenu == '/report/manage'?'btn-item-active':'' " @click="jumpPath('/report/manage')">
            <div class="icon"><el-icon><Setting /></el-icon></div>
            <div class="text">预置字段</div>
        </div>
        <div class="btn-item" :class="activeMenu == '/report/prefield'?'btn-item-active':'' " @click="jumpPath('/report/prefield')">
            <div class="icon"><el-icon><Setting /></el-icon></div>
            <div class="text">预置字段</div>
        </div>
        <div class="btn-item" :class="activeMenu == '/report/template'?'btn-item-active':'' " @click="jumpPath('/report/template')">
            <div class="icon"><el-icon><Warning /></el-icon></div>
            <div class="text">模板管理</div>
        </div>
        <div class="btn-item" :class="activeMenu == '/report/print'?'btn-item-active':'' " @click="jumpPath('/report/print')">
            <div class="icon"><el-icon><Warning /></el-icon></div>
            <div class="text">打印管理</div>
        </div>
        <!--
        <div class="btn-item" v-for="(menu, key) in allRoutes"  :key="key"  @click="jumpPath('')">
            <div class="icon"><el-icon><Warning /></el-icon></div>
            <div class="text">{{ menu }}</div>
        </div>
        -->
    </div>

  </el-scrollbar>
</template>

<script setup>

import { defineComponent, computed, onMounted } from 'vue'

import { useRouter, useRoute } from 'vue-router'
const router = useRouter()


const jumpPath = (path) => {
    router.push({path: path})
}


const route = useRoute()
const activeMenu = computed(() => {
    const { meta, path } = route;
    if (meta.activeMenu) {
      return meta.activeMenu;
    }
    return path;
});

const allRoutes = useRouter().options.routes


/*


import { defineComponent, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useStore } from 'vuex'
import MenuItem from './MenuItem.vue'
export default defineComponent({
  props: {
    mode: {
      type: String,
      default: "vertical"
    }
  },
  components: {
    MenuItem
  },
  setup(props) {
    const { mode } = props
    const store = useStore()
    const isCollapse = computed(() => store.state.app.isCollapse)
    const expandOneMenu = computed(() => store.state.app.expandOneMenu)
    const allRoutes = useRouter().options.routes
    const route = useRoute()
    const activeMenu: any = computed(() => {
      const { meta, path } = route;
      if (meta.activeMenu) {
        return meta.activeMenu;
      }
      return path;
    });
    onMounted(() => {

    })
    return {
      isCollapse,
      expandOneMenu,
      allRoutes,
      activeMenu,
      mode,
    }
  }
})
*/
</script>

<style lang="scss" scoped>
  .el-scrollbar {
    background-color: var(--system-menu-background);
    flex: 1;
  }
  .layout-menu {
    width: 100%;
    border: none;
    &.collapse {
      margin-left: 0px;
    }
    :deep() {
      .el-menu-item, .el-sub-menu {
        background-color: var(--system-menu-background) !important;
      }
      .el-menu-item i, .el-menu-item-group__title, .el-sub-menu__title i {
        color: var(--system-menu-text-color);
      }
      .el-menu-item, .el-sub-menu__title{
        &.is-active {
          background-color: var(--system-primary-color) !important;
          color: var(--system-primary-text-color) !important;
          i {
            color: var(--system-primary-text-color) !important;
          }
          &:hover {
            background-color: var(--system-primary-color) !important;
            color: var(--system-primary-text-color) !important;
          }
        }
        &:hover {
          background-color: var(--system-menu-hover-background) !important;
        }
      }
      .el-sub-menu {
        &.is-active {
          >.el-sub-menu__title, >.el-sub-menu__title i {
            color: var(--system-menu-submenu-active-color) !important;
          }
        }
        .el-menu-item {
          background-color: var(--system-menu-children-background) !important;
          &.is-active {
            background-color: var(--system-primary-color) !important;
            color: var(--system-primary-text-color) !important;
            &:hover {
              background-color: var(--system-primary-color) !important;
              color: var(--system-primary-text-color) !important;
            }
          }
          &:hover {
            background-color: var(--system-menu-hover-background) !important;
          }
        }
        .el-sub-menu {
          .el-sub-menu__title {
            background-color: var(--system-menu-children-background) !important;
            &:hover {
              background-color: var(--system-menu-hover-background) !important;
            }
          }
        }
      }
    }
  }







</style>