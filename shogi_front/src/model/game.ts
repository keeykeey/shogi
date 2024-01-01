
/*
export type KomaArrangement = {
  ID: number,
  ArrangementId: number,
  Arrangement: Arrangement,
  KomaID: number,
  Koma: Koma,
  PositionID: number,
  Position: Position,
  IsFirstMove: boolean,
  IsFront: boolean,
}

export type Arrangement = {
  ID: number,
  Name: string,
  CreatedAt: Date,
  UpdatedAt: Date,
  DeletedAt: Date,
  KomaID: number,
}

export type Koma = {
  ID: number,
  MoveID: number,
  MoveID2: number,
  Name: string,
  Name2: string,
  CreatedAt: Date,
  UpdatedAt: Date,
  DeletedAt: Date,
}

export type Position = {
  ID: number,
  Number: number,
  Name: string,
  CreatedAt: Date,
  UpdatedAt: Date,
  DeletedAt: Date,
}

export type GameConstructor = {
  gameId: number
}

class Game {
  getEndpoint: string;
  komaArrangement: KomaArrangement[] | null;
  constructor(gameConstructor: GameConstructor) {
    this.getEndpoint = "localhost:8080/api/komaArrangements";
    this.komaArrangement= null
  }
  
  function getKomaArrantement() {
    const 

  }
}
*/