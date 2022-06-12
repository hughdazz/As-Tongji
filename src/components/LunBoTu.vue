<template>
	
	<div id="carousel-example-generic" class="carousel slide" data-ride="carousel">
	  <!-- Indicators -->
	  <ol class="carousel-indicators">
	    <li data-target="#carousel-example-generic" data-slide-to="0" class="active"></li>
	    <li data-target="#carousel-example-generic" data-slide-to="1"></li>
	    <li data-target="#carousel-example-generic" data-slide-to="2"></li>
	  </ol>
	
	  <!-- Wrapper for slides -->
	  <div class="carousel-inner" role="listbox">
		  
	    <div class="item active">
	      <img src="https://static.runoob.com/images/mix/img_nature_wide.jpg">
	      <div class="carousel-caption">
	        第1幅图
	      </div>
	    </div>
			
	    <div  v-for="infor in LunBoTu_KeJiQianYan" class="item">
	      <img src='{{infor.imgUrl}}'>
	      <div class="carousel-caption">
	       {{infor.Description}}
	      </div>
	    </div>
		
		
	
	  </div>
	
	  <!-- Controls -->
	  <a class="left carousel-control" href="#carousel-example-generic" role="button" data-slide="prev">
	    <span class="glyphicon glyphicon-chevron-left" aria-hidden="true"></span>
	    <span class="sr-only">Previous</span>
	  </a>
	  <a class="right carousel-control" href="#carousel-example-generic" role="button" data-slide="next">
	    <span class="glyphicon glyphicon-chevron-right" aria-hidden="true"></span>
	    <span class="sr-only">Next</span>
	  </a>
	</div>
			
</template>

<script>
export default{
	name:'LunBoTu',
	
	data(){
		return {

			LunBoTu_KeJiQianYan:[
				{
					imgUrl:'https://static.runoob.com/images/mix/img_nature_wide.jpg',
					Description:'我是猪'
				},
				{
					MimgUrl:'https://img1.baidu.com/it/u=621378781,1734092370&fm=253&fmt=auto&app=138&f=JPEG?w=750&h=500',
					Description:'我是1'
				},
				{
					imgUrl:'https://static.runoob.com/images/mix/img_nature_wide.jpg',
					Description:'我是2'
				}
			]
		
		};
	}
}	
	
</script>

<style>
</style>