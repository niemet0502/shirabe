import { NavLink } from "react-router-dom";

export const Topbar: React.FC = () => {
  return (
    <div className="border-1 border-red-500 w-100 flex items-center justify-center">
      <div className="w-full max-w-[1000px] border-2 flex p-0">
        <NavLink to="/" className="p-4 hover:bg-red-700">
          Home
        </NavLink>
        <NavLink to="/books" className="p-4 hover:bg-red-700">
          My books
        </NavLink>
        <NavLink to="/shelves" className="p-4 hover:bg-red-700">
          My Shelves
        </NavLink>
      </div>
    </div>
  );
};
