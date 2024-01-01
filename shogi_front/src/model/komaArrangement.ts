export type KomaArrangementT = {
  ID: number,
  ArrangementId: number,
  Arrangement: ArrangementT,
  KomaID: number,
  Koma: KomaT,
  PositionID: number,
  Position: PositionT,
  IsFirstMove: boolean,
  IsFront: boolean,
}

export type ArrangementT = {
  ID: number,
  Name: string,
  CreatedAt: Date,
  UpdatedAt: Date,
  DeletedAt: Date,
  KomaID: number,
}

export type KomaT = {
  ID: number,
  MoveID: number,
  MoveID2: number,
  Name: string,
  Name2: string,
  CreatedAt: Date,
  UpdatedAt: Date,
  DeletedAt: Date,
}

export type PositionT = {
  ID: number,
  Number: number,
  Name: string,
  CreatedAt: Date,
  UpdatedAt: Date,
  DeletedAt: Date,
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