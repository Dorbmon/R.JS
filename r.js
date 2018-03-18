//include_c("io");	//RJS
var start = new Date().getTime();
var i = 0
for(var a = 0;a < 100000;a++){
	i = (i + 1) * (i + 2);
}


var end = new Date().getTime(); 


//console.log(end - start);	这是你浏览器用的
output(end - start)//这是RJS
exit(0)