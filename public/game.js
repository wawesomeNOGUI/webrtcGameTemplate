
//Animation and game updates will be triggered once connection made in index.hmtl

var animate = window.requestAnimationFrame || window.webkitRequestAnimationFrame || window.mozRequestAnimationFrame || function (callback) {
        window.setTimeout(callback, 1000/60)
    };
var canvas = document.createElement("canvas");
var width = 1000;
var height = 700;
canvas.width = width;
canvas.height = height;
var context = canvas.getContext('2d');




var keysDown = {};     //This creates var for keysDown event

var playerX = 0;
var playerY = 0;


var render = function () {

   context.fillStyle = "#000000";                 //Example of Drawing a screen & below some text to the canvas
   context.fillRect(0, 0, width, height);

   //Draw Players
   var tempX;

   for (var key in Updates) {     //Updates defined in index.html
      if (Updates.hasOwnProperty(key) && key.includes("X") )  {  //Is X value?
          tempX = Updates[key];
      }else{                                 //Else must have X and Y, so draw
        context.fillStyle = "#FF00FF";
        context.fillRect(tempX, Updates[key], 50, 50);
        console.log(tempX + " " + Updates[key])
      }
   }

   //Local Player
   context.fillStyle = "#F7FF0F";   //yellow
   context.fillRect(playerX, playerY, 50, 50);
   //context.drawImage(pew, playerX, playerY, 200, 100, 100,100,50,50);
  //context.drawImage(pew, playerX, playerY, 200, 100);

};




var update = function() {
keyPress();


};



var step = function() {
update();
render();
animate(step);
};



var keyPress = function() {

for(var key in keysDown) {
    var value = Number(key);

   if (value == 37) {   //37 = left
     playerX = playerX - 4;
     TCPChan.send("X" + playerX);
	    //Put stuff for keypress to activate here
   } else if(value == 39){  //38 = right
     playerX = playerX + 4;
     TCPChan.send("X" + playerX);
   } else if(value == 40){  //40 = down
     playerY = playerY + 4;
     TCPChan.send("Y" + playerY);
   } else if(value == 38){  //39 = up
     playerY = playerY - 4;
     TCPChan.send("Y" + playerY);
   }

 }

};






document.body.appendChild(canvas);



window.addEventListener("keydown", function (event) {
keysDown[event.keyCode] = true;
});


window.addEventListener("keyup", function (event) {
    delete keysDown[event.keyCode];
});