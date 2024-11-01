import { Outlet } from "react-router-dom";
import { Topbar } from "./topbar";

export const Layout: React.FC = () => {
  return (
    <div className="p-0 flex flex-col">
      <Topbar />
      <div className="flex-1 w-full max-w-[1000px] self-center border-2">
        <Outlet />
      </div>
    </div>
  );
};
