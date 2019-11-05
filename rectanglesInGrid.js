const example = [
  [1, 1, 1, 1],
  [1, 1, 0, 1],
  [1, 1, 1, 1],
  [1, 1, 1, 1],
];

const findRects = map => {
  let total = 0;
  // don't check last row
  for (let y = 0; y < map.length - 1; y++) {

    // don't check the last x of y
    for (let x = 0; x < map[y].length - 1; x++) {
      if (map[y][x] !== 0) {
        // Determine all straight line rects
        // unique starting point y,x
        for (let x2 = 1; x + x2 < map[y].length; x2++) {  
          for (let y2 = 1; y + y2 < map.length; y2++) {
            if (map[y][x] && map[y][x + x2] && map[y + y2][x] && map[y + y2][x + x2]) {
              // rect exists
              total++;
            }
          }
        }
      }

      // Determine diagonal rects using center point
      // x and y are not on the edge of grid
      if (x > 0 && y > 0) {
        // center point = map[y][x]

        // check to make sure all sides are not out of range
        // top --> y - dist >= 0
        // bot --> y + dist < map.length
        // left --> x - dist >= 0
        // right --> x + dist < map[y].length
        for (let dist = 1; y - dist >= 0 && y + dist < map.length && x - dist >= 0 && x + dist < map[y].length; dist++) {
          // all points are in range, now check if they all exist
          if (map[y - dist][x] && map[y + dist][x] && map[y][x - dist] && map[y][x + dist]) total++;
        }
      }
    }
  }

  return total;
}

console.log(findRects(example));