/*
	Given an integer k and a string s, find the length of the longest substring that contains at most k distinct characters.

	For example, given s = "abcba" and k = 2, the longest substring with k distinct characters is "bcb".

	This problem was asked by Amazon.
*/
// sliding window moving up string using two pointers (beginning, end)

const pointer = (str, k) => {
  if (str.length <= k || str.length < 2) {
    return str;
  }

  let p = 0;
  let p2 = 1;

  let curChars = str[p] === str[p2] ? 1 : 2;

  let map = {
    [str[p]]: curChars === 1 ? 2 : 1,
    [str[p2]]: curChars === 1 ? 2 : 1,
  };
  
  let longest = {p, p2};

  while (p2 < str.length - 1) {
    p2++;
    if (map[str[p2]]) { // exists && not 0
      map[str[p2]] += 1;
    } else {
      map[str[p2]] = 1;
      curChars++;
      while (curChars > k) {
        // could add check to see if its possible for a longer pair to exist within remainder size of string, and if not, break & return

        map[str[p]] -= 1;
        if (map[str[p]] === 0) {
          curChars -= 1;
        }
        p++;
      }
    }
    // check longest
    if (p2 - p > longest.p2  - longest.p) { 
      longest = {p, p2};
    }
  }

  let answer = "";
  for (let i = 0; i <= longest.p2 - longest.p; i++) {
    answer += str[longest.p + i];
  }

  return answer;
};

// iterative brute force
const bruteForce = (str, k) => {
	if (str.length < k) {
		return str;
  }
  
  let xsl = []; // slice of string (longest)
	for (let i = 0; i < str.length; i++) {
    let m = {
      [str[i]]: true,
    };
    let hit = 1;
    let xsc = []; // slice of string (cur)
    
    for (let j = i; j < str.length; j++) {
      if (!m[str[j]]) {
        if (hit + 1 > k) {
          break;
        }
        hit++;
        m[str[j]] = true;
      }
      xsc.push(str[j]);
    }
    
    if (xsc.length > xsl.length) xsl = xsc;
  }

  return xsl.join("");
}

(function test () {
  let tests = [
    {S: "abcba", K: 2, A: "bcb"},
    {S: "aaabba", K: 1, A: "aaa"},
    {S: "aaabbabbbba", K: 1, A: "bbbb"},
    {S: "abcbacdcbac", K: 3, A: "abcbac"},
  ];

  tests.forEach(t => {
    a = pointer(t.S, t.K);
    if (a !== t.A) {
      console.error(`Fail: for ${t.S}->${t.K} expected ${t.A}, but got ${a}`);
    } else {
      console.log(`Pass: ${t.S}->${t.K}`);
    }
  })
})();
