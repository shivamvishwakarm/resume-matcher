import { Moon, Sun } from "lucide-react";
import useTheme from "./useTheme";

export default function ModeToggle() {
    const { setTheme, theme } = useTheme();

    const toggleTheme = () => {
        setTheme(theme === "dark" ? "light" : "dark");
    };

    return (
        <button
            onClick={toggleTheme}
            className="relative h-9 w-9 flex items-center justify-center cursor-pointer transition-colors hover:bg-slate-100 dark:hover:bg-white/10 rounded-lg"
            title="Toggle Theme"
            type="button"
        >
            {theme === "dark" ? (
                <Sun size={18} className="text-yellow-300 transition-all dark:rotate-0 rotate-90" />
            ) : (
                <Moon size={18} className="text-slate-700 transition-all rotate-0 dark:-rotate-90" />
            )}
        </button>
    );
}