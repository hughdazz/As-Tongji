<template>
	
	<div class="panel panel-mint">
		 <div class="panel-heading">
					
			 <h3 class="panel-title">活动盒子</h3>
		 </div>
		<div class=" pad-all msg_container_base" style="height: 430px ;overflow-y: scroll;">
			
			<div v-for="infor in arr" >
				<div class="panel">
					<div class="panel-body" style="height: 160px ;">
						<div class="nano" style="height: 140px ;">
							<div class="nano-content">
							<div class="row">
								<div class="col-md-9 col-sm-9 col-xs-10">
									<h3 class="mar-no"> <span class="counter ">{{infor.Tittle}}</span></h3>
									<p class="mar-ver-5 text-overflow">{{infor.Time}}</p>
								</div>
								<div class="col-md-3 col-sm-3 col-xs-2"> <i class="fa fa-support fa-3x text-success"></i> </div>
							</div>
							
							<div class="progress progress-striped progress-xs">
								<div style="width: 60%;" aria-valuemax="100" aria-valuemin="0" aria-valuenow="60" role="progressbar" class="progress-bar progress-bar-success"> 
									<span class="sr-only">60% Complete</span> 
								</div>
							</div>
							
							<div class="col-md-9 col-sm-9 col-xs-9">
								<p class="mar-ver-5 text-overflow">{{infor.Position}}</p>
							</div>
								 
							<div class="col-md-3 col-sm-3 col-xs-3"> <button type="button" class="btn btn-info btn-sm">{{infor.is_join}}</button></div>
							</div>
						</div>
							
					</div>
				</div>
			</div>
			
	
		 </div>
	</div>	
		                   
</template>

<script>
export default{
	name:'EventBox',
	
	data(){
		return {

		arr:[
			{
				Tittle:'电院2020级新生开学典礼',
				Time:'2020.9.11',
				Position:'智信楼一楼大会议室',
				is_join:'加入'
			},
			{
				Tittle:'吃瓜大会',
				Time:'2020.9.11',
				Position:'朱建铭的卧室',
				is_join:'必加入'
			},
			{
				Tittle:'电院2020级新生开学典礼',
				Time:'2020.9.11',
				Position:'智信楼一楼大会议室',
				is_join:'加入'
			}
		]
		
		};
		
		
	}
	/*
	mounted (){
		
	  this.axios
      .post('http://v.juhe.cn/toutiao/index?type=top&key=APPKEY')
      .then(response=>{
		  consloe.log(response);
		  response => (this.infors = response.data);
	  })
      .catch(error=>{ // 请求失败处理
        console.log(error);
		console.log('******error');
      });
	}
	*/
	
}
	
	
	
</script>

<style>
</style>