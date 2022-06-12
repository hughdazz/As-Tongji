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
					
	                <div class="row">
					  
						<div class="col-md-3 eq-box-md grid"  style="background-color: mintcream; padding: 0;">
							<EventBox></EventBox>
					   </div>	
				
						<div class="col-md-5 eq-box-md grid">
							
							<RecentlyDdl></RecentlyDdl>
	                        
							<MealsRecommendation LunBoTu='LunBoTu_KeJiQianYan'></MealsRecommendation>
							
	                    </div>
	                  
						<div class="col-md-4 eq-box-md grid">
	                         <!--轮播图-->
							<div class="panel panel-primary">
								<div class="panel-heading">
									 
								    <!-- <div class="panel-control">
								        <button class="btn btn-primary" data-click="panel-expand"><i class="fa fa-expand"></i></button>
								        <button class="btn btn-primary" data-click="panel-reload"><i class="fa fa-refresh"></i></button>
								        <button class="btn btn-primary" data-click="panel-collapse"><i class="fa fa-chevron-down"></i></button>
								        <button class="btn btn-primary" data-dismiss="panel"><i class="fa fa-times"></i></button>
								    </div> -->
									
								    <h3 class="panel-title">学术前沿</h3>
								</div>	
	                            <div class="panel-body">
									<LunBoTu></LunBoTu>		
	                            </div>
	                        </div>
							
							 <!--竞赛消息-->
							<div class="panel panel-success">
								<div class="panel-heading">
									<!-- 
								    <div class="panel-control">
								        <button class="btn btn-primary" data-click="panel-expand"><i class="fa fa-expand"></i></button>
								        <button class="btn btn-primary" data-click="panel-reload"><i class="fa fa-refresh"></i></button>
								        <button class="btn btn-primary" data-click="panel-collapse"><i class="fa fa-chevron-down"></i></button>
								        <button class="btn btn-primary" data-dismiss="panel"><i class="fa fa-times"></i></button>
								    </div>
									-->
								    <h3 class="panel-title">院内简讯</h3>
								</div>	
								<div class="panel-body">
										
									<SchoolNews></SchoolNews>
								</div>
							</div>
	                    </div>
	                </div>
					
					<div class="row" >
						
						<div class="panel panel-warning">
						    <div class="panel-heading">
						        <h3 class="panel-title">竞赛之旅</h3>
						    </div>
						    <!-- Foo Table - Add & Remove Rows -->
						    <!--===================================================-->
						    <div class="panel-body">
						        <Contests></Contests>
							</div>
						    <!--===================================================-->
						    <!-- End Foo Table - Add & Remove Rows -->
						</div>
						                  
								
						
					</div>
					
	                <div class="row">
						<!--美食 胜地图片-->
	                    <div class="col-md-8">
	                        <div class="panel panel-primary">
								<div class="panel-heading">
									<h3 class="panel-title">周边胜地</h3>
								</div>
	                            <div class="panel-body">
									
	                               <!--轮播图-->
	                               <div class="panel">
	                                   <div class="panel-body">
										   
						
	                               		<div id="carousel-example-generic1" class="carousel slide" data-ride="carousel">
	                               		  <!-- Indicators -->
	                               		  <ol class="carousel-indicators">
	                               		    <li data-target="#carousel-example-generic1" data-slide-to="0" class="active"></li>
	                               		    <li data-target="#carousel-example-generic1" data-slide-to="1"></li>
	                               		    <li data-target="#carousel-example-generic1" data-slide-to="2"></li>
	                               		  </ol>
	                               		
	                               		  <!-- Wrapper for slides -->
	                               		  <div class="carousel-inner" role="listbox">
	                               			  
	                               		    <div class="item active">
	                               		      <img src="https://static.runoob.com/images/mix/img_nature_wide.jpg">
	                               		      <div class="carousel-caption">
	                               		        第1幅图
	                               		      </div>
	                               		    </div>
	                               			
	                               		    <div class="item">
	                               		      <img src="https://static.runoob.com/images/mix/img_fjords_wide.jpg">
	                               		      <div class="carousel-caption">
	                               		        第2幅图
	                               		      </div>
	                               		    </div>
	                               			
	                               			<div class="item">
	                               			  <img src="https://img1.baidu.com/it/u=621378781,1734092370&fm=253&fmt=auto&app=138&f=JPEG?w=750&h=500">
	                               			  <div class="carousel-caption">
	                               			    第3幅图
	                               			  </div>
	                               			</div>
	                               		
	                               		  </div>
	                               		
	                               		  <!-- Controls -->
	                               		  <a class="left carousel-control" href="#carousel-example-generic1" role="button" data-slide="prev">
	                               		    <span class="glyphicon glyphicon-chevron-left" aria-hidden="true"></span>
	                               		    <span class="sr-only">Previous</span>
	                               		  </a>
	                               		  <a class="right carousel-control" href="#carousel-example-generic1" role="button" data-slide="next">
	                               		    <span class="glyphicon glyphicon-chevron-right" aria-hidden="true"></span>
	                               		    <span class="sr-only">Next</span>
	                               		  </a>
	                               		</div>
	                               									
	                                   </div>
	                               </div>	
	                            </div>
	                        </div>
	                    
						</div>
						
						<!--官网导航-->
	                    <div class="col-md-4">
									  
	                        <div class="panel panel-mint">
								<div class="panel-heading">
									<h3 class="panel-title">官网导航</h3>
								</div>
									
	                            <div class="panel-body" style="padding-bottom:0px ;">
									<div class="row" style="text-align: center;">
										
										<div class="col-sm-6 table-toolbar-left">
									    </div>
												
									    <div class="col-sm-6 table-toolbar-left">
									        <div class="form-group">
									            <input type="text" placeholder="Search" class="form-control" autocomplete="off">
									        </div>
									    </div>
												
									</div>
									                                                             
									<div class="table-responsive">
									    <table class="table table-striped">
									        <thead>
												<tr>
													<th>名称</th>
													<th>url</th>
													<th></th>
												</tr>
									        </thead>
									        <tbody>
									            <tr>
									                <td>一块钱</td>
									                <td>1.tongji.edu.cn</td>
													<td> <button type="button" class="btn btn-mint">Go</button></td>
									            </tr>
									            <tr>
									                <td>canvas</td>
									                <td>canvas.tongji.edu.cn</td>
													<td> <button type="button" class="btn btn-mint">Go</button></td>
									            </tr>
												<tr>
												    <td>一块钱</td>
												    <td>1.tongji.edu.cn</td>
													<td> <button type="button" class="btn btn-mint">Go</button></td>
												</tr>
												<tr>
												    <td>canvas</td>
												    <td>canvas.tongji.edu.cn</td>
													<td> <button type="button" class="btn btn-mint">Go</button></td>
												</tr>
											</tbody>
											<tfoot>
											    <tr>
											        <td colspan="3" style="padding-bottom:0px ;margin-bottom: 0px;">
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
									</div>
									
								</div>
	                                                                
	                        </div>
							
	                   
	                   
	                    </div>
	             
					</div>
					
					
	                <div class="row">
						
						 <!-- 教材 -->
						<div class="col-lg-4">
	                        <!--Default Tabs (Left Aligned)-->
	                        <!--===================================================-->
	                        <div class="tab-base">
	                            <!--Nav Tabs-->
	                            <ul class="nav nav-tabs">
	                                <li class="">
	                                    <a data-toggle="tab" href="#demo-lft-tab-1" aria-expanded="false">教材</a>
	                                </li>
	                                <li class="">
	                                    <a data-toggle="tab" href="#demo-lft-tab-2" aria-expanded="false">软件</a>
	                                </li>
	                                <li class="active">
	                                    <a data-toggle="tab" href="#demo-lft-tab-3" aria-expanded="true">题集</a>
	                                </li>
	                            </ul>
	                            <!--Tabs Content-->
	                            <div class="tab-content">
	                                <div id="demo-lft-tab-1" class="tab-pane fade">
	                                   
	                                </div>
	                                <div id="demo-lft-tab-2" class="tab-pane fade">
										
	                                  
									</div>
	                                <div id="demo-lft-tab-3" class="tab-pane fade active in">
										
	                                   
									</div>
	                            </div>
	                        </div>
	                        <!--===================================================-->
	                        <!--End Default Tabs (Left Aligned)-->
	                    </div>
						
						<!-- 教 -->
	            
	                   <div class="col-lg-4">
						    <!--===================================================-->
							   <!-- Contact us widget -->
							   <!--===================================================-->
						   <div class="panel panel-success">
							   <div class="panel-heading">
								   <div class="panel-title">
									   <i class="fa fa-question-circle fa-lg fa-fw text-primary"></i> 问题反馈
								   </div>
								</div>
							   <div class="panel-body">
								   <p>感谢您的反馈，我们将及时处理</p>
								   <div class="form-group">
									   <label>您的邮箱（方便后续与您联系）</label>
									   <input type="email" class="form-control" placeholder="youremail@domain.com">
								   </div>
								   <div class="form-group">
									   <label>问题描述</label>
									   <textarea class="form-control" placeholder="Your message" rows="3"></textarea>
								   </div>
								   <button class="btn btn-default"><span class="fa fa-paperclip"></span> Add attachment</button>
								   <button class="btn btn-success pull-right"><span class="fa fa-send"></span> Send</button>
							   </div>
							</div>
						</div>
						
						<!--官网导航-->
						<div class="col-lg-4">
									  
						    <div class="panel panel-info">
								<div class="panel-heading">
									<h3 class="panel-title">社区导航</h3>
								</div>
									
						        <div class="panel-body" style="padding-bottom:0px ;">
									<div class="row" style="text-align: center;">
										
										<div class="col-sm-6 table-toolbar-left">
									    </div>
												
									    <div class="col-sm-6 table-toolbar-left">
									        <div class="form-group">
									            <input type="text" placeholder="Search" class="form-control" autocomplete="off">
									        </div>
									    </div>
												
									</div>
									                                                             
									<div class="table-responsive">
									    <table class="table table-striped">
									        <thead>
												<tr>
													<th>名称</th>
													<th>qq号</th>
													<th></th>
												</tr>
									        </thead>
									        <tbody>
									            <tr>
									                <td>计科闲聊群</td>
									                <td>1346800940</td>
													<td> <button type="button" class="btn btn-mint">Go</button></td>
									            </tr>
									            <tr>
									                <td>高数共学群</td>
									                <td>8946164466</td>
													<td> <button type="button" class="btn btn-mint">Go</button></td>
									            </tr>
												<tr>
												    <td>线代/高代共学群</td>
												    <td>132141236</td>
													<td> <button type="button" class="btn btn-mint">Go</button></td>
												</tr>
												<tr>
												    <td>web技术社团</td>
												    <td>c49406449831</td>
													<td> <button type="button" class="btn btn-mint">Go</button></td>
												</tr>
											</tbody>
											<tfoot>
											    <tr>
											        <td colspan="3" style="padding-bottom:0px ;margin-bottom: 0px;">
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
									</div>
									
								</div>
						                                            
						    </div>
							
						                           
						                           
						</div>
						                     
	                 
	                </div>
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
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup

import HelloWorld from '../components/HelloWorld.vue'
import MyTopNavgation from '../components/MyTopNavgation.vue'
import MyLeftNavgation from '../components/MyLeftNavgation.vue'
import MyFooter from '../components/MyFooter.vue'
import MyPageTittile from '../components/MyPageTittile.vue'
import EventBox from '../components/EventBox.vue'
import RecentlyDdl from '../components/RecentlyDdl.vue'
import MealsRecommendation from '../components/MealsRecommendation.vue'
import LunBoTu from '../components/LunBoTu.vue'
import Contests from '../components/Contests.vue'
import SchoolNews from '../components/SchoolNews.vue'

</script>


<style>


</style>
