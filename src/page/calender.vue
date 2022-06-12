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
	                <section id="page-content">
	                    <div class="row">
	                        <div class="col-md-4 col-lg-3">
	                            <div class="panel">
	                                <div class="panel-heading">
	                                    <h3 class="panel-title">Events</h3>
	                                </div>
									<div class="panel-body">
										<!-- Draggable Events --> 
										<!-- ============================================ -->
										<h4 class="text-thin">Draggable Events</h4>
										<div id="demo-external-events">
											<div class="fc-event">My Event 1</div>
											<div class="fc-event">My Event 2</div>
											<div class="fc-event">My Event 3</div>
											<div class="fc-event">My Event 4</div>
											<div class="fc-event">My Event 5</div>
											<div class="fc-event">My Event 6</div>
											<hr>
											<div>
												<label class="form-checkbox form-normal form-primary">
												<input type="checkbox" id="drop-remove">
												Remove after drop </label>
											</div>
										</div>
										<!-- ============================================ --> 
									</div>
	                            </div>
	                        </div>
	                        <div class="col-md-8 col-lg-9">
	                            <div class="panel">
	                                <div class="panel-body">
	                                    <!-- Calendar placeholder--> 
	                                    <!-- ============================================ -->
	                                    <div id='demo-calendar'></div>
	                                    <!-- ============================================ --> 
	                                </div>
	                            </div>
	                        </div>
	                    </div>
	                </section>
	                <!--===================================================--> 
	                <!--End page content--> 
	            <!--===================================================--> 
	            <!--END CONTENT CONTAINER--> 
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
</script>

<style>
</style>