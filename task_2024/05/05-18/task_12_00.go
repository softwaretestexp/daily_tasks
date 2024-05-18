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

<meta content="j9A.XjJtTKH.OiJyMLD0ACU3FAlpTPkb" r="m"><!--[if lt IE 9]><script r='m'>document.createElement("section")</script><![endif]--><script type="text/javascript" r='m'>$_ts=window['$_ts'];if(!$_ts)$_ts={};$_ts.nsd=38345;$_ts.cd="qhJDrpAlxPLmxqVjqrqmrs7qtG3jrSq5tGVmck7qtGWjqrq5iaqjrPgPJqZ7cqAbiaqXEG94qcgPcG7rlqVctGQ9ck7qtGVjqfqIiaqXEG7jrcLmq1axrcW4E1LktGqAHO7qtG3jrSq5inljr1Lclkg4qcLqtGlAHPLvhqLlDu7qtGEjqSq5iaqjEcLDlkg4qcLltGVAHPLchcLlDk7qtG3jrSq5iaqjr1LclkgjqpqmqPLDnYVDqs3SqGVSpG3O4mcAJLgNVt4qxOsXGw5Q2ddRzyRmnOEur6IooekIYwGqrAEor6WNkZR881VTc09aWoE7YC3yY022p0fCIOzJsbylEkTQtY0fUPNOw6wIQUqBQvlLRGxChvJGQomXlKW7hPZGt6WqQCRwKcyNQfyjhCyaMb7.ooJBMc27Qb78tbTTFKg.JTyLwbSNUn72WsZu3CNGKmGt1OREHV3NsqxCpsVep1VetKfCMD9XFUE8FPeXQDy.tSSuMCaNMox.oPNzwK2Xtuw8FoR7FYZ.xsRvWkxBVsyL10pEWVg0YOZ4MKTjYUW2H12LwbSNhb2nobLBMvrXFn7QQCe7tKSnFSgjMUYNMP76UbfCMD2ItnVO1KYEMompMr2vJsG.wmTVtCxbwmwOxs0NFoR7F1yNQfyjhCyaMb7.ooJBMc27Qb78tbTTFKg.JTyLwbSNUn72EKfx8vR0F6lLWD2TYOYOYsJNibTFp1VetKfCMD9XFUE8FPeXQDy.tSSuMCaNMox.oPNzwK2Xtuw8FoR7FYZ.x1f5R9x78ueBDbpxhPzuwKa6H2x_JVV2H12LwbSNhb2nobLBMvrXFn7QQCe7tKSnFSgjMUYNMP76UbfCMD2ItnVsMuN9Q92iFnpNsO2GMmVCtCeqKsSmxs0NFoR7F1yNQfyjhCyaMb7.ooJBMc27Qb78tbTTFKg.JTyLwbSNUn72DsxzUCmSQsrrHvpsJUf81cYEss2FAnVetKfCMD9XFUE8FPeXQDy.tSSuMCaNMox.oPNzwK2Xtuw8FoR7FYZ.xpSEF2RXMKrPlDyKM6SnM2VpIDeSF9W2H12LwbSNhb2nobLBMvrXFn7QQCe7tKSnFSgjMUYNMP76UbfCMD2ItnVQYKpzA9rIs1zP1KeDW0eXHorcJKxOxs0NFoR7F1yNQfyjhCyaMb7.ooJBMc27Qb78tbTTFKg.JTyLwbSNU6L.tYTzwK2XtC9oFbZ7xuE2tTq5WnENEm25coV.3bf53DzMtDTTRDWN3fRuhKJ.EczNrCNBhkVXFUE8F2g7UbSnFSgjMUYNMP72APABEO0nx1zAKcz0FDN73Ar_wC29Moz0DCYBIoRf86AsFDrLQDJN3qy68UST8veukvR6IoYaMb9D8vJLIomXFGyCF62G8vw4oKfbIDNL8CGKwUeBRoY6wqyLworyMoxbl6e7wUmj8C9r36eXMKe5h2SLwbm9hDxOrnfGMPQXFUE8FPZatKSnFSyMhTz.wCz.UbfCMD9XxswAtPlgJ1q.x1aX3DmGRKz4rDNP3v2T3bgtMC0NtCSnRfWLRbpahDpNVcNzwK2Xtuw8FoR7FYZ.UrSuMCaNMox.oPLfWcQXxsecxcZPKnNuFfNZ3USSwbp4rCwXIoJ.wUWoRCeuMDe4QpTPQDY6wDxTkvrSRKz2Mb9U8vpLwUyaMfSPw6Y.MDx.kvmuhbef86qtMUNgQUyaIppSI6G6Ive9kvyGRbJ5ICqiIbyvF17NFAwv3PyGRvENlKGfhDz6FCz8JPeXQDy.K2ywMUYNMPzNrCNBhcQnx1zAHOWPt1rRhSrZFCzfIbSucbVjhbz6RbRQ3CwTtCmNxYyLwbSNhOQ.ooJBMmGXUD9oFbZ7F6w.FTg9WnENEkZ6VcLfUPT98ClsRK3NtCSnRfWLRbpahDpNVcNzwK2Xtuw8FoR7FYZ.UrSuMCaNMox.oPLfWcQXxsecxcZPKnNPFry_FCmZwbx0rPZzMvrORP9U3v3XRba2tSSuMCaNWPzNrCNBU12RFUE8FPeXQDy.tTqCE1aviOQ2UPpwtoe9wCLUwUe2MDw4FGzP3vp9IbNPobwfIDw.FDWwwKNz3oRTIGyGRbm9RDJ0kbySIDxzFKqrMbLNtCSnRfWLRbpahDpNVcNzwK2Xtuw8FoR7FYZ.UrSuMCaNMox.oPLfWcQXxsecxcZPKnNe3SmTI6JB3KT4rbfC3Kp5MK9oRCxZICRdRY7LMUY93n2PmvWz3C7GtC9oFbZ7J1yNQfyjUPS8Mox.oPNzwK2XtnYcxcZPHsE2tTr3tCm7QDpTlCr6hKpSwDEIRDey8KJOFfzTQD2jMop4lCYXIDpf3CLiRCxBtDwbwmzLwbm9hDxOrnfGMPQXFUE8FPZatKSnFSyMhTz.wCz.UbfCMD9XxswAtPlgJ1q.x1STw6VzRbxNrvLjhbz6RbRQ3CwTtCmNxYyLwbSNhOQ.ooJBMmGXUD9oFbZ7F6w.FTg9WnENEkZ6VcLfUPT4I67DQ1LXF6wbR2SOR6V.3K92UbfCMD9XJn7QQCe7KcyQFAwjMnS.wCz.UPAuEP9GHuwAtPm8hKNLQGmOICSeIvp5qUmb3v26wUwkhnTXQDJbtpwPwPybMcV.ooJBMc9ntC9oFbewtY2NQfyjhCyaMb7.Vk3fhcQ_JnY8xm0Lwvy0MSJ0RP9.MoxblnfbRoq7RD2AtbTTFKg.JTyLwbSNUnzQooJBMc27Qb78tPlaxng2H0E9hnxMtbrekbm_Rv22I67YQCxXFKz43qrvR6Y9wbx4mKpGQo29Iv7VQ6pT8KpuIaR_RKmbwCzucupdRbpfQb7DI6ebRDmnFaRdhbY.IbxOmDzuIDrGMb788bw_IUNGRqyP86reIbJulKYPIDm.wopMtDTTRDWN3fRuhKJ.EczNrCNBhkVXFUE8F2g7UbSnFSgjMUYNMP72APABEO0nx1zAKcz_3Km08SfLwbm9IvJNrCwPIoRLRD0iIDwgRn7NFAwv3PyGRvENlKGfhDz6FCz8JPeXQDy.K2ywMUYNMPzNrCNBhcQnx1zAHOWPt1rRhSf2MC2vwKxXcbezIDefFCgA8b2OM6N4MpmBQCmCQUzzlKpa8vm5MKAiM6r28KzT8SzNw6JbIbyLrOpdFKNCJv7MFUwZMDS08Sz.QoS7MbR2kbePM62jIb7mwvrZMURLRGyf8KxeQvWXUDfC3KA73bWotDrXxnyNQfyjhuqNMox.o27BKCz6FCz8FoR7F1g2JTqjEsGTEc72hcyyQ6eB860rwKRnQCz93GyZQ6SdFUeukvS.8Dx5IKZYMve_3vye3pr.I6NfIvynDKJ63Kz2I6AiI6w2RDe4IGmT3USZFUWNcDxd8oSjICAt8vzyMKNu8aNTRo2y8vpGKnfzwKpOtKEKQnTOFnq.FAwjMnaThb2nobN8hmy7Qb78tbTTFKg.xVE9hnEeWPV.VmG.QDT4wDgpwUe2FKw.RfqBhKya3C3NmCxChbR7x17QQCe7tsE.FAwjMT0NKK2nobLBMvrXFnzAJPl7xuL6xYg9U1fL3KN2o6NXRK2ORUghhnTXQDJbtpwPwPybMcV.ooJBMc9ntC9oFbewtY2NQfyjhCyaMb7.Vk3fhcQ_JnY8xm0LMKwT8afCMP9.MoxblnfbRoq7RD2AtbTTFKg.JTyLwbSNUnzQooJBMc27Qb78tPlaxng2H0E9hnE.tc2NrCwPhbrbQP9pFcl7F6w.FTgChCyaMbzFU2TzwK2XtC9oFbZ7xuE2tTq5WnENEm9BqURP3CqB3vZcFbSnRoRXFrS_wbmuQUrbqDzd3KraFDEQ3KNX3ce.Rfr_F9S93cNdH9wqYVrD8C3UwDS2wUm4Qpz0w6Y9MvpXonZzMvrORP9U3v3XRba2tSSuMCaNWPzNrCNBU12RFUE8FPeXQDy.tTqCE1aviOQ2UPpwtDJ.FKWiFvra8KefMfSPtPy.wCRbUDJvw1zfF1Y8FoR7F1g6tSSuMCS3h2SNrCNBhDz6FCz8xkWPt1qdJTqjEYgzFvw9lDNNIoRfhP9QQCxftCwOQ2S2M1ENMox.oPLuhDz6FC7Bt2fXQDy.tSSuMCaNEkQ2UPA4WcQXxY9cR6wO3DSu3GqXhKJLRCx4KnfzwKpOtKEKQnTOFnq.FAwjMnaThb2nobN8hmy7Qb78tbTTFKg.xVE9hnEeWPV.VmG.3K2SQ6EiIvrvF17NMplLMUY93n2PmvWz3C7GtC9oFbZ7J1yNQfyjUPS8Mox.oPNzwK2XtnYcxcZPHsE2tTr3tCyzFvR4oDyGRU273vWiFCR08KeP3GRBMKwyMby2rbJawo2XMbVqFKSz8KyXRArOQuxyMby2qDJuID2.RKqsQbRZFKz28S9LtUSB8br.UoYGFKr7RDEiMozv8KfLRAy7RoJyFDWCAO3jhbz6RbRQ3CwTtCmNxYyLwbSNhOQ.ooJBMmGXUD9oFbZ7F6w.FTg9WnENEkZ6VcLfUPTfQCliRUqzIKR08SmThbSnQvp4lURd3Cfv3vgcI6eOMvN2Rpe_3Dwy3KxjoPzTQv2f3bgUFKYbMDmP8SmOICJb3CRLqoNG3vR5RKGwMKrnFvybwqR0McTj36p4l6wyQ6T_8C3kMDrZRDynFfSBwoG7hD2nlCQzRKm6tKAQxceXQDy.t0EjMUYNM2NFhaack2T.HClpwlmiHOE6WOwSwkVuiUppiCqaR0mhHD00Yb96HYpEHrJW1lyYwKNt1VJApKQnMYVEFTQnAO2hsPYfFTxAV9VuE00gKUpOQ03QMVmuQbrLHcmtJY9S8DmqxCJvpn7uHOgp3kxPtCf0HrT5K2xDRTVCDDS9sn7uHOgp3kxPtCf0Hka5V2pj3CN9xCpUA17uHOgp3kxPtCf0HqYKw9zHYvw2ECePi1zLwkaVssmCWUf53rmnAYL530pIpDQuiKEnVlVMHUJDw9fqs0wnFCGuJkSzJYJlwlwDQ07p3TlXJk75RrA631yBQsSLpOwbW9msAUAe3CaXJk75RrA631yBQsaNAuZ_3CWuR193wAWqrAW6r6W6qSwzwvJ61oGztPz7Qb3stDR9QcS9FVG5WaqkWug5lKqnqq3uWsYnkkEaWO3aH9EnJuqmrsZ01uA0Wklcmgg38yia9s_QB8bBAC3Gn00o39F5uusjy9I8wqVkrAQmqoVDEZJgfc5U_be6rwBr85pUibXBtn16kTgNZPLUkQWujEuYqqZaHuylJuquJkW5r6W6qaL0Js3SAO36Js3TrnaI3lZCVKyWpqr6WDTdWoyWxvrHVlRTYV3YJmNgMUY2McmYp6YKKuwXhKT1qqVcrawuJGWrWGHgWqGj3bmbhbRjqcNPQCGXRvlMtbmfMPy2RG3j3KTThbYXocNGhDROR17pwUq7RvmStSmNMPSbQvQ.mbw6hDmXFn7KFKq73KTutSRN3oGNR6ePUbJ7wc29FCR83DeOtKpLFYyGMoANRDfNrPN9RDQX3DgVtbp_tKYBFmyXM6qNFvw6Uby63n2.RDf8MCTNtKzN3YyBMUWNF6pLUbSahDeOw173wbA7M6xP3my7QDe.hbTbrDQBMDSGtC7w31ezRoE.FfzPhCTTJPzjqbQBMsmOtCLcJ1ezwKV.Ff2ahCTjRPzjoUqBMKSBtCLcwPezwb7.FGxORcSj3D0.oKY4hDN23c7wFbl7FbyPtSTnR1ST8nz6cCVBQDpSt6QYFneawKZ.waxTh6rTwPz6mbqBQoYCIn7cFDA7wKz9taxjFPSTQKW.qCe2hopCIn7qRnenFcyTRfZjQKTu3czTmClBQbrSt6VD3cenRoq.wr0jQDxNhvpfcPNawCAXwoqVtvrC3Py0IGljwCm6hvJbccN6FKWXQCgV3ce0FvrT3fSnh6pewPzurbqBwvRbt6lrw1eCwvV.QANPh6wjRDl.rK22ROQXQDLltvpzRcyaRGAjworChve2onN4QDGXIC7rtvz7wvL.Iamjh6fLMPzeUvSjQc2OwCS8wbeNwv3.MGxvhCS93qQoQGxQxcm4raWzEbmSrAR3En0CqapFxbmtQGA6JqVqJOynruEqrkWor6ED";if($_ts.lcd)$_ts.lcd();</script><script type="text/javascript" charset="utf-8" src="/LIWghRUPB4QK/oxYRCeR1jjgO.199cf1b.js" r='m'></script><script language="javascript" src="/_js/sudy-jquery-autoload.js" jquery-src="/_js/jquery-1.9.1.min.js" sudy-wp-context="" sudy-wp-siteId="68"></script>
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
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/c9/c6/3e2397ac4208b1357fc8e5380a5c/52d75459-d91c-4da6-bdec-bda3f2c88d79.jpg);"></div>
                          <p><a href='/2024/0412/c11624a335984/page.htm' target='_blank' title='计算机学院邀请南京大学陈道蓄教授来校作名师报告'>计算机学院邀请南京大学陈道蓄教授来校作名师报告</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/48/f1/07f9c38249ad9244bfc7f9b1cb1b/62a83db1-ec5d-44f3-a71b-2384543189f5.jpg);"></div>
                          <p><a href='/2024/0410/c11624a335840/page.htm' target='_blank' title='郑纬民院士作客计算机学院石榴大讲堂'>郑纬民院士作客计算机学院石榴大讲堂</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/80/8d/ecf9b97643dfbc3c0fdef6ee9dfd/db6f014a-a026-44be-acf9-5173345b02ac.jpg);"></div>
                          <p><a href='/2024/0402/c11624a335068/page.htm' target='_blank' title='国家工业信息安全发展研究中心来计算机学院调研交流'>国家工业信息安全发展研究中心来计算机学院调研交流</a></p>
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
                        
                        <li><a href='/2024/0516/c10846a339941/page.htm' target='_blank' title='我院在计算机体系结构领域国际顶级会议发表最新研究成果'>我院在计算机体系结构领域国际顶级会议发表...</a><i>05-16</i></li>
                        
                        <li><a href='/2024/0507/c10846a338414/page.htm' target='_blank' title='我院学生获得ICPC全国邀请赛（武汉）季军'>我院学生获得ICPC全国邀请赛（武汉）季军</a><i>05-07</i></li>
                        
                        <li><a href='/2024/0502/c10846a338078/page.htm' target='_blank' title='计算机学院召开全体教职工大会暨本科教育教学审核评估培训动员会'>计算机学院召开全体教职工大会暨本科教育教...</a><i>05-02</i></li>
                        
                        <li><a href='/2024/0430/c10846a338021/page.htm' target='_blank' title='南京航空航天大学第八届信息安全技能竞赛成功举办'>南京航空航天大学第八届信息安全技能竞赛成...</a><i>04-30</i></li>
                        
                        <li><a href='/2024/0430/c10846a338014/page.htm' target='_blank' title='我院在校“航空杯”教职工乒乓球比赛和江苏省高校计算机学科第六届教职工乒乓球团体赛中荣获佳绩'>我院在校“航空杯”教职工乒乓球比赛和江苏...</a><i>04-30</i></li>
                        
                        <li><a href='/2024/0430/c10846a337908/page.htm' target='_blank' title='2024鲲鹏昇腾开发者大会南航专场站圆满落幕'>2024鲲鹏昇腾开发者大会南航专场站圆满落幕</a><i>04-30</i></li>
                        
                        <li><a href='/2024/0429/c10846a337644/page.htm' target='_blank' title='喜报 │ 学院党委入选 “全国党建工作标杆院系”培育创建单位'>喜报 │ 学院党委入选 “全国党建工作标杆...</a><i>04-29</i></li>
                        
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
                
                <li><a href='/2024/0507/c10847a338384/page.htm' target='_blank' title='计算机学院2024年正高专业技术职务基层推荐评审结果公示'>计算机学院2024年正高专业技术职务基层推荐...</a><i>05-07</i></li>
                
                <li><a href='/2024/0514/c10847a339730/page.htm' target='_blank' title='关于南京航空航天大学本科教育教学审核评估自评报告的公示'>关于南京航空航天大学本科教育教学审核评估...</a><i>04-25</i></li>
                
                <li><a href='/2024/0109/c10847a329258/page.htm' target='_blank' title='关于自编讲义编写人员政治审查公示'>关于自编讲义编写人员政治审查公示</a><i>01-03</i></li>
                
                <li><a href='/2023/1227/c10847a328359/page.htm' target='_blank' title='预备党员转正公示'>预备党员转正公示</a><i>12-27</i></li>
                
                <li><a href='/2023/1215/c10847a327474/page.htm' target='_blank' title='预备党员转正公示'>预备党员转正公示</a><i>12-15</i></li>
                
                <li><a href='/2023/1214/c10847a327326/page.htm' target='_blank' title='2022-2023年度教学优秀奖拟推荐名单公示'>2022-2023年度教学优秀奖拟推荐名单公示</a><i>12-14</i></li>
                
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
                
                <li><a href="/2023/1227/c10853a328359/page.htm" target="_blank"><i>></i>预备党员转正公示</a><span>12.27</span></li>
                
                <li><a href="/2023/1215/c10853a327474/page.htm" target="_blank"><i>></i>预备党员转正公示</a><span>12.15</span></li>
                
                <li><a href="/2023/1128/c10853a325769/page.htm" target="_blank"><i>></i>发展党员公示</a><span>11.28</span></li>
                
                <li><a href="/2023/1121/c10853a325059/page.htm" target="_blank"><i>></i>发展党员公示</a><span>11.21</span></li>
                
                <li><a href="/2023/1013/c10853a321821/page.htm" target="_blank"><i>></i>预备党员转正公示</a><span>10.13</span></li>
                
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
                
                <li><a href="/2024/0511/c10850a339344/page.htm" target="_blank"><i>></i>【毕业设计】2024届本科生毕业设计（论文）...</a><span>05.11</span></li>
                
                <li><a href="/2024/0511/c10850a339341/page.htm" target="_blank"><i>></i>【毕业设计】关于做好2024届本科生毕业设计...</a><span>05.11</span></li>
                
                <li><a href="/2024/0429/c10850a337738/page.htm" target="_blank"><i>></i>【本科生转专业】2024年优秀生类转专业工作...</a><span>04.30</span></li>
                
                <li><a href="/2023/1129/c10850a325842/page.htm" target="_blank"><i>></i>【毕业设计】2024届本科毕业设计工作通知</a><span>11.29</span></li>
                
                <li><a href="/2023/0920/c10850a319935/page.htm" target="_blank"><i>></i>【创新班考核】关于开展2021级、2022级人工...</a><span>09.20</span></li>
                
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
                
                <li><a href="/2024/0511/c10851a338960/page.htm" target="_blank"><i>></i>【博士招生】2024年工程类博士研究生招生工...</a><span>05.11</span></li>
                
                <li><a href="/2024/0416/c10851a336275/page.htm" target="_blank"><i>></i>【博士招生】2024年招收博士研究生“申请考...</a><span>04.16</span></li>
                
                <li><a href="/2024/0330/c10851a334694/page.htm" target="_blank"><i>></i>【硕士招生】2024年专业课笔试、综合面试安...</a><span>03.30</span></li>
                
                <li><a href="/2024/0323/c10851a333892/page.htm" target="_blank"><i>></i>【硕士招生】2024年复试名单</a><span>03.23</span></li>
                
                <li><a href="/2024/0323/c10851a333891/page.htm" target="_blank"><i>></i>【硕士招生】2024年硕士研究生招生复试及录...</a><span>03.23</span></li>
                
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
                
                <li><a href="/2024/0424/c10852a337149/page.htm" target="_blank"><i>></i>翼下山河，守护疆土——飞行员的航空梦想与爱...</a><span>04.24</span></li>
                
                <li><a href="/2024/0325/c10852a334018/page.htm" target="_blank"><i>></i>计算机学院举办“赓续红色血脉，坚定理想信...</a><span>03.25</span></li>
                
                <li><a href="/2024/0320/c10852a333643/page.htm" target="_blank"><i>></i>访企拓岗促就业 | 计算机学院赴华为技术有...</a><span>03.20</span></li>
                
                <li><a href="/2024/0319/c10852a333394/page.htm" target="_blank"><i>></i> 计算机学院开展第十三期“山湖号”红色专...</a><span>03.19</span></li>
                
                <li><a href="/2024/0318/c10852a333263/page.htm" target="_blank"><i>></i>南京师范大学计电学院学工团队来我院调研</a><span>03.18</span></li>
                
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
