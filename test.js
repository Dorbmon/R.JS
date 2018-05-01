include_c("stack");
include_c("io");
stack = new R_Stack("sss");
if(!stack){
    //创建失败
    output("栈创建失败");
}
//stack.Push("ssss");
//output(stack.Pop());
call(test);
function test(){
    output("HAHA");
    return "Nice";
}