include_c("io")
//output("IO_ReadWrite:" )
//output(IO_ReadWrite)
IO = new R_IO();
file = IO.OpenFile("r.txt",IO_ReadWrite);
output(file.File_id);
output("ssss")
if(file.File_id == false){
    output("ERROR OPENING FILE");
}
else{
    file.Write("eeeee");
}
//file.Write("eeeee");
//output("eee")