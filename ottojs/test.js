var global = Function('return this')();

global.test = function() {
    console.log("Hello, World.");
};

function abc() {
    console.log("in function abc");
}
