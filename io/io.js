function R_IO(){
    this.OpenFile = function(FileName,Chmode){
        temp_file_obj = new R_File();
        temp_file_obj.File_id = IO_fopen(FileName,Chmode);
        temp_file_obj.File_name = FileName;
        return temp_file_obj;
    };
}
function R_File(){
    this.File_id = "";
    this.File_name = "";
    this.Close = function(){
        //fclose(this.File_id);
        output(File_id);
        return;
    };
    this.Write = function(data){
        return IO_write(this.File_id,data);
    };
    this.Append = function(data){
        var temp_io = R_IO();
        var temp_file = temp_io.OpenFile(this.File_name,IO_Append);
        var result = temp_file.Write(data);
        temp_file.Close();
        return result;
    };
}