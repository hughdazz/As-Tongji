<template>
	<!--MAIN NAVIGATION-->
	<!--===================================================-->
	<nav id="mainnav-container">
	    <!--Brand logo & name-->
	    <!--================================-->
	    <div class="navbar-header">
	        <a href="index.html" class="navbar-brand">
	            <i class="fa fa-forumbee brand-icon"></i>
	            <div class="brand-title">
	                <span class="brand-text">济网</span>
	            </div>
	        </a>
	    </div>
	    <!--================================-->
	    <!--End brand logo & name-->
	    <div id="mainnav">
	        <!--Menu-->
	        <!--================================-->
	        <div id="mainnav-menu-wrap">
	            <div class="nano">
	                <div class="nano-content">
	                    <ul id="mainnav-menu" class="list-group">
	                        <!--Category name-->
	                        <li class="list-header">导航栏</li>
	                        <!--Menu list item-->
	                        <li> <a href="index.html"> <i class="fa fa-home"></i> <span class="menu-title"> 信息广场 </span> </a> </li>
	                        <!--Category name-->
	                        <li class="list-header">信息分类</li>
	                        <!--Menu list item-->
	                        <li>
	                            <a href="#">
	                            <i class="fa fa-th-large"></i>
	                            <span class="menu-title">
	                               快速导航
	                            </span>
	                            <i class="arrow"></i>
	                            </a>
	                            <!--Submenu-->
	                            <ul class="collapse">
									
									<li><a href="layout-blank.html"><i class="fa fa-caret-right"></i> 1块钱 教学管理信息系统 </a></li>
									<li><a href="layout-boxed.html"><i class="fa fa-caret-right"></i> Canvas 控制面板 </a></li>
									<li><a href="layout-collapsed-sidebar.html"><i class="fa fa-caret-right"></i> 同济云课堂 </a></li>
									<li><a href="layout-push-menu.html"><i class="fa fa-caret-right"></i> 同济大学综合服务门户 </a></li>
									<li><a href="layout-push-menu.html"><i class="fa fa-caret-right"></i> 同济大学迎新网 </a></li>
								
	                            </ul>
	                        </li>
							
							<li>
							    <a href="#">
							    <i class="fa fa-list-ul"></i>
							    <span class="menu-title">
							       时事之窗
							    </span>
							    <i class="arrow"></i>
							    </a>
							    <!--Submenu-->
							    <ul class="collapse">
									
									<li> <a href="index.html"> <i class="fa fa-building"></i> <span class="menu-title"> 美食 </span> </a> </li>
									<li> <a href="index.html"> <i class="fa fa-list"></i> <span class="menu-title"> 竞赛 </span> </a> </li>
									<li> <a href="index.html"> <i class="fa fa-building"></i> <span class="menu-title"> 胜地 </span> </a> </li>
									<li> <a href="index.html"> <i class="fa fa-building"></i> <span class="menu-title"> 活动 </span> </a> </li>
									<li> <a href="index.html"> <i class="fa fa-building"></i> <span class="menu-title"> 新闻 </span> </a> </li>
							    </ul>
							</li>
							
				
							<li>
								<router-link to="/calender">
							    <i class="fa fa-calendar"></i>
							    <span class="menu-title">
							    日程安排
							    </span>
							    </router-link>
							</li>
							
							<li>
							    <a href="#">
							    <i class="fa fa-wechat"></i>
							    <span class="menu-title">
							       社交区
							    </span>
							    <i class="arrow"></i>
							    </a>
							    <!--Submenu-->
							    <ul class="collapse">
									
									<li><a href="layout-blank.html"><i class="fa fa-caret-right"></i> 开黑</a></li>
									<li><router-link to="/chat"><i class="fa fa-caret-right"></i> 聊天广场 </router-link></li>
								
							    </ul>
							</li>
							
							<!--Menu list item-->
							<li>
							    <a href="#">
							    <i class="fa fa-map-marker"></i>
							    <span class="menu-title">
							    地图
							    <span class="label label-mint pull-right">New</span>
							    </span>
							    </a>
							    <!--Submenu-->
							    <ul class="collapse">
							        <li><a href="maps-gmap.html">Google Maps</a></li>
							        <li><a href="maps-vectormap.html">Vector Maps</a></li>
							    </ul>
							</li>
						
							
							<!--Menu list item-->
							<li>
							    <a href="#">
							    <i class="fa fa-arrow-up"></i>
							    <span class="menu-title">
							    一键上课
							    <span class="label label-mint pull-right">New</span>
							    </span>
							    </a>
							</li>
							
							
	                    </ul>
	                    <!--Widget-->
	                    <!--================================-->
	                    <div class="mainnav-widget">
	                        <!-- Show the button on collapsed navigation -->
	                        <div class="show-small">
	                            <a href="#" data-toggle="menu-widget" data-target="#demo-wg-server">
	                            <i class="fa fa-desktop"></i>
	                            </a>
	                        </div>
	                        <!-- Hide the content on collapsed navigation -->
	                        <div id="demo-wg-server" class="hide-small mainnav-widget-content">
	                            <ul class="list-group">
	                                <li class="list-header pad-no pad-ver">一些status</li>
	                                <li class="mar-btm">
	                                    <span class="label label-primary pull-right">15%</span>
	                                    <p>本学期 进度</p>
	                                    <div class="progress progress-sm">
	                                        <div class="progress-bar progress-bar-primary" style="width: 15%;">
	                                            <span class="sr-only">15%</span>
	                                        </div>
	                                    </div>
	                                </li>
	                                <li class="mar-btm">
	                                    <span class="label label-purple pull-right">75%</span>
	                                    <p>本科 进度</p>
	                                    <div class="progress progress-sm">
	                                        <div class="progress-bar progress-bar-purple" style="width: 75%;">
	                                            <span class="sr-only">75%</span>
	                                        </div>
	                                    </div>
	                                </li>
	                            </ul>
	                        </div>
	                    </div>
	                    <!--================================-->
	                    <!--End widget-->
	                </div>
	            </div>
	        </div>
	        <!--================================-->
	        <!--End menu-->
	    </div>
	</nav>
	<!--===================================================-->
	<!--END MAIN NAVIGATION-->
</template>

<script>
</script>

<style>
</style>