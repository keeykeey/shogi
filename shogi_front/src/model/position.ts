import { BoardC } from "@/model/board";

export interface PositionCConstructor {
  width: number,
  height: number,
  board: BoardC,
}

export class PositionC {
  width: number;
  height: number;
  top: number | null = null;   // css top
  left: number | null = null;  // css left
  constructor(pc: PositionCConstructor) {
    this.width = pc.width;
    this.height = pc.height;
    this.setTop();
    this.setLeft();
  }
  setTop() {
    this.top = 65 + 32 * (this.height -1); // TODO :今は仮

  }
  setLeft() {
    this.left = 41 + 32 * (this.width -1); // TODO :今は仮
  }
};
