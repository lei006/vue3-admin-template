<template>
    <div class="system-button-col text-no-select" :style="{  'color': color, }" @click="onClick">
        <div class="system-button-icon"><slot>ICON</slot></div><div class="system-button-text">{{text}}</div>
    </div>
</template>

<script>

export default {

    props: {
        text: {
          type: String,
          default: ''
        },
        color: {
          type: String,
          default: '#222'
        },
    },



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


.system-button-col {

    width: 60px;
    min-height:34px;
    padding: 5px;

    border-radius: 3px;

    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;

    cursor: pointer;   
}

.system-button-col:hover {
  background-color: #ccc;
  color: #fff;
}


.system-button-icon{
    font-size: 24px;
}

.system-button-text{
    font-size: 14px;
    font-weight: 400;
    white-space:nowrap;
    font-family: 'Microsoft YaHei';
}


</style>
