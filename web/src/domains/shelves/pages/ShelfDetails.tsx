import { useParams } from "react-router-dom";

export const ShelfDetails: React.FC = () => {
  const { slug } = useParams<{ slug: string }>();
  return <div>ShelfDetails {slug} </div>;
};
