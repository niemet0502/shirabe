import { createBrowserRouter } from "react-router-dom";
import { Layout } from "./components/layout/Layout";
import { BooksDetails } from "./domains/books/pages/BooksDetails";
import { BooksList } from "./domains/books/pages/BooksList";
import { NewBook } from "./domains/books/pages/NewBook";
import { Home } from "./domains/home/Home";
import { NewShelf } from "./domains/shelves/pages/NewShelf";
import { ShelfDetails } from "./domains/shelves/pages/ShelfDetails";
import { ShelvesList } from "./domains/shelves/pages/ShelvesList";

export const router = createBrowserRouter([
  {
    path: "/",
    element: <Layout />,
    children: [
      {
        index: true,
        element: <Home />,
      },
      {
        path: "/books",
        element: <BooksList />,
      },
      {
        path: "/books/:slug",
        element: <BooksDetails />,
      },
      {
        path: "/books/new",
        element: <NewBook />,
      },
      {
        path: "/shelves",
        element: <ShelvesList />,
      },
      {
        path: "/shelves/new",
        element: <NewShelf />,
      },
      {
        path: "/shelves/:slug",
        element: <ShelfDetails />,
      },
    ],
  },
]);
