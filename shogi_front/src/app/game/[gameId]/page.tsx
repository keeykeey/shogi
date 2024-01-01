import Board from "@/app/game/[gameId]/components/board";
import Koma from "@/app/game/[gameId]/components/koma";
import { KomaArrangement, KomaArrangementT } from "@/model/komaArrangement";

export default async function Page() {
  try {
    const model = new KomaArrangement();
    const baseUrl = process.env.NEXT_PUBLIC_API_HOST_URL;
    const endpoint = `${baseUrl}/api/komaArrangements`;
    const komaArrangements = await model.get(1, endpoint) as KomaArrangementT[];

    return (
      <>
        <Board
          width={400}
          height={400}
        />
        { komaArrangements && 
          komaArrangements.map( item => {
            const h = item.Position.Number % 10
            const w = Math.floor(item.Position.Number/10)
            const top = 65 + 32 * (h - 1);
            const left = 41 + 32 * (w - 1);
            const style = {
              top: `${top}px`,
              left: `${left}px`
            };
            return (
              <Koma
                key={item.ID}
                isFirstMove={true}
                isFront={true}
                label={item.Koma.Name}
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
