<template>


<el-dialog v-model="layer.show" title="Shipping address" width="800">
    <el-form :model="form" :rules="rules" ref="ruleForm" label-width="120px" style="margin-right:30px;">
      <el-form-item label="用户名：" prop="name">
        管理员
      </el-form-item>
      <el-form-item label="原密码：" prop="old">
        <el-input v-model="form.old" placeholder="请输入原密码" show-password></el-input>
      </el-form-item>
			<el-form-item label="新密码：" prop="new">
			  <el-input v-model="form.new" placeholder="请输入新密码" show-password></el-input>
			</el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">Create</el-button>
        <el-button>Cancel</el-button>
      </el-form-item>      
    </el-form>
</el-dialog>



</template>

<script setup>

import { defineComponent, ref } from 'vue'
import { ElMessage } from 'element-plus'

  defineProps({
    layer: {
      type: Object,
      default: () => {
        return {
          show: false,
          title: '',
          showButton: true
        }
      }
    }})


  const ruleForm = ref(null)
    const layerDom = ref(null)
    let form = ref({
      userId: '123465',
      name: '',
      old: '',
      new: ''
    })
    const rules = {
      old: [{ required: true, message: '请输入原密码', trigger: 'blur' }],
      new: [{ required: true, message: '请输入新密码', trigger: 'blur' }],
    }
    function onSubmit() {
      if (ruleForm.value) {
        ruleForm.value.validate((valid) => {
          if (valid) {
            let params = {
              id: form.value.userId,
              old: form.value.old,
              new: form.value.new
            }
            /*
            passwordChange(params).then(res => {
              ElMessage({type: 'success',message: '密码修改成功，即将跳转到登录页面'})
              layerDom.value && layerDom.value.close()
              setTimeout(() => {
                //store.dispatch('user/loginOut')
              }, 2000)
            })
            */
          } else {
            return false;
          }
        });
      }
    }









</script>

<style scoped>
  
</style>