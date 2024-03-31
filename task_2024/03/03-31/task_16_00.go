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

<meta content="{qr0hIIIapPKk9aHPpjqXTkhpIJSLPKOtWZAvPcHkWg32S6n1wGc64qqqqr7r0Ddee59cll00ZqqqqqhOT3jCzdeOhwr0qqqqqr0YqV4lWEcvMEa9vamHNv8lxr1c64qqqq.AwKdrFgUjM1W0ezWDI337NIL_XY6jLkLTOSvqOPNZgWl4096~RphVxyGIqxzqQ66dEWe3KncOJlg3uPGHjrjmebU5jQgAVjanqmB0BZniGrR05O6dghE9AOPMdQdAo7aivJh3NXo5vhyWAy1tHQ8WLPqj.x4abbv5oW_aCfpHsEeLT9chKEJlq0o56JEliP1c4cFGz.A4VWz0hTAiclx0sPoCpERAG7P_yqxwdvPXdr.Q8LrD1q7LtSPK0qFVGRkeGt9l.sUgTwUxZscpcJDl9I19SkV3a8Ol4QDWuFpwNF6VGw6qjxPJ.Iu3zw0ldFrm4lVZcVuQBF69G86qbxfw.1u3Rh0rnlGGAcbVcIaJtIbenQOSqh2g.D13FtrA9crSRQGWJUuxUEbQqqqqqqh5yFY8ebn4NFqkt99qqqqr1k162QH3atBqKZc40iza1OcFQHLAnOrtViHFHpRaYkWK7yRaYqqqq!p7z,aac,amr,asm,avi,bak,bat,bmp,bin,c,cab,css,csv,com,cpp,dat,dll,doc,dot,docx,exe,eot,fla,flc,fon,fot,font,gdb,gif,gz,gho,hlp,hpp,htc,ico,ini,inf,ins,iso,js,jar,jpg,jpeg,json,java,lib,log,map,mid,mp4,mpa,m4a,mp3,mpg,mkv,mod,mov,mim,mpp,msi,mpeg,obj,ocx,ogg,olb,ole,otf,py,pyc,pas,pgm,ppm,pps,ppt,pdf,pptx,png,pic,pli,psd,qif,qtx,ra,rm,ram,rmvb,reg,res,rtf,rar,so,sbl,sfx,swa,swf,svg,sys,tar,taz,tif,tiff,torrent,txt,ttf,vsd,vss,vsw,vxd,woff,woff2,wmv,wma,wav,wps,xbm,xpm,xls,xlsx,xsl,xml,z,zip,apk,plist,ipa,labVEL3s7piCD.Zs7ridmMapt1083179048`qX.lvNxqsrxFuWfsKYAAfyR0YO9VWfO52BHU5zsNcOQ1wWaTu6.3B7d5CDPRI0e4bhGH5x6gDs_QIleXuP0HXJG4UIls4RZ4v4aHBJvZv_uIJY5us4WRZe5PouAI7NHPuk.ty2SSUk0swp8bmjKYLmRaDh6xgNY91OBIWr5CO4RRep52UuHIzQ50bIpINxZnv_7Hjev.v_8HXRKNco08xA.NnXhKIzRdVH0K_RZNlt5WiNxBC85IBeZLU_yH.fvX640ItS5zO_TR4r5ZUn.I.JHukcaU3pFPbDn83yu0c.xRhxr7oFHR.2uuPPCIxz5Gu_8RyT56vnDING5bC8FIzzZOU_YHNpv26_fHzer9mPm3oRdw0908URqW6wUVb9jF9SHWsJERDRvInre8oePHCmbIoZeIKwdIue_Rv2divTPI6piJsmqF6ypwKL286RSYs0i32JgRnznFppVrn2eIYxdpkeoR9Rdr6TRI0ldADRkISSepoe8HTTbKoeLHSpnUPNQRoxFY9YVUGYCV0lH12RAD1LDWUfEhK3LIC2eW6Z.HnRbxUebIsydEkZORPxd3602IuTiIOlgxuNWWKqGxDA.h0rNFAJQkm01tGx1Ic9lIpYdksZWRGed1o0AIaLdcK3lI0Yeo6ZoHaxbqUZuH0TeHvLv8rT0mvlQpfNjrvL2Qqxdcoa4WVmElb3pITxecvd7H5eb4v5BIImd_sdLRXzdNo6ZI5RiNuoPAiN5nYoUKQ26C6upEzz_CbUVxBmnz1vkI8SdeOdTRBrd7U6.IdWdubMJIgyeTvdEHZzbnvd_MLpEZv6IIZQdabMrIgESvuDQW39O9k5psQTNSOBVI7z0TkjUIQenCO6RwZGe6UCwH_yb.U5IHyfe7U5NH_fbj6d9WI9SyOouE5gd406bR_x4X6KOM_R_7vCL3hSaXb59QtpGXoCy8_yn7v1f3jpB.CX3RyTO2b5s33Te9K5K8yyGGU_Y3gxjObXQMgxePC5q83SZ2KciIxx9fDcxI3mOP6CJQNfj2KFH8greOvFcIRwy8s0B8vwT3ONfQoq_RUL2iDRfRUfUimr_8UffiDrf36GZJPlCHO3.xk3gJ0L5w6RSI6f23vRT8Kf.QUp7w6LCi9yfAoVWQVwfcvwhJGN_Vvfl1AYHpDmo8qr_Kofciq09lOGshqECouRjHlmyUbyYFlSGKbfHwYRL1vTH3lzTVCQN8sTPhUqaQsRPED3_3OYSxvWG8PeLHoaZFkNeEoq43sy_JoGX3nW_HvQN3nydhba7FcyXHDWvMDyLxUQPMkSnrbGRQSy7kvWEHVz_l6VDiGTaokLx899gqblr8kWgJoLm8r0goDlI8mGCpsRFJ9Z9KOfM12xCqUaUQqmdcbaQHFy_eoUOiXRaZs.b8jAg_Ko98JAg0U.58iqggCoP8tWCLkFeJjL97sPC1.Tn_vhCwd0ez6.BwIL_5UkLiBT_BsOi3QTgduOg8zmgnsOE3QzgfkFRJL3C9n1iiEetfvoq8WSZ0D6DQQrGvs.M3wwTbUoQ3Rw6TkFU872g0Vb1semf.oOz8.wg.uObxiQgycu6JXAC2unaQ.zyjK8SQFwS76bNQhrG7CP5QMfTzCP0IMfazC8a8MeGjKCGQ.fTB6t4RySZ9CPm8gTu26nWHQf_9vKxiLSavujE87ggGCDc8hEgj6jl8xggvKD38J9CSOMMJ7q9CuO31zJXbo8iFET5voeZ3KeG36pdICNGRop_MCJzFOav86e0FOzNQUG4MU7ziUyGIU2HiVf43U2yiUfG8690Jn96JOQOxsgyH07CwoRvACwqQT3gAbRUQlf61bSMwlSgADRRIlp4UC2Ww9JNsvRV8lT0V6L8QmYboozkwm34UUmWiT24Ds213mmy3u298uYyhs973ceyFkwTJCL6InYjiKzhiKl5H1faEClnHDm4xUAaiuwSxO7686QyiDq48AVyqv7B8sQyWbqS8UE6MuwWJT92Ak2r10S7oK9139g4rbqEwlpakCrcwGW6Av7F8pSEUVgQ3reyKo7q8pLyY1YFiY7z1OwUJnEemULhRfNN4D6eRBzbeokbFIrveP1Xi5RG_o19QieGLvMOJ_x45vvZ1hTidDkd8Mf4gov_iMV2XOb4hFL6zuIcH8yN4U65RZVZO6BUwWE4aUoAiZY4fsvD3WYyeuva8e2yasvq3WJyvk8rJyg6PnKKiRwhvvhJFwe76CvwMQwnnUss8wwNCD8fQMp0zvcdQjRCzKbGMhzuBKtZFFSf7CDaQ_SCX6b4MixNz61S3hef.KbuwhYGXCDfihS4NvcdQ4m2LUcZIhJyvoDMMRNnGDbqMRYzbvDswLz02UDkRNGZavX3wEZ4PvDmiyx40kbw3xfyjkbz8gxy2kbU3xRyTuIWJel6SPUwiPxEtKSbwUrjMvfawCwjQ6fewCfj8USLIor7FKy7HnmgRUwaiownROf68kQ_8DJ489V_AvfB8vQ_QbJS8OE0huAWJr9bqkLr1qSZpvNHwmruAKyq8mRZ1oNmwmmZsUSUF2NTKDRsRVrjUvYURVSP16SMF2r2UCN1RlrdVbxJwVJdhCxuFPzyJnQgInNPivZNRspdxDVCInw6hbL9wnmjxCZZ31m2hvgv8sYjWba0MPwPQsG63uz6wUWg36JTWkpu8rY_kVzisATOcoNx8rJ_cuNtx97_1cyMJlq0WugVRpY4kDacRpw9rUZUQa27rbLJIArPrKa1MfSjrvgFIGx7DD._QiJ4dbuvFBTO4vHnw52SZDUTMdyu_vBaMdy24D.CI8Ag_6hXMXN54oUzFXyBd647FXm2dKdvMXeSzkOa3dy6OvHm3gSTus1E8Jr_n9XhsQROSUjD8Jy_akjHxLa_O1_qJx004k4kI7T26oUmFzx_6U.3wxWgTKHwwyJuCbipwFg04UP68jeh_lC.3_N__6Pe8jV__cMNiXgXdssPJxl5BUKZwhejBvigw4wB.udy8tT6XkPdQjggB6Odi.rO76jei.Ng26jHiLNOPU.tJJQ0Skspx7G_nVOlILf99bnI8e9gvCisweTu9DHMwRG0C6PJ8NYhT9CY3yr_0UPh8NW4EPQOinqBWuVLJ17eto2nwb9yMoJdiKfyhOzO3bx4VOzh8Kp4IOze3vf4isVbJs7TEcQ4iufxWCp7wvynW6xvR0SgV6RKI9SyKDxl3YYGpbYmF0fnKofi8Vz5A6Ek80yTsoyew2q_1TzTYDr1QbzKQ9NgUDwxF0T5sbJEw2x9YvNNi1ePE6RzQPrPF6ALJKzyHU721bRWxKEG8Cmyi67Oi6EvQuzahCWT8kVhHur5iUW98nNvxvLeIOxuwk9a3OyCsvEm39S0ksYE8pr4c9NhsqRPrUzD8py4qkzHx0a4s1eqJY0ThkZMRSwTDU0tRrA5Yv9Wwf7yovWxipyyUk7p3XS46kBR8IR4_kB.3XN4juKgJFVT7P8Ni_yxdUM_wiNyyohOH4S0dvhSF4ay5DhCw.RndKt0wHVT7ovL8dxt70ui37T4OvvK8ZG4fn8ligVBbkKmJIAdnvBI3ef.TbsE8y9yC61Q8QRCTsvAQZWyTobFiewPTvBbi5wynoB1iemP.oBzJIET4uUZx5Q4ylvB8FrdXDX9I_fn7obL34x_jbi78.pnNv_eQtf4BoKPQjpONoXawFfgzojG84pT.otBxRS4PoKAQNJ.Po_oQEwOGbXlILz.GsvrFgRBPvBqRES52Ub1QLyeP6cEIRNTPvCAREpev1MtxeEy9btQwzrnbC0X8CZ_hvRCwDfmtVJLMKz0xoRZwD30t1GXJcZGHOLOi9gaQbS0Ror.iKrSFCN4RkY0MC2.Rvz.RoeT8oTvFDR9A6JR89rTU6zkRlRe16AtM2N91vTkIVp51vmiIYx41bps8TryUKNH8TR.1UNRwTS41vwm8V271UQpwTS.Av2i3OJGiKAOF124io9NM1zChbZaWKwCH69jJP2Zik3zwb00Hb0ewqemkfElpHx.qqqq"><!--[if lt IE 9]><script r='m'>document.createElement("section")</script><![endif]--><script type="text/javascript" charset="utf-8" src="/aJEZedzaoCnN/2VzpTbTUGHr3.dee59c7.js" r='m'></script><script type="text/javascript" r='m'>(function(){var _$I6=16,_$pR=[[13,8,2,5,7,1,13,13,10,3,6,12,11,1,3,9,6],[0,60,4,51,123,12,44,101,32,80,101,107,116,24,101,76,5,28,106,70,24,46,18,43,57,93,88,24,103,102,98,103,71,17,115,130,5,61,37,113,52,43,15,16,10,114,81,86,13,50,67,112,27,29,134,90,2,112,66,3,117,75,26,113,65,101,68,44,121,51,102,56,101,65,51,101,27,110,50,7,97,80,87,39,105,83,135,48,106,114,69,122,134,26,92,87,66,10,99,46,88,0,60,34,53,76,120,44,51,43,115,118,29,101,32,89,13,3,108,42,25,108,62,65,5,103,76,23,122,38,63,6,63,46,27,65,16,19,105,32,9,14,86,39,87,78,127,31,58,90,125,74,82,118,37,37,120,36,109,55,95,117,0,69,110,92,115,18,89,19,73,69,50,48,128,63,62,41,77,5,2,88,11,32,65,112,83,101,131,10,111,48,87,52,44,40],[7,5,30,8,16,6,6,1,25,31,3,30,18,17,33,3,5,24,34,20,0,31,14,23,9,20,9,11,13,7,9,2,18,9,12,6,6,7,27,17,20,28,3],[4,31,17,4,6,44,44,11,7,25,9,38,0,23,43,34,47,24,15,37,25,26,2,48,10,24,16,41,54,32,0,30,37,30,29,36,4,25,23,37,51,20,33,27,25,46,6,14,12,45,13,15,43,28,55,18,32,39,3,29,10,12,28,6,10,47,44,0,38],[18,1,47,19,19,36,47,21,40,14,7,8,7,2,17,24,11,10,32,33,7,28,45,30,6,44,14,23,11,30,22,0,27,33,11,39,8,27,10,28,5,28,37,9,5,36,13,20,17,2,27,43,25,16,7]];function _$Pb(_$Du,_$ph){return _$8o[_$o5[27]].abs(_$Du)%_$ph;}function _$ea(_$Du){var _$RT=_$Du.length;var _$Pb,_$lx=new _$kL(_$RT-1),_$oS=_$Du.charCodeAt(0)-97;for(var _$aD=0,_$w9=1;_$w9<_$RT; ++_$w9){_$Pb=_$Du.charCodeAt(_$w9);if(_$Pb>=40&&_$Pb<92){_$Pb+=_$oS;if(_$Pb>=92)_$Pb=_$Pb-52;}else if(_$Pb>=97&&_$Pb<127){_$Pb+=_$oS;if(_$Pb>=127)_$Pb=_$Pb-30;}_$lx[_$aD++ ]=_$Pb;}return _$wt.apply(null,_$lx);}function _$w9(_$Du){var _$RT=_$wt(96);var _$Pb=_$ea(_$Du).split(_$RT);for(var _$lx=0;_$lx<_$Pb.length;_$lx++ ){_$zw.push(Number(_$Pb[_$lx]));}}function _$1q(_$Du){var _$RT=_$wt(96);_$o5=_$ea(_$Du).split(_$RT);}function _$31(_$qm){_$qm[14]=_$uQ()-_$qm[_$Pb(_$xg(),16)];if(_$DX()-_$qm[_$Pb(_$Bf(),16)]){if(_$uQ()+_$w0()){_$cH(_$qm);}}_$9J(_$qm);_$qm[6]=_$DX()-_$qm[_$Pb(_$Bf(),16)];_$qm[_$Pb(_$uQ()-_$qm[_$Pb(_$xg(),16)],16)]=_$qm[_$Pb(_$DX()+_$T$(),16)];_$qm[14]=_$uQ()-_$qm[_$Pb(_$xg(),16)];return _$qm[_$Pb(_$DX()-_$qm[_$Pb(_$Bf(),16)],16)];}function _$uQ(){return 5}function _$xg(){return 8}function _$DX(){return 13}function _$Bf(){return 0}function _$w0(){return 11}function _$cH(_$qm){var _$cA=_$T$();var _$cA=_$yG();_$qm[15]=_$uQ();var _$cA=_$hY();return _$Sv();}function _$T$(){return 3}function _$yG(){return 9}function _$hY(){return 1}function _$Sv(){return 7}function _$ha(){return 2}function _$0z(){return 12}function _$km(_$qm){_$qm[8]=_$Ug();_$qm[_$Pb(_$DX(),16)]=_$T$();_$qm[9]=_$4I();return _$uQ();}function _$Ug(){return 6}function _$4I(){return 15}function _$9J(_$qm){_$qm[1]=_$qm[_$Pb(_$cz(),16)];_$qm[_$Pb(_$qm[_$Pb(_$Bf(),16)],16)]=_$Z3(_$qm);var _$$Z=_$w0();var _$$Z=_$hY();var _$$Z=_$DX();if(_$yG()+_$4I()){_$wj(_$qm);}_$qm[_$Pb(_$uQ()+_$w0(),16)]=_$7F(_$qm);return _$rp(_$qm);}function _$cz(){return 4}function _$Z3(_$qm){_$qm[_$Pb(_$uQ(),16)]=_$w0();_$qm[_$Pb(_$Ug(),16)]=_$cz();_$qm[2]=_$Bf();return _$nZ();}function _$nZ(){return 14}function _$wj(_$qm){var _$cA=_$Sv();var _$cA=_$DX();var _$$Z=_$nZ();var _$$Z=_$0z();_$qm[10]=_$xg();return _$Ug();}function _$n1(_$qm){var _$cA=_$yG();var _$cA=_$4I();_$qm[_$Pb(_$VQ(),16)]=_$xg();var _$$Z=_$DX();var _$$Z=_$T$();return _$yG();}function _$VQ(){return 10}function _$7F(_$qm){_$qm[4]=_$ha();_$qm[_$Pb(_$yG(),16)]=_$4I();var _$$Z=_$xg();var _$$Z=_$Ug();return _$cz();}function _$rp(_$qm){_$qm[_$Pb(_$yG(),16)]=_$4I();var _$$Z=_$xg();var _$wT=_$Ug();if(_$DX()){_$qm[3]=_$yG();}return _$qm[_$Pb(_$xg(),16)];}var _$o5,_$8o,_$fT,_$rJ,_$31,_$GL,_$wt,_$KJ,_$mk,_$Kw,_$zw,_$kL;var _$k5,_$hS,_$RT=_$I6,_$aD=_$pR[0];while(1){_$hS=_$aD[_$RT++];if(38+_$hS>41&&_$hS+126<134){if(_$hS+128===134){_$vY(148);_$RT=0;}else if(76-_$hS===71){_$w9('}Hzhhhh`JMN`J`IHJL`NMMKN`LHQN`O`IN`HFM`L`IHH`NL`INKPL');}else if(_$hS*31===124){_$o5=[],_$zw=[],_$wt=String.fromCharCode;}else{_$GL=[_$zw[9],_$zw[7],_$zw[11],_$zw[1],_$zw[3],_$zw[5],_$zw[12],_$zw[4]];}}else if(116-_$hS>112){if(_$hS+62===64){_$vY(0);}else if(111-_$hS===110){_$1q('b|`$_sr`+`q~mcnl`rtarsq`du~k`(z`cdd48b6`|dkrd he[`<`bg~qBncd@s`Z\\q\\m\\r]`qdok~bd`vghkd[0(z`[`fdsShld`nmqd~cxrs~sdbg~mfd`Lhbqnrnes-WLKGSSO`;`9`rdmc`rokhs`nodm`eqnlBg~qBncd`[etmbshnm[(zu~q `@bshudWNaidbs`rvhsbg[`L~sg`dwdbRbqhos`he[`f`:`etmbshnm `<$_sr-rbi+`aqd~j:`-~ookx[sghr+`rbqhosr`qd~cxRs~sd`<$_sr-~dah:`([(:`tmrghes`|dkrdz`]:@qq~x-oqnsnsxod-otrg-~ookx[`~dah`**]:`etmbshnmdu~k[(zZm~shudbncd]|`qdronmrdSdws`b~kk`b~rd `u~q `snRsqhmf`Z`+~qftldmsr(:qdstqm `eknnq`WLKGssoQdptdrs`[(zu~q ');}else if(_$hS*90===0){_$k5= !_$fT;}else{_$fT=_$8o[_$o5[1]]={};}}else if(18-_$hS<11&&50-_$hS>38){if(_$hS+107===117){_$RT+=5;}else if(31-_$hS===22){_$kL=Array;}else if(_$hS*84===672){return;}else{_$fT=_$8o[_$o5[1]];}}else{if(_$hS+118===131){_$RT+=-5;}else if(63-_$hS===51){_$8o=window,_$mk=String,_$kL=Array,_$KJ=document,_$Kw=Date;}else{if( !_$k5)_$RT+=1;}}}function _$vY(_$cA,_$Du){function _$jA(){return _$2v[_$o5[10]](_$qU++ );}var _$hS,_$aD,_$qm,_$lx,_$r$,_$F8,_$2v,_$74,_$lh,_$qU,_$YK,_$$9,_$1q,_$w9,_$Pb,_$I6,_$RT,_$k5,_$oS;var _$wT,_$xg,_$$Z=_$cA,_$DX=_$pR[1];while(1){_$xg=_$DX[_$$Z++];if(_$xg+76<140){if(102+_$xg>117&&_$xg+40<72){if(92+_$xg>111&&_$xg+68<92){if(_$xg+50===72){_$qU+=_$qm;}else if(103-_$xg===82){_$$Z+=-13;}else if(_$xg*120===2400){_$Du._$KQ="_$hS";}else{_$Du[_$vY(146,_$vY(161))]=_$vY(126);}}else if(_$xg+37>52&&50-_$xg>30){if(_$xg+68===86){return _$vY(173);}else if(15-_$xg===-2){var _$74=_$jA();}else if(_$xg*23===368){var _$YK=_$fT._$wh;}else{return 8;}}else if(121-_$xg<98&&83-_$xg>55){if(_$xg+117===143){return _$oS;}else if(106-_$xg===81){_$RT+="a$3wAD5SHwHYcWhiSTEfWR7Cv0ArhvRxWcdI8QtH$1BQwNNuVybf5OrZxP_M45Wq$EybotfkOAKIjdzi_2N1BdVO86GRYCta2oHZmRHGW8yBpNjxCmayhT7M";}else if(_$xg*44===1056){_$Du._$1z="_$EM";}else{_$vY(73,_$1q);}}else{if(_$xg+118===148){var _$RT;}else if(20-_$xg===-9){_$RT+="RvKvYRZ5Wln1qw9PbI6RTk5oShSaDqmlxr$cA$ZwTuQxgDXBfw0cHT$yGhYSvha0zkmUg4I9JczZ3nZwjn1VQ7FrpzBZpBv7$fEMNG";}else if(_$xg*53===1484){var _$I6=[];}else{return 0;}}}else if(21-_$xg>5){if(35+_$xg>38&&_$xg+26<34){if(_$xg+60===66){_$RT=_$8o[_$o5[28]](_$Du);}else if(115-_$xg===110){_$hS=_$vY(162);}else if(_$xg*104===416){_$wT=_$r$>0;}else{_$wT=_$fT[_$o5[7]];}}else if(111-_$xg>107){if(_$xg+23===25){_$fT._$3v=_$vY(21)-_$RT;}else if(20-_$xg===19){_$Du._$v2="raGwlrnhz3a";}else if(_$xg*44===0){return 15;}else{for(_$k5=0;_$k5<_$r$;_$k5++ ){_$I6.push(_$o5[0]);}}}else if(106-_$xg<99&&17-_$xg>5){if(_$xg+38===48){return 5;}else if(55-_$xg===46){_$fT._$wh=_$vY(15);}else if(_$xg*68===544){var _$Pb=_$vY(66);}else{_$Du._$1M="_$Ra";}}else{if(_$xg+79===93){var _$$9=_$jA();}else if(75-_$xg===62){_$$Z+=1;}else if(_$xg*117===1404){_$RT+="uYo58ofTrJ31GLwtKJmkKwzweakLDuphjAaHF82v74lhqUYK$9_G$fkvuxTZ4SpgqSzZpVnTT0znXqmEutZIWQzu3pUv6W1MhFpm7_p";}else{var _$w9=_$jA();}}}else if(77-_$xg<46&&50-_$xg>2){if(107+_$xg>142&&_$xg+118<158){if(_$xg+93===131){_$Du[0]=_$vY(195);}else if(125-_$xg===88){var _$aD=_$jA();}else if(_$xg*78===2808){_$vY(127,_$Du);}else{return _$RT;}}else if(_$xg+58>89&&14-_$xg>-22){if(_$xg+59===93){var _$qU=0;}else if(106-_$xg===73){_$RT+="RAw3vdxZtsM_tc0ftMToXurRBOdxDZf6LQ8dPBS_Dz6xxyLefzTPUGA2UKGHEUDOfxeSN7tD_vcfCjDgoxWajTNPrDwtWXo1wRWQZ";}else if(_$xg*44===1408){return 13;}else{_$Du._$sk="bmzB0ZhypOXDgvvpyiu7rq";}}else if(40-_$xg<1&&91-_$xg>47){if(_$xg+121===163){_$RT+="4skV$KQuGs_1U5ody8nBtNjrcglJD18zPwJZHIxpND4UnDAYP2KNJ5_juMwVBngIu4xbGm_sYXCS5U09CWak$darS65wMAfuxoLPxY";}else if(108-_$xg===67){_$Du._$RB="_$gy";}else if(_$xg*93===3720){return 9;}else{_$RZ(0);}}else{if(_$xg+30===76){_$Du._$sm="f9qrggj.toq";}else if(79-_$xg===34){_$Du._$oX="_$4t";}else if(_$xg*81===3564){_$RT=_$RT[_$o5[12]](RegExp(_$o5[11],_$o5[30]),"");}else{return new _$Kw()[_$o5[15]]();}}}else{if(55+_$xg>106&&_$xg+51<107){if(_$xg+107===161){_$Du._$EP="_$1X";}else if(47-_$xg===-6){var _$oS=_$jA();}else if(_$xg*78===4056){_$Du._$0R=130;}else{_$aD=_$vY(137);}}else if(_$xg+32>79&&76-_$xg>24){if(_$xg+85===135){if( !_$wT)_$$Z+=1;}else if(62-_$xg===13){_$Du._$Me=3;}else if(_$xg*99===4752){return 10;}else{if( !_$wT)_$$Z+=2;}}else if(6-_$xg<-49&&90-_$xg>30){if(_$xg+88===146){_$hS=[];}else if(89-_$xg===32){return 0;}else if(_$xg*20===1120){_$RT+="Q6$jNB$F8N8rlbLShxdSh$zU$oqafixLbu2Hsa3r4X4lSb8yv6oPr8UaTs707bKoiFH9$lzfVLi5GznxYBtQGDcF2m103agsFWaE";}else{_$Du._$Ui=_$31;}}else{if(_$xg+26===88){_$Du[4]=_$vY(114);}else if(24-_$xg===-37){_$vY(85,_$fT);}else if(_$xg*96===5760){var _$lx=_$2v.length;}else{return _$vY(187);}}}}else if(_$xg+112>175&&123-_$xg>-5){if(87+_$xg>166&&_$xg+56<152){if(97+_$xg>180&&_$xg+117<205){if(_$xg+66===152){_$Du._$3q="_$nY";}else if(48-_$xg===-37){_$Du._$Qb="XKZr_h3LV2GIbUjDMBE2yV";}else if(_$xg*20===1680){_$RT=_$Pb[_$o5[47]](_$8o,_$Du);}else{var _$qm=_$jA()*_$zw[0]+_$jA();}}else if(_$xg+81>160&&11-_$xg>-73){if(_$xg+126===208){_$Du._$00=51;}else if(111-_$xg===30){_$Pb=_$8o[_$o5[5]];}else if(_$xg*48===3840){_$r$=_$jA();}else{_$RT+="jMRmpzloVAx4lt_0EIvuvPJWqN8SEUcv3ygeh4Fdovm_o31aRvghnp4tSD0OifoOvOYFQyF6qQoHVdEC$tSQAkt24BvtIkbwJw";}}else if(95-_$xg<8&&110-_$xg>18){if(_$xg+34===124){for(_$qm=0;_$qm<16;_$qm++)_$hS[_$qm]=1;}else if(80-_$xg===-9){_$RT+="6xIV$38WBRjENYerbRS2_xTP3YtzmSbcSXVzL5aQmZbkhnvU7IBBVnqn6Hu0KMWKnSadhNFcNUtFIE5LSpXsyiYGa1M1S7Dl3RIS";}else if(_$xg*23===2024){return 12;}else{for(_$k5=0;_$k5<_$r$;_$k5++ ){_$RZ(16,_$k5,_$I6);}}}else{if(_$xg+63===157){_$$Z+=13;}else if(127-_$xg===34){_$Du._$un="_$DX";}else if(_$xg*11===1012){_$Du[9]=_$vY(187);}else{return _$vY(122,_$Du);}}}else if(_$xg+50>113&&73-_$xg>-7){if(11+_$xg>78&&_$xg+32<104){if(_$xg+3===73){_$RT+="sm_aXjDcZR98b5whKs7Q3n5FlPOuB3WbMewPKQUiLz2lLxcU2ZabLMtuldyyexun5Jsk7cG_7u4KgydLP6nklgv2EMwnRaTyNFTUbm0";}else if(100-_$xg===31){return 3;}else if(_$xg*85===5780){_$Du._$Ty="_$wn";}else{var _$RT=_$8o[_$o5[5]][_$o5[50]]();}}else if(_$xg+54>117&&39-_$xg>-29){if(_$xg+6===72){_$Du._$ur="_$qa";}else if(6-_$xg===-59){_$Du[3]=_$vY(147);}else if(_$xg*30===1920){_$F8=_$2v[_$o5[4]](_$qU,_$qm)[_$o5[21]](_$mk[_$o5[23]](1));}else{var _$1q=_$I6.join('');}}else if(58-_$xg<-13&&43-_$xg>-33){if(_$xg+58===132){for(_$qm=0;_$qm<_$pR.length;_$qm++){_$aD=_$pR[_$qm];for(_$lx=0;_$lx<_$aD.length;_$lx++){_$aD[_$lx]^=_$hS[Math.abs(_$lx)%16];}}return;}else if(89-_$xg===16){var _$hS=_$vY(21);}else if(_$xg*74===5328){_$fT._$zB=_$vY(21)-_$RT;}else{_$Du._$7_="_$2l";}}else{if(_$xg+122===200){_$fT[_$o5[7]]=_$rJ;}else if(14-_$xg===-63){var _$RT='';}else if(_$xg*90===6840){return 7;}else{_$wT=_$Du===undefined||_$Du==="";}}}else if(12-_$xg<-83&&37-_$xg>-75){if(122+_$xg>221&&_$xg+12<116){if(_$xg+78===180){var _$2v=_$fT[_$o5[7]];}else if(17-_$xg===-84){return 14;}else if(_$xg*98===9800){return;}else{return _$vY(181);}}else if(_$xg+23>118&&45-_$xg>-55){if(_$xg+55===153){return _$vY(11,_$RT);}else if(103-_$xg===6){return 1;}else if(_$xg*96===9216){_$Du._$Ks="xdVX8PSHyYq";}else{var _$RT=_$vY(21);}}else if(8-_$xg<-95&&74-_$xg>-34){if(_$xg+3===109){var _$RT,_$Pb,_$lx=_$Du.length,_$oS=new _$kL(_$lx/_$zw[2]),_$aD='_$';}else if(78-_$xg===-27){_$RT+="d3bjt5lqh3kxd6SozEPSHD7t3EaaaMpdk3q83jFPfkO00dMrSl2wEiwZ6O4ze6NXr26ZDXNjX1Xs3K$HynYBs1zcKVumpQbuG7yV$Hq";}else if(_$xg*53===5512){_$$Z+=2;}else{_$Du._$O4="_$sM";}}else{if(_$xg+25===135){_$Du[2]=_$vY(163);}else if(85-_$xg===-24){_$Du[_$vY(146,_$vY(181))]=_$vY(189);}else if(_$xg*84===9072){_$wT=_$8o[_$o5[28]];}else{return 2;}}}else{if(114+_$xg>229&&_$xg+56<176){if(_$xg+57===175){_$vY(30);}else if(44-_$xg===-73){return 6;}else if(_$xg*30===3480){_$wT=_$RT!==_$o5[45];}else{_$Du._$MT="_$uY";}}else if(_$xg+90>201&&28-_$xg>-88){if(_$xg+96===210){_$Du._$w9="_$LM";}else if(41-_$xg===-72){for(_$RT=0,_$Pb=0;_$Pb<_$lx;_$Pb+=_$zw[2]){_$oS[_$RT++ ]=_$aD+_$Du[_$o5[4]](_$Pb,_$zw[2]);}}else if(_$xg*2===224){}else{_$Du[12]=_$vY(189);}}else if(15-_$xg<-104&&15-_$xg>-109){if(_$xg+73===195){_$Du[_$vY(146,_$vY(114))]=_$vY(190);}else if(91-_$xg===-30){_$Du._$nk="_$zw";}else if(_$xg*25===3000){_$vY(164,_$Du);}else{_$Du._$3k="_$Zt";}}else{if(_$xg+110===236){_$I6.push(_$o5[39]);}else if(110-_$xg===-15){_$vY(156,_$hS);}else if(_$xg*99===12276){_$Du._$26="_$tu";}else{return Math.abs(arguments[1]) % 16;}}}}else{if(_$xg+82===212){var _$lh=_$fT[_$o5[43]]=[];}else if(13-_$xg===-116){var _$r$=_$jA();}else if(_$xg*121===15488){_$Du._$7Q=189;}else{_$Du._$cz="_$cK";}}}function _$RZ(_$I6,_$_G,_$$f){function _$kv(){var _$oS=[0];Array.prototype.push.apply(_$oS,arguments);return _$5W.apply(this,_$oS);}var _$qS,_$zZ,_$pV,_$nT,_$T0,_$zn,_$Xq,_$mE,_$ut,_$ZI,_$1q,_$w9,_$Pb,_$TZ,_$4S,_$pg;var _$k5,_$hS,_$RT=_$I6,_$aD=_$pR[2];while(1){_$hS=_$aD[_$RT++];if(_$hS+45<61){if(92+_$hS>99&&_$hS+40<52){if(_$hS+44===54){var _$ZI=_$RZ(11);}else if(22-_$hS===13){var _$Xq=_$jA();}else if(_$hS*109===872){}else{_$k5=_$Pb;}}else if(_$hS+85>88&&66-_$hS>58){if(_$hS+61===67){_$RT+=-33;}else if(123-_$hS===118){var _$T0=_$RZ(11);}else if(_$hS*97===388){var _$Pb=_$w9>1?_$KJ[_$o5[36]][_$w9-_$zw[2]].src:_$rJ;}else{if( !_$k5)_$RT+=4;}}else if(61-_$hS>57){if(_$hS+49===51){return;}else if(44-_$hS===43){_$zn=_$8o[_$o5[25]]?new _$8o[_$o5[25]](_$o5[17]):new _$8o[_$o5[54]]();}else if(_$hS*107===0){_$zn[_$o5[22]]('GET',_$Pb,false);}else{var _$w9=_$RZ(11);}}else{if(_$hS+15===29){var _$w9=_$KJ[_$o5[36]].length;}else if(48-_$hS===35){var _$nT=_$jA();}else if(_$hS*19===228){var _$zn=_$jA();}else{var _$qS=_$RZ(11);}}}else if(_$hS+69>84&&16-_$hS>-16){if(62+_$hS>85&&_$hS+121<149){if(_$hS+40===66){var _$mE=_$RZ(11);}else if(103-_$hS===78){var _$TZ=_$jA();}else if(_$hS*116===2784){_$zn[_$o5[16]]=_$kv;}else{_$lh[_$_G]=_$w9;}}else if(_$hS+93>112&&38-_$hS>14){if(_$hS+46===68){var _$pV=_$jA();}else if(27-_$hS===6){for(_$1q=0;_$1q<_$Pb;_$1q++ ){_$pg[_$1q]=_$RZ(11);}}else if(_$hS*57===1140){var _$4S=_$jA();}else{var _$Pb=new _$kL(_$w9);}}else if(97-_$hS<82&&81-_$hS>61){if(_$hS+121===139){_$5W(7,_$$f);}else if(7-_$hS===-10){_$zn[_$o5[20]]();}else if(_$hS*123===1968){var _$pg=[];}else{for(_$1q=0;_$1q<_$w9;_$1q++ ){_$Pb[_$1q]=_$jA();}}}else{if(_$hS+55===85){var _$ut=_$jA();}else if(107-_$hS===78){_$RT+=33;}else if(_$hS*4===112){var _$Pb=_$jA();}else{var _$w9=_$jA();}}}else{if(_$hS+77===109){return _$Pb;}else{var _$zZ=_$jA();}}}function _$5W(_$Pb,_$WQ){var _$1q,_$w9,_$3p,_$Uv;var _$RT,_$oS,_$I6=_$Pb,_$hS=_$pR[3];while(1){_$oS=_$hS[_$I6++];if(52-_$oS<21&&52-_$oS>4){if(92-_$oS<53&&95-_$oS>51){if(_$oS*21===840){var _$w9=[];}else if(_$oS*38===1558){if( !_$RT)_$I6+=10;}else if(_$oS+8===50){_$WQ.push(_$YK[_$$9]);}else{return;}}else if(16-_$oS<-19&&30+_$oS<70){if(_$oS*121===4356){_$WQ.push(_$o5[49]);}else if(_$oS*7===259){_$WQ.push(_$o5[32]);}else if(_$oS+107===145){_$WQ.push(_$YK[_$74]);}else{_$WQ.push(_$o5[24]);}}else if(67+_$oS>98&&_$oS+31<67){if(_$oS*18===576){_$RT=_$ZI.length;}else if(_$oS*19===627){_$RT=_$pg.length;}else if(_$oS+50===84){for(_$1q=0;_$1q<_$w9.length;_$1q++ ){_$ln(0,_$w9[_$1q][0],_$w9[_$1q][1],_$WQ);}}else{_$WQ.push(_$o5[33]);}}else{if(_$oS*97===4268){_$WQ.push(_$o5[2]);}else if(_$oS*94===4230){_$vY(30);}else if(_$oS+56===102){_$Uv=_$pg.length;}else{_$ln(10,0,_$pg.length);}}}else if(100-_$oS<85&&35+_$oS<67){if(78-_$oS<55&&3-_$oS>-25){if(_$oS*114===2736){for(_$1q=1;_$1q<_$ZI.length;_$1q++ ){_$WQ.push(_$o5[2]);_$WQ.push(_$YK[_$ZI[_$1q]]);}}else if(_$oS*106===2650){_$WQ.push(_$o5[38]);}else if(_$oS+103===129){_$I6+=-10;}else{_$WQ.push(_$o5[14]);}}else if(53-_$oS<34&&43+_$oS<67){if(_$oS*27===540){_$WQ.push(_$YK[_$nT]);}else if(_$oS*4===84){_$WQ.push("=0,");}else if(_$oS+11===33){for(_$1q=0;_$1q<_$mE.length;_$1q++ ){_$WQ.push(_$o5[2]);_$WQ.push(_$YK[_$mE[_$1q]]);}}else{_$RT=_$_G==0;}}else if(99+_$oS>114&&_$oS+6<26){if(_$oS*75===1200){_$WQ.push(_$o5[31]);}else if(_$oS*100===1700){_$WQ.push(_$o5[6]);}else if(_$oS+13===31){_$vY(73,_$zn[_$o5[46]]);}else{_$WQ.push("];");}}else{if(_$oS*34===952){_$WQ.push(_$YK[_$Xq]);}else if(_$oS*95===2755){_$WQ.push(_$o5[51]);}else if(_$oS+37===67){if( !_$RT)_$I6+=4;}else{_$I6+=10;}}}else if(_$oS+13<29){if(26-_$oS<19&&87-_$oS>75){if(_$oS*93===744){if( !_$RT)_$I6+=8;}else if(_$oS*71===639){_$WQ.push(_$YK[_$ZI[0]]);}else if(_$oS+9===19){var _$1q,_$3p=_$zw[9];}else{if( !_$RT)_$I6+=1;}}else if(112-_$oS<109&&37+_$oS<45){if(_$oS*124===496){_$RT=_$fT[_$o5[7]];}else if(_$oS*2===10){_$WQ.push(_$YK[_$TZ]);}else if(_$oS+126===132){var _$Uv=0;}else{_$WQ.push(_$YK[_$zn]);}}else if(_$oS+75<79){if(_$oS*86===0){_$WQ.push(_$o5[0]);}else if(_$oS*84===84){_$ln(49);}else if(_$oS+41===43){_$WQ.push(_$o5[13]);}else{_$WQ.push(_$o5[44]);}}else{if(_$oS*47===564){_$I6+=8;}else if(_$oS*65===845){_$RT=_$zn[_$o5[37]]==_$zw[9];}else if(_$oS+12===26){_$WQ.push(_$YK[_$pV]);}else{_$WQ.push(_$o5[9]);}}}else{if(_$oS*85===4080){_$WQ.push(_$_G);}else if(_$oS*108===5292){_$RT=_$mE.length;}else if(_$oS+125===175){_$WQ.push(_$YK[_$zZ]);}else{for(_$1q=0;_$1q<_$T0.length;_$1q+=_$zw[2]){if(_$zw[8]<Math[_$o5[3]]()){_$w9.push([_$T0[_$1q],_$T0[_$1q+1]]);}else{_$w9[_$o5[40]]([_$T0[_$1q],_$T0[_$1q+1]]);}}}}}function _$ln(_$lx,_$6W,_$1M,_$hF){var _$hS,_$aD,_$qm,_$1q,_$w9,_$Pb,_$I6,_$RT,_$k5,_$oS;var _$cA,_$wT,_$r$=_$lx,_$uQ=_$pR[4];while(1){_$wT=_$uQ[_$r$++];if(_$wT+98<114){if(59+_$wT>62&&_$wT+113<121){if(_$wT+121===127){var _$RT,_$hS,_$aD,_$Pb=_$1M-_$6W;}else if(22-_$wT===17){_$aD=0;}else if(_$wT*21===84){for(;_$6W+_$aD<_$1M;_$6W+=_$aD){_$WQ.push(_$hS);_$WQ.push(_$YK[_$Xq]);_$WQ.push(_$o5[18]);_$WQ.push(_$6W+_$aD);_$WQ.push(_$o5[6]);_$ln(10,_$6W,_$6W+_$aD);_$hS=_$o5[8];}}else{_$qm=Math[_$o5[53]]((Math[_$o5[3]]()*_$zw[10])+1);}}else if(67-_$wT>63){if(_$wT+41===43){if( !_$cA)_$r$+=2;}else if(87-_$wT===86){for(_$RT=0;_$RT<_$Pb-1;_$RT++ ){if(_$RT==_$oS){_$I6=_$1M;if(_$6W>1&&_$qm%_$zw[2]==0){_$I6=_$6W-1;}_$WQ.push(_$hS);_$WQ.push(_$YK[_$Xq]);_$WQ.push(_$1q);_$WQ.push(_$I6);_$WQ.push(_$o5[6]);_$ln(2,_$qm%_$Uv);_$hS=_$o5[8];}_$WQ.push(_$hS);_$WQ.push(_$YK[_$Xq]);_$WQ.push(_$1q);_$WQ.push(_$k5[_$RT]);_$WQ.push(_$o5[6]);_$ln(2,_$k5[_$RT]);_$hS=_$o5[8];}}else if(_$wT*24===0){return;}else{var _$RT=_$qS.length;}}else if(37-_$wT<30&&127-_$wT>115){if(_$wT+30===40){if( !_$cA)_$r$+=15;}else if(44-_$wT===35){_$cA=_$Pb==0;}else if(_$wT*114===912){_$WQ.push(_$o5[0]);}else{_$ln(2,_$6W);}}else{if(_$wT+64===78){_$k5[_$RT]=_$w9;}else if(40-_$wT===27){}else if(_$wT*102===1224){_$r$+=4;}else{_$w9=_$k5[0];}}}else if(_$wT+121>136&&73-_$wT>41){if(111+_$wT>130&&_$wT+45<69){if(_$wT+25===47){_$k5[0]=_$k5[_$RT];}else if(83-_$wT===62){_$RT-=_$RT%_$zw[2];}else if(_$wT*88===1760){if( !_$cA)_$r$+=1;}else{_$oS=_$qm%_$3p;}}else if(_$wT+27>42&&14-_$wT>-6){if(_$wT+117===135){_$WQ.push(_$o5[41]);}else if(46-_$wT===29){_$WQ.push(_$F8[_$qS[_$RT]]);}else if(_$wT*37===592){_$r$+=29;}else{var _$hS=_$RT.length;}}else if(10-_$wT<-13&&106-_$wT>78){if(_$wT+79===105){for(_$RT=1;_$RT<_$zw[6];_$RT++ ){if(_$Pb<=_$GL[_$RT]){_$aD=_$GL[_$RT-1];break;}}}else if(127-_$wT===102){_$cA=_$Pb==1;}else if(_$wT*7===168){_$r$+=-4;}else{_$hF.push([_$o5[32],_$YK[_$6W],_$o5[55],_$YK[_$4S],"=[",_$1M,_$o5[42],_$YK[_$4S],_$o5[52],_$YK[_$ut],_$o5[35],_$YK[_$4S],");}"].join(''));}}else{if(_$wT+21===51){_$hS-=_$hS%_$zw[2];}else if(37-_$wT===8){_$hS=_$o5[29];}else if(_$wT*58===1624){_$1q="===";}else{for(_$RT=0;_$RT<_$Pb;_$RT++ ){_$k5[_$RT]=_$6W+_$RT;}}}}else{if(92+_$wT>127&&_$wT+66<106){if(_$wT+88===126){for(_$hS=0;_$hS<_$RT;_$hS+=_$zw[2]){_$WQ.push(_$F8[_$qS[_$hS]]);_$WQ.push(_$YK[_$qS[_$hS+1]]);}}else if(48-_$wT===11){for(_$aD=0;_$aD<_$hS;_$aD+=_$zw[2]){_$WQ.push(_$F8[_$RT[_$aD]]);_$WQ.push(_$YK[_$RT[_$aD+1]]);}}else if(_$wT*15===540){_$ln(10,_$6W,_$1M);}else{_$r$+=8;}}else if(_$wT+47>78&&66-_$wT>30){if(_$wT+33===67){_$RT=_$qm%_$Pb;}else if(107-_$wT===74){_$cA=_$Pb<=_$3p;}else if(_$wT*75===2400){_$ln(2,_$k5[_$RT]);}else{_$r$+=25;}}else if(36-_$wT<-3&&30-_$wT>-14){if(_$wT+82===124){_$k5=[];}else if(61-_$wT===20){_$WQ.push(_$F8[_$RT[_$hS]]);}else if(_$wT*23===920){_$cA=_$RT.length!=_$hS;}else{_$cA=_$qS.length!=_$RT;}}else{var _$RT=_$pg[_$6W];}}}}}}}})()</script><script language="javascript" src="/_js/sudy-jquery-autoload.js" jquery-src="/_js/jquery-1.9.1.min.js" sudy-wp-context="" sudy-wp-siteId="68"></script>
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
      
      <li><img src="/_upload/article/images/8e/5b/f9e805164267b723013414e8d7ca/e2397f85-7e06-4124-97ba-0ea7a35e1a05.png" /></li>
      
      <li><img src="/_upload/article/images/ee/3f/c23c09c74990b6c6bf72de29cc9c/535d31e3-7de2-485f-993a-811042e8da7f.png" /></li>
      
      <li><img src="/_upload/article/images/d6/28/c0272c7b4ef98fdcf3de3f3e8990/8d7ab9a8-25ac-4bb1-b7d9-edf26d58cc92.png" /></li>
      
      <li><img src="/_upload/article/images/20/35/0e7ca14f46ea844bd475df8ebf37/48f77e80-bf7c-474d-a504-5ea301c1b032.png" /></li>
      
      <li><img src="/_upload/article/images/74/5b/9dda4c0446fd93b0326d50a71b45/91732d13-5f6e-4d78-96ef-7b4c3c2f565b.png" /></li>
      
      <li><img src="/_upload/article/images/21/54/835e7c8d47a3ac33755a44a53147/a96926c2-fd26-4c54-b1e5-cace00df4e3c.jpg" /></li>
      
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
                          <div style="background-image:url(/_upload/article/images/ad/1e/723319cd4dfab53ebf23bd8c4a1b/c9e64f6b-1bc5-4d41-876f-073aefc948dd.jpg);"></div>
                          <p><a href='/2024/0315/c11624a333101/page.htm' target='_blank' title='中国电子科技集团第三十八研究所来我校调研交流'>中国电子科技集团第三十八研究所来我校调研交流</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/8e/8d/0803e99a4fffa1bcc7b4a4badfa9/24703103-b16e-4f85-880a-a7179a4c26bd.jpg);"></div>
                          <p><a href='/2024/0313/c11624a332826/page.htm' target='_blank' title='计算机学院召开党委（扩大）会暨2023年度教工党支部书记述职评议会议'>计算机学院召开党委（扩大）会暨2023年度教工党支部书记述职评议...</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/f7/c7/ce01ee414cb8938df2e15d69bf31/c6259543-2e90-4314-9bd2-f08eae92cbb4.jpg);"></div>
                          <p><a href='/2024/0301/c11624a331881/page.htm' target='_blank' title='信息安全与隐私保护前沿技术研讨会暨FCS青年编委会走进南航活动成功举行'>信息安全与隐私保护前沿技术研讨会暨FCS青年编委会走进南航活动...</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/48/34/5026a499413487c59afc048d13b5/0cea9a29-c77c-47e9-bcc6-4c5175e54a34.jpg);"></div>
                          <p><a href='/2023/1211/c11624a326983/page.htm' target='_blank' title='计算机学院程序设计课程组开展“乐享”专项课程建设活动'>计算机学院程序设计课程组开展“乐享”专项课程建设活动</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/bf/42/d382bce94bdf874bb30bb1ff16c6/e2be4433-f405-4508-8ec3-b1de83ae0f69.jpg);"></div>
                          <p><a href='/2023/1211/c11624a326958/page.htm' target='_blank' title='第七届江苏省计算机类专业建设与课程改革研讨会顺利举行'>第七届江苏省计算机类专业建设与课程改革研讨会顺利举行</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/f3/e7/a345bacd45a1ad3b3c884063dff8/3bc1db35-7317-49c1-90c3-5297fc276378.png);"></div>
                          <p><a href='/2023/1209/c11624a326887/page.htm' target='_blank' title='我院李鑫老师入选2022-2023年度“高校计算机专业优秀教师奖励计划”'>我院李鑫老师入选2022-2023年度“高校计算机专业优秀教师奖励计...</a></p>
                        </li>
                        
                        <li>
                          <div style="background-image:url(/_upload/article/images/ba/86/8cdd614e4b528be8b72aae91d822/df34ee89-02a8-4515-8f9b-4ed0cc1ae8d3.png);"></div>
                          <p><a href='/2023/1127/c11624a325630/page.htm' target='_blank' title='计算机学院举办“踔厉奋发新征程 凝心聚力‘毅’起向未来”教职工徒步活动'>计算机学院举办“踔厉奋发新征程 凝心聚力‘毅’起向未来”教职...</a></p>
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
                        
                        <li><a href='/2024/0331/c10846a334731/page.htm' target='_blank' title='清华大学出版社来计算机学院开展交流研讨'>清华大学出版社来计算机学院开展交流研讨</a><i>03-31</i></li>
                        
                        <li><a href='/2024/0331/c10846a334725/page.htm' target='_blank' title='计算机学院顺利开展本科教育教学审核评估预评估工作'>计算机学院顺利开展本科教育教学审核评估预...</a><i>03-31</i></li>
                        
                        <li><a href='/2024/0331/c10846a334730/page.htm' target='_blank' title='我校成功举行第十八届程序设计竞赛'>我校成功举行第十八届程序设计竞赛</a><i>03-30</i></li>
                        
                        <li><a href='/2024/0327/c10846a334151/page.htm' target='_blank' title='我院在程序设计语言领域顶级国际会议OOPSLA 2024发表论文'>我院在程序设计语言领域顶级国际会议OOPSLA...</a><i>03-27</i></li>
                        
                        <li><a href='/2024/0325/c10846a334018/page.htm' target='_blank' title='计算机学院举办“赓续红色血脉，坚定理想信念”主题升旗仪式'>计算机学院举办“赓续红色血脉，坚定理想信...</a><i>03-25</i></li>
                        
                        <li><a href='/2024/0321/c10846a333662/page.htm' target='_blank' title='计算机学院召开教育教学审核评估校内预评估工作启动会'>计算机学院召开教育教学审核评估校内预评估...</a><i>03-21</i></li>
                        
                        <li><a href='/2024/0320/c10846a333643/page.htm' target='_blank' title='访企拓岗促就业 | 计算机学院赴华为技术有限公司开展访企拓岗促就业专项行动'>访企拓岗促就业 | 计算机学院赴华为技术有...</a><i>03-20</i></li>
                        
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
                
                <li><a href='/2024/0109/c10847a329258/page.htm' target='_blank' title='关于自编讲义编写人员政治审查公示'>关于自编讲义编写人员政治审查公示</a><i>01-03</i></li>
                
                <li><a href='/2023/1227/c10847a328359/page.htm' target='_blank' title='预备党员转正公示'>预备党员转正公示</a><i>12-27</i></li>
                
                <li><a href='/2023/1215/c10847a327474/page.htm' target='_blank' title='预备党员转正公示'>预备党员转正公示</a><i>12-15</i></li>
                
                <li><a href='/2023/1214/c10847a327326/page.htm' target='_blank' title='2022-2023年度教学优秀奖拟推荐名单公示'>2022-2023年度教学优秀奖拟推荐名单公示</a><i>12-14</i></li>
                
                <li><a href='/2023/1129/c10847a325851/page.htm' target='_blank' title='计算机科学与技术学院/人工智能学院/软件学院 2023-2024学年第二学期本科生教材选用公示'>计算机科学与技术学院/人工智能学院/软件学...</a><i>11-29</i></li>
                
                <li><a href='/2023/1128/c10847a325769/page.htm' target='_blank' title='发展党员公示'>发展党员公示</a><i>11-28</i></li>
                
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
                
                <li><a href="/2024/0102/c10849a328638/page.htm" target="_blank"><i>></i>关于开展第十八届中国青年科技奖候选人提名...</a><span>01.02</span></li>
                
                <li><a href="/2023/1227/c10849a328354/page.htm" target="_blank"><i>></i>关于开展第二十届中国青年女科学家奖和第九...</a><span>12.27</span></li>
                
                <li><a href="/2023/1213/c10849a327172/page.htm" target="_blank"><i>></i>关于推荐全国科普工作先进集体和先进工作者...</a><span>12.13</span></li>
                
                <li><a href="/2023/1212/c10849a327155/page.htm" target="_blank"><i>></i>关于组织申报2024年度“江苏科技智库青年人...</a><span>12.12</span></li>
                
                <li><a href="/2022/0831/c10849a290883/page.htm" target="_blank"><i>></i>生物大分子动态修饰与化学干预重大研究计划...</a><span>08.31</span></li>
                
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
                
                <li><a href="/2023/1129/c10850a325842/page.htm" target="_blank"><i>></i>【毕业设计】2024届本科毕业设计工作通知</a><span>11.29</span></li>
                
                <li><a href="/2023/0920/c10850a319935/page.htm" target="_blank"><i>></i>【创新班考核】关于开展2021级、2022级人工...</a><span>09.20</span></li>
                
                <li><a href="/2023/0919/c10850a319894/page.htm" target="_blank"><i>></i>【推免工作】计算机科学与技术学院/人工智...</a><span>09.19</span></li>
                
                <li><a href="/2023/0916/c10850a319663/page.htm" target="_blank"><i>></i>【推免工作】计算机科学与技术学院/人工智...</a><span>09.16</span></li>
                
                <li><a href="/2023/0913/c10850a319313/page.htm" target="_blank"><i>></i>【本科生转专业】2023年特殊类型转专业工作...</a><span>09.13</span></li>
                
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
                
                <li><a href="/2024/0330/c10851a334694/page.htm" target="_blank"><i>></i>【硕士招生】2024年专业课笔试、综合面试安...</a><span>03.30</span></li>
                
                <li><a href="/2024/0323/c10851a333892/page.htm" target="_blank"><i>></i>【硕士招生】2024年复试名单</a><span>03.23</span></li>
                
                <li><a href="/2024/0323/c10851a333891/page.htm" target="_blank"><i>></i>【硕士招生】2024年硕士研究生招生复试及录...</a><span>03.23</span></li>
                
                <li><a href="/2024/0321/c10851a333724/page.htm" target="_blank"><i>></i>【优博优硕】拟推荐申报2024年校优秀博士、...</a><span>03.21</span></li>
                
                <li><a href="/2024/0321/c10851a333718/page.htm" target="_blank"><i>></i>【导师】2024年研究生导师招生资格补充公示</a><span>03.21</span></li>
                
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
                
                <li><a href="/2024/0325/c10852a334018/page.htm" target="_blank"><i>></i>计算机学院举办“赓续红色血脉，坚定理想信...</a><span>03.25</span></li>
                
                <li><a href="/2024/0320/c10852a333643/page.htm" target="_blank"><i>></i>访企拓岗促就业 | 计算机学院赴华为技术有...</a><span>03.20</span></li>
                
                <li><a href="/2024/0319/c10852a333394/page.htm" target="_blank"><i>></i> 计算机学院开展第十三期“山湖号”红色专...</a><span>03.19</span></li>
                
                <li><a href="/2024/0318/c10852a333263/page.htm" target="_blank"><i>></i>南京师范大学计电学院学工团队来我院调研</a><span>03.18</span></li>
                
                <li><a href="/2023/1213/c10852a327192/page.htm" target="_blank"><i>></i> 计算机学院举办“让诚信记心间，安全需牢...</a><span>12.13</span></li>
                
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
                <div class="date">2024-03-08</div>
              </div>
              <div class="title"> <a href='/2024/0307/c11617a332243/page.htm' target='_blank' title='【石榴大讲堂】信息论与编码理论'>【石榴大讲堂】信息论与编码理论</a>
                <p class="p1">题目：信息论与编码理论</p>
                <p class="p2">报告人：符方伟</p>
                <p class="p3">时间： <span class="t-time">2024-03-08</span> 15:00-16:00</p>
                <p class="p4">地点：计算机学院505会议室</p>
              </div>
            </div>
          </div>
          
          <div class="col xs-12 sm-4 md-4 lg-4">
            <div class="info-item">
              <div class="rili">
                <div class="date">2024-01-18</div>
              </div>
              <div class="title"> <a href='/2024/0113/c11617a329540/page.htm' target='_blank' title='【石榴大讲堂】全局和局部低秩假设下的缺失多标签学习'>【石榴大讲堂】全局和局部低秩假设下的缺失...</a>
                <p class="p1">题目：全局和局部低秩假设下的缺失多标签学习</p>
                <p class="p2">报告人：马忠臣</p>
                <p class="p3">时间： <span class="t-time">2024-01-18</span> 上午10:30-11:30</p>
                <p class="p4">地点：计算机学院507会议室</p>
              </div>
            </div>
          </div>
          
          <div class="col xs-12 sm-4 md-4 lg-4">
            <div class="info-item">
              <div class="rili">
                <div class="date">2024-01-27</div>
              </div>
              <div class="title"> <a href='/2024/0113/c11617a329539/page.htm' target='_blank' title='【石榴大讲堂】广义光滑非凸优化与光滑非凸优化一样高效'>【石榴大讲堂】广义光滑非凸优化与光滑非凸...</a>
                <p class="p1">题目：广义光滑非凸优化与光滑非凸优化一样高效</p>
                <p class="p2">报告人：周义</p>
                <p class="p3">时间： <span class="t-time">2024-01-27</span> 下午14:00-15:00</p>
                <p class="p4">地点：计算机学院507会议室</p>
              </div>
            </div>
          </div>
          
          <div class="col xs-12 sm-4 md-4 lg-4">
            <div class="info-item">
              <div class="rili">
                <div class="date">2024-01-05</div>
              </div>
              <div class="title"> <a href='/2024/0104/c11617a329012/page.htm' target='_blank' title='【石榴大讲堂】迈向更高层次智能化的生成式软件开发'>【石榴大讲堂】迈向更高层次智能化的生成式...</a>
                <p class="p1">题目：迈向更高层次智能化的生成式软件开发</p>
                <p class="p2">报告人：彭鑫</p>
                <p class="p3">时间： <span class="t-time">2024-01-05</span> 下午16：00</p>
                <p class="p4">地点：计算机学院507会议室</p>
              </div>
            </div>
          </div>
          
          <div class="col xs-12 sm-4 md-4 lg-4">
            <div class="info-item">
              <div class="rili">
                <div class="date">2024-01-07</div>
              </div>
              <div class="title"> <a href='/2024/0104/c11617a329002/page.htm' target='_blank' title='【石榴大讲堂】扩散概率模型及其应用'>【石榴大讲堂】扩散概率模型及其应用</a>
                <p class="p1">题目：扩散概率模型及其应用</p>
                <p class="p2">报告人：朱军</p>
                <p class="p3">时间： <span class="t-time">2024-01-07</span> 上午10:15</p>
                <p class="p4">地点：计算机学院楼507会议室</p>
              </div>
            </div>
          </div>
          
          <div class="col xs-12 sm-4 md-4 lg-4">
            <div class="info-item">
              <div class="rili">
                <div class="date">2024-01-06</div>
              </div>
              <div class="title"> <a href='/2024/0104/c11617a328965/page.htm' target='_blank' title='【石榴大讲堂】代码大模型的高效测试与即时修复'>【石榴大讲堂】代码大模型的高效测试与即时...</a>
                <p class="p1">题目：代码大模型的高效测试与即时修复</p>
                <p class="p2">报告人：陈俊洁</p>
                <p class="p3">时间： <span class="t-time">2024-01-06</span> 下午2:00</p>
                <p class="p4">地点：计算机学院511会议室</p>
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
