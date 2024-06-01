<template>
  <Layer :layer="layer" @confirm="submit">
    <el-form :model="ruleForm" :rules="rules" ref="form" label-width="120px" style="margin-right:30px;">
      <el-form-item label="名称：" prop="username">
        <el-input v-model="ruleForm.username" placeholder="请输入名称"></el-input>
      </el-form-item>
      <el-form-item label="昵称：" prop="nickname">
        <el-input v-model="ruleForm.nickname" placeholder="请输入昵称"></el-input>
      </el-form-item>
      <el-form-item label="密码：" prop="password">
        <el-input v-model="ruleForm.password" placeholder="请输入密码"></el-input>
      </el-form-item>
      <el-form-item label="确认密码:" prop="password1">
        <el-input v-model="ruleForm.password1" placeholder="请输入密码"></el-input>
      </el-form-item>
			<el-form-item label="选择器：" prop="select">
			  <el-select v-model="ruleForm.select" placeholder="请选择" clearable>
					<el-option
						v-for="item in options"
						:key="item.value"
						:label="item.label"
						:value="item.value">
					</el-option>
				</el-select>
			</el-form-item>
    </el-form>
  </Layer>
</template>

<script lang="ts">
import { defineComponent, ref, reactive } from 'vue'
import Layer from '@/components/layer/index.vue'
import apiUser from "@/api/user.js"

export default defineComponent({
  components: {
    Layer
  },
  props: {
    layer: {
      type: Object,
      default: () => {
        return {
          show: false,
          title: '',
          row:{},
          showButton: true
        }
      }
    }
  },
  setup(props, ctx) {
    let ruleForm = reactive({
      username: '',
      nickname: '',
      password: '',
      password1: '',
    })
    const rules = {
      username: [{ required: true, message: '请输入姓名', trigger: 'blur' }],
      nickname: [{ required: true, message: '请输入昵称', trigger: 'blur' }],
      sort: [{ required: true, message: '请输入数字', trigger: 'blur' }],
      select: [{ required: true, message: '请选择', trigger: 'blur' }],
      radio: [{ required: true, message: '请选择', trigger: 'blur' }]
    }

    Object.assign(ruleForm, props.layer.row)

    const options = [
      { value: 1, label: '运动'},
      { value: 2, label: '健身'},
      { value: 3, label: '跑酷'},
      { value: 4, label: '街舞'},
    ]
    return {
      ruleForm,
      rules,
      options
    }
  },
  methods: {
    submit() {
      this.$refs['form'].validate((valid: boolean) => {
        if (valid) {
          let params = this.ruleForm
          if (this.layer.row) {
            this.updateForm(params)
          } else {
            this.addForm(params)
          }
        } else {
          return false;
        }
      });
    },
    // 新增提交事件
    addForm(params: object) {
      apiUser.Create(params).then(res => {
        this.$message({
          type: 'success',
          message: '新增成功'
        })
        this.layer.show = false
        this.$emit('getTableData', true)
      })
    },
    // 编辑提交事件
    updateForm(params: object) {
      apiUser.PutOne(params).then(res => {
        this.$message({type: 'success',message: '编辑成功'})
        this.$emit('getTableData', false)
      })
    }
  }
})
</script>

<style lang="scss" scoped>
  
</style>