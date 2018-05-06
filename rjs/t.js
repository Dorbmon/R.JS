include("ww.js");
test();
var net = new brain.NeuralNetwork();
net.train([{input:[0],output:[1]},{input:[1],output:[0]}]);
output(net.run([0]));