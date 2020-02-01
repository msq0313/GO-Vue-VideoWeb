<template>
    <div class="login-container">
        <el-form ref="form" :model="form" label-width="80px" label-position="left" class="login-page">
            <h3 class="title">用户登录</h3>
            <el-form-item label="Username">
                <el-input type="text" v-model="form.user_name" placeholder="请输入用户名"/>
            </el-form-item>

            <el-form-item label="Password">
                <el-input type="password" v-model="form.password" placeholder="请输入密码"/>
            </el-form-item>

            <el-form-item>
                <el-button type="primary" @click="onsubmit" >立即登录</el-button>
            </el-form-item>

        </el-form>
    </div>
</template>

<script>
    import * as API from '@/api/login/';

    export default {
        name: 'Login',
        data() {
            return {
                form: {
                    user_name: '',
                    password: '',
                },
            };
        },
        methods: {
            onsubmit() {
                API.login(this.form).then((res) => {
                    if(res.code > 0) {
                        this.$notify.error({
                            title: '登录失败',
                            message: res.msg
                        });
                    } else {
                        this.$notify({
                            title: '登录成功',
                            type:'success'
                        });
                        this.$router.push({path: '/about'});
                    }
                }).catch((error) => {
                    this.$notify.error({
                        title: '登录失败',
                        message: error
                    });
                });
            },
        },
        components: {

        },
    };
</script>

<style>
    .login-container {
        width: 100%;
        height: 100%;
    }
    .login-page {
        -webkit-border-radius: 5px;
        border-radius: 5px;
        margin: 180px auto;
        width: 350px;
        padding: 35px 35px 15px;
        background: #fff;
        border: 1px solid #eaeaea;
        box-shadow: 0 0 25px #cac6c6;
    }
</style>
