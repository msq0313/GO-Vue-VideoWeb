<template>
    <div class="login-container">
        <el-form ref="form" :model="form" label-width="80px" label-position="left" class="login-page">
            <h3 class="title">用户登出</h3>
            <el-form-item>
                <el-button type="primary" @click="onsubmit">登出</el-button>
            </el-form-item>

        </el-form>
    </div>
</template>

<script>
    import * as API from '@/api/login/';

    export default {
        name: "Logout",
        methods: {
            onsubmit() {
                API.logout(this.form).then((res) => {
                    if(res.status > 0) {
                        this.$notify.error({
                            title: '登出失败',
                            message: res.msg
                        });
                    } else {
                        this.$notify({
                            title: '登出成功',
                            type:'success'
                        });
                    }
                }).catch((error) => {
                    this.$notify.error({
                        title: '网络错误，或者服务器宕机',
                        message: error
                    });
                });
            },
        },
        components: {

        },
    }

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