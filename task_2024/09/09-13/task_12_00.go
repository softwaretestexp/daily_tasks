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

<meta content="1aoda_4NM.f3zvlf8246LskS953hgvA6" r="m"><!--[if lt IE 9]><script r='m'>document.createElement("section")</script><![endif]--><script type="text/javascript" r='m'>$_ts=window['$_ts'];if(!$_ts)$_ts={};$_ts.nsd=51871;$_ts.cd="qt2DrrAltu7qtGWjqrq5iaqjrPLolkg4qcLntGEAHPLrlqVciaqXEGG4JnLDhPEYknaErcq4qcLmtOVAHPLrlqVciaqjrnLllkgWqcLctGWAHPLqhq9loO7qoAVjqfq5tGlArAE4qcLotOWAHO7qtGEXEfVHhqLlWO76tGljrfq5tGA9ck7qtGEjJrq5iaqjq1LDlkgjrpqmqGaqtGWjqrq5iaqjrPgPmAZMxtqrJOqcrsE6LcHhl3b7.i3Kk8v_DbDNopr7QOuMP0.RRwO90laoVuDWLaqlqGQlqLZHS2eyxsR_4uqaQnSURVOS1KYD1oRe1XziM9VGHYe35nm8tCJ6wc4UQceuQ1f2QeqXQCmaQKTGS1aXhPq.QooYR2TRhb2SMZaBMvmXFnfzTKe7tDSnFD5WMUYLMP70hZyCMDfItnRqbsJOFCr3K6B7J9GZpP2HQermJsSmxsy3uvR7FcyNQKdWhCySMb7LMwYBMcf7QbfBODTTFDg.JcdiwbSLUn7bpZlaWbypHb7_g9GSsupKHl.hMKNVQnV5KWyCMDGXFUmBuceXQKy.tD6YMCaLMoxLMgSzwKfXtuABuvR7FmZ.x2hBRVfzQYmBFWlTicN0UYwFSDR0YbW2H2OiwbSLhb2SMZaBMvmXFnfzTKe7tDSnFD5WMUYLMP70hZyCMDfItnRAu92_w6p.Q2bKFUYUJKJKA7Z_MYzmxsy3uvR7FcyNQKdWhCySMb7LMwYBMcf7QbfBODTTFDg.JcdiwbSLUn7bKwfbAbT4HCS.a99XtvJTM0F3VvL0p1V5KWyCMDGXFUmBuceXQKy.tD6YMCaLMoxLMgSzwKfXtuABuvR7FmZ.xbhxibpu1VyLYQeiiKm7YsYFnVrRHVV2H2OiwbSLhb2SMZaBMvmXFnfzTKe7tDSnFD5WMUYLMP70hZyCMDfItnRL2DTI3omSWuc3wmW08myhYiyJiYzkxsy3uvR7FcyNQKdWhCySMb7LMwYBMcf7QbfBODTTFDg.JcdiwbSLUn7bM8yjYbez3DmfnmRBIUxXYbkFM6m.AnV5KWyCMDGXFUmBuceXQKy.tD6YMCaLMoxLMgSzwKfXtuABuvR7FmZ.xb6uRDTkAmeiK7pxM0QusCTJT9ET3CW2H2OiwbSLhb2SMZaBMvmXFnfzTKe7tDSnFD5WMUYLMP70hZyCMDfI8Pf35CTTFDg.FvMWMnabWPVLERGuEPGGKUTy0nzPFUyGFcjiMUY23n2fReVz3CLGtCNCuDZ7JcyNQKdWUPSIMoxLMgSzwKfXtn3uGPZPHkE2tco5t6pNFbSPQwS6FKY786wjaCeZwCm4wChi3DfS3C2OIemyIop_I6rd0owZworXFvIQQCfeQKyNIer.8Km_w6f7ubRZFCf4FCIsICe2wUw6IZy0Qof7QbYPCKSTQDN4FvkmICy7M6gzhQyC3KV73brCObrXxPyNQKdWhuqLMoxLMd0BKCN6FCGBuvR7Fcg2JcDWEsG0Ec7bU3fG3Km2F6fTuDxP8oYPMUBxM19jMox23LybRol7RDLfODTTFDg.JcdiwbSLUnzIMwYBMcf7QbfBOclaxPg2HkichnxFtvJNFeNfIorCRofCaKyZQbzTQoMrM6JzM6znMWpaRUR03vRdTvm2FbYXFbMQwDf0IvrzMQp6wKN73bfdT6WXMbm4Q6Bh8vfaIvr5Rwwd8kR5ICwdCDrb3vyd36d3MbwLtn2jw7mPhbmbQPNGuPl7Fvw.Fc5AhCySMbzRhdzzwKfXtCNCuDZ7xOE2tcD3WnELEm973wN.FCr5Mor4anLXFvwbR16mR6Vj3K9bhZyCMDGXJnfzTKe7K1yQFvMWMnSjwCzLhgEuEPGGHuAfOcm8hDp43bhkRP9jMox23LybRol7RDLfODTTFDg.JcdiwbSLUnzIMwYBMcf7QbfBOclaxPg2HkichnxFtbx.MeS.3UTC3vw6P1TXQKJbtbMlwPyvMcVLMwYBMcGntCNCuDewtm2NQKdWhCySMb7LERqfhc3_Jn3BG20LIbp6FKMsICYzR6zN87pPwDY5FbmBaKmZRbzL3DusFbTfwvYeIZRb3KY936wdnomZ3CTN36kEFP9jMox23LybRol7RDLfODTTFDg.JcdiwbSLUnzIMwYBMcf7QbfBOclaxPg2HkichnxFtvTG3JJdQCyfMUf6uvRfRUyzFvMr3vS5RvebtLyzwKYOtKmvT1TOFPq.FvMWMna0hb2SMZS8hme7QbfBODTTFDg.xOichnEZWPVLEIg.3KTaRoRPSvQzRUm03K4cM62g3CJNFzJaFKz7wofPaCyZRKmOFUdr3bZNRCR6t7yC3KV73brCObrXxPyNQKdWhuqLMoxLMd0BKCN6FCGBuvR7Fcg2JcDWEsG0Ec7bUwJ6w1Sb3bN66cLXFvwbR16mR6Vj3K9bhZyCMDGXJnfzTKe7K1yQFvMWMnSjwCzLhgEuEPGGHuAfOcm8hozd8oIUtPyjwCR2hQYvw1NfF13BuvR7Fcg6tD6YMCSRh2Sjw7SBhDN6FCGBGOWPtcqdJcDWEYg7FbfSQWYdMoS5woTaTCRP8owTwktxhKyS3C3jR7pChbw7x1fzTKe7tkE.FvMWMT0LKK2SMZaBMvmXFnGf9cl7xOL6xP5cU1f6Ivp73zxbt1N7QbYPObR9Q1S9FPDWMUYLMP70hZyCMDfItTezTKe7tDSnFD5WEsqbhPVZWgEBE2Lj3oydS6Nb8DYd8DsY3byLF6zOQQmvwKYC3vf9abra8Dp58DUYwoYgRoJZweS23Kw6F6r_9vebRKmnFoIRICR23UxLwe9XRKN53br2uUJZ3KrXFDdQR6NZ8brvIZp4QoS536rGaCxZ3Dz0wnjiMUY23n2fReVz3CLGtCNCuDZ7JcyNQKdWUPSIMoxLMgSzwKfXtn3uGPZPHkE2tco5t6NG3UpgFQyC3KY5QCNCaKxZwCf9MvdFR6fbtn2jw7mPhbmbQPNGuPl7Fvw.Fc5AhCySMbzRhdzzwKfXtCNCuDZ7xOE2tcD3WnELEm97FQJBFKRT3bT4nKTZMbm.MKoQFKJX8vzX3W2u3UwS8CyGavr_QoyzR6dxQDYgF6YgF7e6QCw5Mbz69veNF63C8D.Hw6SzMopgF7TaIDTXRbRdnKxB8DN5FosAQoS7wbf2IZwy3vSnQn0zuvRfR1SP3oFi3Dgbhb2SMZaBWcf7QbfBd1eMFvw.FcdiwbSLhPV0E3afikAGtnRwPo2uIbe4IvksRvxnFCpfIeNuIoyNI6rdCbzg36yeMDsEI6NOIvTP3QTd8br5IvmLSUwfFbYdw6dRwCY2M6zeQJJfIoTNQnNySoegIKNdR6BQ8627FvJg8eJ98Uf_wDWjObTTRKWN3KIYhKJjEczjw7SBhkAXFUmBumg7UCSnFD5WMUYLMP7bWgEBEOanx1GfdPzaMoz0MKksICYLRCz2339zMvmORPNbSo3XRCa2tD6YMCaLWPzjw7SBU1fRFUmBuceXQKy.tcDAE1abiOQbhgxwtD2fFbRX6Dy2FDJ2IUjxhKyS3C3jR7pChbw7x1fzTKe7tkE.FvMWMT0LKK2SMZaBMvmXFnGf9cl7xOL6xP5cU1f7R6Yg8wrXt1N7QbYPObR9Q1S9FPDWMUYLMP70hZyCMDfItTezTKe7tDSnFD5WEsqbhPVZWgEBEPLLtKNCaKEX3KRntbkiE1SjwCzLhyqBMvmXFT9B5CTTFDg.FvMWMnabWPVLERGuEPGGK1Sa0DxO31ePIoHWMox2wby.MwSC3UxSQDYSuUef3UrL3K6oFbyfhCz23wSjsDYGtvSHgVrYAVR4RKM1MDY0QUznFzx6wKYXwDTXP1TXQKJbtbMlwPyvMcVLMwYBMcGntCNCuDewtm2NQKdWhCySMb7LERqfhc3_Jn3BG20L3bzN3odWQorgMCmzMQAjhbN6RbVzSKwTtbmNxPdiwbSLhOQLMwYBMm9XUDNCuDZ7Fvw.Fc5cWnELEkZ0E3afUPSjwCwfuD2ZwClXtb6Y3bljRCJShQJzEPf7QbfBOkW7Fvw.Fm_WKDySMb7LMwYBMcGGJn3BGO7axPg2KvHrwCJfMoJfQ3Gz3C2v3vGjObTTRKWN3KIYhKJjEczjw7SBhkAXFUmBumg7UCSnFD5WMUYLMP7bWgEBEOanx1GfdPzfFomuQUd3QDwLtn2X3EyzwKYOtKmvT1TOFPq.FvMWMna0hb2SMZS8hme7QbfBODTTFDg.xOichnEZWPVLEIg.MbSjwvfznDr28DSPQodJRvpgMCxfwZ2zR6fXMbR6SUr08DyXRvoHMDTgMbybQQYuWvfXMbRSSUJZFDz23ChVRvSLFCYgFEgNIDy_3DGz0CrN3KS93Udt8CwgFDfbIZN9Q6fL3uVu9cLXFvwbR16mR6Vj3K9bhZyCMDGXJnfzTKe7K1yQFvMWMnSjwCzLhgEuEPGGHuAfOcm8hDmu3vdkQcTZRvpg3JWXMoJnwofG0oeOFKwPIUHRICJB8bYbMzSGR6ff3b2BOUYn8DmPMKMHRKRz3Kxg3WYd3CwORbzS6DrPw6y2FCuh3ox.IbR6wex7hKzOwofP0U2uIoN4RUht3oS2MvxNMQ208c07FUmPa1T23oQNRCCchCySMb7LWgSzwKfXK228jGEiUD7d3bk1AVgdWOQSpzw0JsEywmxtSOrvAl7ZM9UaFsQZYlGd3XNEs2pTF99gyYrmROxzYvvJYuxkiVzJYWw.V2m1psr1_ufRwKJupK6hAUJu3DGdViaTKsrg3VJKSKwotOW5HKkDWbEjFoAdMJfFV9wvYsYNnbpHtOW5HKkDWbEjFoAdsRf1YDzOFbwKa2wrtOW5HKkDWbEjFoAdwHrC10yKwCR1nKEytbf0HCU.WU3a8oyO3JprKOeOpmaTasWy3uxApbjMQ0R61lrWJzp.FkE0HDyleVp0p9Ru1DkmV1gTiuyvRRmfhbySHD769K3nAmJrw0umMcgTiuyvRRmfhbySHcLnbsNO3OJ2tbvsqaqlquQlqNlkmCSCQ6Ay6P0jtbSnRKtiRbpShDpjiRGuqqEuHOTGSOEqruWSJGijWOquJOqeWyA6WqAmHuQ_9OAaWqEUOKvw94cAHNn5bq2kRPohsAAsYF0dF5Mips8OLaWlrGVkr_qOrsXN8kwIiO6oyaQynvlfGbx1yGlsH5DO8stETYAYfa7qkOl6JODpWsQZiuqlqNlkcOVTJsEnbsLTJk3ohCs9i9YkMvJkidY4WU3_sTlSCV3XRORfw0Bz3KzDVUYYJQmCVC0_H0JX9OWqrqEorqtAqaluqdGu83SP3KQXRb2SODxuF1ybQbjW3KmBhbY2wgSfMUAXRKTzODq7RCJ2tDksQcSvQUlL3WeXhDwnwnfvaUQ73Dy.tDIH3cSGMUWLRZeG8cfvIClBSKS0tDp.R1dDMCWLRDfjhZR7wcf9FDN6ODp9RPyGMKKWRoLLRKe.hZfBQcfjw6ABnowOtDz9FndxMK9LFC2PhZ2zQnf.wD7Bnvq7MbJTtDvARcSXQbxOhZNaMCLXMKYTa1e7MKq.FDuohCT2QPzNF7ABMUACtC2ua1ezWDW.FUishCT0R1zNFJ3BMKzbtC27TPezMKZ.FUHAhCT6FnzNQZY9hDzGM1f7a677FCYGtDuW31S.MbELMJp2hop4t6p_aneaRUl.wDsJh6r0Mnz0QeWBQopCt6pvSPeawoRdtoHiRcS0FCALQZSjhopS3nfnnKV7wURdtoorh6xNhvY2MLSSMUxGt6RbSnen3Ul.wvIDh6x2Qcz6MESa3bGXwDx4Oor6R1y0QCKWQoROhvpeQES63U3XQCYyOowN3PyuMKUDh6p.QoYfMwABwoSCt6r6SPeCwC3.Qvksh6R6w1zn8ZABwCz931f0uCYbJPyaFUKWwDT2hvr2w3S0QoQXICRXOozaF1ydFocW8CS68PzZQWaB8D2Xt6gBCbNatDJ6MndAMC26wPzzQZlBMDYGraAkNma9K09oqmnlFOqlr2GGRdWkrm7bMKgkLOEnrqquHAifWGqmWaQlq_q";if($_ts.lcd)$_ts.lcd();</script><script type="text/javascript" charset="utf-8" src="/LIWghRUPB4QK/oxYRCeR1jjgO.199cf1b.js" r='m'></script><script language="javascript" src="/_js/sudy-jquery-autoload.js" jquery-src="/_js/jquery-1.9.1.min.js" sudy-wp-context="" sudy-wp-siteId="68"></script>
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
      
      <li><img src="/_upload/article/images/de/1f/d41dead342b98bcd2fc4628e1067/681fed19-48c5-4eed-9828-01b0a69e130d.jpg" /></li>
      
      <li><img src="/_upload/article/images/86/88/f2ef68b44c5e9a69ddd5b32f819f/bbfa325d-5e6c-4e38-8d5b-531597e40432.png" /></li>
      
      <li><img src="/_upload/article/images/d4/9c/2ba2b6a34b1d9fc89fcc925c3bb7/00fab851-97ca-4ec2-b88a-f3d3c6a84d78.png" /></li>
      
      <li><img src="/_upload/article/images/45/45/a46e74f24558b84ce2c9d7e02af5/24004537-f4bf-40aa-83b1-0f9e7afcb15e.png" /></li>
      
      <li><img src="/_upload/article/images/e4/10/fc65960e495594ed53c0251205cd/a8f964df-b3a8-4a50-b217-be146078f351.png" /></li>
      
      <li><img src="/_upload/article/images/8e/5b/f9e805164267b723013414e8d7ca/e2397f85-7e06-4124-97ba-0ea7a35e1a05.png" /></li>
      
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
                          <div style="background-image:url(/_upload/article/images/11/37/c1f6a51c4a39be9ed36445def052/247564a3-a3a4-4146-a40f-b1939b35691a.png);"></div>
                          <p><a href='/2024/0903/c11624a352365/page.htm' target='_blank' title='计遇·智算未来 | 学院2024级新生迎新工作圆满完成'>计遇·智算未来 | 学院2024级新生迎新工作圆满完成</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/96/93/313cc7ea498e94224e738e21a160/55d0b3fa-fb5d-4459-bb82-8cb6f9549557.jpg);"></div>
                          <p><a href='/2024/0827/c11624a351892/page.htm' target='_blank' title='第十三届“中国软件杯”大学生软件设计大赛全国总决赛在我校成功举办'>第十三届“中国软件杯”大学生软件设计大赛全国总决赛在我校成功...</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/7c/20/14dc95fa4f8cbbea4d5aecf8bfdf/9a7c1f14-bc09-4207-a673-d64a6778c454.png);"></div>
                          <p><a href='/2024/0819/c11624a351709/page.htm' target='_blank' title='学院赴中国矿业大学调研交流'>学院赴中国矿业大学调研交流</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/6e/de/95994b044d4ab5aee33c6a3a87af/3f78bdba-ef3b-4421-88b3-21fd4d3fd2ba.png);"></div>
                          <p><a href='/2024/0802/c11624a351432/page.htm' target='_blank' title='计算机学院皮德常教授团队荣获第四届全国高校教师教学创新大赛一等奖'>计算机学院皮德常教授团队荣获第四届全国高校教师教学创新大赛一...</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/e9/ed/3ea4333a4a2c94a554dfc36c1e4d/2fdc93a5-11c2-4bf6-aa06-71358596fcb9.png);"></div>
                          <p><a href='/2024/0801/c11624a351423/page.htm' target='_blank' title='学院“云游非遗”实践团队获国家级媒体报道'>学院“云游非遗”实践团队获国家级媒体报道</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/34/9c/fc344b4b42bf93b77028623fe755/b641212a-84ba-4d45-978b-71c03fb1fc60.jpg);"></div>
                          <p><a href='/2024/0729/c11624a351367/page.htm' target='_blank' title='南航学子探寻非遗足迹，助力传承文化精髓'>南航学子探寻非遗足迹，助力传承文化精髓</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/05/ba/ac1677b1462894a7deff393d0d17/8ee56baa-7c84-4975-8743-cf03bd6cc6d7.jpg);"></div>
                          <p><a href='/2024/0729/c11624a351355/page.htm' target='_blank' title='访企拓岗促就业 | 计算机科学与技术学院/软件学院、人工智能学院赴四川开展走访调研'>访企拓岗促就业 | 计算机科学与技术学院/软件学院、人工智能学院...</a></p>
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
                        
                        <li><a href='/2024/0903/c10846a352365/page.htm' target='_blank' title='计遇·智算未来 | 学院2024级新生迎新工作圆满完成'>计遇·智算未来 | 学院2024级新生迎新工作圆...</a><i>09-03</i></li>
                        
                        <li><a href='/2024/0827/c10846a351892/page.htm' target='_blank' title='第十三届“中国软件杯”大学生软件设计大赛全国总决赛在我校成功举办'>第十三届“中国软件杯”大学生软件设计大赛...</a><i>08-27</i></li>
                        
                        <li><a href='/2024/0821/c10846a351745/page.htm' target='_blank' title='深圳力维智联技术有限公司董事长徐明一行来院交流'>深圳力维智联技术有限公司董事长徐明一行来...</a><i>08-21</i></li>
                        
                        <li><a href='/2024/0819/c10846a351709/page.htm' target='_blank' title='学院赴中国矿业大学调研交流'>学院赴中国矿业大学调研交流</a><i>08-19</i></li>
                        
                        <li><a href='/2024/0802/c10846a351431/page.htm' target='_blank' title='计算机学院皮德常教授团队荣获第四届全国高校教师教学创新大赛一等奖'>计算机学院皮德常教授团队荣获第四届全国高...</a><i>08-02</i></li>
                        
                        <li><a href='/2024/0801/c10846a351423/page.htm' target='_blank' title='学院“云游非遗”实践团队获国家级媒体报道'>学院“云游非遗”实践团队获国家级媒体报道</a><i>08-01</i></li>
                        
                        <li><a href='/2024/0729/c10846a351367/page.htm' target='_blank' title='南航学子探寻非遗足迹，助力传承文化精髓'>南航学子探寻非遗足迹，助力传承文化精髓</a><i>07-29</i></li>
                        
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
                
                <li><a href='/2024/0911/c10847a353216/page.htm' target='_blank' title='2024级人工智能创新班拟录取学生名单公示'>2024级人工智能创新班拟录取学生名单公示</a><i>09-11</i></li>
                
                <li><a href='/2024/0906/c10847a352631/page.htm' target='_blank' title='预备党员转正公示'>预备党员转正公示</a><i>09-06</i></li>
                
                <li><a href='/2024/0630/c10847a348271/page.htm' target='_blank' title='预备党员转正公示'>预备党员转正公示</a><i>06-30</i></li>
                
                <li><a href='/2024/0625/c10847a347930/page.htm' target='_blank' title='2023级“卓越工程师教育培养计划”专业 “卓越班”拟录取公示'>2023级“卓越工程师教育培养计划”专业 “...</a><i>06-25</i></li>
                
                <li><a href='/2024/0621/c10847a347604/page.htm' target='_blank' title='预备党员转正公示'>预备党员转正公示</a><i>06-21</i></li>
                
                <li><a href='/2024/0606/c10847a342432/page.htm' target='_blank' title='计算机科学与技术学院/人工智能学院/软件学院2024年志趣专长类转专业拟录取学生名单的公示'>计算机科学与技术学院/人工智能学院/软件学...</a><i>06-06</i></li>
                
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
                
                <li><a href="/2024/0906/c10853a352631/page.htm" target="_blank"><i>></i>预备党员转正公示</a><span>09.06</span></li>
                
                <li><a href="/2024/0630/c10853a348271/page.htm" target="_blank"><i>></i>预备党员转正公示</a><span>06.30</span></li>
                
                <li><a href="/2024/0621/c10853a347604/page.htm" target="_blank"><i>></i>预备党员转正公示</a><span>06.21</span></li>
                
                <li><a href="/2024/0603/c10853a341740/page.htm" target="_blank"><i>></i>关于2022-2024年度学院先进基层党支部、优...</a><span>06.03</span></li>
                
                <li><a href="/2024/0524/c10853a340864/page.htm" target="_blank"><i>></i>发展党员公示</a><span>05.24</span></li>
                
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
                
                <li><a href="/2024/0912/c10850a353274/page.htm" target="_blank"><i>></i>【推免工作】2025届本科毕业生推荐免试攻读...</a><span>09.12</span></li>
                
                <li><a href="/2024/0902/c10850a352230/page.htm" target="_blank"><i>></i>【创新班选拔】2024级人工智能创新班选拔报...</a><span>09.02</span></li>
                
                <li><a href="/2024/0606/c10850a342416/page.htm" target="_blank"><i>></i>【毕业设计】2024届本科生毕业生毕业设计归...</a><span>06.06</span></li>
                
                <li><a href="/2024/0524/c10850a340772/page.htm" target="_blank"><i>></i>南京航空航天大学计算机学院 关于征集集中...</a><span>05.24</span></li>
                
                <li><a href="/2024/0524/c10850a340765/page.htm" target="_blank"><i>></i>【本科生实习】关于开展2024年度本科生实习...</a><span>05.24</span></li>
                
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
                
                <li><a href="/2024/0621/c10851a347584/page.htm" target="_blank"><i>></i><font style='font-weight:bold;color:#FF050D;'>【硕士招生】2025年招收推免生预报名通知</font></a><span>06.21</span></li>
                
                <li><a href="/2024/0912/c10851a353271/page.htm" target="_blank"><i>></i>【博士答辩】虞建 博士答辩公告</a><span>09.12</span></li>
                
                <li><a href="/2024/0910/c10851a352992/page.htm" target="_blank"><i>></i>【博士答辩】孔令军 博士答辩公告</a><span>09.10</span></li>
                
                <li><a href="/2024/0904/c10851a352380/page.htm" target="_blank"><i>></i>【博士答辩】邱厚明 博士答辩公告</a><span>09.04</span></li>
                
                <li><a href="/2024/0904/c10851a352379/page.htm" target="_blank"><i>></i>【博士答辩】严大鹏 博士答辩公告</a><span>09.04</span></li>
                
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
                
                <li><a href="/2024/0712/c10852a350851/page.htm" target="_blank"><i>></i>我院举办“生涯辅导谈话主题工作坊”暨班主...</a><span>07.12</span></li>
                
                <li><a href="/2024/0613/c10852a342859/page.htm" target="_blank"><i>></i>计算机学院开展“浓情端午‘粽’享美好”端...</a><span>06.13</span></li>
                
                <li><a href="/2024/0529/c10852a341300/page.htm" target="_blank"><i>></i>计算机学院举办“导学引领，乐享运动”2024...</a><span>05.29</span></li>
                
                <li><a href="/2024/0528/c10852a341072/page.htm" target="_blank"><i>></i>计算机学院举办“生涯筑梦 计遇未来” 职业...</a><span>05.28</span></li>
                
                <li><a href="/2024/0424/c10852a337149/page.htm" target="_blank"><i>></i>翼下山河，守护疆土——飞行员的航空梦想与爱...</a><span>04.24</span></li>
                
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
      <div class="in_title in_title_c in-info-span"><span class="active">学术信息</span><a class="in-news-more" href="/10854/list.htm">更多>></a></div>
      <div class="in-info-con">
        <div class="row" style="display: block;" frag="窗口13" portletmode="simpleNews" configs="{'c8':'1','c42':'320','c25':'320','c30':'0','c29':'1','c22':'0','c23':'1','c34':'300','c20':'0','c31':'0','c16':'1','c3':'6','c2':'标题,扩展字段11,扩展字段12,扩展字段13,扩展字段14,扩展字段15','c27':'480','c43':'0','c17':'0','c5':'_blank','c24':'240','c32':'','c26':'1','c37':'1','c28':'640','c40':'1','c15':'0','c14':'1','c44':'0','c33':'500','c10':'30','c18':'yyyy-MM-dd','c36':'0','c1':'1','c6':'20','c19':'yyyy-MM-dd','c21':'0','c4':'1','c35':'-1:-1','c39':'300','c38':'100','c7':'1','c12':'0','c9':'0','c11':'0','c13':'200','c41':'240'}" contents="{'c2':'0', 'c1':'/学术信息'}"><div id="wp_news_w13"> 

          
          <div class="col xs-12 sm-4 md-4 lg-4">
            <div class="info-item">
              <div class="rili">
                <div class="date">2024-07-02</div>
              </div>
              <div class="title"> <a href='/2024/0628/c11617a348195/page.htm' target='_blank' title='【石榴大讲堂】Exploring Trustworthy Foundation Models under Imperfect Data'>【石榴大讲堂】Exploring Trustworthy Foun...</a>
                <p class="p1">题目：Exploring Trustworthy Foundation Models under Imperfect Data</p>
                <p class="p2">报告人：Bo Han</p>
                <p class="p3">时间： <span class="t-time">2024-07-02</span> 下午2:00</p>
                <p class="p4">地点：计算机学院楼507会议室</p>
              </div>
            </div>
          </div>
          
          <div class="col xs-12 sm-4 md-4 lg-4">
            <div class="info-item">
              <div class="rili">
                <div class="date">2024-06-21</div>
              </div>
              <div class="title"> <a href='/2024/0619/c11617a346041/page.htm' target='_blank' title='【石榴大讲堂】Blockchain-based Federated Learning for Wireless Metaverses'>【石榴大讲堂】Blockchain-based Federated...</a>
                <p class="p1">题目：Blockchain-based Federated Learning for Wireless Metaverses</p>
                <p class="p2">报告人：康嘉文</p>
                <p class="p3">时间： <span class="t-time">2024-06-21</span> 15:50-16:20</p>
                <p class="p4">地点：计算机学院楼507会议室</p>
              </div>
            </div>
          </div>
          
          <div class="col xs-12 sm-4 md-4 lg-4">
            <div class="info-item">
              <div class="rili">
                <div class="date">2024-06-21</div>
              </div>
              <div class="title"> <a href='/2024/0619/c11617a346040/page.htm' target='_blank' title='【石榴大讲堂】口令猜测的新路线：基于经典机器学习的猜测技术'>【石榴大讲堂】口令猜测的新路线：基于经典...</a>
                <p class="p1">题目：口令猜测的新路线：基于经典机器学习的猜测技术</p>
                <p class="p2">报告人：汪定</p>
                <p class="p3">时间： <span class="t-time">2024-06-21</span> 16:20-16:50</p>
                <p class="p4">地点：计算机学院楼507会议室</p>
              </div>
            </div>
          </div>
          
          <div class="col xs-12 sm-4 md-4 lg-4">
            <div class="info-item">
              <div class="rili">
                <div class="date">2024-06-20</div>
              </div>
              <div class="title"> <a href='/2024/0607/c11617a342522/page.htm' target='_blank' title='【石榴大讲堂】云原生智能弹性缓存研究及其开源发展'>【石榴大讲堂】云原生智能弹性缓存研究及其...</a>
                <p class="p1">题目：云原生智能弹性缓存研究及其开源发展</p>
                <p class="p2">报告人：顾荣</p>
                <p class="p3">时间： <span class="t-time">2024-06-20</span> 下午3点</p>
                <p class="p4">地点：计算机学院507会议室</p>
              </div>
            </div>
          </div>
          
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
              <p>版权所有：南京航空航天大学计算机科学与技术学院/软件学院 ALL RIGHTS RESERVED 苏ICP备05070685号 <a href="http://site.nuaa.edu.cn/" target="_blank">后台管理</a> <a href="mailto:njliujr@nuaa.edu.cn" target="_blank"><img src="/_upload/tpl/04/02/1026/template1026/images/email.png"> 书记信箱</a> <a href="mailto:huangsj@nuaa.edu.cn" target="_blank"><img src="/_upload/tpl/04/02/1026/template1026/images/email.png"> 院长信箱</a></p>
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
