/*
  This problem was asked by Google.

  The area of a circle is defined as πr^2. 
  Estimate π to 3 decimal places using a Monte Carlo method.

  Hint: The basic equation of a circle is x2 + y2 = r2.
*/

const generateRandom = () => {
  return [(Math.random() * 2) - 1, (Math.random() * 2) - 1];
};

const checkPoint = (c) => {
  return (c[0] * c[0]) + (c[1] * c[1]) <= 1;
};

(() => {
  const iterations = 1000000;
  let inside = 0;
  for (let i = 0; i < iterations; i++) {
    inside += checkPoint(generateRandom()) ? 1 : 0;
  }
  // 1/4th of pi, (* and / for correct amount of decimal places)
  const p4 = Math.floor((inside / iterations) * 1000) / 1000;
  const pi = p4 * 4;
  console.log(pi);
  return pi;
})();