(function(e){function t(t){for(var n,s,l=t[0],i=t[1],u=t[2],c=0,m=[];c<l.length;c++)s=l[c],Object.prototype.hasOwnProperty.call(a,s)&&a[s]&&m.push(a[s][0]),a[s]=0;for(n in i)Object.prototype.hasOwnProperty.call(i,n)&&(e[n]=i[n]);p&&p(t);while(m.length)m.shift()();return o.push.apply(o,u||[]),r()}function r(){for(var e,t=0;t<o.length;t++){for(var r=o[t],n=!0,s=1;s<r.length;s++){var i=r[s];0!==a[i]&&(n=!1)}n&&(o.splice(t--,1),e=l(l.s=r[0]))}return e}var n={},a={app:0},o=[];function s(e){return l.p+"js/"+({about:"about"}[e]||e)+"."+{about:"e50fa587"}[e]+".js"}function l(t){if(n[t])return n[t].exports;var r=n[t]={i:t,l:!1,exports:{}};return e[t].call(r.exports,r,r.exports,l),r.l=!0,r.exports}l.e=function(e){var t=[],r=a[e];if(0!==r)if(r)t.push(r[2]);else{var n=new Promise((function(t,n){r=a[e]=[t,n]}));t.push(r[2]=n);var o,i=document.createElement("script");i.charset="utf-8",i.timeout=120,l.nc&&i.setAttribute("nonce",l.nc),i.src=s(e);var u=new Error;o=function(t){i.onerror=i.onload=null,clearTimeout(c);var r=a[e];if(0!==r){if(r){var n=t&&("load"===t.type?"missing":t.type),o=t&&t.target&&t.target.src;u.message="Loading chunk "+e+" failed.\n("+n+": "+o+")",u.name="ChunkLoadError",u.type=n,u.request=o,r[1](u)}a[e]=void 0}};var c=setTimeout((function(){o({type:"timeout",target:i})}),12e4);i.onerror=i.onload=o,document.head.appendChild(i)}return Promise.all(t)},l.m=e,l.c=n,l.d=function(e,t,r){l.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:r})},l.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},l.t=function(e,t){if(1&t&&(e=l(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var r=Object.create(null);if(l.r(r),Object.defineProperty(r,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var n in e)l.d(r,n,function(t){return e[t]}.bind(null,n));return r},l.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return l.d(t,"a",t),t},l.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},l.p="/starter-gin-devtools/",l.oe=function(e){throw console.error(e),e};var i=window["webpackJsonp"]=window["webpackJsonp"]||[],u=i.push.bind(i);i.push=t,i=i.slice();for(var c=0;c<i.length;c++)t(i[c]);var p=u;o.push([0,"chunk-vendors"]),r()})({0:function(e,t,r){e.exports=r("56d7")},"034f":function(e,t,r){"use strict";r("4b77")},1106:function(e,t,r){},"210a":function(e,t,r){},3021:function(e,t,r){"use strict";r("210a")},"4b77":function(e,t,r){},"50c8":function(e,t,r){},"56d7":function(e,t,r){"use strict";r.r(t);var n=r("430a"),a=r("edcc"),o=r.n(a),s=(r("0db3"),function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",{attrs:{id:"app"}},[r("div",{attrs:{id:"nav"}},[r("router-link",{attrs:{to:"/"}},[e._v("首页")]),e._v(" | "),r("router-link",{attrs:{to:"/context"}},[e._v(" 环境 ")]),e._v(" | "),r("router-link",{attrs:{to:"/requests"}},[e._v(" 请求统计 ")]),e._v(" | "),r("router-link",{attrs:{to:"/about"}},[e._v("关于")])],1),r("router-view")],1)}),l=[],i=(r("034f"),r("cba8")),u={},c=Object(i["a"])(u,s,l,!1,null,null,null),p=c.exports,m=r("a5e4"),d=function(){var e=this,t=e.$createElement;e._self._c;return e._m(0)},h=[function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",{staticClass:"home"},[r("h1",[e._v("Starter Gin 开发助手")]),r("div",[e._v(" 可以通过设置属性：web.devtools.enable=true (或false), 来启用（或关闭）开发助手。 ")])])}],f={name:"Home",components:{}},b=f,v=Object(i["a"])(b,d,h,!1,null,null,null),y=v.exports,T=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",{staticClass:"about"},[r("h1",[e._v("环境信息")]),r("el-button",{staticClass:"btn-refresh el-icon-refresh",attrs:{type:"primary"},on:{click:e.refresh}},[e._v(" 刷新 ")]),r("el-collapse",{attrs:{accordion:""},model:{value:e.activeName,callback:function(t){e.activeName=t},expression:"activeName"}},[r("el-collapse-item",{attrs:{title:"Arguments",name:"1"}},[r("el-table",{attrs:{data:e.myArguments}},[r("el-table-column",{attrs:{prop:"index",label:"Index",width:"100px"}}),r("el-table-column",{attrs:{prop:"value",label:"Value"}})],1)],1),r("el-collapse-item",{attrs:{title:"Environment",name:"2"}},[r("el-table",{attrs:{data:e.myEnvironment}},[r("el-table-column",{attrs:{prop:"index",label:"Index",width:"100px"}}),r("el-table-column",{attrs:{prop:"name",label:"Name"}}),r("el-table-column",{attrs:{prop:"value",label:"Value"}})],1)],1),r("el-collapse-item",{attrs:{title:"Properties",name:"3"}},[r("el-table",{attrs:{data:e.myProperties}},[r("el-table-column",{attrs:{prop:"index",label:"Index",width:"100px"}}),r("el-table-column",{attrs:{prop:"name",label:"Name"}}),r("el-table-column",{attrs:{prop:"value",label:"Value"}})],1)],1)],1)],1)},g=[],_={name:"debug-context",data(){return{myArguments:[],myEnvironment:[],myProperties:[],activeName:1}},mounted(){this.refresh()},methods:{refresh(){this.$store.dispatch("api/httpGet",{type:"context"}).then(e=>{this.handleHttpOk(e.data)}).catch(e=>{})},handleHttpOk(e){this.myArguments=this.convertStringArrayToList(e.arguments),this.myEnvironment=this.convertTableToList(e.environment),this.myProperties=this.convertTableToList(e.properties)},convertTableToList(e){let t=[],r=[];for(var n in e)r.push(n);for(var a in r.sort(),r){let n=r[a],o=e[n];t.push({index:a,name:n,value:o})}return t},convertStringArrayToList(e){let t=[];for(var r in e){let n=r+"",a=e[r];t.push({key:n,index:n,value:a})}return t}}},x=_,k=(r("b63f"),Object(i["a"])(x,T,g,!1,null,"39099d59",null)),w=k.exports,O=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",{staticClass:"about"},[r("h1",[e._v("HTTP 请求统计")]),r("el-button",{staticClass:"btn-reset el-icon-reset",attrs:{type:"default"},on:{click:e.reset}},[e._v(" 复位 ")]),r("el-button",{staticClass:"btn-refresh el-icon-refresh",attrs:{type:"primary"},on:{click:e.refresh}},[e._v(" 刷新 ")]),r("span",[e._v(" 自动刷新 : "),r("el-switch",{model:{value:e.enableRefreshTimer,callback:function(t){e.enableRefreshTimer=t},expression:"enableRefreshTimer"}}),e._v(" "+e._s(e.timerCount)+" 次 ")],1),r("el-table",{attrs:{data:e.myEndpoints}},[r("el-table-column",{attrs:{prop:"index",label:"序号"}}),r("el-table-column",{attrs:{prop:"Method",label:"方法"}}),r("el-table-column",{attrs:{prop:"Path",label:"路径"}}),r("el-table-column",{attrs:{prop:"Count",label:"请求计数"}}),r("el-table-column",{attrs:{prop:"TimeSpan",label:"持续时间",formatter:e.formatTimeSpan}})],1)],1)},E=[],P={name:"debug-requests",data(){return{myEndpoints:[],timerCount:0,timerInterval:1e3,timerThreadId:0,enableRefreshTimer:!1,netWorking:!1}},watch:{enableRefreshTimer(e){e?this.startTimer():this.stopTimer()}},mounted(){this.refresh()},deactivated(){this.stopTimer()},methods:{refresh(){this.netWorking=!0,this.$store.dispatch("api/httpGet",{type:"requests"}).then(e=>{this.handleHttpOk(e.data),this.netWorking=!1}).catch(e=>{this.netWorking=!1})},handleHttpOk(e){let t=e.requests;for(var r in t){let e=t[r];e.index=r,e.key=e.Method+":"+e.Path,e.TimeSpan=e.TimeEnd-e.TimeBegin}this.myEndpoints=t},reset(){this.$store.dispatch("api/httpDelete",{type:"requests"}).then(e=>{}).catch(e=>{})},formatTimeSpan(e,t,r,n){let a=r/1e3;return a+"秒"},onTimer(e){if(e!=this.timerThreadId)return;let t=this,r=this.timerInterval;setTimeout(()=>{t.onTimer(e)},r),this.netWorking||(this.refresh(),this.timerCount++)},startTimer(){let e=new Date,t=e.getTime();this.timerThreadId=t,this.timerCount=0,this.onTimer(t)},stopTimer(){this.timerThreadId=0}}},j=P,C=(r("3021"),Object(i["a"])(j,O,E,!1,null,"2c09dfd0",null)),S=C.exports,I=r("a0c0");n["default"].use(m["a"]);const $=[{path:"/",name:"Home",component:y},{path:"/context",name:"debug-context",component:w},{path:"/requests",name:"debug-requests",component:S},{path:"/modules",name:"debug-modules",component:I["a"]},{path:"/about",name:"About",component:function(){return r.e("about").then(r.bind(null,"f820"))}}],N=new m["a"]({mode:"history",base:"/starter-gin-devtools/",routes:$});var H=N,M=r("7736"),q=r("73ef"),A=r.n(q);const L="/api/devtools.v1/";function G(e,t,r){let n=t.url,a=t.params,o=t.data,s=t.type,l=t.id;null==n&&(n=L+s,null!=l&&(n=n+"/"+l));let i=A()({method:r,url:n,params:a,data:o});return i}const R={},V={},W={},D={httpGet(e,t){return G(e,t,"GET")},httpPost(e,t){return G(e,t,"POST")},httpDelete(e,t){return G(e,t,"DELETE")},httpPut(e,t){return G(e,t,"PUT")}};var J={name:"rest",namespaced:!0,state:R,getters:V,mutations:W,actions:D};n["default"].use(M["a"]);var B=new M["a"].Store({state:{},mutations:{},actions:{},modules:{api:J}});n["default"].config.productionTip=!1,n["default"].use(o.a),new n["default"]({router:H,store:B,render:function(e){return e(p)}}).$mount("#app")},"973a":function(e,t,r){"use strict";r("1106")},a0c0:function(e,t,r){"use strict";var n=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("div",{staticClass:"about"},[r("h1",[e._v("模块信息")]),r("el-button",{staticClass:"btn-refresh el-icon-refresh",attrs:{type:"primary"},on:{click:e.refresh}},[e._v(" 刷新 ")]),r("el-table",{attrs:{data:e.myModules}},[r("el-table-column",{attrs:{prop:"Index",label:"序号"}}),r("el-table-column",{attrs:{prop:"Name",label:"名称"}}),r("el-table-column",{attrs:{prop:"Version",label:"版本"}}),r("el-table-column",{attrs:{prop:"Revision",label:"版次"}})],1)],1)},a=[],o={name:"debug-modules",data(){return{myModules:[]}},mounted(){this.refresh()},methods:{refresh(){this.$store.dispatch("api/httpGet",{type:"modules"}).then(e=>{this.handleHttpOk(e.data)}).catch(e=>{})},handleHttpOk(e){let t=e.modules;for(var r in t){let e=t[r];e.key=e.Name+"@"+e.Version}t.sort((e,t)=>this.cmpstr(e.Index,t.Index)),this.myModules=t},cmpstr(e,t){let r=Number(e),n=Number(t);return r-n}}},s=o,l=(r("973a"),r("cba8")),i=Object(l["a"])(s,n,a,!1,null,"2c6fa758",null);t["a"]=i.exports},b63f:function(e,t,r){"use strict";r("50c8")}});
//# sourceMappingURL=app.74197211.js.map