let digits;
let searchBox;
let indexP;

function preload() {
  raw = loadStrings("https://www.angio.net/pi/digits/10000.txt");
}

function setup() {
  digits = raw[0];

  noCanvas();
  searchBox = createInput("");
  searchBox.input(search);
  indexP = createP("searching");
}

function search() {
  let num = searchBox.value();
  let index = digits.indexOf(num) - 1;

  if (index == -2) {
    let url = `/search?q=${num}`;
    loadJSON(url, (data) => {
      indexP.html(`Found at index ${data.index}`);
    });
  }
  indexP.html(`Found at index ${index}`);
}
