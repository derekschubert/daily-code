class Tree {
  constructor() {
    this.root = null
  }

  add(n) {
    const node = new Node(n)
    if (this.root === null) {
      this.root = node;
    } else {
      this.root.add(node);
    }
  }

  log() { this.root.visit((val) => console.log(val)); }
  search(val) { return this.root.search(val); }
  getMinHeight() { return this.root.getMinHeight(); }
  getMaxHeight() { return this.root.getMaxHeight(); }

  traverse(method) {
    if (this.root === null) return null;
    let result = new Array();
    function inOrder(n) {
      if (n.left !== null) inOrder(n.left);
      result.push(n.val);
      if (n.right !== null) inOrder(n.right);
    }
    function preOrder(n) {
      result.push(n.val);
      if (n.left !== null) preOrder(n.left);
      if (n.right !== null) preOrder(n.right);
    }
    function postOrder(n) {
      if (n.left !== null) postOrder(n.left);
      if (n.right !== null) postOrder(n.right);
      result.push(n.val);
    }
    function levelOrder(n) {
      let queue = [n];
      while (queue.length > 0) {
        const node = queue.shift();
        result.push(node.val);
        if (node.left !== null) queue.push(node.left);
        if (node.right !== null) queue.push(node.right);
      }
    }

    switch (method) {
      case 'in' || 'inOrder': 
        inOrder(this.root); 
        break;
      case 'pre' || 'preOrder':
        preOrder(this.root);
        break;
      case 'post' || 'postOrder':
        postOrder(this.root);
        break;
      case 'level' || 'levelOrder':
        levelOrder(this.root);
        break;
      default:
        console.error("Invalid method called in traverse:", method)
    }

    return result;
  }

  getValsAtLevel(level) {
    if (this.root === null) return null;
    let queue = [this.root];
    let result = [];
    let lvRight = this.root.val; // last in lvl;
    let height = 0;
    let skip = true;

    while (queue.length > 0) {
      let n = queue.shift();
      if (n.val === lvRight) height++;
      if (height === level) {
        if (skip) skip = false; // skip last node of prev height
        else result.push(n.val);
      }
      if (height === level+1) {
        if (n.val === lvRight) result.push(n.val);
        return result;
      }
      if (n.left !== null) queue.push(n.left);
      if (n.right !== null) {
        queue.push(n.right);
        if (n.right.val > lvRight) lvRight = n.right.val;
      }
    }

    return result;
  }

  rotateRight(n = this.root) {
    if (n === null) return null;
    const n2 = n.left;
    n.left = n2.right;
    n2.right = n;
    if (n === this.root) this.root = n2;
  }

  rotateLeft(n = this.root) {
    if (n === null) return null;
    const n2 = n.right;
    n.right = n2.left;
    n2.left = n;
    if (n === this.root) this.root = n2;
  }
}

class Node {
  constructor(val) {
    this.val = val;
    this.left = null;
    this.right = null;
  }
  
  add(n) {
    if (n.val < this.val) {
      if (this.left === null) {
        this.left = n;
      } else {
        this.left.add(n);
      }
    } else {
      if (this.right === null) {
        this.right = n;
      } else {
        this.right.add(n)
      }
    }
  }

  visit(fn) {
    if (this.left !== null) this.left.visit(fn);
    fn(this.val);
    if (this.right !== null) this.right.visit(fn);
  }

  search(val) {
    if (this.val === val) return this;
    else if (this.val > val && this.left !== null) return this.left.search(val);
    else if (this.val < val && this.right !== null) return this.right.search(val);
    return null;
  }

  getMinHeight() {
    const left = this.left === null ? -1 : this.left.getMinHeight();
    const right = this.right === null ? -1 : this.right.getMinHeight();
    if (left < right) return left + 1;
    return right + 1;
  }

  getMaxHeight() {
    const left = this.left === null ? -1 : this.left.getMaxHeight();
    const right = this.right === null ? -1 : this.right.getMaxHeight();
    if (left > right) return left + 1;
    return right + 1;
  }

};

(function main() {
  let tree = new Tree();
  tree.add(9);
  tree.add(4);
  tree.add(3);
  tree.add(6);
  tree.add(5);
  tree.add(7);
  tree.add(17);
  tree.add(22);
  tree.add(20);
  console.log(tree.traverse('level'));
  console.log(tree.getValsAtLevel(0));
  console.log(tree.getValsAtLevel(1));
  console.log(tree.getValsAtLevel(2));
  console.log(tree.getValsAtLevel(3));
  console.log();
  console.log();
  tree.rotateRight();
  console.log(tree.traverse('level'));
  console.log(tree.getValsAtLevel(0));
  console.log(tree.getValsAtLevel(1));
  console.log(tree.getValsAtLevel(2));
  console.log(tree.getValsAtLevel(3));
})();
