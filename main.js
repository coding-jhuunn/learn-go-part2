let hello = "a";

const modified = (letter) => {
  letter = "b";
  return letter;
};

console.log(modified(hello));
