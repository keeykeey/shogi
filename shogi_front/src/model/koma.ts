import { PositionC, PositionCConstructor } from "./position"

export interface PlainKomaCConstructor {
  label: string,
  label2: string,
} 

export class PlainKomaC {
  label: string;
  label2: string;
  constructor(c: PlainKomaCConstructor) {
    this.label = c.label;
    this.label2 = c.label2;
  }
}

export interface KomaCConstructor {
  plainKoma: PlainKomaCConstructor
  position: PositionCConstructor,
  isFirstMove: boolean,
  isFront: boolean, 
}

export class KomaC extends PlainKomaC{
  position: PositionC;
  isFirstMove: boolean;
  isFirstMoveStyle: string;
  isFront: boolean;
  constructor (c: KomaCConstructor) {
    super(c.plainKoma);
    this.position = new PositionC(c.position);
    this.isFirstMove = c.isFirstMove;
    this.isFirstMoveStyle = this.isFirstMove ? "scale(1, 1)" : "scale(1, -1)";
    this.isFront = c.isFront;
  }
}
