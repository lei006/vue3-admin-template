<template>
    <div class="system-button text-no-select" @click="onClick">
        111
    </div>
</template>

<script>

export default {


    data() {
        return {
        }
    },
    methods: {
        onClick(e){
            this.$emit("click", e);
        }
    }
}
</script>

<style lang="scss" scoped>


.system-button {

    min-width: 34px;
    height:34px;
    line-height: 34px;
  
    padding: 5px;

    border-radius: 3px;

    margin-right: 5px;

    display: flex;
    justify-content: center;
    align-items: center;

  cursor: pointer;   
}

.system-button:hover {
  background-color: #ccc;
  color: #fff;
}


.system-button-icon{
    font-size: 24px;
    display: flex;
}

.system-button-text{
    margin-left: 5px;
}


</style>
