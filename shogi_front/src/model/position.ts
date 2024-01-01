export interface PositionCConstructor {
  w: number,
  h: number,
}

export class PositionC {
  w: number;
  h: number;
  top: number | null = null;   // css top
  left: number | null = null;  // css left
  constructor(pc: PositionCConstructor) {
    this.w = pc.w;
    this.h = pc.h;
    this.setTop();
    this.setLeft();
  }
  setTop() {
    this.top = 65 + 32 * (this.h -1); // TODO :今は仮
  }
  setLeft() {
    this.left = 41 + 32 * (this.w -1); // TODO :今は仮
  }
};
