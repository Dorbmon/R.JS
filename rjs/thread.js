function Thread(){
	var a = 0;
	for(;;){
		//output("哈哈哈哈哈")
		console.log("HHHH");
		//runtime.Wait(0.5 * Second);
	}
	return a;
}
function OnERROR(data){
	output(data);
	return;
}
function Finished(value){
	output(value);
}
output("Hello");
var ThreadID = RunThread(Thread,OnERROR,Finished);
//runtime.Wait(Second);
//output("ID：",ThreadID);
output("Wait()");
runtime.Wait(2 * Second);
/*
output("1254345645345345");
if(!PauseThread(ThreadID)){
	output("ERROR PAUSE");
}
runtime.Wait(3 * Second);
if(!ContinueThread(ThreadID)){
	output("无法继续");
	exit(0);
}*/
output("Will");
for(;;){
	output("你好！\n");
	//runtime.Gosched();
}