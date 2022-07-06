
<template>

    <router-view v-slot="{ Component }">
        <transition>
            <component :is="Component" ></component>
        </transition>
    </router-view>

</template>

<script setup>

/*
let routePath = computed(()=>{
    let {currentRoute } = useRouter()
      return currentRoute.value.path;
})

onMounted(()=>{

})
*/
</script>