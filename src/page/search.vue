<template>
	<div id="container" class="effect mainnav-sm navbar-fixed mainnav-fixed">
	
		<MyTopNavgation></MyTopNavgation>
	 
	    <div class="boxed">
	        <!--CONTENT CONTAINER-->
	        <!--===================================================-->
	        <div id="content-container">
	         
				<MyPageTittile></MyPageTittile>
				
	            <!--Page content-->
	            <!--===================================================-->
	            <div id="page-content">
					
					<Gallery></Gallery>
				</div>
	            <!--===================================================-->
	            <!--End page content-->
	        </div>
	        <!--===================================================-->
	        <!--END CONTENT CONTAINER-->
	        
			<MyLeftNavgation></MyLeftNavgation>
			<MyFooter></MyFooter>
	    </div>
	    
	    <!-- SCROLL TOP BUTTON -->
	    <!--===================================================-->
	    <button id="scroll-top" class="btn"><i class="fa fa-chevron-up"></i></button>
	    <!--===================================================-->
	</div>
	
	
</template>

<script setup>
import MyTopNavgation from '../components/MyTopNavgation.vue'
import MyLeftNavgation from '../components/MyLeftNavgation.vue'
import MyFooter from '../components/MyFooter.vue'
import MyPageTittile from '../components/MyPageTittile.vue'
import Gallery from '../components/Gallery.vue'
</script>

<style>
	

</style>