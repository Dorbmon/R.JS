function R_File(){
    this.File_id = 0;
    this.File_name = "";
    this.Close = function(){
        fclose(this.File_id);
        //output(this.File_id);
        return;
    };
    this.Write = function(data){
        return IO_write(this.File_id,data);
    };
    this.Append = function(data){
        var temp_io = R_IO();
        var temp_file = IO_OpenFile(this.File_name,IO_Append);
        var result = temp_file.Write(data);
        temp_file.Close();
        return result;
    };
}
function R_IO(){
    this.OpenFile = function(FileName,Chmode){
        var temp_file_obj = new R_File();
        temp_file_obj.File_id = IO_fopen(FileName,Chmode);
        temp_file_obj.File_name = FileName;
        return temp_file_obj;
    };
}
function IO_OpenFile (FileName,Chmode){
    var temp_file_obj = new R_File();
    temp_file_obj.File_id = IO_fopen(FileName,Chmode);
    temp_file_obj.File_name = FileName;
    return temp_file_obj;
}
function Print(){
    var data;
    for( var i = 0; i < arguments.length; i++ ){
        data += arguments[i];
    }
    output(data);
}
function Exit(code){
    exit(code);
}
function TCP_SERVER(){
    this.TCP_LISTENER_ID = 0;
    this.Onconnection = function(){

    };
    this.SetOnConnectFunc = function(func){
        this.OnConnection = func;
    };
    this.Listener = function(ip,port){
        IO_Start_TCP_Server(ip,port,this.Onconnection);
    };
}
function XMLHttpRequest(){
    this.readyState = 0;
    this.status = 0;
    this.mode = "GET";
    this.url = "https://rxues.site";
    this.open = function(mode,url,asynchronous){
        this.mode = mode;
        this.url = url;
        var RandName = OnlyRand();
        if(NETWORK_AJAX(url,mode,RandName) == false){
            //生成失败
            return false;
        }
        this.asynchronous = asynchronous;
        this.ObjName = RandName;
        return true;
    };
    this.Set = function(key,value){
        return NETWORK_AJAX_SET(this.ObjName,key,value);
    };
    this.Send = function(CallBackFunction,ErrorFunction){
        //发送消息
        NET_AJAX_ON_MESSAGE_FUNCTION(CallBackFunction,this.ObjName);
        NET_AJAX_ON_ERROR_FUNCTION(ErrorFunction,this.ObjName);
        NETWORK_AJAX_SEND(this.ObjName,this.asynchronous);
        return;
    };
    this.setRequestHeader = function(key,value){
        return NETWORK_AJAX_SET_REQUEST_HEADER(this.ObjName,key,value);
    };
}