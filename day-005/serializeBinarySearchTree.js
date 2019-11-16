/*

Given the root to a binary tree, implement serialize(root), which serializes the tree into a string, and deserialize(s), which deserializes the string back into the tree.

*/

class Node {
  constructor(val) {
    this.val = val;
    this.left = null;
    this.right = null;
  }
}

const serialize = (node) => {
  if (node === null) return "";
  const ls = serialize(node.left);
  const rs = serialize(node.right);

  return `{val: ${node.val}, left: ${ls || null}, right: ${rs || null}}`;
};

const deserialize = (srl) => {
  let cobj = srl.substr(1, srl.length - 2);
  const val = cobj.split(": ", 2)[1].split(", ", 1)[0];
  let left = cobj.indexOf("left: ");
  let leftEndIndex = left + 6 + 4;

  if (cobj[left + 6] == "n") left = null;
  else {
    leftEndIndex = findIndexObjectEnds(cobj, left + 6);
    const leftNode = cobj.substring(left + 6, leftEndIndex + 1);
    left = deserialize(leftNode);
  }
  
  let rightStr = cobj.substring(leftEndIndex + 1, cobj.length);
  let right = rightStr.indexOf("right: ");
  if (rightStr[right + 7] === "n") right = null;
  else {
    rightEndIndex = findIndexObjectEnds(rightStr, right + 7);
    right = deserialize(rightStr.substring(right + 7, rightEndIndex + 1));
  }

  const n = new Node(parseInt(val));
  n.left = left;
  n.right = right;
  return n;
};

// starting must point to opening { -> returns index of closing }
const findIndexObjectEnds = (str, starting) => {
  let open = 0;
  for (let i = starting; i < str.length; i++) {
    if (str[i] == "{") open++;
    else if (str[i] == "}") open--;
    if (open === 0) return i;
  }
  return null;
}