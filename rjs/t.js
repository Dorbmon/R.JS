output("Server Started");
var Web = WebServer();
Web.Bind("/",function(writer,request){
	writer.Write("<html>");
	writer.Write("Hello World\n");
	//writer.Write("U URL is:",request.RequestUrl);
	writer.Write("data is:",request.GetForm("data"));
	writer.Write("\
	<script>\
		//alert('SSSS');\
		\
		</script>\
	")
	writer.Writer("</html>");
	return;
});
Web.Listen("0.0.0.0:5020");
for(;;){
	
}