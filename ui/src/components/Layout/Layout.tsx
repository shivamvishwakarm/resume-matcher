import { Outlet } from "react-router";
import Navbar from "../LandingPage/Navbar";

const Layout = () => {
    return (
        <div className="min-h-screen flex flex-col">
            <Navbar />
            <main className="flex-1 flex flex-col">
                <Outlet />
            </main>
        </div>
    );
};

export default Layout;