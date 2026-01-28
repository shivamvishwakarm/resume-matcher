import ThemeProvider from "./components/ThemeProvider/ThemeProvider";

import Hero from "./components/LandingPage/Hero";

function App() { 
  return (
    <ThemeProvider defaultTheme="system" storageKey="vite-ui-theme">
      <Hero />
    </ThemeProvider>
  );
}

export default App;