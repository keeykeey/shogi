import Board from "@/app/game/[gameId]/components/board";
import Koma from "@/app/game/[gameId]/components/koma";
import { KomaArrangement, KomaArrangementT } from "@/model/komaArrangement";
import { createKomaOnBoard, createKomaStyle } from "./interface";
import { BoardC, BoardCConstructor } from "@/model/board";

export default async function Page() {
  try {
    const model = new KomaArrangement();
    const baseUrl = process.env.NEXT_PUBLIC_API_HOST_URL;
    const endpoint = `${baseUrl}/api/komaArrangements`;
    const komaArrangements = await model.get(1, endpoint) as KomaArrangementT[];

    // TODO: create endpoint and fetch from api
    const boardCC: BoardCConstructor = {
      boardTop: 1,
      boardLeft: 1,
      squareHeight: 1,
      squareWidth: 1,
      lineWidth: 1,
    }
    const board = new BoardC(boardCC);

    return (
      <>
        <Board
          width={400}
          height={400}
        />
        { komaArrangements && 
          komaArrangements.map( item => {
            const koma = createKomaOnBoard(item, board);
            const style = createKomaStyle(koma);
            return (
              <Koma
                key={item.id}
                label={koma.label}
                label2={koma.label2}
                isFirstMove={koma.isFirstMove}
                isFront={koma.isFront}
                width={35}
                height={35}
                style={style}
              />
            )
          })
        }
      </>
    )
  } catch {
    return (
      <>
        page not found
      </>
    );
  }

}
