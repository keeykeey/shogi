import { KomaArrangementT } from "@/model/komaArrangement";
import { KomaC, KomaCConstructor, PlainKomaCConstructor } from "@/model/koma";
import { PositionCConstructor } from "@/model/position";
import { KomaStyle } from "./components/koma";

export interface CreateKomaOnBoard {
  (komaArrangement: KomaArrangementT): KomaC
}

export interface CreateKomaStyle {
  (koma: KomaC): KomaStyle;
}

export const createKomaOnBoard: CreateKomaOnBoard = (komaArrangement) => {
  const h = komaArrangement.Position.Number % 10
  const w = Math.floor(komaArrangement.Position.Number/10)
  const positionCConstructor: PositionCConstructor = {
    w: w,
    h: h
  };

  const plainKomaCConstructor: PlainKomaCConstructor = {
    label: komaArrangement.Koma.Name,
    label2: komaArrangement.Koma.Name2,
  }

  const komaCConstructor: KomaCConstructor = {
    plainKoma: plainKomaCConstructor,
    position: positionCConstructor,
    isFirstMove: komaArrangement.IsFirstMove,
    isFront: komaArrangement.IsFront,
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
