<template>
	<div class="bg-light">
	    <ul class="nav nav-section nav-justified">
	        <li v-for="infor in arr">
	            <div class="section">
					<div>
					
						<h4 class="text-center" style="display: inline-block">{{infor.MealName}}<i class="fa fa-caret-up text-success"></i> </h4>  
						
						<button class="btn btn-sm pull-right" data-click="panel-reload"><i class="fa fa-refresh"></i></button>
					
					</div>
					
					<div>
						
						<p class="text-center text-overflow"> {{infor.Details}} </p>
						<div class="progress progress-striped progress-xs">
							<div style="width: 60%;" aria-valuemax="100" aria-valuemin="0" aria-valuenow="60" role="progressbar" class="progress-bar"> 
								<span class="sr-only">60% Complete</span> 
							</div>
						</div>
					</div>
	                <p class="text-left"> <span class="text-success">{{infor.Description}}</span></p>
	            </div>
	        </li>
			
	       
	    </ul>
	</div>
	
</template>

<script>
export default{
	name:'MealsRecommendation',
	
	data(){
		return {

			arr:[
				{
					MealName:'早餐',
					Details:'两个肉包+一杯豆浆+一个鸡蛋',
					Description:'营养丰富'
				},
				{
					MealName:'午餐',
					Details:'红烧排骨+牛肉+葡萄酒',
					Description:'补充能量'
				},
				{
					MealName:'晚餐',
					Details:'小唐馄饨',
					Description:'爱的味道~'
				}
			]
		
		};
	}
}
	
</script>

<style>
</style>