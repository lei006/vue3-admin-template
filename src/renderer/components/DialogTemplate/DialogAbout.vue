<template>
    <DialogBase01>
       我是测试
    </DialogBase01>
</template>

<script>


import DialogBase01 from './DialogBase01.vue'

export default {
    components:{
        DialogBase01
    },
    data() {

        function isRegisterUserName(s)
        {
            var patrn=/^([a-zA-Z0-9]|[._@]){4,32}$/;
            if (!patrn.exec(s)) return false
            return true
        }

        var validateUsername = (rule, value, callback) => {
            if (isRegisterUserName(value) === false) {
            callback(new Error('4-32个字母,数字、“_”、“.”、“@”组成的字串'));
            }
            callback();
        };
        var validatePass = (rule, value, callback) => {
            if (value === '') {
            callback(new Error('请输入密码'));
            } else {
            callback();
            }
        };
        
        return {
        dialogVisible:false,
        dialog_modal:false,
        click_dialog_modal_close:false,
        ruleForm: {
            password: '',
            username: ''
        },
        rules: {
            password: [
                { validator: validatePass, trigger: 'blur' }
            ],
            username: [
                { validator: validateUsername, trigger: 'blur' }
            ]
        },      
        sendsms_disabled:false,
        sendsms_text:"发送验证码",
        sendsms_space: "",
        }
    },
    created() {
    },
    methods: {
        handleLink(item) {

        },
        show(){
            this.dialogVisible = true;
        },
        async onBtnSendSms() {
        let username = this.ruleForm.username;
        ipcRenderer.once("send-sms-ack", (event, ret) => {
            if ( ret.code == 20000 ) {
                this.sendsmsSuccess();
            }else{
                this.$message({ message: ret.message, type: 'error'});
            }
        });
        ipcRenderer.send("send-sms", username);
    
        },

    }
}
</script>

<style lang="scss" scoped>


</style>
