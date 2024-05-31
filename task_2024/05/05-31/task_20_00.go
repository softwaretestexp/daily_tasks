package main
import "fmt"
func main() {
	s := "<!DOCTYPE html>
<html lang="zh-CN">

<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge,Chrome=1" />
  <meta name="viewport" content="width=device-width,initial-scale=1.0,minimum-scale=1.0,maximum-scale=1.0,user-scalable=no" />
  <meta name="format-detection" content="telephone=no" />
  <title>计算机科学与技术学院</title>
  <meta name="keywords" content="" />
  <meta name="description" content="" />
  
<link type="text/css" href="/_css/_system/system.css" rel="stylesheet"/>
<link type="text/css" href="/_upload/site/1/style/2/2.css" rel="stylesheet"/>
<link type="text/css" href="/_upload/site/00/44/68/style/128/128.css" rel="stylesheet"/>
<link type="text/css" href="/_js/_portletPlugs/sudyNavi/css/sudyNav.css" rel="stylesheet" />
<link type="text/css" href="/_js/_portletPlugs/datepicker/css/datepicker.css" rel="stylesheet" />
<link type="text/css" href="/_js/_portletPlugs/simpleNews/css/simplenews.css" rel="stylesheet" />

<meta content="3evcRkb2l4JcqXoQOIS_3lgpS5DOYdoK" r="m"><!--[if lt IE 9]><script r='m'>document.createElement("section")</script><![endif]--><script type="text/javascript" r='m'>$_ts=window['$_ts'];if(!$_ts)$_ts={};$_ts.nsd=53092;$_ts.cd="qtxDrrAlWk7qtGEjqSq5tGl9ck7qtGQjrrqLiaqjqPLklkgjq1AEiaqjl1LllkgjqSqmrcLrhqGlhn0qhPExtGq9ck7qtG3jrSqLiaqXEGL4qcLltGVAHO7qHqWjqrq5iaqjrPgPmAZ7cpEYtGWArAA4qcLrtG3AHPLslqVctGE7cGAGtGqArAV4mPLDtGQAHO7qtGVjqfq5iaqjqcLMlkgjrPabvalCWAEmWOKIE_m9BjRu4qpcZ7qyevpTzJhDs9YVeY5LbAS2LOArCxFGqqVcraAq7qdDUv92JvOAWkrnhmw217yxRVwEwvTETlyXpcqeUT.lEYgj36w0hifah6xStDYS4nyu3orSMKh_hcgBEczaQ5NbUYGXFUmzZneXQDy.tDdjMCaNMoxLMMJzwK2XtuAzZUR7FYZ.x9tbWCJj3mfM8wWC1kTDtVeS40VTHVV2H2.SwbSNhb2SMIWBMvrXFnfB_be7tKSnFDCuMUYNMP70hImCMD2ItnRmdOqnMTl5FuOw1kmiJm3Z1HpzF2YuxsyFZUR7F1yNQK6uhCyaMb7LM5NBMc27Qbfz7CTTFKg.Jc6SwbSNUn7bVJr21DSSVKejdOVgt6pwYTO4RvpU3nV5KBmCMD9XFUmzZneXQDy.tDdjMCaNMoxLMMJzwK2XtuAzZUR7FYZ.x2hS1UNCwDznK_2NwY3T3TRqZkNzK9V2H2.SwbSNhb2SMIWBMvrXFnfB_be7tKSnFDCuMUYNMP70hImCMD2ItnRMyCRcMo7dMDve11g.Q6YBptw18OYmxsyFZUR7F1yNQK6uhCyaMb7LM5NBMc27Qbfz7CTTFKg.Jc6SwbSNUn7b3H3_RoxxsbfUZlgy3DSYJYOCAm9Zp1V5KBmCMD9XFUmzZneXQDy.tDdjMCaNMoxLMMJzwK2XtuAzZUR7FYZ.xbj.MYefQUluWjw0VuYgU9zK0V0yK9W2H2.SwbSNhb2SMIWBMvrXFnfB_be7tKSnFDCuMUYNMP70hImCMD2ItnRB0KNKMbTG3KH6YbeZQbyK3XYBQKNkxsyFZUR7F1yNQK6uhCyaMb7LM5NBMc27Qbfz7CTTFKg.Jc6SwbSNUn7bMZS9MVxqU0y3502BpOJHMVdNAOYP3nV5KBmCMD9XFUmzZneXQDy.tDdjMCaNMoxLMMJzwK2XtuAzZUR7FYe_tm.lMUYNMPzjwHJBhcQnx1GONuWPt1rRIUB7tCxLIbrLthmzwKpOtKm2_PTOFnq.FvsuMnaThb2SMIJ8hmy7Qbfz7CTTFKg.xODehnEeWPVLEZl.wDfjMDxug6wNRDS4wCOZM6SC3Uz63HmGForOFKJe4v2ZwUNdQo6Bw6STQDyjw8JuFoSSMb2e4Czy3oN68D4SRvSjFoz.R8SdMCpTw6weZUpa8KSnRKkbMoYSFvzjQBNdMbTBIP0BZURfRcSP3o1S3Dgvhb2SMIWBWc27QbfzaPeMF6w.Fc6SwbSNhPV0E4WfikVGtnRIzCrf3DY.8oIu3bxywKxz8IRzt1z7QbYG7KR9QcS9FPiuMUYNMP70hImCMD2ItTeB_be7tKSnFDCuEsqvhPVZWM0BE27jQC2NyKmZQUR98os5MvSuF6YawHyBQCSB86xX56r2w6pPwv6NQKY.RKyjRiJ0FoY5QDTB56wTFCSPFo6LQnyB3Uzu8Ip_8om5Qoyv_oegJ6ydR66v3DROIveOI8wXRC9.tKN0dbEX3DRntbhSE1S.wCzLhFZBMvrXFT9zSDTTFKg.FvsuMnavWPVLE_QuEP9GK1SOyKz.3Uy7QovZtPy.wCR2hdNvw1zfF13zZUR7F1g6tDdjMCS3h2SjwHJBhDz6FCGzBsWPt1qdJciuEYgzRozP3HSbt1z7QbYG7KR9QcS9FPiuMUYNMP70hImCMD2ItTeB_be7tKSnFDCuEsqvhPVZWM0BE27j3bzzgCzfIoRPw6KChKya3C3jRHzChbR7x1fB_be7tsE.FvsuMT0NKK2SMIWBMvrXFnGO.nl7xuL6xPCeU1fdRowNRiSdRKSv8C2g5Cx0Royj3K653USvFCfGMXSjMKrCwUTe5KRfRDpOw666QUSPMK2OQBwjt1z7QbYG7KR9QcS9FPiuMUYNMP70hImCMD2ItTeB_be7tKSnFDCuEsqvhPVZWM0BE27jIKrP4oeuMCmz8oUSwbm9IbTjwHyPIoebIC3L7KTTRDWN3KUjhKJ.EczjwHJBhkVXFUmzZTg7UbSnFDCuMUYNMP7bWM0BEO0nx1GOa1zfMor9wbkywnT9QUpfMHTB8U2O3C2L4orNFbS08DkZMvS93KJNIIyPMnfvRvQLZURfRcSP3o1S3Dgvhb2SMIWBWc27QbfzaPeMF6w.Fc6SwbSNhPV0E4WfikVGtnRI4owThKRPFvUOtPy.wCR2hdNvw1zfF13zZUR7F1g6tDdjMCS3h2SjwHJBhDz6FCGzBsWPt1qdJciuEYgz86egw8Gjhbz6RbVB5bwTtCmNxP6SwbSNhOQLM5NBMmGXUDN0ZCZ7F6w.FcCeWnENEkZ0E4WfUPTjFomn5ve7IUy0IUhLRbxyw6Y0Wh3zMvrORPN9563XRba2tDdjMCaNWPzjwHJBU12RFUmzZneXQDy.tciBE1aviOQbhMTwtoR5wDSG4K3NtCSnRKcSRbpahDpjE4JzwK2XtuAzZUR7FYZ.UCdjMCaNMoxLMMWfWcQXxsaSB1ZPKnNGIv6g8bRyRUegRXNPMb2.8CJudbwTRoRP8DMe3oryRoygR5N0wU29Q6STgCYfRvw.QoPzICR93UxLw8RdRbpfQbfTyPy2F6yP3DInQ6SG3DyLM8Jv8vT_3Dwe56zaIUyOQDhZ3vSfF6p6thmzwKpOtKm2_PTOFnq.FvsuMnaThb2SMIJ8hmy7Qbfz7CTTFKg.xODehnEeWPVLEZl.8bmfwof.ZURfRoyuFvs53vSCFDpXI8Yv8DQ.tKN0dbEX3DRntbhSE1S.wCzLhFZBMvrXFT9zSDTTFKg.FvsuMnavWPVLE_QuEP9GK1S.dDeNR6YPMUvCMUSB3Kzz35JN3Ce_8CyPevJfwvm4MbheQoNSIbTvIIRaRU2.wUfLZowuRvyXF6KzIC2jwORgFHr6IDS7wofLZvrZMDybRv6C3beyFvyLwXeaIDTCFDYe5o2PIUxuh1dSwbm9hDxGwhmGMPQXFUmzZnZatKSnFD6chTz.wCzLhImCMD9XxsAO7nlgJ1q.x2568UJdM6z5QXSbQvJ.RDmeyUJZIC2dQo62F6fPIvT7RBwd8bx5IKxOZve_3Uy5QK._wCm.RUe6I8R6RKpB86Tn4DmZID2utvB4I6fZFvev88J48KTjQ6fy4Dpy8UN03nPSMUY93n2fR8Lz3C7GtCN0ZCZ7J1yNQK6uUPS8MoxLMMJzwK2Xtn3SB1ZPHsE2tcHrt6rz86pz3XSdRK2vFCYOzPTXQDJbtbs4wPybMcVLM5NBMc9ntCN0ZCewtY2NQK6uhCyaMb7LE_ZfhcQ_Jn3zBY0LFDmjRbuOMbYN3CYe8h3zMvrORPN9563XRba2tDdjMCaNWPzjwHJBU12RFUmzZneXQDy.tciBE1aviOQbhMTwtDTvwUf44CgNtCSnRKcSRbpahDpjE4JzwK2XtuAzZUR7FYZ.UCdjMCaNMoxLMMWfWcQXxsaSB1ZPtnGNFvs53PyGRvEj3BlfhDz6FCGz.neXQDy.K16lMUYNMPzjwHJBhcQnx1GONuWPt1rRhohB3bJGhCxZQIJ7QbpCMbzBg6RfQ6maRUHnICmGQDffMdfjMbqBFCYOgCNWRDqjIl68AmYqpvz2RiT7RUYS86xL4UwTRDy0MKnChKya3C3jRHzChbR7x1fB_be7tsE.FvsuMT0NKK2SMIWBMvrXFnGO.nl7xuL6xPCeU1fOFC2GIIJaQo2B3KTB5nLXF6wbR1d_R6V.3K9bhImCMD9XJnfB_be7KcyQFvsuMnS.wCzLhM0uEP9GHuAO7nm8hKN6RCHuFUSC31gjM5NP31zv36lBdD0PtKSnFDCuWnS.wCzLUhJ3MvrXFnfB_be7t1q6xPCeiuqvhPYFQIy63Cr7QCmuznTOFDwP8cPSMUY93n2fR8Lz3C7GtCN0ZCZ7J1yNQK6uUPS8MoxLMMJzwK2Xtn3SB1ZPHsE2tcHrtCmNQUJSI8waRC9.tKyP7KTTRDWN3KUjhKJ.EczjwHJBhkVXFUmzZTg7UbSnFDCuMUYNMP7bWM0BEO0nx1GOa1zXMKNC8Dd63DYyMDxaIIxbwo2B3bmTebTv8KyXRvU_QopyMbybQdr7MU2XMbRu5vWn8KyXRvH_Q6SNFCYO3izbID2.RUfj712ZM6NGFcd73D2GMDpfIIY4R62LFDReeKpu8KfOJsDBtPy.wCR2hdNvw1zfF13zZUR7F1g6tDdjMCS3h2SjwHJBhDz6FCGzBsWPt1qdJciuEYgz3UJPIISahUTbwofP4cy7Qvx08DhBICJLRCxeQ8Rd3Cy_RKR7gCrv8KmPFKCnwUxy3KxzRHr2RbSf3vfP5veORbJbF6HO3DxCIbY.MBpGQbN5RvwT4KazFbJ08DkX8UJe8vz2QHYGIDpXQb2BevpghcSNQKk5hKYfwn2vM40BMvrXFnGS7CTTFKyFKYnYqGywFueP3XTrsOgnJuml_oATWu20V9OyWDwr1uSXpZeNJuTV1keGTVfiYUYj1s.FVlYvWbTK8dxUW9Ey10TK5Dz1VmxmWmkVim2036JmMdprQ6xGFke10uYRWUffATs_R0Q.Wuge3B9n3PzLwkeXyTN1pbwYJKBTRlZ.Wuge3B9n3PzLwkeHNTxVFbJjRTseY0l.Wuge3B9n3PzLwkeCu6RhsTR6R2kC3O9.FoAdReASwOmgIbJP_0mFHbJlUkI5Wu2PW2rDFtpupvREAla0_CzLWuAZM9MtpopopvJE3BNphOA4Hbwv.blXM6lZFuKXROxrV0m6sBN7hOA4Hbwv.blXM6lZtOcbibJOWCVjF5akqqVkJaAq.GWUMURuwkBfh1L.Mox23hmbRoq7RDL4NuWqquW5HKhyWGqDWulCqRVnWkECWk0SjuQarAVdJuPzJkqSqf8JF5ts7xl5vh_FJ9JbEWzJrrlRTxfpFLc_nx_QqaVorAWmvnJbP9kzjaSdjY6tRE9OyN6tnVxZ74RkcNTKnyHLMbqtlqqHWkg0it0_WO3_WGAq.GWtJkVCWucjJsV0WGQBMB3CwDwOFbxy0K2lVsWaJKCjF9psRc9ehZfDFUSiRV20.0x9RTraYmcLYAqmqGQlqtZkqsEk5kJ47CxfRPybFUiu3vJLhbRnFhJf3KgXRKYT7Cmzw1y2MK5u3cSb3CVL3XSahDRSQ1fPZDg7Rvx6tDU5wnSfMb7LRIrGhDmzQnf2ZDrgtKwd316_MoANRDz2hI2B3n29FDLz5KS0tKpLFvKuRDpvhbrzR4J98P22MC7zeCeatKN6wc66wCWNFCp.hIRzF12.FKEzebTutKz0Fn6TQcSB36VLF5e9hDen3bWzeUrBFnyzRUI5hCS73czLMBGBMKpntC2L5nezws3.FUo5hC0S3nzNQFaBMUY2tC2j_1ezFb3.FK4NhCT7MnzNQ8ZBMUR.tC2S5bA7FbrztD4Z8nSjRKqLMBJfhDNX3PfX_CV7wU7.woOZh6r9Q1z0RBEBQoYBt6pS4ceawU3.wDUyh6rTwvZLQIm9hoY.RcfSZCL7wUmOtokCR1S0wvZLQdgBQbGXwKY776mzQCq.wbsgh6xGQ1zCwI9BQbpat6wX76rPF1y03vbuQow9hvpuR4JawbEXwoTn76wfwnyuRUXuwC2OhvJzRd9BwDNawKmB_ne0IU3.QoUyh6RCRPznQXaBwvRTt6xy5ne6FbpftoMaRK36hvrNw4J0MKAXQDY676pawPydRbnu86rLhveLQ.J4MoR_t6SnZnegFDg.IP62FvqN36wBh8eBFURCtCTSdPe7RDqorqDMUcpzMqQqU.zQJAVDKcr90AWlKPRf1qDwWOWmqk3TqRgnqqAuraArva";if($_ts.lcd)$_ts.lcd();</script><script type="text/javascript" charset="utf-8" src="/LIWghRUPB4QK/oxYRCeR1jjgO.199cf1b.js" r='m'></script><script language="javascript" src="/_js/sudy-jquery-autoload.js" jquery-src="/_js/jquery-1.9.1.min.js" sudy-wp-context="" sudy-wp-siteId="68"></script>
<script language="javascript" src="/_js/jquery-migrate.min.js"></script>
<script language="javascript" src="/_js/jquery.sudy.wp.visitcount.js"></script>
<script type="text/javascript" src="/_js/_portletPlugs/sudyNavi/jquery.sudyNav.js"></script>
<script type="text/javascript" src="/_js/_portletPlugs/datepicker/js/jquery.datepicker.js"></script>
<script type="text/javascript" src="/_js/_portletPlugs/datepicker/js/datepicker_lang_HK.js"></script>
<link href="/_upload/tpl/04/02/1026/template1026/css/base.css" rel="stylesheet" type="text/css">
  <link href="/_upload/tpl/04/02/1026/template1026/css/11ml.css" rel="stylesheet" type="text/css">
  
  <script type="text/javascript" src="/_upload/tpl/04/02/1026/template1026/js/jQuery.meanMenu.js"></script>
  <script type="text/javascript" src="/_upload/tpl/04/02/1026/template1026/js/jquery.bxslider.min.js"></script>
  <link href="/_upload/tpl/04/02/1026/template1026/extends/extends.css" rel="stylesheet">
  <script type="text/javascript" src="/_upload/tpl/04/02/1026/template1026/extends/extends.js"></script>
  <!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
  <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
  <!--[if lt IE 9]>
  <script src="https://cdn.bootcss.com/html5shiv/3.7.3/html5shiv.min.js"></script>
  <script src="https://cdn.bootcss.com/respond.js/1.4.2/respond.min.js"></script>
<![endif]-->
  <link rel="stylesheet" href="/_upload/tpl/04/02/1026/template1026/css/owl.carousel.min.css">
  <link rel="stylesheet" href="/_upload/tpl/04/02/1026/template1026/css/owl.theme.default.css">
  <script src="/_upload/tpl/04/02/1026/template1026/js/owl.carousel.min.js"></script>
</head>

<body>
  <div class="top-se">
    <div class="container">
      <div class="t_se clearfix" frag="窗口0" portletmode="search">            <form method="post" action="/_web/_search/api/search/new.rst?locale=zh_CN&request_locale=zh_CN&_p=YXM9NjgmdD0xMDI2JmQ9NDE2MCZwPTEmbT1TTiY_" target="_blank">
                <input name="keyword" type="text" class="se_txt">
                <input name="submit" type="submit" class="se_sub" value="搜 索">
                <a class="se_close" href="javascript:;"></a>
            </form>
        </div>
    </div>
  </div>
  <div class="top">
    <div class="container">
      <div class="row">
        <div class="col xs-12 sm-12 md-4 lg-5">
          <div class="logo">
            <div class="logo_table">
              <div class="logo_cell"> <a href="/main.htm"><img src="/_upload/tpl/04/02/1026/template1026/images/logo.png"></a> </div>
            </div>
          </div>
        </div>
        <div class="col xs-12 sm-12 md-8 lg-7">
          <div class="top_r">
            <div class="top_r_a"><a class="a1" href="http://www.nuaa.edu.cn/" target="_blank">南航主页</a><a class="a2" href="javascript:void(0)">ENGLISH</a><a class="a3" href="javascript:;"></a></div>
            <div class="top_nav hidden-xs" frag="窗口1" portletmode="simpleSudyNavi" configs="{'c1':'1','c7':'2','c4':'_self','c3':'30','c6':'0','c8':'2','c9':'0','c2':'1','c5':'2'}" contents="{'c2':'0', 'c1':'/学院概况,/师资队伍,/科学研究,/教学工作,/学生工作,/党群工作,/工会与校友工作,/常用下载'}">
              
              <ul class="clearfix pc_menuCon">
                <li class="active"><a href="/main.htm">首页</a></li>
                
                <li><a href="/1954/list.htm" target="_self">学院概况</a>
                  
                  <ul>
                    
                    <li><a href="/xyjj/list.htm" target="_self">学院简介</a></li>
                    
                    <li><a href="/1965/list.htm" target="_self">Introduction of the College</a></li>
                    
                    <li><a href="/1966/list.htm" target="_self">领导班子</a></li>
                    
                    <li><a href="/1967/list.htm" target="_self">学院概况</a></li>
                    
                    <li><a href="/1968/list.htm" target="_self">历史沿革</a></li>
                    
                  </ul>
                  
                </li>
                
                <li><a href="/1956/list.htm" target="_self">师资队伍</a>
                  
                  <ul>
                    
                    <li><a href="http://graduate.nuaa.edu.cn/gmis5/dsfc/dsfc_yx_new" target="_self">硕导博导信息</a></li>
                    
                    <li><a href="/7382/list.htm" target="_self">学院教师</a></li>
                    
                    <li><a href="/7383/list.htm" target="_self">办公室人员</a></li>
                    
                  </ul>
                  
                </li>
                
                <li><a href="/1957/list.htm" target="_self">科学研究</a>
                  
                  <ul>
                    
                    <li><a href="/1975/list.htm" target="_self">科研动态</a></li>
                    
                    <li><a href="/1976/list.htm" target="_self">表格下载</a></li>
                    
                  </ul>
                  
                </li>
                
                <li><a href="/1958/list.htm" target="_self">教学工作</a>
                  
                  <ul>
                    
                    <li><a href="/zyjs/list.htm" target="_self">专业介绍</a></li>
                    
                    <li><a href="/1977/list.htm" target="_self">教学动态</a></li>
                    
                    <li><a href="/jxzd/list.htm" target="_self">教学制度</a></li>
                    
                    <li><a href="/1978/list.htm" target="_self">本科生教学</a></li>
                    
                    <li><a href="/1979/list.htm" target="_self">研究生教学</a></li>
                    
                    <li><a href="/1980/list.htm" target="_self">教学改革与成果</a></li>
                    
                    <li><a href="/1981/list.htm" target="_self">资料下载</a></li>
                    
                  </ul>
                  
                </li>
                
                <li><a href="/1959/list.htm" target="_self">学生工作</a>
                  
                  <ul>
                    
                    <li><a href="/1982/list.htm" target="_self">学生活动</a></li>
                    
                    <li><a href="/1993/list.htm" target="_self">招生信息</a></li>
                    
                    <li><a href="/1983/list.htm" target="_self">研究生会</a></li>
                    
                    <li><a href="/1994/list.htm" target="_self">就业信息</a></li>
                    
                    <li><a href="/1984/list.htm" target="_self">共青团</a></li>
                    
                    <li><a href="/1985/list.htm" target="_self">学生党建</a></li>
                    
                    <li><a href="/1986/list.htm" target="_self">获奖情况</a></li>
                    
                    <li><a href="/1987/list.htm" target="_self">通知公告</a></li>
                    
                  </ul>
                  
                </li>
                
                <li><a href="/1960/list.htm" target="_self">党群工作</a>
                  
                  <ul>
                    
                    <li><a href="/qgdjgzybzb/list.htm" target="_self">全国党建工作样板支部</a></li>
                    
                    <li><a href="/1988/list.htm" target="_self">党员管理</a></li>
                    
                    <li><a href="/1989/list.htm" target="_self">主题活动</a></li>
                    
                    <li><a href="/1990/list.htm" target="_self">组织建设</a></li>
                    
                    <li><a href="/1991/list.htm" target="_self">党员发展</a></li>
                    
                    <li><a href="/1992/list.htm" target="_self">党校工作</a></li>
                    
                  </ul>
                  
                </li>
                
                <li><a href="/1962/list.htm" target="_self">工会与校友工作</a>
                  
                  <ul>
                    
                    <li><a href="/7384/list.htm" target="_self">教师工作指南</a></li>
                    
                    <li><a href="/tmhxqgzshzn/list.htm" target="_self">天目湖校区工作生活指南</a></li>
                    
                    <li><a href="/16162/list.htm" target="_self">教职工之家</a></li>
                    
                    <li><a href="/16163/list.htm" target="_self">校友返校指南</a></li>
                    
                    <li><a href="/16164/list.htm" target="_self">基金工作指南</a></li>
                    
                  </ul>
                  
                </li>
                
                <li><a href="/cyxz/list.htm" target="_self">常用下载</a>
                  
                  <ul>
                    
                    <li><a href="/aqgz/list.htm" target="_self"> 安全工作</a></li>
                    
                    <li><a href="/rsxz/list.htm" target="_self">人事工作</a></li>
                    
                    <li><a href="/bkjx/list.htm" target="_self">本科生培养</a></li>
                    
                    <li><a href="/yjspy/list.htm" target="_self">研究生培养</a></li>
                    
                    <li><a href="/kxyj/list.htm" target="_self">科学研究</a></li>
                    
                    <li><a href="/xsgz/list.htm" target="_self">学生工作</a></li>
                    
                    <li><a href="/djqt/list.htm" target="_self">党建群团</a></li>
                    
                    <li><a href="/xygz/list.htm" target="_self">校友工作</a></li>
                    
                    <li><a href="/xyLogo/list.htm" target="_self">学院Logo</a></li>
                    
                  </ul>
                  
                </li>
                
              </ul>
              
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div class="phone_menu_t visible-xs">
    <div class="phone_menu" style="display: none" frag="窗口01" portletmode="simpleSudyNavi" configs="{'c1':'1','c7':'2','c4':'_self','c3':'30','c6':'0','c8':'2','c9':'0','c2':'1','c5':'2'}" contents="{'c2':'0', 'c1':'/学院概况,/师资队伍,/科学研究,/教学工作,/学生工作,/党群工作,/工会与校友工作,/常用下载'}">
      
      <ul>
        
        <li><a href="/1954/list.htm" target="_self">学院概况</a>
          
          <ul>
            
            <li><a href="/xyjj/list.htm" target="_self">学院简介</a></li>
            
            <li><a href="/1965/list.htm" target="_self">Introduction of the College</a></li>
            
            <li><a href="/1966/list.htm" target="_self">领导班子</a></li>
            
            <li><a href="/1967/list.htm" target="_self">学院概况</a></li>
            
            <li><a href="/1968/list.htm" target="_self">历史沿革</a></li>
            
          </ul>
          
        </li>
        
        <li><a href="/1956/list.htm" target="_self">师资队伍</a>
          
          <ul>
            
            <li><a href="http://graduate.nuaa.edu.cn/gmis5/dsfc/dsfc_yx_new" target="_self">硕导博导信息</a></li>
            
            <li><a href="/7382/list.htm" target="_self">学院教师</a></li>
            
            <li><a href="/7383/list.htm" target="_self">办公室人员</a></li>
            
          </ul>
          
        </li>
        
        <li><a href="/1957/list.htm" target="_self">科学研究</a>
          
          <ul>
            
            <li><a href="/1975/list.htm" target="_self">科研动态</a></li>
            
            <li><a href="/1976/list.htm" target="_self">表格下载</a></li>
            
          </ul>
          
        </li>
        
        <li><a href="/1958/list.htm" target="_self">教学工作</a>
          
          <ul>
            
            <li><a href="/zyjs/list.htm" target="_self">专业介绍</a></li>
            
            <li><a href="/1977/list.htm" target="_self">教学动态</a></li>
            
            <li><a href="/jxzd/list.htm" target="_self">教学制度</a></li>
            
            <li><a href="/1978/list.htm" target="_self">本科生教学</a></li>
            
            <li><a href="/1979/list.htm" target="_self">研究生教学</a></li>
            
            <li><a href="/1980/list.htm" target="_self">教学改革与成果</a></li>
            
            <li><a href="/1981/list.htm" target="_self">资料下载</a></li>
            
          </ul>
          
        </li>
        
        <li><a href="/1959/list.htm" target="_self">学生工作</a>
          
          <ul>
            
            <li><a href="/1982/list.htm" target="_self">学生活动</a></li>
            
            <li><a href="/1993/list.htm" target="_self">招生信息</a></li>
            
            <li><a href="/1983/list.htm" target="_self">研究生会</a></li>
            
            <li><a href="/1994/list.htm" target="_self">就业信息</a></li>
            
            <li><a href="/1984/list.htm" target="_self">共青团</a></li>
            
            <li><a href="/1985/list.htm" target="_self">学生党建</a></li>
            
            <li><a href="/1986/list.htm" target="_self">获奖情况</a></li>
            
            <li><a href="/1987/list.htm" target="_self">通知公告</a></li>
            
          </ul>
          
        </li>
        
        <li><a href="/1960/list.htm" target="_self">党群工作</a>
          
          <ul>
            
            <li><a href="/qgdjgzybzb/list.htm" target="_self">全国党建工作样板支部</a></li>
            
            <li><a href="/1988/list.htm" target="_self">党员管理</a></li>
            
            <li><a href="/1989/list.htm" target="_self">主题活动</a></li>
            
            <li><a href="/1990/list.htm" target="_self">组织建设</a></li>
            
            <li><a href="/1991/list.htm" target="_self">党员发展</a></li>
            
            <li><a href="/1992/list.htm" target="_self">党校工作</a></li>
            
          </ul>
          
        </li>
        
        <li><a href="/1962/list.htm" target="_self">工会与校友工作</a>
          
          <ul>
            
            <li><a href="/7384/list.htm" target="_self">教师工作指南</a></li>
            
            <li><a href="/tmhxqgzshzn/list.htm" target="_self">天目湖校区工作生活指南</a></li>
            
            <li><a href="/16162/list.htm" target="_self">教职工之家</a></li>
            
            <li><a href="/16163/list.htm" target="_self">校友返校指南</a></li>
            
            <li><a href="/16164/list.htm" target="_self">基金工作指南</a></li>
            
          </ul>
          
        </li>
        
        <li><a href="/cyxz/list.htm" target="_self">常用下载</a>
          
          <ul>
            
            <li><a href="/aqgz/list.htm" target="_self"> 安全工作</a></li>
            
            <li><a href="/rsxz/list.htm" target="_self">人事工作</a></li>
            
            <li><a href="/bkjx/list.htm" target="_self">本科生培养</a></li>
            
            <li><a href="/yjspy/list.htm" target="_self">研究生培养</a></li>
            
            <li><a href="/kxyj/list.htm" target="_self">科学研究</a></li>
            
            <li><a href="/xsgz/list.htm" target="_self">学生工作</a></li>
            
            <li><a href="/djqt/list.htm" target="_self">党建群团</a></li>
            
            <li><a href="/xygz/list.htm" target="_self">校友工作</a></li>
            
            <li><a href="/xyLogo/list.htm" target="_self">学院Logo</a></li>
            
          </ul>
          
        </li>
        
      </ul>
      
    </div>
  </div>
  <script type="text/javascript">
  $(function() {
    $('.pc_menuCon li').hover(function() {
      $('ul', this).slideDown(200);
    }, function() {
      $('ul', this).slideUp(100);
    });
    $('.a3').click(function() {
      $('.top-se').stop().slideToggle();
    });
    $('.se_close').click(function() {
      $('.top-se').stop().slideUp(200);;
    });
  });
  jQuery('.phone_menu').meanmenu();
  $(window).scroll(function() {
    var scrollHeight = $(document).scrollTop();
    if (scrollHeight > 340) {
      $('.phone_menu_t').addClass('lighted-fixed');
    } else {
      $('.phone_menu_t').removeClass('lighted-fixed');
    }
  });
  </script>
  <div class="banner hidden-xs" frag="窗口2" portletmode="simpleNews" configs="{'c8':'1','c42':'320','c25':'320','c30':'0','c29':'1','c22':'0','c23':'1','c34':'300','c20':'0','c31':'0','c16':'1','c3':'6','c2':'图片路径','c27':'480','c43':'0','c17':'0','c5':'_blank','c24':'240','c32':'','c26':'1','c37':'1','c28':'640','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'50','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'15','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'1','c13':'200','c41':'240'}" contents="{'c2':'1', 'c1':'/PC首页幻灯'}"><div id="wp_news_w2"> 

    <ul class="bxslider">
      
      <li><img src="/_upload/article/images/45/45/a46e74f24558b84ce2c9d7e02af5/24004537-f4bf-40aa-83b1-0f9e7afcb15e.png" /></li>
      
      <li><img src="/_upload/article/images/e4/10/fc65960e495594ed53c0251205cd/a8f964df-b3a8-4a50-b217-be146078f351.png" /></li>
      
      <li><img src="/_upload/article/images/8e/5b/f9e805164267b723013414e8d7ca/e2397f85-7e06-4124-97ba-0ea7a35e1a05.png" /></li>
      
      <li><img src="/_upload/article/images/ee/3f/c23c09c74990b6c6bf72de29cc9c/535d31e3-7de2-485f-993a-811042e8da7f.png" /></li>
      
      <li><img src="/_upload/article/images/d6/28/c0272c7b4ef98fdcf3de3f3e8990/8d7ab9a8-25ac-4bb1-b7d9-edf26d58cc92.png" /></li>
      
      <li><img src="/_upload/article/images/20/35/0e7ca14f46ea844bd475df8ebf37/48f77e80-bf7c-474d-a504-5ea301c1b032.png" /></li>
      
    </ul>
  </div> 
</div>
  <script type="text/javascript">
  $('.bxslider').bxSlider({
    mode: 'fade',
    auto: true,
    autoControls: true,
    autoHover: true,
    controls: false
  });
  </script>
  <!-- pc-banner -->
  <div class="banner visible-xs" frag="窗口02" portletmode="simpleNews" configs="{'c8':'1','c42':'320','c25':'320','c30':'0','c29':'1','c22':'0','c23':'1','c34':'300','c20':'0','c31':'0','c16':'1','c3':'6','c2':'图片路径','c27':'480','c43':'0','c17':'0','c5':'_blank','c24':'240','c32':'','c26':'1','c37':'1','c28':'640','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'50','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'15','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'1','c13':'200','c41':'240'}" contents="{'c2':'1', 'c1':'/手机首页幻灯'}"><div id="wp_news_w02"> 

    <ul class="bxslider02">
      
      <li><img src="/_upload/article/images/20/35/0e7ca14f46ea844bd475df8ebf37/48f77e80-bf7c-474d-a504-5ea301c1b032.png" /></li>
      
      <li><img src="/_upload/article/images/74/5b/9dda4c0446fd93b0326d50a71b45/91732d13-5f6e-4d78-96ef-7b4c3c2f565b.png" /></li>
      
      <li><img src="/_upload/article/images/21/54/835e7c8d47a3ac33755a44a53147/a96926c2-fd26-4c54-b1e5-cace00df4e3c.jpg" /></li>
      
      <li><img src="/_upload/article/images/70/99/35cbc4664cb48bdf524a41fac7a3/669f44d5-c735-4e4b-a6a4-d306f826e255.png" /></li>
      
      <li><img src="/_upload/article/images/08/98/6c53849b4957a1216a245e5c91d4/8a16e31c-8887-4390-ae2c-807a4c763fc0.png" /></li>
      
      <li><img src="/_upload/article/images/2e/c2/c5683ded49faafa8ca32ba4bb670/5a4d4182-bc67-4584-904d-a7e6ced069f6.jpg" /></li>
      
    </ul>
  </div> 
</div>
  <script type="text/javascript">
  $('.bxslider02').bxSlider({
    mode: 'horizontal',
    auto: true,
    autoControls: true,
    autoHover: true,
    controls: false
  });
  </script>
  <!-- phone-banner -->
  <div class="type">
    <div class="container">
      <div class="row">
        <div class="col xs-4 sm-4 md-2 lg-2">
          <div class="type-a type-p"> <a href="/xyjj/list.htm">
              <p>学院简介</p>
            </a> <span class="an"></span> </div>
        </div>
        <div class="col xs-4 sm-4 md-2 lg-2">
          <div class="type-b type-p"> <a href="/10850/list.htm">
              <p>本科生</p>
            </a> <span class="an"></span> </div>
        </div>
        <div class="col xs-4 sm-4 md-2 lg-2">
          <div class="type-c type-p"> <a href="/10851/list.htm">
              <p>研究生</p>
            </a> <span class="an"></span> </div>
        </div>
        <div class="col xs-4 sm-4 md-2 lg-2 col-d">
          <div class="type-d type-p"> <a href="http://bsh.nuaa.edu.cn/">
              <p>博士后</p>
            </a> <span class="an"></span> </div>
        </div>
        <div class="col xs-4 sm-4 md-2 lg-2 col-e">
          <div class="type-e type-p"> <a href="_s68/2017/0928/c11649a148148/page.psp">
              <p>校友与基金工作</p>
            </a> <span class="an"></span> </div>
        </div>
        <div class="col xs-4 sm-4 md-2 lg-2 col-f">
          <div class="type-f type-p"> <a href="/2017/1115/c11649a177680/page.htm" target="_blank">
              <p>人才引进</p>
            </a> <span class="an"></span> </div>
        </div>
        <!--<div class="col xs-6 sm-4 md-2 lg-2">
                <div class="type-a type-p"> <a href="javascript:void(0)">
                    <p>UG</p>
                    <p><i>本科</i></p>
                    </a> <span class="an"></span> </div>
            </div>
            <div class="col xs-6 sm-4 md-2 lg-2">
                <div class="type-b type-p"> <a href="javascript:void(0)">
                    <p>PG</p>
                    <p><i>研究生</i></p>
                    </a> <span class="an"></span> </div>
            </div>
            <div class="col xs-6 sm-4 md-2 lg-2">
                <div class="type-c type-p"> <a href="javascript:void(0)">
                    <p>MBA</p>
                    <p><i>工商管理硕士</i></p>
                    </a> <span class="an"></span> </div>
            </div>
            <div class="col xs-6 sm-4 md-2 lg-2 col-d">
                <div class="type-d type-p"> <a href="javascript:void(0)">
                    <p>M.Eng </p>
                    <p><i>工程硕士</i></p>
                    </a> <span class="an"></span> </div>
            </div>
            <div class="col xs-6 sm-4 md-2 lg-2 col-e">
                <div class="type-e type-p"> <a href="javascript:void(0)">
                    <p>MPAcc</p>
                    <p><i>专业会计硕士</i></p>
                    </a> <span class="an"></span> </div>
            </div>
            <div class="col xs-6 sm-4 md-2 lg-2 col-f">
                <div class="type-f type-p"> <a href="javascript:void(0)">
                    <p>EDP</p>
                    <p><i>高级经理人发展课程</i></p>
                    </a> <span class="an"></span> </div>
            </div>-->
      </div>
    </div>
  </div>
  <div class="section">
    <div class="container container-pd">
      <div class="row">
        <div class="col xs-12 sm-8 md-8 lg-8">
          <div class="section-l">
            <div class="in_title in_title_b"><span>热点新闻</span><a href="/10846/list.htm">更多>></a></div>
            <div class="section-l-con">
              <div class="row">
                <div class="col xs-12 sm-6 md-6 lg-6">
                  <div class="section-column">
                    <div class="list-img" frag="窗口3" portletmode="simpleNews" configs="{'c8':'1','c42':'320','c25':'320','c30':'0','c29':'1','c22':'0','c23':'1','c34':'300','c20':'0','c31':'0','c16':'1','c3':'7','c2':'图片路径,标题','c27':'480','c43':'0','c17':'0','c5':'_blank','c24':'240','c32':'','c26':'1','c37':'1','c28':'640','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'50','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'30','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'1','c13':'200','c41':'240'}" contents="{'c2':'1', 'c1':'/热点图片'}"><div id="wp_news_w3"> 

                      <ul class="bxslider03 bxslider03-pic">
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/d2/7a/35622189448888f4317d9d27ab96/02a7ec1b-6984-408f-b218-10527500b393.png);"></div>
                          <p><a href='/2024/0531/c11624a341535/page.htm' target='_blank' title='计算机学院教师在教师教学创新大赛中取得突破性成绩'>计算机学院教师在教师教学创新大赛中取得突破性成绩</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/41/1c/4fb160ae44e6856e82b64c6b886e/d0d524fe-b512-47e8-8283-d2a74bb8d7ac.jpg);"></div>
                          <p><a href='/2024/0523/c11624a340746/page.htm' target='_blank' title='中国电科第三十研究所来访我院'>中国电科第三十研究所来访我院</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/f9/d5/aac92e124b189ab73e2cd1288995/78c6b1c1-93ae-49bc-ab9f-a8ada182b90c.jpg);"></div>
                          <p><a href='/2024/0521/c11624a340419/page.htm' target='_blank' title='计算机学院、数学学院、经管学院、人文学院党委理论学习中心组举行联合学习会'>计算机学院、数学学院、经管学院、人文学院党委理论学习中心组举...</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/8c/c4/fc491e1f4c94bb4d880f013b29e2/4938d553-203d-4f84-ad98-cb1e5764760a.jpg);"></div>
                          <p><a href='/2024/0507/c11624a338414/page.htm' target='_blank' title='我院学生获得ICPC全国邀请赛（武汉）季军'>我院学生获得ICPC全国邀请赛（武汉）季军</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/f4/a5/b5a235914085bda09a2d386cfc9d/c0677153-05d1-41ca-8d2c-4d100b66ec30.jpg);"></div>
                          <p><a href='/2024/0502/c11624a338079/page.htm' target='_blank' title='计算机学院召开全体教职工大会暨本科教育教学审核评估培训动员会'>计算机学院召开全体教职工大会暨本科教育教学审核评估培训动员会</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/5c/cb/e617dc0f4743a643cd9b287bfbcf/3c348944-5a16-45ce-abb2-4ce56828af8b.jpg);"></div>
                          <p><a href='/2024/0430/c11624a338021/page.htm' target='_blank' title='南京航空航天大学第八届信息安全技能竞赛成功举办'>南京航空航天大学第八届信息安全技能竞赛成功举办</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/70/9d/87f04c5f4f468004ba9de5e0bc2d/fd1322a9-545c-4507-a39a-714ac6106ff6.jpg);"></div>
                          <p><a href='/2024/0429/c11624a337637/page.htm' target='_blank' title='中北大学软件学院来访计算机学院'>中北大学软件学院来访计算机学院</a></p>
                        </li>
                        
                      </ul>
                    </div> 
</div>
                  </div>
                </div>
                <div class="col xs-12 sm-6 md-6 lg-6">
                  <div class="section-column">
                    <div class="list-ul" frag="窗口4" portletmode="simpleNews" configs="{'c8':'1','c42':'320','c25':'320','c30':'0','c29':'1','c22':'0','c23':'1','c34':'300','c20':'0','c31':'0','c16':'1','c3':'7','c2':'标题,发布时间','c27':'480','c43':'0','c17':'0','c5':'_blank','c24':'240','c32':'','c26':'1','c37':'1','c28':'640','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'50','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'20','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'1','c13':'200','c41':'240'}" contents="{'c2':'0', 'c1':'/热点新闻'}"><div id="wp_news_w4"> 

                      <ul class="list-ul-li">
                        
                        <li><a href='/2024/0531/c10846a341552/page.htm' target='_blank' title='计算机学院举办“乐享教学”沙龙活动'>计算机学院举办“乐享教学”沙龙活动</a><i>05-31</i></li>
                        
                        <li><a href='/2024/0531/c10846a341535/page.htm' target='_blank' title='计算机学院教师在教师教学创新大赛中取得突破性成绩'>计算机学院教师在教师教学创新大赛中取得突...</a><i>05-31</i></li>
                        
                        <li><a href='/2024/0531/c10846a341491/page.htm' target='_blank' title='我院三名学生入选腾讯犀牛鸟精英人才计划'>我院三名学生入选腾讯犀牛鸟精英人才计划</a><i>05-31</i></li>
                        
                        <li><a href='/2024/0529/c10846a341300/page.htm' target='_blank' title='计算机学院举办“导学引领，乐享运动”2024年研究生导学运动会'>计算机学院举办“导学引领，乐享运动”2024...</a><i>05-29</i></li>
                        
                        <li><a href='/2024/0528/c10846a341072/page.htm' target='_blank' title='计算机学院举办“生涯筑梦 计遇未来” 职业生涯嘉年华活动'>计算机学院举办“生涯筑梦 计遇未来” 职业...</a><i>05-28</i></li>
                        
                        <li><a href='/2024/0523/c10846a340746/page.htm' target='_blank' title='中国电科第三十研究所来访我院'>中国电科第三十研究所来访我院</a><i>05-23</i></li>
                        
                        <li><a href='/2024/0521/c10846a340419/page.htm' target='_blank' title='计算机学院、数学学院、经管学院、人文学院党委理论学习中心组举行联合学习会'>计算机学院、数学学院、经管学院、人文学院...</a><i>05-21</i></li>
                        
                      </ul>
                    </div> 
</div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="col xs-12 sm-4 md-4 lg-4">
          <div class="section-column">
            <div class="in_title"><span>公示</span><a href="/10847/list.htm">更多>></a></div>
            <div class="list-ul list-ul-b" frag="窗口5" portletmode="simpleNews" configs="{'c8':'1','c42':'320','c25':'320','c30':'0','c29':'1','c22':'0','c23':'1','c34':'300','c20':'0','c31':'0','c16':'1','c3':'6','c2':'标题,发布时间','c27':'480','c43':'0','c17':'0','c5':'_blank','c24':'240','c32':'','c26':'1','c37':'1','c28':'640','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'50','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'20','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'1','c13':'200','c41':'240'}" contents="{'c2':'0', 'c1':'/公示'}"><div id="wp_news_w5"> 

              <ul class="list-ul-li list-ul-li-b">
                
                <li><a href='/2024/0524/c10847a340864/page.htm' target='_blank' title='发展党员公示'>发展党员公示</a><i>05-24</i></li>
                
                <li><a href='/2024/0522/c10847a340578/page.htm' target='_blank' title='2024级研究生新生党员组织关系转入相关说明'>2024级研究生新生党员组织关系转入相关说明</a><i>05-22</i></li>
                
                <li><a href='/2024/0507/c10847a338384/page.htm' target='_blank' title='计算机学院2024年正高专业技术职务基层推荐评审结果公示'>计算机学院2024年正高专业技术职务基层推荐...</a><i>05-07</i></li>
                
                <li><a href='/2024/0514/c10847a339730/page.htm' target='_blank' title='关于南京航空航天大学本科教育教学审核评估自评报告的公示'>关于南京航空航天大学本科教育教学审核评估...</a><i>04-25</i></li>
                
                <li><a href='/2024/0109/c10847a329258/page.htm' target='_blank' title='关于自编讲义编写人员政治审查公示'>关于自编讲义编写人员政治审查公示</a><i>01-03</i></li>
                
                <li><a href='/2023/1227/c10847a328359/page.htm' target='_blank' title='预备党员转正公示'>预备党员转正公示</a><i>12-27</i></li>
                
              </ul>
            </div> 
</div>
            <div class="zt">
              <div class="zt-title">
                <p>专</p>
                <p>题</p>
              </div>
              <div class="zt-con" frag="窗口18" portletmode="simpleNews" configs="{'c8':'1','c42':'320','c25':'320','c30':'0','c29':'1','c22':'0','c23':'1','c34':'300','c20':'0','c31':'0','c16':'1','c3':'1','c2':'图片路径','c27':'480','c43':'0','c17':'0','c5':'_blank','c24':'240','c32':'','c26':'1','c37':'1','c28':'640','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'50','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'15','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'1','c13':'200','c41':'240'}" contents="{'c2':'0', 'c1':'/专题'}"><div id="wp_news_w18"> 

                
                <a href="http://47.100.138.90/SpaCCS2020/" target="_blank"><img src="/_upload/article/images/d7/71/813ceda34f738b57ac391a1734b7/6912482e-b750-4e38-bb29-dac7727f56d3.jpg"></img></a>
                
              </div> 
</div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <script type="text/javascript">
  $('.bxslider03').bxSlider({
    mode: 'horizontal',
    auto: true,
    autoControls: true,
    autoHover: true,
    pager: false,
    pause: 3000
  });
  </script>
  <div class="in-news">
    <div class="container container-pd">
      <div class="in_title"><span>通知公告</span><a href="http://cs.nuaa.edu.cn/tzgg/list.htm">更多>></a></div>
      <div class="row">
        <div class="col xs-12 sm-6 md-4 lg-4">
          <div class="main-c-item">
            <div class="main-c-t"> <span>党委行政</span><a href="/dwxz/list.htm" target="_blank">MORE</a> </div>
            <div class="main-c-list" frag="窗口7" portletmode="simpleNews" configs="{'c8':'0','c42':'320','c25':'320','c30':'0','c29':'1','c22':'0','c23':'1','c34':'300','c20':'0','c31':'0','c16':'1','c3':'5','c2':'标题,发布时间','c27':'480','c43':'0','c17':'0','c5':'_blank','c24':'240','c32':'','c26':'1','c37':'1','c28':'640','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'50','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'20','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'1','c13':'200','c41':'240'}" contents="{'c2':'0', 'c1':'/通知公告/党委行政'}"><div id="wp_news_w7"> 

              <ul>
                
                <li><a href="/2024/0524/c10853a340864/page.htm" target="_blank"><i>></i>发展党员公示</a><span>05.24</span></li>
                
                <li><a href="/2024/0522/c10853a340578/page.htm" target="_blank"><i>></i>2024级研究生新生党员组织关系转入相关说明</a><span>05.22</span></li>
                
                <li><a href="/2023/1227/c10853a328359/page.htm" target="_blank"><i>></i>预备党员转正公示</a><span>12.27</span></li>
                
                <li><a href="/2023/1215/c10853a327474/page.htm" target="_blank"><i>></i>预备党员转正公示</a><span>12.15</span></li>
                
                <li><a href="/2023/1128/c10853a325769/page.htm" target="_blank"><i>></i>发展党员公示</a><span>11.28</span></li>
                
              </ul>
            </div> 
</div>
          </div>
        </div>
        <div class="col xs-12 sm-6 md-4 lg-4">
          <div class="main-c-item">
            <div class="main-c-t"> <span>人事</span><a href="/10848/list.htm" target="_blank">MORE</a> </div>
            <div class="main-c-list" frag="窗口8" portletmode="simpleNews" configs="{'c8':'0','c42':'320','c25':'320','c30':'0','c29':'1','c22':'0','c23':'1','c34':'300','c20':'0','c31':'0','c16':'1','c3':'5','c2':'标题,发布时间','c27':'480','c43':'0','c17':'0','c5':'_blank','c24':'240','c32':'','c26':'1','c37':'1','c28':'640','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'50','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'20','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'1','c13':'200','c41':'240'}" contents="{'c2':'0', 'c1':'/通知公告/人事'}"><div id="wp_news_w8"> 

              <ul>
                
                <li><a href="/2023/1214/c10848a327388/page.htm" target="_blank"><i>></i>2023年度教职工考核工作安排</a><span>12.14</span></li>
                
                <li><a href="/2023/0302/c10848a303998/page.htm" target="_blank"><i>></i>计算机科学与技术学院/人工智能学院/软件学...</a><span>03.02</span></li>
                
                <li><a href="/2022/1210/c10848a300564/page.htm" target="_blank"><i>></i>2022年度教职工考核工作安排</a><span>12.10</span></li>
                
                <li><a href="/2022/0922/c10848a293100/page.htm" target="_blank"><i>></i>2022年拟申报初级专业技术职务人员公示</a><span>09.22</span></li>
                
                <li><a href="/2021/1214/c10848a272172/page.htm" target="_blank"><i>></i>【年终考核】2021年度计算机科学与技术学院...</a><span>12.14</span></li>
                
              </ul>
            </div> 
</div>
          </div>
        </div>
        <div class="col xs-12 sm-6 md-4 lg-4">
          <div class="main-c-item">
            <div class="main-c-t"> <span>学科科研</span><a href="/10849/list.htm" target="_blank">MORE</a> </div>
            <div class="main-c-list" frag="窗口9" portletmode="simpleNews" configs="{'c8':'0','c42':'320','c25':'320','c30':'0','c29':'1','c22':'0','c23':'1','c34':'300','c20':'0','c31':'0','c16':'1','c3':'5','c2':'标题,发布时间','c27':'480','c43':'0','c17':'0','c5':'_blank','c24':'240','c32':'','c26':'1','c37':'1','c28':'640','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'50','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'20','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'1','c13':'200','c41':'240'}" contents="{'c2':'0', 'c1':'/通知公告/学科科研'}"><div id="wp_news_w9"> 

              <ul>
                
                <li><a href="/2024/0420/c10849a336622/page.htm" target="_blank"><i>></i>关于组织模式分析与机器智能工业和信息化部...</a><span>04.20</span></li>
                
                <li><a href="/2024/0412/c10849a335982/page.htm" target="_blank"><i>></i>关于组织高安全系统的软件开发与验证技术工...</a><span>04.12</span></li>
                
                <li><a href="/2024/0410/c10849a335810/page.htm" target="_blank"><i>></i>关于组织脑机智能技术教育部重点实验室2024...</a><span>04.10</span></li>
                
                <li><a href="/2024/0102/c10849a328638/page.htm" target="_blank"><i>></i>关于开展第十八届中国青年科技奖候选人提名...</a><span>01.02</span></li>
                
                <li><a href="/2023/1227/c10849a328354/page.htm" target="_blank"><i>></i>关于开展第二十届中国青年女科学家奖和第九...</a><span>12.27</span></li>
                
              </ul>
            </div> 
</div>
          </div>
        </div>
        <div class="col xs-12 sm-6 md-4 lg-4">
          <div class="main-c-item">
            <div class="main-c-t"> <span>本科生培养</span><a href="/10850/list.htm" target="_blank">MORE</a> </div>
            <div class="main-c-list" frag="窗口10" portletmode="simpleNews" configs="{'c8':'0','c42':'320','c25':'320','c30':'0','c29':'1','c22':'0','c23':'1','c34':'300','c20':'0','c31':'0','c16':'1','c3':'5','c2':'标题,发布时间','c27':'480','c43':'0','c17':'0','c5':'_blank','c24':'240','c32':'','c26':'1','c37':'1','c28':'640','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'50','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'20','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'1','c13':'200','c41':'240'}" contents="{'c2':'0', 'c1':'/通知公告/本科生培养'}"><div id="wp_news_w10"> 

              <ul>
                
                <li><a href="/2024/0524/c10850a340772/page.htm" target="_blank"><i>></i>南京航空航天大学计算机学院 关于征集集中...</a><span>05.24</span></li>
                
                <li><a href="/2024/0524/c10850a340765/page.htm" target="_blank"><i>></i>【本科生实习】关于开展2024年度本科生实习...</a><span>05.24</span></li>
                
                <li><a href="/2024/0511/c10850a339344/page.htm" target="_blank"><i>></i>【毕业设计】2024届本科生毕业设计（论文）...</a><span>05.11</span></li>
                
                <li><a href="/2024/0511/c10850a339341/page.htm" target="_blank"><i>></i>【毕业设计】关于做好2024届本科生毕业设计...</a><span>05.11</span></li>
                
                <li><a href="/2024/0429/c10850a337738/page.htm" target="_blank"><i>></i>【本科生转专业】2024年优秀生类转专业工作...</a><span>04.30</span></li>
                
              </ul>
            </div> 
</div>
          </div>
        </div>
        <div class="col xs-12 sm-6 md-4 lg-4">
          <div class="main-c-item">
            <div class="main-c-t"> <span>研究生培养</span><a href="/10851/list.htm" target="_blank">MORE</a> </div>
            <div class="main-c-list" frag="窗口11" portletmode="simpleNews" configs="{'c8':'0','c42':'320','c25':'320','c30':'0','c29':'1','c22':'0','c23':'1','c34':'300','c20':'0','c31':'0','c16':'1','c3':'5','c2':'标题,发布时间','c27':'480','c43':'0','c17':'0','c5':'_blank','c24':'240','c32':'','c26':'1','c37':'1','c28':'640','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'50','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'20','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'1','c13':'200','c41':'240'}" contents="{'c2':'0', 'c1':'/通知公告/研究生培养'}"><div id="wp_news_w11"> 

              <ul>
                
                <li><a href="/2024/0531/c10851a341546/page.htm" target="_blank"><i>></i>【博士答辩】吴新泉 博士答辩公告</a><span>05.31</span></li>
                
                <li><a href="/2024/0530/c10851a341427/page.htm" target="_blank"><i>></i>【博士答辩】谢明昆 博士答辩公告</a><span>05.30</span></li>
                
                <li><a href="/2024/0530/c10851a341426/page.htm" target="_blank"><i>></i>【博士答辩】唐英鹏 博士答辩公告</a><span>05.30</span></li>
                
                <li><a href="/2024/0530/c10851a341417/page.htm" target="_blank"><i>></i>【博士答辩】‪汤宇豪 博士答辩公告</a><span>05.30</span></li>
                
                <li><a href="/2024/0530/c10851a341405/page.htm" target="_blank"><i>></i>【博士答辩】‪祁树仁 博士答辩公告</a><span>05.30</span></li>
                
              </ul>
            </div> 
</div>
          </div>
        </div>
        <div class="col xs-12 sm-6 md-4 lg-4">
          <div class="main-c-item">
            <div class="main-c-t"> <span>学生工作</span><a href="/10852/list.htm" target="_blank">MORE</a> </div>
            <div class="main-c-list" frag="窗口12" portletmode="simpleNews" configs="{'c8':'0','c42':'320','c25':'320','c30':'0','c29':'1','c22':'0','c23':'1','c34':'300','c20':'0','c31':'0','c16':'1','c3':'5','c2':'标题,发布时间','c27':'480','c43':'0','c17':'0','c5':'_blank','c24':'240','c32':'','c26':'1','c37':'1','c28':'640','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'50','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'20','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'1','c13':'200','c41':'240'}" contents="{'c2':'0', 'c1':'/通知公告/学生工作'}"><div id="wp_news_w12"> 

              <ul>
                
                <li><a href="/2024/0529/c10852a341300/page.htm" target="_blank"><i>></i>计算机学院举办“导学引领，乐享运动”2024...</a><span>05.29</span></li>
                
                <li><a href="/2024/0528/c10852a341072/page.htm" target="_blank"><i>></i>计算机学院举办“生涯筑梦 计遇未来” 职业...</a><span>05.28</span></li>
                
                <li><a href="/2024/0424/c10852a337149/page.htm" target="_blank"><i>></i>翼下山河，守护疆土——飞行员的航空梦想与爱...</a><span>04.24</span></li>
                
                <li><a href="/2024/0325/c10852a334018/page.htm" target="_blank"><i>></i>计算机学院举办“赓续红色血脉，坚定理想信...</a><span>03.25</span></li>
                
                <li><a href="/2024/0320/c10852a333643/page.htm" target="_blank"><i>></i>访企拓岗促就业 | 计算机学院赴华为技术有...</a><span>03.20</span></li>
                
              </ul>
            </div> 
</div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <script src="/_upload/tpl/04/02/1026/template1026/js/owl.carousel.min.js"></script>
  <script type="text/javascript">
  jQuery("#owl-demo1D").owlCarousel({
    loop: false,
    margin: 15,
    autoWidth: true,
    items: 3,
    dots: false,
    mouseDrag: false
  });
  $(".owl-item").on('click', '.item .news_title_item', function() {
    $('.news_title_item').removeClass('active');
    $(this).addClass('active');
    var index = $(this).parents(".owl-item").index();
    $('.in-news-con .row').eq(index).stop().slideDown().siblings().stop().slideUp();
  });
  </script>
  <div class="in-info">
    <div class="container container-pd">
      <div class="in_title in_title_c in-info-span"><span class="active">学术信息</span></div>
      <div class="in-info-con">
        <div class="row" style="display: block;" frag="窗口13" portletmode="simpleNews" configs="{'c8':'1','c42':'320','c25':'320','c30':'0','c29':'1','c22':'0','c23':'1','c34':'300','c20':'0','c31':'0','c16':'1','c3':'6','c2':'标题,扩展字段11,扩展字段12,扩展字段13,扩展字段14,扩展字段15','c27':'480','c43':'0','c17':'0','c5':'_blank','c24':'240','c32':'','c26':'1','c37':'1','c28':'640','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'30','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'20','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'0','c13':'200','c41':'240'}" contents="{'c2':'0', 'c1':'/学术信息'}"><div id="wp_news_w13"> 

          
          <div class="col xs-12 sm-4 md-4 lg-4">
            <div class="info-item">
              <div class="rili">
                <div class="date">2024-05-15</div>
              </div>
              <div class="title"> <a href='/2024/0511/c11617a338940/page.htm' target='_blank' title='青年学者论坛'>青年学者论坛</a>
                <p class="p1">题目：1、信息抽取任务中弱监督数据的生成和利用；2、高性能非易失性内存存储栈关键技术研究</p>
                <p class="p2">报告人：戴洪良、蔡淼</p>
                <p class="p3">时间： <span class="t-time">2024-05-15</span> 中午12:00</p>
                <p class="p4">地点：计算机学院楼113会议室</p>
              </div>
            </div>
          </div>
          
          <div class="col xs-12 sm-4 md-4 lg-4">
            <div class="info-item">
              <div class="rili">
                <div class="date">2024-05-13</div>
              </div>
              <div class="title"> <a href='/2024/0509/c11617a338771/page.htm' target='_blank' title='【石榴大讲堂】Theories of contracts and their applications'>【石榴大讲堂】Theories of contracts and ...</a>
                <p class="p1">题目：Theories of contracts and their applications</p>
                <p class="p2">报告人：Jean-Pierre Talpin, INIRIA</p>
                <p class="p3">时间： <span class="t-time">2024-05-13</span> 上午9：30-12：00</p>
                <p class="p4">地点：计算机学院511会议室</p>
              </div>
            </div>
          </div>
          
          <div class="col xs-12 sm-4 md-4 lg-4">
            <div class="info-item">
              <div class="rili">
                <div class="date">2024-05-17</div>
              </div>
              <div class="title"> <a href='/2024/0508/c11617a338634/page.htm' target='_blank' title='【石榴大讲堂】可信量子机器学习算法———鲁棒性和公平性'>【石榴大讲堂】可信量子机器学习算法———鲁...</a>
                <p class="p1">题目：可信量子机器学习算法———鲁棒性和公平性</p>
                <p class="p2">报告人：官极</p>
                <p class="p3">时间： <span class="t-time">2024-05-17</span> 10：30</p>
                <p class="p4">地点：计算机学院507会议室</p>
              </div>
            </div>
          </div>
          
          <div class="col xs-12 sm-4 md-4 lg-4">
            <div class="info-item">
              <div class="rili">
                <div class="date">2024-04-29</div>
              </div>
              <div class="title"> <a href='/2024/0423/c11617a336884/page.htm' target='_blank' title='【石榴大讲堂】Generate Anything: Learning to Generate Things and Stuff'>【石榴大讲堂】Generate Anything: Learnin...</a>
                <p class="p1">题目：Generate Anything: Learning to Generate Things and Stuff</p>
                <p class="p2">报告人：唐浩</p>
                <p class="p3">时间： <span class="t-time">2024-04-29</span> 上午10点</p>
                <p class="p4">地点：计算机学院楼507会议室</p>
              </div>
            </div>
          </div>
          
          <div class="col xs-12 sm-4 md-4 lg-4">
            <div class="info-item">
              <div class="rili">
                <div class="date">2024-05-17</div>
              </div>
              <div class="title"> <a href='/2024/0417/c11617a336381/page.htm' target='_blank' title='【石榴大讲堂】神经网络的可证明原像下近似'>【石榴大讲堂】神经网络的可证明原像下近似</a>
                <p class="p1">题目：神经网络的可证明原像下近似</p>
                <p class="p2">报告人：张喜悦</p>
                <p class="p3">时间： <span class="t-time">2024-05-17</span> 上午9点30</p>
                <p class="p4">地点：计算机学院507会议室</p>
              </div>
            </div>
          </div>
          
          <div class="col xs-12 sm-4 md-4 lg-4">
            <div class="info-item">
              <div class="rili">
                <div class="date">2024-04-21</div>
              </div>
              <div class="title"> <a href='/2024/0416/c11617a336237/page.htm' target='_blank' title='【石榴大讲堂】A Cryptography-Specific Language for Security Analysis of Block Ciphers against Differential Cryptanalysis'>【石榴大讲堂】A Cryptography-Specific La...</a>
                <p class="p1">题目：A Cryptography-Specific Language for Security Analysis of Block Ciphers against Differential Cryptanalysis</p>
                <p class="p2">报告人：Taolue Chen</p>
                <p class="p3">时间： <span class="t-time">2024-04-21</span> 上午9：30</p>
                <p class="p4">地点：计算机学院507会议室</p>
              </div>
            </div>
          </div>
          
          <a class="in-news-more" href="/10854/list.htm">更多>></a>
        </div> 
</div>
      </div>
    </div>
  </div>
  <script type="text/javascript">
  $(function() {
    // 文章日历
    $(".rili").sudyPubdate({
      target: ".date", // 日期元素
      lang: "en", //    月份语言，支持"en"英文，"cn"中文, "num"数字，默认为数字
      separator: "-", // 文章日期之间分隔符的，默认是后台输出的"-"号
      format: "年月日", // 文章日期格式，默认为年月日
      tpl: '<h3>%m%</h3><p>%d%</span></p>'
      // 自定义输出格式  %Y%代表年，%m%代表月, %d%代表日, %w%代表星期, %M%代表翻译后月份，%W%代表翻译后星期
    });
    // 文章日历
    $(".p3").sudyPubdate({
      target: ".t-time", // 日期元素
      lang: "", //  月份语言，支持"en"英文，"cn"中文, "num"数字，默认为数字
      separator: "-", // 文章日期之间分隔符的，默认是后台输出的"-"号
      format: "年月日", // 文章日期格式，默认为年月日
      tpl: '%m%月%d%日'
      // 自定义输出格式  %Y%代表年，%m%代表月, %d%代表日, %w%代表星期, %M%代表翻译后月份，%W%代表翻译后星期
    });
  });
  </script>
  <script type="text/javascript">
  $('.in-info-span span').on('click', function() {
    var index = $(this).index();
    $(this).addClass('active').siblings().removeClass('active');
    $('.in-info-con .row').eq(index).show().siblings().hide();
  });
  </script>
  <div class="lnk">
    <div class="container">
      <div class="row">
        <div class="col xs-12 sm-12 md-5 lg-5">
          <div class="lnk-l">
            <div class="row">
              <div class="col xs-4 sm-2 md-2 lg-2">
                <div class="lnk-l-item"> <a href="http://newsweb.nuaa.edu.cn/"> <span><img src="/_upload/tpl/04/02/1026/template1026/images/a1.png"></span>
                    <p>新闻网</p>
                  </a> </div>
              </div>
              <div class="col xs-4 sm-2 md-2 lg-2">
                <div class="lnk-l-item"> <a href="http://i.nuaa.edu.cn/"> <span><img src="/_upload/tpl/04/02/1026/template1026/images/a2.png"></span>
                    <p>智慧校园</p>
                  </a> </div>
              </div>
              <div class="col xs-4 sm-2 md-2 lg-2">
                <div class="lnk-l-item"> <a href="http://lib.nuaa.edu.cn/"> <span><img src="/_upload/tpl/04/02/1026/template1026/images/a3.png"></span>
                    <p>图书馆</p>
                  </a> </div>
              </div>
              <div class="col xs-4 sm-2 md-2 lg-2">
                <div class="lnk-l-item"> <a href="https://mail.nuaa.edu.cn/coremail/index.jsp"> <span><img src="/_upload/tpl/04/02/1026/template1026/images/a4.png"></span>
                    <p>电邮系统</p>
                  </a> </div>
              </div>
              <div class="col xs-4 sm-2 md-2 lg-2">
                <div class="lnk-l-item"> <a href="http://hqjt.nuaa.edu.cn/bcsk/list.htm"> <span><img src="/_upload/tpl/04/02/1026/template1026/images/a5.png"></span>
                    <p>班车时刻</p>
                  </a> </div>
              </div>
              <div class="col xs-4 sm-2 md-2 lg-2">
                <div class="lnk-l-item"> <a href="https://oas.nuaa.edu.cn"> <span><img src="/_upload/tpl/04/02/1026/template1026/images/a6.png"></span>
                    <p>OA办公网</p>
                  </a> </div>
              </div>
            </div>
          </div>
        </div>
        <div class="col xs-12 sm-12 md-7 lg-7">
          <div class="lnk-r">
            <div class="row">
              <div class="col xs-6 sm-3 md-3 lg-3">
                <div class="lnk-r-item"> <a href="/2017/1115/c11649a177680/page.htm" target="_blank"><img src="/_upload/tpl/04/02/1026/template1026/images/b1.jpg"></a> </div>
              </div>
              <div class="col xs-6 sm-3 md-3 lg-3">
                <div class="lnk-r-item"> <a href="http://keyselab.nuaa.edu.cn/"><img src="/_upload/tpl/04/02/1026/template1026/images/b2.jpg"></a> </div>
              </div>
              <div class="col xs-6 sm-3 md-3 lg-3">
                <div class="lnk-r-item"> <a href="http://pami.nuaa.edu.cn/"><img src="/_upload/tpl/04/02/1026/template1026/images/b3.jpg"></a> </div>
              </div>
              <div class="col xs-6 sm-3 md-3 lg-3">
                <div class="lnk-r-item"> <a href="javascript:void(0)"><img src="/_upload/tpl/04/02/1026/template1026/images/b4.jpg"></a> </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div class="foot">
    <div class="container container-pd">
      <div class="fo-con clearfix">
        <div class="row">
          <div class="col xs-12 sm-7 md-8 lg-9">
            <div class="fo-l">
              <p>地址：江苏省南京市江宁区将军大道29号</p>
              <p>邮政编码: 211106 </p>
              <p>版权所有：南京航空航天大学计算机科学与技术学院/人工智能学院 ALL RIGHTS RESERVED 苏ICP备05070685号 <a href="http://site.nuaa.edu.cn/" target="_blank">后台管理</a> <a href="mailto:njliujr@nuaa.edu.cn" target="_blank"><img src="/_upload/tpl/04/02/1026/template1026/images/email.png"> 书记信箱</a> <a href="mailto:cb_china@nuaa.edu.cn" target="_blank"><img src="/_upload/tpl/04/02/1026/template1026/images/email.png"> 院长信箱</a></p>
            </div>
          </div>
          <div class="col xs-12 sm-5 md-4 lg-3">
            <div class="fo-r">
              <h3>友情链接</h3>
              <div class="fo-r-con clearfix">
                <div class="select fl" frag="窗口15" portletmode="simpleNews" configs="{'c8':'1','c42':'320','c25':'320','c30':'0','c29':'1','c22':'0','c23':'1','c34':'300','c20':'0','c31':'0','c16':'1','c3':'30','c2':'标题','c27':'480','c43':'0','c17':'0','c5':'_blank','c24':'240','c32':'','c26':'1','c37':'1','c28':'640','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'50','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'-1','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'1','c13':'200','c41':'240'}" contents="{'c2':'0', 'c1':'/友情链接'}"><div id="wp_news_w15"> 

                  <p>校外导航链接<i></i></p>
                  <ul>
                    
                    <li><a href='javascript:void(0)' target='_blank' title='友情链接'>友情链接</a></li>
                    
                  </ul>
                </div> 
</div>
                <div class="fo-lnk clearfix fl"> <a class="wx" href="javascript:;">
                    <div class="ewm"><img src="/_upload/tpl/04/02/1026/template1026/images/ewm.png" /></div>
                  </a> <a class="wb" href="javascript:void(0)"></a> </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <script>
  $(function() {
    $(".select p").on("click", function(e) {
      var objDiv = $(this).next('ul');
      if (objDiv.css('display') == 'none') {
        objDiv.css('display', 'block');
        objDiv.parent().siblings().find('ul').css('display', 'none');
        event.stopPropagation();
      } else {
        objDiv.css('display', 'none');
      }
      $(document).one("click", function() {
        objDiv.hide();
      });
      e.stopPropagation();
    });
    $(".select p").next('ul').on("click", function(e) {
      e.stopPropagation();
    });
    $('.wx').on('click', function() {
      $('.ewm').toggle();
    });
  });
  </script>
</body>

</html>
 <img src="/_visitcount?siteId=68&type=1&columnId=1952" style="display:none" width="0" height="0"/>"
	fmt.Println(s)
}
