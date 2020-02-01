<template>
	<el-menu :default-active="this.$route.path" class="top-bar" mode="horizontal" router>
		<el-menu-item index="/">首页</el-menu-item>
		<el-menu-item index="/postvideo">投稿视频</el-menu-item>
		<el-menu-item index="/about">关于我们</el-menu-item>
<!-- 		<el-submenu index="2">
			<template slot="title">投稿</template>
			<el-menu-item index="2-1">选项1</el-menu-item>
			<el-menu-item index="2-2">选项2</el-menu-item>
			<el-menu-item index="2-3">选项3</el-menu-item>
			<el-submenu index="2-4">
				<template slot="title">选项4</template>
				<el-menu-item index="2-4-1">选项1</el-menu-item>
				<el-menu-item index="2-4-2">选项2</el-menu-item>
				<el-menu-item index="2-4-3">选项3</el-menu-item>
			</el-submenu>
		</el-submenu> -->
		<el-menu-item index="3" disabled>消息中心</el-menu-item>
		<el-menu-item index="4"><a href="https://www.ele.me" target="_blank">订单管理</a></el-menu-item>
		<el-menu-item index="/login">用户登录</el-menu-item>
		<el-menu-item index="/logout">登出</el-menu-item>
	</el-menu>
</template>

<script>
	export default{
		name:'NavBar',
	}
</script>

<style scoped lang="scss">
	.top-bar {
		margin-bottom: 30px;
	}
</style>
