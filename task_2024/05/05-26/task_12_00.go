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

<meta content="umq7X_UjVu0ubz3p6oS.UmlcyrZYyrwL" r="m"><!--[if lt IE 9]><script r='m'>document.createElement("section")</script><![endif]--><script type="text/javascript" r='m'>$_ts=window['$_ts'];if(!$_ts)$_ts={};$_ts.nsd=23234;$_ts.cd="qEpDrrAlxk7qtGQjrrq5tGE9ck7qtGqjqpq4tGAArAV4qcLctGWAHPLrtAGlm1LmlqVciaqXEG94qcZctGWAHPLrhq9lEk7qtGlBrfq5iaqjrcLmlkgjqrqmqk7qtG3jrSq5tGA7cGAwtGqArAEjqrqmrs7qtG3jrSq4tGEArAWjqcLrq1aObGlTWAVmWOcFEn9OFkdi2D_zTsC8YFJWNh4ATsDKyqqX1P.TNOXPiZeWqaVmraAmzAdmUvg2Jv6YWkJntTw2sZzxR9wHwvTJ0VyzpcQeUTdrEY0jRow0t8TahoxCtDY0_Pya3owSMKIdhc0BxnzawXebU2GNFUmLePezQD2.tD.BMCZNF6xLFhRzwb2NtuALeoRBFYG.x98OWCrjRTfMIJ3C1uTrtVe0_9VCHVE2H2d0wbeNtK2SFH3BMUrNFnfj4CeBtKznFDPCMURNF170tHpCMC23tnRq5uqSMT35Fu631kxiWT3ZsImzFYY0xsyIeoRBF12NQKOChCTaFK7LFXeBMn2.QbfLzbTCFK9.JcO0wbeNKc7bYwY21CSCVKeB5uV4t6JwYT65RvwURcV5UdpCMC9NFUmLePezQD2.tD.BMCZNF6xLFhRzwb2NtuALeoRBFYG.x2I01U2CQCznUjSNw23n3TRmesNXK9E2H2d0wbeNtK2SFH3BMUrNFnfj4CeBtKznFDPCMURNF170tHpCMC23tnR8gbRrMoadMDug110.woYBAFx18sYcxsyIeoRBF12NQKOChCTaFK7LFXeBMn2.QbfLzbTCFK9.JcO0wbeNKc7bRIW_R6xisbf1eVg_3DzYJY6uAmLZAPV5UdpCMC9NFUmLePezQD2.tD.BMCZNF6xLFhRzwb2NtuALeoRBFYG.xb5XMYSfwvluJ_x0VkYdU9zsTl0_K9A2H2d0wbeNtK2SFH3BMUrNFnfj4CeBtKznFDPCMURNF170tHpCMC23tnRjTDNYMbNG3KMnYbSZwKyKR5rBQbNlxsyIeoRBF12NQKOChCTaFK7LFXeBMn2.QbfLzbTCFK9.JcO0wbeNKc7bF729M9xoU0ywd927pOpHMV.7AORPRcV5UdpCMC9NFUmLePezQD2.tD.BMCZNF6xLFhRzwb2NtuALeoRBFYf_tmdrMURNF1zjQIRBhnQTx1GbLOWft1wRIU4NtCmL8KrLhMpzwbp9tKmG4nTGFnQ.FvoCMnZTtK2SFHR8hTy.QbfLzbTCFK9.xO1ghnleJ1VLx7A.wCfzMDxCyvwjRDz4wC6yM6eCRvz6RIpGF6r9FKJg_62dwUTdQoOjw6eTwCyjQiRuF6SCMb2g_bz_3oT68DB0RvejM6z.3i2dMDpnw6wgeopu8KznRKUOMoRSMUzjwdedMKTLIP0jeoRPRczP3oD03D0vtK2SFH3BWn2.QbfLSne3F6r.FcO0wbeNt1V0x.3fiuVvtnRF7brP3Dx.8ohC3bmyQbxzIHJztPz.QbY2zDRvQcz9FPFCMURNF170tHpCMC23tTej4CeBtKznFDPCEsWvt1VZJhGBEY7zQC27gDmdQUm98oo4MveuMoYaQIzBQDSL86x.dvrbw6JPwvO7QKR.3byj38R0F6YyQDTjdvwCFCzPFoOzQnTBRvzuIHm_86myQoyP4Ue4J62dR6OP3DYO8UeO8ixXRD97tKNS5CEz3DmntbI0E1e.QDzLttLBMUrNFT9LaKTCFK9.FvoCMnZvJ1VLxjEuE19vK1SbgDzL3U27QouytPT.QDR2tBevwPzbF13LeoRBF196tD.BMCe3tYSjQIRBhCzaFCGLXkWft1QdJcFCEY0z36zPRI2btPz.QbY2zDRvQcz9FPFCMURNF170tHpCMC23tTej4CeBtKznFDPCEsWvt1VZJhGBEY7z3bzLybzPIomPw6cuhKTaRD3j3IyChKR.x1fj4CeBtsV.FvoCMTgNUb2SFH3BMUrNFnGbjPlBxu06xPPgU1zd36wN382dRbSG8C2edbx6Ro2j3KO43UevMDfGF52jMbrSwUTgdDRPRDJOw6OnQUePFb2OwdxjtPz.QbY2zDRvQcz9FPFCMURNF170tHpCMC23tTej4CeBtKznFDPCEsWvt1VZJhGBEY7zIKrv_UeaMCRz8ok0wbx98KTjQIzPI6efIC3zzDTCRDAN3KkBhKr.xnzjQIRBhuVNFUmLe2gBUbznFDPCMURNF17bJhGBEs0Tx1GbSczPMow9wbUZwny9wvpfFIfB8v293C2z_UrjFbz08DUyMve9RbJN8HzPMcfGRvQzeoRPRczP3oD03D0vtK2SFH3BWn2.QbfLSne3F6r.FcO0wbeNt1V0x.3fiuVvtnRF_UwChKmPFvkbtPT.QDR2tBevwPzbF13LeoRBF196tD.BMCe3tYSjQIRBhCzaFCGLXkWft1QdJcFCEY0zIoegQi0jhKzaRbVjdCwCtCRNxPO0wbeNtsQLFXeBMTGNUDNSebZBF6r.FcPgWnlNxuZ0x.3fU1TzFom6d6eBIU20IUIzRbmyQoY0JMWzMUr9RPNfdv3zRb72tD.BMCZNJ1zjQIRBUP28FUmLePezQD2.tcFjE1ZvHsQbthfwt6RywDS2_D3jtCznRKK0RbwatCpjx.Rzwb2NtuALeoRBFYG.UC.BMCZNF6xLFh3fWnQNxsa0XcZfKnTGIvOe8bYy3veg35ePMK278CJC5CwCRomP8DHg3oJy36yg3Xe0wv2OQ6SaybYPRvr.QoCLICY9RvxLQiJdRKpbQbfagnybF62P3Dh6Q6eGRCyLFiRv8UTe3DwgdvzuIU2OQDIy3vefMop6hMpzwbp9tKmG4nTGFnQ.FvoCMnZTtK2SFHR8hTy.QbfLzbTCFK9.xO1ghnleJ1VLx7A.8KmbwofXeoRPRo2uFvo43veCMCpX8irv8CQ7tKNS5CEz3DmntbI0E1e.QDzLttLBMUrNFT9LaKTCFK9.FvoCMnZvJ1VLxjEuE19vK1SX5KejR6xPMUuuMUeBRbzzRXRN3Dee8CyvZ6JPwvR4MbIgQo2S8KTv8HJaRv27wUfzeUwaRv2XF6cLICNjQsRgMIY6ICS.wofze6rdMD2bRvOu3bSyMUyLQ5NaICTSFDYgdU2fIUYuh1.0wbx9tCxGQMpGM1QNFUmLePZutKznFDOohTf.QDzLtHpCMC9NxsAbzPl4J1Q.x2jn8UrdFoz5w52bQUJ7RDmggoJdICydQoOGF6zP8UT73dxd8KxyIKxbe6ey3U25QKddwCx.3ve68iJ6RbpL86T6_KmdIDyutv45I6zZMUevIiR48bTzQ6fZ_Kp_8UT03nC0MUR9Rc2f3iZz3D7vtCNSebZBJ12NQKOCUPe8F6xLFhRzwb2Ntn30XcZfHsV2tcMlt6JzIopzR52dRb2GFCYb7nTzQDpbtbo5wPTbFnVLFXeBMn9TtCNSebe8tYyNQKOChCTaFK7LxjLfhnQeJn3LXm0.FDRjRbvbMbRNRDYeIMWzMUr9RPNfdv3zRb72tD.BMCZNJ1zjQIRBUP28FUmLePezQD2.tcFjE1ZvHsQbthfwtCTGwUf5_bgjtCznRKK0RbwatCpjx.Rzwb2NtuALeoRBFYG.UC.BMCZNF6xLFh3fWnQNxsa0XcZftnZNFvo43PTG3UEjRdAfhCzaFCGLjPezQD2.K1OrMURNF1zjQIRBhnQTx1GbLOWft1wRhoIj3brGtDxZwHR7QKpSMbzjyvRPQ6RaRUM6ICxGwCffFBTjMKqLFCYbybNHRDQjIlOMAmRqAUz238f7RvYC86xz_owCRD20MKbuhKTaRD3j3IyChKR.x1fj4CeBtsV.FvoCMTgNUb2SFH3BMUrNFnGbjPlBxu06xPPgU1zOMD2G8HRaQ62L3KTjdPLzF6rbR1.dR63.Rb9btHpCMC9NJnfj4CeBKc2QFvoCMne.QDzLthGuE19vHuAbzPmwhKT6RCMCFUeCRPgjFXeP3PzG36lj5K0ftKznFDPCWne.QDzLKMR3MUrNFnfj4CeBt1Q6xPPgiuWvt1YFwHz63Dr.QCmC7PTGFDrP8cC0MUR9Rc2f3iZz3D7vtCNSebZBJ12NQKOCUPe8F6xLFhRzwb2Ntn30XcZfHsV2tcMltCxNwvJS8ixaRD97tKyvzDTCRDAN3KkBhKr.xnzjQIRBhuVNFUmLe2gBUbznFDPCMURNF17bJhGBEs0Tx1GbSczzMKTC8D.n3DRyFCxa8Hwbw62L3bmaZCT98K2XRvkdQowyFKybwBY7Mv2NMbRCd6WS8K2XRvMdQ6eNMDYOR8ybIC27RUfBzc2dM6TGFc.N3DNGFCpf8Hr4Ro2BFDRgZDpa8KeOJs1jtPT.QDR2tBevwPzbF13LeoRBF196tD.BMCe3tYSjQIRBhCzaFCGLXkWft1QdJcFCEY0zRvJP8H2ahvTfwofv_1yBQvY08DIjICrL3DxewiJd3DyeRKRNybr98KRPFKP6wUmyRbxz3IY2RKSb3vfvd6eGRbpbF6Mb3DmC8KY.FdmGQKNyRvwa_DaXFbp08DU.8UreIUz2wIrGICpNQb2jZ6p4hczNQKU4hKRfQc2vF.GBMUrNFnG0zbTCFK2FKYbAqGTwMkePR5frssgTJumr4UACWuy0V96ZWDprskSXA7NNJkTs1ke20lfJYUxj1sdIVlRvJKTKIBwUWVE510TsdKzpVmYmWmUpimN0RoJmFBmrQoxvFkeUTOYFWUefATodR0A.JkgeRdan31zBwke.g2NppbrYJK4aRla.JkgeRdan31zBwketL2xUFbpjRTogY0E.JkgeRdan31zBwkeunvREsTm6R2Uu3OL.M6Ad3zlSwsmdIbJv49mRHbplUkh4WuNPJYrDMFmupURHAlaS4bz.WuWZM9HHpowoAUJERdephsAZHbwPjClzM63ZFuc.ROmrYlm61de7hsAZHbwPjClzM63ZtOKOibrOJDVjMX9kqaVlJaAmjaWVMUmuwk49h19.F6x2RMpbR6q.RDL5LOWkquA5HKIZWGWDJklCrxqnWuESWk00.OQurAEdJuCLJkWSrp8JMX8s7ylyvh_IW0J2EWSJrrsndUHKZmwLtN_Rqq3qqGWqbPW9iiApkH0zB2duEOXqou5wqlx_fCGxaBepMBhvvknUqGWHJug0J_anJuAeHGAmjaWxJkECWOtOJOA6JaQBVQeiYVRCIvp1amJ1RoRKMVh0FYxw36eLW5SvQVaeH0gSa9Qa1nz219uHWqWmrAQlrFLkqOEl5kJ5zbxPRP2bFUFC3vrLtKRnMMRf3bgNRKYazbmXw122MKjC3cebRDVLR52ahCRCQ1fveKgBRvY6tDk4wnefFK7L3HYGhCmjQnfGeKr4tKrd31OdMoQN3Cz2tHSB3c2OFDLLdDS6tKJLFvcCRDwvtKrz3.R9812PMC7LZbeutKT6wcOnwCqNMDp.tHJzFP27FKELZCTatKS0FnOaQceBRoVLMXN9hCeT3bWLZor7Fn2zRUh4hCe7RnzLFd0BMbpTtC2zdPeXwsl.FUs4hCgSRczNwt9BMvYPtC2B4ceXFbl.FKB7hCy7FczNwiLBMvR7tC20dCABFbwztDBy8nej3bqLFdRfhCNN3Pf.4bVBwUa.wo6yh6J9wPz03dQBQ6YLt6p0_1euwUl.wDkZh6JTQUZLwHp9h6Y7Rcf0ebLBwUROtoUuR1e0QUZLwB7BQKGNwKYNzvmXQCQ.wboeh6mGwPzCQHaBQKp6t6w.zvrfF1203vnCQop9tUpu3.RawKENwoT6zvwPwn2uRU_CwCNOtUJz3BaBwCN6wKmj4Pe6IUl.QokZh6YC31znw59BwURnt6xZdPe0FbJftoHTRKV6tUrNQ.R0MbANQDYnzvpuwP2dRbbC86JLtUeLw4R4M6Ret6S6ePe4FD9.IPOGFvWNRowBtiNBFvRStCT05neBRDQorq18UcwFJqQqK4y.JGVrKcrMuqWoKPm9RG13WOEmrugZrx7nqaA0raAlbG";if($_ts.lcd)$_ts.lcd();</script><script type="text/javascript" charset="utf-8" src="/LIWghRUPB4QK/oxYRCeR1jjgO.199cf1b.js" r='m'></script><script language="javascript" src="/_js/sudy-jquery-autoload.js" jquery-src="/_js/jquery-1.9.1.min.js" sudy-wp-context="" sudy-wp-siteId="68"></script>
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
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/c9/c6/3e2397ac4208b1357fc8e5380a5c/52d75459-d91c-4da6-bdec-bda3f2c88d79.jpg);"></div>
                          <p><a href='/2024/0412/c11624a335984/page.htm' target='_blank' title='计算机学院邀请南京大学陈道蓄教授来校作名师报告'>计算机学院邀请南京大学陈道蓄教授来校作名师报告</a></p>
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
                        
                        <li><a href='/2024/0523/c10846a340746/page.htm' target='_blank' title='中国电科第三十研究所来访我院'>中国电科第三十研究所来访我院</a><i>05-23</i></li>
                        
                        <li><a href='/2024/0521/c10846a340419/page.htm' target='_blank' title='计算机学院、数学学院、经管学院、人文学院党委理论学习中心组举行联合学习会'>计算机学院、数学学院、经管学院、人文学院...</a><i>05-21</i></li>
                        
                        <li><a href='/2024/0507/c10846a338414/page.htm' target='_blank' title='我院学生获得ICPC全国邀请赛（武汉）季军'>我院学生获得ICPC全国邀请赛（武汉）季军</a><i>05-07</i></li>
                        
                        <li><a href='/2024/0502/c10846a338078/page.htm' target='_blank' title='计算机学院召开全体教职工大会暨本科教育教学审核评估培训动员会'>计算机学院召开全体教职工大会暨本科教育教...</a><i>05-02</i></li>
                        
                        <li><a href='/2024/0430/c10846a338021/page.htm' target='_blank' title='南京航空航天大学第八届信息安全技能竞赛成功举办'>南京航空航天大学第八届信息安全技能竞赛成...</a><i>04-30</i></li>
                        
                        <li><a href='/2024/0430/c10846a338014/page.htm' target='_blank' title='我院在校“航空杯”教职工乒乓球比赛和江苏省高校计算机学科第六届教职工乒乓球团体赛中荣获佳绩'>我院在校“航空杯”教职工乒乓球比赛和江苏...</a><i>04-30</i></li>
                        
                        <li><a href='/2024/0430/c10846a337908/page.htm' target='_blank' title='2024鲲鹏昇腾开发者大会南航专场站圆满落幕'>2024鲲鹏昇腾开发者大会南航专场站圆满落幕</a><i>04-30</i></li>
                        
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
                
                <li><a href="/2024/0521/c10851a340477/page.htm" target="_blank"><i>></i>【博士招生】2024年工程博士复试名单</a><span>05.21</span></li>
                
                <li><a href="/2024/0511/c10851a338960/page.htm" target="_blank"><i>></i>【博士招生】2024年工程类博士研究生招生工...</a><span>05.11</span></li>
                
                <li><a href="/2024/0416/c10851a336275/page.htm" target="_blank"><i>></i>【博士招生】2024年招收博士研究生“申请考...</a><span>04.16</span></li>
                
                <li><a href="/2024/0330/c10851a334694/page.htm" target="_blank"><i>></i>【硕士招生】2024年专业课笔试、综合面试安...</a><span>03.30</span></li>
                
                <li><a href="/2024/0323/c10851a333892/page.htm" target="_blank"><i>></i>【硕士招生】2024年复试名单</a><span>03.23</span></li>
                
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
