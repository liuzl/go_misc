var Algebrite = require("./algebrite.js");

var x = Algebrite.run('x + x') // => "2 x"
console.log(x);

x = Algebrite.factor('10!').toString() // => "2^8 3^4 5^2 7"
console.log(x);

x = Algebrite.eval('integral(x^2)').toString() // => "1/3 x^3"
console.log(x);

// composing...
x = Algebrite.integral(Algebrite.eval('x')).toString() // => "1/2 x^2"
console.log(x);

