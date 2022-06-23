<template>
	
	
	<div id="container" style="position: relative;">
		<div class="snow-container">
		  <div class="snow foreground"></div>
		  <div class="snow foreground layered"></div>
		  <div class="snow middleground"></div>
		  <div class="snow middleground layered"></div>
		  <div class="snow background"></div>
		  <div class="snow background layered"></div>
		</div>

		<div class="top-buttons-agileinfo">
		<a href="index.html"  class="active">登录</a><a href="">注册</a>
		</div>
		<h1>登陆</h1>
		<div class="main-agileits">
		<!--form-stars-here-->
				<div class="form-w3-agile">
					
					<form v-if="state==0" action="#" method="post">
							<input type="email" name="Eamil" placeholder="Email" required="" />
							<input type="password" name="Password" placeholder="Password" required="" />
							<a href="#" class="forgot-w3layouts">忘记密码 ?</a>
						<div class="submit-w3l">
							<input type="submit" @click="submit()" value="登录">
						</div>
						<p class="p-bottom-w3ls"><a href="signup.html">这里注册</a>如果你还没有一个账号.</p>
					</form>
					
					<div v-if="state==1" >
						
						<div ckass="verifyPicture"  style=" ">
							<img :src="IdetifyImgUrl" v-on:click="verify_click($event)">
						</div>
						 <h3 style="background-color: bisque;display: inline-block">{{verify_Tishi}}</h3>
						<button class="btn btn-success"  v-on:click="verify_refresh"><i class="fa fa-refresh"></i></button>
					</div>
					
					
					
				</div>
				</div>
		<!--//form-ends-here-->
		<!-- copyright -->
			<div class="copyright w3-agile">
				<p> © 2022 一如既往 . All rights reserved | Design By <a href="">Levoisa</a></p>
			</div>
			<!-- //copyright --> 
	</div>
	
	
</template>

<script> 
	import '../static/css/login_style.css';
	import '../static/css/login_snow.css';
	
	import MyIdentifyImg from '../components/MyIdentifyImg.vue'

	export default{
		name:'login',
		
	/* 	setup(props,ctx){
				  const movement = MouseMovement();
				  return {movement};
			  }, */
		
		data(){
			return {
				state:0,//int 表示状态
				//0:未登录；1：点击登录从后端返回验证码并显示；2：登录成功（应该具有session功能） 
				
				verify_Tishi:"请依次点击 7 3 4",
				clickNum:0,
				clickXY: [ 
						{
							x:'',
							y:' '
						},
						{
							x:'',
							y:' '
						},
						{
							x:'',
							y:' '
						}
				],
				
				IdetifyCode_input:"",//../static/img/验证码-测试用.JPG
				IdetifyImgUrl:"../static/img/验证码-测试用.JPG",//src:"/static/img/验证码-测试用.JPG"';
				Username:"",
				secret:"",
				is_Jizhu:''
			};
		},
		methods:{
			submit() {
				if(this.state == 0){
					this.IdetifyImgUrl='https://img0.baidu.com/it/u=2673243120,3070690751&fm=253&fmt=auto&app=138&f=JPEG?w=417&h=211';
					this.state = 1;
					return ;
				}
				if (this.state == 1 ) {
					if (!this.IdetifyCode_input) {
					  this.$message({
						type: 'warning',
						duration: '2000',
						message: '验证码不能为空！'
					  })
					  return
					}
					if (this.IdetifyCode_input.toLowerCase() !='abcd') {
					  
					}
				}
			},
			
			verify_click(event) {
				if(this.clickNum>=3)
					return ;
				this.clickXY[this.clickNum].x=event.offsetX; 
				this.clickXY[this.clickNum].y=event.offsetY; 
				console.log(this.clickXY,this.clickNum);
				this.clickNum++;
			},
			verify_refresh(){
				this.clickNum=0;
			}
		}
		
	}	



</script>

<style>
	
	
	
</style>