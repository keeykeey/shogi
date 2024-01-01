
interface PositionConstructor {
  w: number,
  h: number,
}

class Position {
  w: number;
  h: number;
  constructor(pc: PositionConstructor) {
    this.w = pc.w;
    this.h = pc.h;
  }
};
