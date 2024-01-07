export type KomaArrangementT = {
  id: number,
  arrangementId: number,
  arrangement: ArrangementT,
  komaId: number,
  koma: KomaT,
  positionId: number,
  position: PositionT,
  isFirstMove: boolean,
  isFront: boolean,
}

export type ArrangementT = {
  id: number,
  name: string,
  komaId: number,
}

export type KomaT = {
  id: number,
  moveId: number,
  moveId2: number,
  name: string,
  name2: string,
}

export type PositionT = {
  id: number,
  number: number,
  height: number,
  width: number,
  name: string,
}

export class KomaArrangement {
  komaArrangement: KomaArrangementT[];
  constructor() {
    this.komaArrangement = [];
  }

  async get(gameId: number, baseEndpoint: string) {
    const res = await fetch(`${baseEndpoint}/${gameId}`, {
      cache: "no-store"
    });
    return res.json();
  }
}