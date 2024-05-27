import { useAuth } from "@/hooks";

export default function HomePage() {
  const data = useAuth();
  console.log(data);

  return (
    <div className={styles.container}>
        <h2 className={styles.title}>Home Page here</h2>
    </div>
  )
}
