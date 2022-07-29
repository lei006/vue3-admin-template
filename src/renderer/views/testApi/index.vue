<template>
    <div class="test-box">
        <div class="test-aside">
            <TestAside></TestAside>
        </div>
        <div class="test-contain">
            <TestContainer></TestContainer>
        </div>
    </div>
</template>

<script>

//const Container = require('./container.vue')

import TestAside from './aside'
import TestContainer from './container'


export default {
  components:{
    TestAside, TestContainer
  },
  data() {
    return {
      list: null,
      listLoading: true
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

.test-box {
  width: 100%;
  height: 100%;
  color: red;
  background-color:#222;

  display: flex;
  flex-direction: row;
}

.test-aside{
  height: 100%;
  width: 200px;

  background-color:#ddd;

}

.test-contain{
  height: 100%;
  flex:1;
  background-color: bisque;
}



</style>

