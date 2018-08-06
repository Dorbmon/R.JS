var consts = ImportGTKConst();
var Window = NewWindow(consts.WINDOW_TOPLEVEL);
Window.SetSizeRequest(500,500);
var image = NewImage();
var pixbuf = NewPixbuf("fruit.png",50,50,false);
if(!pixbuf){
	output("Error Pixbuf");
	exit(0);
}
image.SetSizeRequest(50,50);
if(!image.SetFromPixbuf(pixbuf)){
	output("Error set");
	exit(0);
}
Window.SetPosition(consts.WIN_POS_CENTER);
Window.SetTitle("RJS!!!!!!");
var Fixed = NewGtkFixed();
var button = NewButton();
if(!button.SetImage(image)){
	output("设置Image出错");
	exit(0);
}
button.SetLable("试试事件链接。");
button.Connect("clicked",function(){
	output("clicked....");
});
Window.Add(Fixed);
Fixed.Put(button,50,50);
//Fixed.Put(image,0,0);
Window.Connect("destroy",function(){
	MainQuit();
	output("lol");
});
Window.ShowAll();
GTKMain();
