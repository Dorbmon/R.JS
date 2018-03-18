function R_File(){
    this.File_id = 0;
    this.File_name = "";
    this.Close = function(){
        //fclose(this.File_id);
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
        temp_file_obj = new R_File();
        temp_file_obj.File_id = IO_fopen(FileName,Chmode);
        temp_file_obj.File_name = FileName;
        return temp_file_obj;
    };
}
function IO_OpenFile (FileName,Chmode){
    temp_file_obj = new R_File();
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
        IO_Start_TCP_Server(ip,port);
    };
}