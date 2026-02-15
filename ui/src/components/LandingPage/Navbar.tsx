import { Zap } from 'lucide-react';
import ModeToggle from '../ThemeProvider/ModeToggle';

const Navbar = () => {
    return (
        <nav className='fixed top-0 left-0 right-0 z-50 px-6 py-4 transition-colors duration-300'>
            <div className="max-w-7xl mx-auto flex items-center justify-between">
                <div className="flex items-center gap-2">
                    <div className="bg-blue-600 p-1.5 rounded-lg">
                        <Zap className="w-5 h-5 text-white fill-white" />
                    </div>
                    <span className='text-xl font-bold tracking-tight text-slate-900 dark:text-white'>
                        Resume Matcher
                    </span>
                </div>

                <div className="flex items-center gap-4">
                    <ModeToggle />
                    <button type="button"
                        className='px-4 py-2 text-lg font-bold rounded-lg transition-all active:scale-95 bg-blue-600 text-white hover:bg-blue-700 cursor-pointer'
                    >
                        Sign In
                    </button>
                </div>
            </div>
        </nav>
    );
};

export default Navbar;