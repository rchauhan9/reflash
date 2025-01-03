import { createBrowserRouter } from "react-router-dom";

import { Applayout } from "./components/layouts/AppLayout";

import NoMatch from "./pages/NoMatch";
import Dashboard from "./pages/Dashboard";
import Work from "./pages/Work";
import Calendar from "./pages/Calendar";
import Notifications from "./pages/Notifications";

export const router = createBrowserRouter([
    {
        path: "/",
        element: <Applayout />,
        children: [
            {
                path: "home",
                element: <Dashboard />,
            },
            {
                path: "work",
                element: <Work />,
            },
            {
                path: "calendar",
                element: <Calendar />,
            },
            {
                path: "notifications",
                element: <Notifications />,
            },
        ],
    },
    {
        path: "*",
        element: <NoMatch />,
    },
], {
    basename: global.basename
})
