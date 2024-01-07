import { KomaArrangementT } from "@/model/komaArrangement";
import { KomaC, KomaCConstructor, PlainKomaCConstructor } from "@/model/koma";
import { PositionCConstructor } from "@/model/position";
import { KomaStyle } from "./components/koma";
import { BoardC } from "@/model/board";

export interface CreateKomaOnBoard {
  (komaArrangement: KomaArrangementT,
   board: BoardC ): KomaC
}

export interface CreateKomaStyle {
  (koma: KomaC): KomaStyle;
}

export const createKomaOnBoard: CreateKomaOnBoard = (komaArrangement, board) => {
  const h = komaArrangement.position.height;
  const w = komaArrangement.position.width;
  const positionCConstructor: PositionCConstructor = {
    width: w,
    height: h,
    board: board,
  };

  const plainKomaCConstructor: PlainKomaCConstructor = {
    label: komaArrangement.koma.name,
    label2: komaArrangement.koma.name2,
  }

  const komaCConstructor: KomaCConstructor = {
    plainKoma: plainKomaCConstructor,
    position: positionCConstructor,
    isFirstMove: komaArrangement.isFirstMove,
    isFront: komaArrangement.isFront,
  }
  return new KomaC(komaCConstructor);
};

export const createKomaStyle: CreateKomaStyle = (koma) => {
  return {
    top: `${koma.position.top}px`,
    left: `${koma.position.left}px`,
    transform: koma.isFirstMoveStyle,
  };
}
