<template>
	<div class="pageheader hidden-xs">
		<h3><i class="fa fa-home"></i> 主页 </h3>
		
		<div class="row" style="padding: 0px; margin: 0px;">
			
			<div class="col-md-3"></div>
			
			<div class="col-md-6" >
		  
			<div class="well" style="padding: 0px; margin: 0px;">
				<div class="input-group">
					<input type="text" class="form-control search-input" placeholder="Search...">
						<div class="input-group-btn">
							<button class="btn btn-info search-btn" type="button"><router-link to="search">Search</router-link></button>
						</div>
				</div>
			</div>
			
			<div class="col-md-3"></div>
			
		</div>
		
		</div>
		
	    <div class="breadcrumb-wrapper">
	      <span class="label">你的位置:</span>
	         <ol class="breadcrumb">
	            <li> <a href="#"> 同济大学 </a> </li>
	            <li class="active"> 嘉定校区 </li>
	         </ol>
	    </div>
	</div>
</template>

<script>
</script>

<style>
</style>