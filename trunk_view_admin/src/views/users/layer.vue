<template>
  <Layer :layer="layer" @confirm="submit">
    <el-form :model="ruleForm" :rules="rules" ref="form" label-width="120px" style="margin-right:30px;">
      <el-form-item label="名称：" prop="username">
        <el-input v-model="ruleForm.username" placeholder="请输入名称"></el-input>
      </el-form-item>
      <el-form-item label="昵称：" >
        <el-input v-model="ruleForm.nickname" placeholder="请输入昵称"></el-input>
      </el-form-item>
      <el-form-item label="密码：" prop="password">
        <el-input v-model="ruleForm.password" placeholder="请输入密码"></el-input>
      </el-form-item>
      <el-form-item label="重录密码" prop="password1">
        <el-input v-model="ruleForm.password1" placeholder="请再次输入密码"></el-input>
      </el-form-item>

      <el-form-item label="签名：" >
        <el-input v-model="ruleForm.usersign" placeholder="请输入签名"></el-input>
      </el-form-item>

      <el-form-item label="是否停用:" >
        <el-checkbox v-model="ruleForm.is_disable" label="停用" />
      </el-form-item>

      <el-form-item label="描述：" >
        <el-input v-model="ruleForm.desc" placeholder="请输入描述"></el-input>
      </el-form-item>
    </el-form>
  </Layer>
</template>

<script setup lang="ts">
import { defineComponent, ref, reactive } from 'vue'
import Layer from '@/components/layer/index.vue'
import { add, update } from '@/api/table'
import apiUsers from "@/api/system/user";
import { ElMessage } from 'element-plus'
import stringRandom  from 'string-random'


let layer = ref({
          show: false,
          title: '',
          showButton: true
        })

let is_edit = false;


        
let ruleForm = reactive({
  username: '',
  nickname: '',
  password: '',
  password1: '',
  usersign: '',
  is_disable: false,
  desc: '',
})
const rules = {
  username: [{ required: true, message: '请输入姓名', trigger: 'blur' }],
  password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
  password1: [{ required: true, message: '确认密码', trigger: 'blur' }],
  sort: [{ required: true, message: '请输入数字', trigger: 'blur' }],
  select: [{ required: true, message: '请选择', trigger: 'blur' }],
  radio: [{ required: true, message: '请选择', trigger: 'blur' }]
}
const options = [
  { value: 1, label: '运动'},
  { value: 2, label: '健身'},
  { value: 3, label: '跑酷'},
  { value: 4, label: '街舞'},
]



const submit = () => {

    let params = ruleForm
    // is_disable 必需是 bool,true
    params.is_disable = (params.is_disable != 0);

    if (is_edit) {
        updateForm(params)
    }else{
        addForm(params)
    }
}

// 新增提交事件
const addForm = (params: object) => {
    apiUsers.AddOne(params).then(res => {
        ElMessage({type: 'success',message: '新增成功'})
        layer.value.show = false;
    })
}

// 编辑提交事件
const updateForm = (params: object) => {

  apiUsers.PutOne(params).then(res => {
        ElMessage({type: 'success',message: '编辑成功'})
        layer.value.show = false;
    })
}



const showEdit = (row) => {

    console.log("", row);

    Object.assign(ruleForm, row)


    is_edit = true;
    layer.value.title = "编辑数据";
    layer.value.show = true;    

}

const showAdd = () => {
    is_edit = false;


    ruleForm.username = stringRandom()
    ruleForm.nickname = stringRandom()
    ruleForm.password = stringRandom()
    ruleForm.password1 = stringRandom()
    ruleForm.usersign = stringRandom()

    console.log("ruleForm", ruleForm);

    layer.value.title = "新增数据";
    layer.value.show = true;
    //delete layer.row;
}


defineExpose({showEdit,showAdd})








</script>

<style lang="scss" scoped>
  
</style>