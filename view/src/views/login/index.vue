<template>
  
    <div data-v-18439d2c="" class="login-container">
      <div data-v-18439d2c="" class="login-form">
          <img data-v-18439d2c="" alt="" class="left-img" src="@\assets\image\login_left.jpg">
          <div class="form-box">
          <el-form
              ref="loginForm"
              :model="loginFormData"
              :rules="rules"
              :validate-on-rule-change="false"
              @keyup.enter="submitForm"
            >

            <el-form-item class="mb-6">
              <div class="title">Hello</div>
            </el-form-item>
            <el-form-item class="mb-6">
              <div class="sub-title">欢迎使用本系统</div>
            </el-form-item>

              <el-form-item
                prop="username"
                class="mb-6"
              >
                <el-input
                  v-model="loginFormData.username"
                  size="large"
                  placeholder="请输入用户名"
                  suffix-icon="user"
                />
              </el-form-item>
              <el-form-item
                prop="password"
                class="mb-6"
              >
                <el-input
                  v-model="loginFormData.password"
                  show-password
                  size="large"
                  type="password"
                  placeholder="请输入密码"
                />
              </el-form-item>
              <el-form-item
                v-if="loginFormData.openCaptcha"
                prop="captcha"
                class="mb-6"
              >
                <div class="flex w-full justify-between">
                  <el-input
                    v-model="loginFormData.captcha"
                    placeholder="请输入验证码"
                    size="large"
                    class="flex-1 mr-5"
                  />
                  <div class="w-1/3 h-11 bg-[#c3d4f2] rounded">
                    <img
                      v-if="picPath"
                      class="w-full h-full"
                      :src="picPath"
                      alt="请输入验证码"
                      @click="loginVerify()"
                    >
                  </div>
                </div>
              </el-form-item>
              <el-form-item class="mb-6">
                <el-button
                  class="shadow shadow-blue-600 h-11 w-full"
                  type="primary"
                  size="large"
                  @click="submitForm"
                >登 录</el-button>
              </el-form-item>
              <!--
              <el-form-item class="mb-6">
                <el-button
                  class="shadow shadow-blue-600 h-11 w-full"
                  type="primary"
                  size="large"
                  @click="checkInit"
                >前往初始化</el-button>
              </el-form-item>
              -->
          </el-form>
          </div>

      </div>
    </div>

</template>

<script setup>
import apiAuth from '@/api/auth'
import { reactive, ref } from 'vue'
import { ElMessage } from 'element-plus'
import { useRouter } from 'vue-router'

import { useStore } from 'vuex'
const store = useStore()


const router = useRouter()



defineOptions({
  name: 'Login',
})




// 验证函数
const checkUsername = (rule, value, callback) => {
  if (value.length < 5) {
    return callback(new Error('请输入正确的用户名'))
  } else {
    callback()
  }
}
const checkPassword = (rule, value, callback) => {
  if (value.length < 6) {
    return callback(new Error('请输入正确的密码'))
  } else {
    callback()
  }
}

// 获取验证码
const loginVerify = async () => {
  
  apiAuth.captcha({}).then(async(ele) => {
    rules.captcha.push({
      max: ele.data.length,
      min: ele.data.length,
      message: `请输入${ele.data.length}位验证码`,
      trigger: 'blur',
    })
    picPath.value = ele.data.captcha
    loginFormData.captchaId = ele.data.id
    loginFormData.openCaptcha = ele.data.enable
  })
}




loginVerify()

// 登录相关操作
const loginForm = ref(null)
const picPath = ref('')
const loginFormData = reactive({
    username: 'admin',
    password: '123456',
    captcha: '',
    captchaId: '',
    openCaptcha: false,
})
const rules = reactive({
  username: [{ validator: checkUsername, trigger: 'blur' }],
  password: [{ validator: checkPassword, trigger: 'blur' }],
  captcha: [
    {
      message: '验证码格式不正确',
      trigger: 'blur',
    },
  ],
})


const login = async(data) => {

  //let ret = await apiAuth.login(loginFormData)
  let ret = store.dispatch('user/login', data)

  return ret;
}
const submitForm = () => {

  loginForm.value.validate(async(v) => {
    if (v) {
      const ret = await login(loginFormData)
      if (ret.code == 200) {
        router.push({ path: "/" })
        console.log("====================")
      }else{
        console.log("===11======", ret)

        loginVerify()
      }
    } else {
      ElMessage({type: 'error',message: '请正确填写登录信息',showClose: true,})
      loginVerify()
      return false
    }
  })
}


</script>
<style scoped>

.login-app {
  width: 100%;
  height: 100%;


  display: flex;
  justify-content: center;
  align-items: center;
}


.login-container {

  width: 100%;
  height: 100%;

  background-color: rgb(119, 112, 242);

  display: flex;
  flex-direction: column;

  align-items: center;
  justify-content:center;  
}

.login-form {

  width: 1000px;
  height: 560px;

  display: flex;
  flex-direction: row;

  align-items: center;
  justify-content:center;  

  background-color: rgb(242, 242, 242);
  border-radius: 10px;

  overflow: hidden;
}


.login-form .left-img {
  flex:1;

  width: 500px;
  height: 458px;

  padding: 30px;
  border-radius: 50px;
}

.login-form .form-box {
  flex:1;

  display: flex;
  justify-content: center;
}

.title {
  font-size: 48px;
}
.sub-title {
  font-size: 32px;
}


.justify-between {
  display: flex;
  height: 40px;
}

.justify-between img{
  height: 40px;
}



</style>
