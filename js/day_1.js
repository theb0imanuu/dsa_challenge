//Hello world

// 'const' for variable
const message = "Hello World";
console.log(message);

// Arrow fucntion
const add = (a, b) => {
    return a + b;
};

const result = add(5, 10);
console.log("The sum of 5 and 10 is", result);

//FizzBuzz
for (i = 1; i<= 100; i++){
    if (i % 15 ===0){
        console.log("FizzBuzz");
    }else if (i % 3 ===0){
        console.log("Fizz");
    }else if ( i % 5 ===0){
        console.log("Buzz");
    }else{
        console.log(i);
    }
}
