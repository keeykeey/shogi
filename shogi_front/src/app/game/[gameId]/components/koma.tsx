import Image from "next/image";
import style from "@/style/game/[gameId]/components/koma.module.css"

export interface IKoma {
  label: string,
  label2: string,
  isFirstMove: boolean,
  isFront: boolean
  width: number,
  height: number,
  style?: KomaStyle,
}

export interface KomaStyle {
  top: string,
  left: string,
  backgroundColor?: string,
}

export default function Koma(props: IKoma) {
  const label = props.isFront ? props.label : props.label2;
  return (
    <>
      <Image
        src={`/shogi/${label}.svg`}
        alt="image not found"
        width={props.width}
        height={props.height}
        style={props.style}
        className={style.koma}
      />
    </>
  )
}
