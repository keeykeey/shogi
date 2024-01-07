export interface BoardCConstructor {
  boardTop: number,
  boardLeft: number,
  squareHeight: number;
  squareWidth: number;
  lineWidth: number;
}

export class BoardC {
  boardTop: number;
  boardLeft: number;
  squareHeight: number;
  squareWidth: number;
  lineWidth: number;
  constructor(c: BoardCConstructor) {
    this.boardTop = c.boardTop;
    this.boardLeft = c.boardLeft;
    this.squareHeight = c.squareHeight;
    this.squareWidth = c.squareWidth;
    this.lineWidth = c.lineWidth;
  }
}
