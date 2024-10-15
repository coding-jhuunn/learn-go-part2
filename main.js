let hello = "a";

const modified = (hello) => {
  hello = "b";
  return;
};

console.log(hello);
modified(hello);
console.log(hello);
