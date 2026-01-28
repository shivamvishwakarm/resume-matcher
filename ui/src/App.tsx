import ThemeProvider from "./components/ThemeProvider/ThemeProvider";

import Hero from "./components/LandingPage/Hero";
import Navbar from "./components/LandingPage/Navbar";

function App() { 
  return (
    <ThemeProvider defaultTheme="system" storageKey="vite-ui-theme">
      <Navbar />
      <Hero />
    </ThemeProvider>
  );
}

export default App;