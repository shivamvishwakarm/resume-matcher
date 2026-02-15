import { RouterProvider } from "react-router";
import { router } from "./components/Layout/BrowserRouter";
import ThemeProvider from "./components/ThemeProvider/ThemeProvider";

function App() {
  return (
    <ThemeProvider defaultTheme="system" storageKey="vite-ui-theme">
      <RouterProvider router={router} />
    </ThemeProvider>
  );
}

export default App;