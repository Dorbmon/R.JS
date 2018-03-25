function R_Stack(Stack_name){
    //this.Stack_name = "";
    if(!NewStack(Stack_name)){
        return false;
    }
    this.Stack_name = Stack_name;
    this.Push = function(data){
        return Stack_Push(this.Stack_name,data);
    };
    this.Pop = function(){
        return Stack_Pop(this.Stack_name);
    };
}