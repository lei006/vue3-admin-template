<template>
  <component :is="type" v-bind="linkProps(to)" @click="hideMenu" >
    <slot></slot>
  </component>
</template>

<script setup>

import { ref,defineComponent, computed } from 'vue'


const props = defineProps({
    to: {
        type: String,
        required: true
    }
	});

const type = ref("router-link")

const isCollapse = computed(() => true);
const linkProps = (to) => {
  return {
    to: to
  } 
}
const hideMenu = () => {
  if (document.body.clientWidth <= 1000 && !isCollapse.value) {
    store.commit("app/isCollapseChange", true);
  }
};





/*
import { useStore } from "vuex";
export default defineComponent({
  name: 'appLink',
  props: {
    to: {
      type: String,
      required: true
    }
  },
  setup(props) {
    const store = useStore();
    const isCollapse = computed(() => store.state.app.isCollapse);
    const linkProps = (to) => {
     return {
       to: to
     } 
    }
    const hideMenu = () => {
      if (document.body.clientWidth <= 1000 && !isCollapse.value) {
        store.commit("app/isCollapseChange", true);
      }
    };
    return {
      type: "router-link",
      linkProps,
      hideMenu
    }
  }
})
*/
</script>


