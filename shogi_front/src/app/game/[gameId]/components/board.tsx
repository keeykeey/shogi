import Image from 'next/image'
import styles from "@/style/game/[gameId]/components/board.module.css";

export interface IBoardProps {
  width: number,
  height: number,
}

export default function Board(props: IBoardProps) {
  return (
    <>
      <Image
        src="/shogi/ban.svg"
        alt="image not found"
        width={props.width}
        height={props.height}
        className={styles.board}
      />
    </>
  );
};
