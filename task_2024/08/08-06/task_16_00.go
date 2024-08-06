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

<meta content="bn_Xu5pW9pZne6.sdFQ6Lfajo7ZJ.Cda" r="m"><!--[if lt IE 9]><script r='m'>document.createElement("section")</script><![endif]--><script type="text/javascript" r='m'>$_ts=window['$_ts'];if(!$_ts)$_ts={};$_ts.nsd=49382;$_ts.cd="qoNDrrAlDk7qtGqjqpq5iaqjrnLllkgjqSqtrcLrhqGlxs7qhPExiaqjknLrlkgjrrqmrs7qtGEjqSZ5iaqXEGLjqcLrq1axrrE2qSqmrcLDlqVktGVArAEjk1atrr04qcLotGAAHO7qtG0jqSq5tG3ArAWjrpqmHcLchciMqsVarAAnWx7OdL088DLGFGh_DX80C5kXtl7qWz_b4Sytjr.w700fqalkrqVDrqK0kZYIIPACIwAaW6W.YbYHn92bp9e6IV6AsbTmxuSIUZqfU1z9wvAN4oq7Q63NRUs.hvrfw6rzRdg7h1gvtvrC4bR8K12LQKO6hCTSFKLLFXyBMnf.QCfLzCTCFD9jJcOTwbeLKcLbp_Eu3DzvKYzZTuRhHllLs6s.ps3ZAPA5UdYCMCGNFomLenezQK2jtD.XMCZLF6JLFhwzwbfNtOALeURBFmGjx9k_WkmXYOe.Hw2EW9ZuYuSJZDTNYoA9H2dTwbeLtKfSFHQBMUmNFPfj4beBtDzuFDP6MURLF1L0tHYCMCf3tPRYTDYhMURAMCd_Js7jQTSUUIebwTR9xkyIeURBFc2LQKO6hCTSFKLLFXyBMnf.QCfLzCTCFD9jJcOTwbeLKcLbYdYx8UwuFvx8.K2CYuxPY9U0ibyMAPA5UdYCMCGNFomLenezQK2jtD.XMCZLF6JLFhwzwbfNtOALeURBFmGjx24OR9mzIky7MH2xh1N0wDerLTxyJlE9H2dTwbeLtKfSFHQBMUmNFPfj4beBtDzuFDP6MURLF1L0tHYCMCf3tPR2ZONvQ0yHFmh0sONfFTATUIxqKOTcxkyIeURBFc2LQKO6hCTSFKLLFXyBMnf.QCfLzCTCFD9jJcOTwbeLKcLbM_ezUDrCQkW6L6pAJoeI1THYssNMpcA5UdYCMCGNFomLenezQK2jtD.XMCZLF6JLFhwzwbfNtOALeURBFmGjxb.YF2YBFbmfRBJKMoTTMTRvgKenF0A9H2dTwbeLtKfSFHQBMUmNFPfj4beBtDzuFDP6MURLF1L0tHYCMCf3tPRjnDpXA0w8s26d1KSoJlyz1XTcJbJ9xkyIeURBFc2LQKO6hCTSFKLLFXyBMnf.QCfLzCTCFD9jJcOTwbeLKo7LUZrzwbfNtbNSeCZBxOV9tcFOWnlLxTfeIXG.3K2y3KGzzKTCRKAL3KkXhKrjxnNjQIwBhuANFomLeTgBUCzuFDP6MURLF1LbJh9BEsaTxcGbS1z6FKTz3vMvwCN2F6N6MIfBI6wb8vw2eKr.QKpL36Oj8Ue0IUya8iz6I6p6MCNay6J.IURBFUO.F6NfIURgFdYbICzB8bzG_oe7RUxCw6OTwoJgF6J2R8x7wvrz8bN6dvezMDfdh1.Twbx2tCJGQMYGM13NFomLenZutDzuFDODhTfjQDNLtHYCMCGNxkAbznl4JcQjx2ju3Dxf3bNgQBwP3Ufn3CTZZb0jtbzuRKKTRbwStCYjx.wzwbfNtOALeURBFmGjUC.XMCZLF6JLFhQfWn3Nxka0X1ZfKPTnFKuf3UeaQKYgQINXI6x7worS5beaMKf_QbBdQDRCQCJC8iTSRbNPMCNfy6p.wo2SMK.dw6RjFCJL8iSuhKyb8vJZZoN4Qo2SIbh7I67C8Uyv8iJGRKxyIbJggCy9FcaLFvo_3PTf3UWjRdVfhCNaFbGLjnezQK2jK1OqMURLF1NjQIwBhn3TxcGbLuWftcw3hDMfFCfG8KTaIHGjhKNaRCVjdbwCtbRLxPOTwbeLts3LFXyBMT9NUKNSeCZBFvrjFcPyWnlLxug0x.QfU1SO8bx25D3jtbzuRKKTRbwStCYjx.wzwbfNtOALeURBFmGjUC.XMCZLF6JLFhQfWn3Nxka0X1ZfKPTOFCOvFCxeQKJ6QhEzMUm9RnNfd63zRC79tD.XMCZLJ1NjQIwBUPf8FomLenezQK2jtcF.E1ZbHs3bth2wt6yOwb2f_oebMKr_FU6d3vw28KzfFHNfICR7FKr._DNX3Um0IUOeRbx23Cx68HJSICJjFDJ6ZCLjtbzuRKKTRbwStCYjx.wzwbfNtOALeURBFmGjUC.XMCZLF6JLFhQfWn3Nxka0X1ZfKPTZ3DILI6rXRbSgQHYC3bYyMDNS5bxdIbm5RPCTMUR2Rcff3igz3DLvtbNSeCZBJc2LQKO6UPeIF6JLFhwzwbfNtP30X1ZfHkV9tcMmtCxzwCYCRIT6hbYCwKmN5Ke_8DpPFK6LQDN.F6YgRIfXICYb3b2g5bx7tKrvwn6Twbx2tCJGQMYGM13NFomLenZutDzuFDODhTfjQDNLtHYCMCGNxkAbznl4JcQjx2.Lw6373KJjQiQjhKNaRCVjdbwCtbRLxPOTwbeLts3LFXyBMT9NUKNSeCZBFvrjFcPyWnlLxug0x.QfU1SZIvfa4cLzFvrvR1.5R63jRbGbtHYCMCGNJPfj4beBK12wFvo6MnejQDNLth9uE1GvHOAbznmwhDTNQUI5ICeZ8UYew5Sb3UfawoAT7PTzQKpvtbodwPTvFnALFXyBMnGTtbNSeCe8tmyLQKO6hCTSFKLLxj7fhn3eJP3LXY0.w62TMDUNRPLjF6J2RMYbR6l.RKLbzCTCFD9jJcOTwbeLKcNIFXyBMnf.QCfLznluxP99Hk1yhnmFhKm58HS_RUfPIvfO4bxzFDS_36M_R6R2QKJg3d2GQ6fOI6fP4vpC8DJnIokvRKxvQDNaIF2dRKYbQCfagve2RKRuFokPhbRj8KJG3BRuICmvMCfLyCwyIoTfR6Od86JZ8KxaRdfPICr7wUQzzKTCRKAL3KkXhKrjxnNjQIwBhuANFomLeTgBUCzuFDP6MURLF1LbJh9BEsaTxcGbS1zy3DRT8D4Twbx28UxjQINPI6wBRKyggKw4RPaLFvo_3PTf3UWjRdVfhCNaFbGLjnezQK2jK1OqMURLF1NjQIwBhn3TxcGbLuWftcw3hD4gMCNbQbJzIHxzICybFbTbyC2GMvT_MbInQCx6wvNXRd2a8UryMDwgZvrb8DS08D60w6rv8Ke.Qt2dFbzSJ6fzeowdMKzT8D6CQoezFKwb8HxPMofzICfu_6rdMomNRUOZ8KmZwUEztBYC3bV.3CrSzKrzxP2LQKO6huWLF6JLFW3BKDNaFbGLeURBFc99JcF6Es70xnLbK.JyQoyL8vy6_DRSQbS23UOfQ6e5Mvya8im.8CJyIDSOZ6ey362Z3bMCI62G8UeSMdy63bNPIvwggvwbRKf_IUIL3UeeMvEjIBed86TzIbwZy6z_MDTn8ouLRoNgIUYOhMYzwbY9tDmG4PTGFPQjFvo6MnZ0tKfSFHw8hTe.QCfLzCTCFD9jxO1yhnlZJ1ALx7V.QCSZwKTv_oebFDrjRKFnhKTSRDQj3IeChKw.xcfj4beBtkVjFvo6MTgLUbfSFHQBMUmNFPGbjnlBxO0CxPPyU1zNRbzbF8wXRbf9RoTe7PTzQKpvtbodwPTvFnALFXyBMnGTtbNSeCe8tmyLQKO6hCTSFKLLxj7fhn3eJP3LXY0.MDr08o4.MPLjF6J2RMYbR6l.RKLbzCTCFD9jJcOTwbeLKcNIFXyBMnf.QCfLznluxP99Hk1yhnljhnfjQINPhKmfQnNve1lBFvrjFcP.hCTSFKNRtWrzwbfNtbNSeCZBxOV9tcFOWnlLxTG7w5zP3DlL36S0eCSSRUmBFC.vwbxnwvm2wBRd3bm6FKmjdDNz31fjRKMvF9e2RnzZ1JNqY9mr8bYf_KSbwoR_Qb6Nw6R2FUYzFMEzMUm9RnNfd63zRC79tD.XMCZLJ1NjQIwBUPf8FomLenezQK2jtcF.E1ZbHs3bth2wtCx7FDrge6ru8DfGMK.dtPTjQDw2tByvwPNbFc3LeURBFc9CtD.XMCeRtYTjQIwBhCNaFbGLXsWftcQ5JcF6EY07MURvRBwNI6wbhnNj4bxPtbrPQ1.gM1lLF6JLFhQuhCNaFbfRzTfzQK2jtD.XMCZLxu3bth94Wn3NxmN05vwG3Kzn3UFuhKrN3DJghMYzwbY9tDmG4PTGFPQjFvo6MnZ0tKfSFHw8hTe.QCfLzCTCFD9jxO1yhnlZJ1ALx7V.3bfCQvmgg6r9FcaLMb8TMUR2Rcff3igz3DLvtbNSeCZBJc2LQKO6UPeIF6JLFhwzwbfNtP30X1ZfHkV9tcMmtCT7MUwgFBJGRvf.36rgebR68DfO3UknMKpgFKebQHyaw6fNMCRCeDSX8D2BRvM5QumgFKebwByuICf7RDJ24CRdFDS98D5TtUeXIKmLtXfGFbm.RKmgZUz98DeNRvOSRorgMCETJt7jhKNaRCVjdbwCtbRLxPOTwbeLts3LFXyBMT9NUKNSeCZBFvrjFcPyWnlLxug0x.QfU1SbQbxg5oqXIDmT8DILhbeuwUYgR5zd3D2G36T0gveGM6T9Rbvv3DpgRbJNFhRTQUfb3CTfeDY2MKRO8DI5ICrvRDw.wXwG3UwyRDz.ZDrSF62vw6kNMcy.RoYgR8NyQoSe8bYTZKrdRK2uFK.nwo7ztCfSRILzRbratDwjX1ezQK2jtk16MURLFYzRKRlckYS7Hbxv_VmJHuVCWVo7wk3nHvY1sI0aRlrWHKycnC90HmJxHCUp1lTVQbzxHeyApb3TMmR5e2QSAuytsYHZFTmpYVAaYwqgKvY9Q9YjZlmaQCwNHTIUJYLaICrkVIyvpcL0HuTvdsxftbeTHCBOK2mo3mATMBm9scL0HuTvdsxftbeTH0jOV2w.RDzvVI2UAPL0HuTvdsxftbeTH6Hhw9fiVURbYIxPiPNBwseP0kmTWoed3CIBAY9dRlYQWBLuibWTVVRzLoJmw0erskoBFC7nWuTXAZylw0RrQ9fvd2lzJsadRCtj31TXwOT.WtNbWVrVAowhdbazJsadRCtj31TXwO0jJFE_3DE0RcNX_qWkrqACrqKjqSp7QUx0HXVzt1N.QCY2zKRvQ1z2FO_OWaWcJkZeRd0nqaQ0Wk3l9sEuWulSHs1BJuWlqOg6HF90WuqmmLTXyNiu9kbwBNZnACVfulalmmgLf_ORkKgcvyGqralkqGcMES38nNKPfOmKaMQAoF_FOxvzFiAtgYrs0i1TG.GXMiRQqag6HuWn.s7nWuWnrqKjqa9TWOACHFGuJkaTrPeY.2xQJ0JQR24shVNtR2wX1I2R12rIw9mhLvxpQ2enpT6PRvJtJ0yjYjJlqaAmrGAmjaWcWasyW6_63bxvtKwNw.wPQD9NR6xzzCmPMn29RUc63Ky0tKpzF.wGhCw9Rcfv_oqBR6RatDI0MPevwU3L3HN6hCrNFPfGeDqB3DNntDk03o7L3oyftHy7wnfOFbVLdKeGtDJNFPOeMoQL3C2jQhw9RC3N3KTPzCpytDxXFnOuM6WLMUR0tHJ63cf7RK7LZbTjtDSL3POnMUqLMoY.tHmahCy9wcfX_CABMvYO3nOSQDSjtKS2QBLBMCTvtbf.dceXRUVjFK6dhCy0W1NNwHLBMOr9tb20jceXwDEjFKdzhCy.31NNF50BMbTLtb20_neXwCajFUs5Rce.RCaLFdf4hCzP31f.eClBFC2OtDBBR1e0IcN0IIGBQCYCtvpOePeuwDGjwosLh6J0Q1N03H0BQ6pSIPf0eKABwDS2tos6FPe0wbELwIx2h6YSIPfC5PeSF120RKb6QKynRnNC3IaBQKmCtvRad1eSRUQjwCX6QDmLtUYPIhwawDVNwUJPz6rT3n2TIU86wCxCtUx2I.w6FbENQbTPd1e6F6w03K.Bh6wZQ1NaQH0BwUwftvx6_ceTw6EjQvudh6p.3CqLQdp2Rs3NQK2nz6pXR12SRUt6woJ6tUybFMw4QC9NIbf6z6zBw60jIoI6h6zNF1N5timjQnf9wbZL_Cejw6ljMUs_hCe2Ra3lrxeQxCRprGrMBC2xrqmRED4sqawMEY2CrxLaqa3aHGAmjGWcJAVDrqA";if($_ts.lcd)$_ts.lcd();</script><script type="text/javascript" charset="utf-8" src="/LIWghRUPB4QK/oxYRCeR1jjgO.199cf1b.js" r='m'></script><script language="javascript" src="/_js/sudy-jquery-autoload.js" jquery-src="/_js/jquery-1.9.1.min.js" sudy-wp-context="" sudy-wp-siteId="68"></script>
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
      
      <li><img src="/_upload/article/images/d4/9c/2ba2b6a34b1d9fc89fcc925c3bb7/00fab851-97ca-4ec2-b88a-f3d3c6a84d78.png" /></li>
      
      <li><img src="/_upload/article/images/45/45/a46e74f24558b84ce2c9d7e02af5/24004537-f4bf-40aa-83b1-0f9e7afcb15e.png" /></li>
      
      <li><img src="/_upload/article/images/e4/10/fc65960e495594ed53c0251205cd/a8f964df-b3a8-4a50-b217-be146078f351.png" /></li>
      
      <li><img src="/_upload/article/images/8e/5b/f9e805164267b723013414e8d7ca/e2397f85-7e06-4124-97ba-0ea7a35e1a05.png" /></li>
      
      <li><img src="/_upload/article/images/ee/3f/c23c09c74990b6c6bf72de29cc9c/535d31e3-7de2-485f-993a-811042e8da7f.png" /></li>
      
      <li><img src="/_upload/article/images/d6/28/c0272c7b4ef98fdcf3de3f3e8990/8d7ab9a8-25ac-4bb1-b7d9-edf26d58cc92.png" /></li>
      
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
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/2d/96/e71cab884ec8b344c69ef8fd19b1/5019fb7e-8b80-4a76-a0bf-1b023f4324c3.jpg);"></div>
                          <p><a href='/2024/0719/c11624a351080/page.htm' target='_blank' title='计算机学院组织开展党纪学习教育暨  “石榴赋能”干部教师能力提升培训'>计算机学院组织开展党纪学习教育暨  “石榴赋能”干部教师能力提...</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/b5/32/637922e4472a9754f98bb71acddb/32a66fd3-0f8b-4e4c-86a4-9e19c6639c8a.png);"></div>
                          <p><a href='/2024/0708/c11624a350667/page.htm' target='_blank' title='计算机学院党委与南京市残联机关党总支签署党建共建合作协议'>计算机学院党委与南京市残联机关党总支签署党建共建合作协议</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/14/56/a46fbd4f49b6bd0ee7ffa5201823/fecf3081-c0f7-45f0-9583-e794180c0351.jpg);"></div>
                          <p><a href='/2024/0611/c11624a342685/page.htm' target='_blank' title='计算机学院工会举办“心有所期，忙而不茫”  芒种节气主题活动'>计算机学院工会举办“心有所期，忙而不茫”  芒种节气主题活动</a></p>
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
                        
                        <li><a href='/2024/0802/c10846a351431/page.htm' target='_blank' title='计算机学院皮德常教授团队荣获第四届全国高校教师教学创新大赛一等奖'>计算机学院皮德常教授团队荣获第四届全国高...</a><i>08-02</i></li>
                        
                        <li><a href='/2024/0801/c10846a351423/page.htm' target='_blank' title='学院“云游非遗”实践团队获国家级媒体报道'>学院“云游非遗”实践团队获国家级媒体报道</a><i>08-01</i></li>
                        
                        <li><a href='/2024/0729/c10846a351367/page.htm' target='_blank' title='南航学子探寻非遗足迹，助力传承文化精髓'>南航学子探寻非遗足迹，助力传承文化精髓</a><i>07-29</i></li>
                        
                        <li><a href='/2024/0729/c10846a351355/page.htm' target='_blank' title='访企拓岗促就业 | 计算机科学与技术学院/软件学院、人工智能学院赴四川开展走访调研'>访企拓岗促就业 | 计算机科学与技术学院/软...</a><i>07-29</i></li>
                        
                        <li><a href='/2024/0719/c10846a351081/page.htm' target='_blank' title='计算机学院赴电子科技大学调研交流'>计算机学院赴电子科技大学调研交流</a><i>07-19</i></li>
                        
                        <li><a href='/2024/0719/c10846a351080/page.htm' target='_blank' title='计算机学院组织开展党纪学习教育暨  “石榴赋能”干部教师能力提升培训'>计算机学院组织开展党纪学习教育暨  “石榴...</a><i>07-19</i></li>
                        
                        <li><a href='/2024/0712/c10846a350851/page.htm' target='_blank' title='我院举办“生涯辅导谈话主题工作坊”暨班主任培训会'>我院举办“生涯辅导谈话主题工作坊”暨班主...</a><i>07-12</i></li>
                        
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
                
                <li><a href='/2024/0630/c10847a348271/page.htm' target='_blank' title='预备党员转正公示'>预备党员转正公示</a><i>06-30</i></li>
                
                <li><a href='/2024/0625/c10847a347930/page.htm' target='_blank' title='2023级“卓越工程师教育培养计划”专业 “卓越班”拟录取公示'>2023级“卓越工程师教育培养计划”专业 “...</a><i>06-25</i></li>
                
                <li><a href='/2024/0621/c10847a347604/page.htm' target='_blank' title='预备党员转正公示'>预备党员转正公示</a><i>06-21</i></li>
                
                <li><a href='/2024/0606/c10847a342432/page.htm' target='_blank' title='计算机科学与技术学院/人工智能学院/软件学院2024年志趣专长类转专业拟录取学生名单的公示'>计算机科学与技术学院/人工智能学院/软件学...</a><i>06-06</i></li>
                
                <li><a href='/2024/0604/c10847a342057/page.htm' target='_blank' title='计算机科学与技术学院/人工智能学院/软件学院 2024-2025-1学期教材选用公示'>计算机科学与技术学院/人工智能学院/软件学...</a><i>06-04</i></li>
                
                <li><a href='/2024/0603/c10847a341740/page.htm' target='_blank' title='关于2022-2024年度学院先进基层党支部、优秀共产党员和优秀党务工作者拟表彰对象的公示'>关于2022-2024年度学院先进基层党支部、优...</a><i>06-03</i></li>
                
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
                
                <li><a href="/2024/0630/c10853a348271/page.htm" target="_blank"><i>></i>预备党员转正公示</a><span>06.30</span></li>
                
                <li><a href="/2024/0621/c10853a347604/page.htm" target="_blank"><i>></i>预备党员转正公示</a><span>06.21</span></li>
                
                <li><a href="/2024/0603/c10853a341740/page.htm" target="_blank"><i>></i>关于2022-2024年度学院先进基层党支部、优...</a><span>06.03</span></li>
                
                <li><a href="/2024/0524/c10853a340864/page.htm" target="_blank"><i>></i>发展党员公示</a><span>05.24</span></li>
                
                <li><a href="/2024/0522/c10853a340578/page.htm" target="_blank"><i>></i>2024级研究生新生党员组织关系转入相关说明</a><span>05.22</span></li>
                
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
                
                <li><a href="/2024/0606/c10850a342416/page.htm" target="_blank"><i>></i>【毕业设计】2024届本科生毕业生毕业设计归...</a><span>06.06</span></li>
                
                <li><a href="/2024/0524/c10850a340772/page.htm" target="_blank"><i>></i>南京航空航天大学计算机学院 关于征集集中...</a><span>05.24</span></li>
                
                <li><a href="/2024/0524/c10850a340765/page.htm" target="_blank"><i>></i>【本科生实习】关于开展2024年度本科生实习...</a><span>05.24</span></li>
                
                <li><a href="/2024/0612/c10850a342731/page.htm" target="_blank"><i>></i>【卓越计划】 2023级“卓越工程师教育培养...</a><span>05.22</span></li>
                
                <li><a href="/2024/0511/c10850a339344/page.htm" target="_blank"><i>></i>【毕业设计】2024届本科生毕业设计（论文）...</a><span>05.11</span></li>
                
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
                
                <li><a href="/2024/0719/c10851a351073/page.htm" target="_blank"><i>></i>【博士答辩】乔塨哲 博士答辩公告</a><span>07.19</span></li>
                
                <li><a href="/2024/0719/c10851a351072/page.htm" target="_blank"><i>></i>【博士答辩】高增 博士答辩公告</a><span>07.19</span></li>
                
                <li><a href="/2024/0719/c10851a351071/page.htm" target="_blank"><i>></i>【博士答辩】范利利 博士答辩公告</a><span>07.19</span></li>
                
                <li><a href="/2024/0717/c10851a350983/page.htm" target="_blank"><i>></i>【博士答辩】孟祥伟 博士答辩公告</a><span>07.17</span></li>
                
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
