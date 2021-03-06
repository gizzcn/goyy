// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplNewProjAdmTemplatesCoreCommBreadcrumb = `<!-- 导航栏 -->
<ol id="{%.param%}Breadcrumb" class="breadcrumb"></ol>
<script type="text/javascript">
	// 导航栏点击或列表查找子集数据时调用函数
	var {%.param%}BreadcrumbPage=function(id){
		$("#{%.param%}ListSform").find("input[name='sParentId']").val(id);
		{%.param%}Page(1);
	}
	// 导航栏菜单生成函数 在page函数调用
	var {%.param%}Breadcrumb=function(datas){
		if($.isNotBlank(datas)){
			if(typeof(datas)=="string"){
				datas=$.parseJSON(datas);
			}
			var html="";
			for(var i=0;i<datas.length;i++){
				var data=datas[i];
				if(data.active){
					html+='<li class="active">'+data.name+'</li>';
				}else{
					html+='<li><a href="#list-content" onclick="{%.param%}BreadcrumbPage(\''+data.id+'\')">'+data.name+'</a></li>';
				}
			}
			$("#{%.param%}Breadcrumb").html(html);// 导航栏菜单
		}
	}
</script>
`
