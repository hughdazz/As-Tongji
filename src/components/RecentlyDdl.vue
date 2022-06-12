<template>
	<div class="panel panel-danger">
		<div class="panel-heading">
		    <h3 class="panel-title">最近的ddl</h3>
		</div>
	    <div class="panel-body">
			<div class="jumbotron" style="padding: 20px;">
			  <h1 >{{infor.Tittle}}</h1>
			  <p style="text-align: right;">{{infor.Ddl_time}}</p>
			  <p style="text-align: right;"><a class="btn btn-primary btn-lg" href="#" role="button">Do Now!</a></p>
			</div>
			  
			
	    
	    </div>
	</div>
</template>

<script>
export default{
	name:'RecentlyDdl',
	
	data(){
		return {

			infor :{
				Tittle: '51条指令多周期CPU',
				Ddl_time:'6.29-哎我'
			}
		
		};
	}
}
	
	
</script>

<style>
</style>