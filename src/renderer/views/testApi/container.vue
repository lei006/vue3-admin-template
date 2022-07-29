
<template>
  <el-container>
  1111
  </el-container>
</template>

<script>

export default {

  data() {
    return {
      list: null,
    }
  },
  created() {
    //this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      getList().then(response => {
        this.list = response.data.items
        this.listLoading = false
      })
    }
  }
}
</script>

<style lang="scss" scoped>

.box-container {
  width: 100%;
  height: 100%;
  color: red;
  background-color:#222;
}


</style>
