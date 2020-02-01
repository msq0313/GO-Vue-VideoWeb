<template>
	<div class="home">
		<div class="">
			<el-row :gutter="20">
				<el-col :span="4" v-for="video in videos" :key="video.id">
					<el-card>
						<img src="../../public/timg.jpg" class="video-avatar">
						<div>
							<div class="video-title">{{video.title}}</div>
							<div class="video-bottom clearfix">
								<span class="video-info">{{video.info}}</span>
							</div>
						</div>
					</el-card>
				</el-col>
			</el-row>
		</div>
	</div>
</template>

<script>
import * as API from '@/api/video/';

export default {
	name: 'home',
	data() {
		return {
			videos: [],
		};
	},
	methods: {
		load() {
			API.getVideos().then((res) => {
				this.videos = res.data;
			});
		},
	},
	components: {
	},
	beforeMount() {
		this.load();
	},
};
</script>

<style>
	.video-avatar {
		width: 100%;
	}
	.video-title {
		margin: 10px 0 4px 0;
	}
	.video-bottom {
		margin-top: 4px;
	}
	.video-info {
		color: #999;
	}
</style>