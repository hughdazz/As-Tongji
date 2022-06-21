<template>
	<div id="carousel-example-vertical1" class="carousel vertical slide" data-ride="carousel">
	        <div class="carousel-inner" role="listbox">
	            <div class="item" v-for="infor in arr">
	                <div class="ticker-headline">
	                    <div class="media">
	                        <span class="pull-left"><i class="fa fa-twitter fa-4x text-azure"></i></span>
	                        <div class="media-body">
	                            <div class="h4"><strong>{{infor.Tittle}}</strong> <small>{{infor.Time}}</small></div>
	                            <p >
									{{infor.Description}}
	                            </p>
	                        </div>
	                    </div>
	                </div>
	            </div>
	            <div class="item active">
	                <div class="ticker-headline">
	                    <div class="media">
	                        <span class="pull-left"><i class="fa fa-facebook fa-4x text-primary"></i></span>
	                        <div class="media-body">
	                            <div class="h4"><strong>牛顿发现了万有引力</strong> <small>2020.8.3</small></div>
	                            <p >
	                            	中国大学生程序设计竞赛（China Collegiate Programming Contest， CCPC）是由教育部高等学校计算机类专业教学指导委员会主办的面向全国高校大学生的年度学科竞赛。为青年学子提供了一个展示数学基本功和数学思维的舞台。
	                            </p>
	                        </div>
	                    </div>
	                </div>
	            </div>
	        </div>
	        <!-- Controls -->
	        <a class="up carousel-control" href="#carousel-example-vertical1" role="button" data-slide="prev">
	        <span class="glyphicon glyphicon-chevron-up" aria-hidden="true" ></span>
	        <span class="sr-only">Previous</span>
	        </a>
	        <a class="down carousel-control" href="#carousel-example-vertical1" role="button" data-slide="next">
	        <span class="glyphicon glyphicon-chevron-down" aria-hidden="true"></span>
	        <span class="sr-only">Next</span>
	        </a>
	</div>
</template>

<script>
export default{
	name:'SchoolNews',
	
	data(){
		return {
			
			arr:[
				{
					Tittle:'牛顿吃苹果',				
					Time:'2020.6.7',
					Description:'很棒'
				},
				{
					Tittle:'李白不要再喝酒',
					Time:'2020.6.123',
					Description:'很123棒'
				},
				{
					Tittle:'杜甫想吃炒饭',
					Time:'2020.6.1414',
					Description:'很棒141'
				}
			]
		
		};
	}

}	
	
</script>

<style>
</style>