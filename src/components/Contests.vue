<template>
	<table id="demo-foo-addrow" class="table table-bordered table-hover toggle-circle" data-page-size="7">
	    <thead>
	        <tr>
	            <th data-sort-initial="true" data-toggle="true">名称</th>
	            <th>级别</th>
	            <th data-hide="phone, tablet">简介</th>
	            <th data-hide="phone, tablet">时间</th>
	            <th data-hide="phone, tablet">推荐指数</th>
	            <th data-sort-ignore="true" class="min-width">是否加入日程</th>
	        </tr>
	    </thead>
	    <tbody>
	        <tr v-for="infor in arr">
	            <td>{{infor.Name}}</td>
	            <td>{{infor.Degree}}</td>
	            <td>{{infor.Description}}</td>
	            <td>{{infor.Time}}</td>
	            <td><span class="label label-table label-success">{{infor.Advice}}</span></td>
	            <td>
	                <button class="btn btn-danger fa fa-trash"></button>
	            </td>
	        </tr>
	    </tbody>
	    <tfoot>
	        <tr>
	            <td colspan="6">
	                <div class="text-right">
	                    <ul class="pagination">
							<li class="footable-page-arrow disabled">
								<a data-page="first" href="#first">«</a>
							</li>
							<li class="footable-page-arrow disabled">
								<a data-page="prev" href="#prev">‹</a>
							</li>
							<li class="footable-page active">
								<a data-page="0" href="#">1</a>
							</li>
							<li class="footable-page-arrow disabled">
								<a data-page="next" href="#next">›</a>
							</li>
							<li class="footable-page-arrow disabled">
								<a data-page="last" href="#last">»</a>
							</li>
						</ul>
	                </div>
	            </td>
	        </tr>
	    </tfoot>
	</table>
							    
</template>

<script>
export default{
	name:'Contests',
	
	data(){
		return {
			
			arr:[
				{
					Name:'同济大学数学竞赛',
					Degree:'校级',
					Time:'2020.6.7',
					Description:'很棒',
					Advice:'超级推荐'
				},
				{
					Name:'同济大学123',
					Degree:'校级123',
					Time:'2020.6.7',
					Description:'很棒123',
					Advice:'超级推荐123'
				},
				{
					Name:'同济大学数学竞赛qwe',
					Degree:'校级qwe',
					Time:'2020.6.7qe',
					Description:'很棒qwe',
					Advice:'超级推qwe荐'
				}
			]
		
		};
	}
}
		
	
</script>

<style>
</style>