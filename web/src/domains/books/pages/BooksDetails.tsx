import { useParams } from "react-router-dom";

export const BooksDetails: React.FC = () => {
  const { slug } = useParams<{ slug: string }>();
  return <div>books details {slug}</div>;
};
