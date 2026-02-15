import { createBrowserRouter } from "react-router";
import Layout from "./Layout";
import Hero from "../LandingPage/Hero";

export const router = createBrowserRouter([
    {
        path: "/",
        element: <Layout />,
        children: [
            {
                path: "/",
                element: <Hero />
            }
        ]
    }
]);