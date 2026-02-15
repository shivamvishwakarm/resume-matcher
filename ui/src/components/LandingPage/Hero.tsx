const Hero = () => {
  return (
    <section className="w-full min-h-screen relative pt-32 pb-20 overflow-hidden">
      {/* Background Decorative Diagonal Split */}
      <div className="absolute inset-0 -z-10 transition-colors duration-500 bg-white dark:bg-slate-900">
        <div className="absolute inset-0 opacity-10 transition-opacity duration-500 dark:bg-[radial-gradient(circle_at_bottom_right,var(--tw-gradient-stops))] dark:from-violet-500 dark:via-transparent dark:to-transparent" />
      </div>

      <div className="max-w-7xl mx-auto px-6 grid lg:grid-cols-2 gap-12 items-center">
        <div className="animate-in fade-in slide-in-from-left-5 duration-700">
          <h1 className="text-5xl lg:text-7xl font-bold leading-[1.1] mb-6 tracking-tight text-slate-900 dark:text-white">
            Stop Guessing. <br />
            <span className="text-blue-600 dark:text-transparent dark:bg-clip-text dark:bg-linear-to-r dark:from-cyan-400 dark:to-violet-400">
              The Right Resume
            </span> for Every App.
          </h1>
          <p className="text-xl mb-10 max-w-lg leading-relaxed text-slate-600 dark:text-slate-400">
            Use AI to instantly match your best-fit resume version against any job description. Increase your interview callback rate with semantic precision.
          </p>
          <div className="flex flex-wrap gap-4">
            <button className="px-8 py-4 rounded-lg font-bold text-lg transition-all transform hover:-translate-y-1 bg-blue-600 text-white hover:shadow-xl hover:shadow-blue-200 dark:bg-linear-to-r dark:from-cyan-500 dark:to-violet-500 dark:hover:shadow-none cursor-pointer">
              Try the Demo
            </button>
            <button className="px-8 py-4 rounded-lg font-bold text-lg border transition-all border-slate-200 text-slate-600 hover:bg-slate-50 dark:border-slate-700 dark:text-white dark:hover:bg-slate-800 cursor-pointer">
              Build Your Vault
            </button>
          </div>
        </div>

        {/* Hero Visual Animation */}
        <div className="relative animate-in fade-in zoom-in duration-1000 delay-200">
          <div className="relative p-8 rounded-3xl overflow-hidden bg-slate-50 border border-slate-100 dark:bg-slate-800/50 dark:glassmorphism dark:border-none">
            <div className="flex items-center justify-between gap-4 h-[300px]">
              {/* Job Description Side */}
              <div className="w-1/3 h-full p-4 rounded-xl flex flex-col gap-2 bg-white border border-blue-100 dark:bg-slate-900/80 dark:border-cyan-500/30">
                <div className="h-4 w-3/4 bg-blue-500/20 rounded mb-2"></div>
                <div className="h-3 w-full bg-slate-500/10 rounded"></div>
                <div className="h-3 w-5/6 bg-slate-500/10 rounded"></div>
                <div className="h-3 w-4/6 bg-slate-500/10 rounded"></div>
                <div className="mt-auto h-8 w-full bg-blue-600/10 rounded flex items-center justify-center">
                  <span className="text-[10px] font-bold text-blue-600 uppercase tracking-widest">Job Description</span>
                </div>
              </div>

              {/* Connecting Vector Line */}
              <div className="flex-1 relative h-full flex items-center justify-center">
                <svg className="w-full h-full" viewBox="0 0 200 100">
                  <path
                    d="M 10 50 Q 100 50 190 30"
                    fill="none"
                    strokeWidth="2"
                    strokeDasharray="4 4"
                    className="animate-pulse-line stroke-blue-600 dark:stroke-cyan-400"
                  />
                  <circle cx="10" cy="50" r="4" className="fill-blue-600 dark:fill-cyan-400" />
                  <circle cx="190" cy="30" r="4" className="fill-blue-600 dark:fill-cyan-400" />
                </svg>
              </div>

              {/* Resume Stack Side */}
              <div className="w-1/3 relative h-full flex flex-col gap-3 justify-center">
                {[1, 2, 3].map((i) => (
                  <div
                    key={i}
                    className={`h-24 p-3 rounded-lg border transition-all duration-500 ${i === 1
                      ? 'bg-white border-blue-400 shadow-lg scale-105 -translate-x-2 dark:bg-slate-700 dark:border-cyan-400 dark:shadow-[0_0_20px_rgba(34,211,238,0.2)]'
                      : 'bg-white border-slate-100 opacity-40 dark:bg-slate-800/40 dark:border-slate-700'
                      }`}
                  >
                    <div className={`h-2 w-1/2 rounded mb-2 ${i === 1 ? 'bg-blue-600/20 dark:bg-cyan-400/30' : 'bg-slate-500/10'}`}></div>
                    <div className="h-1.5 w-full bg-slate-500/10 rounded mb-1"></div>
                    <div className="h-1.5 w-3/4 bg-slate-500/10 rounded"></div>
                    {i === 1 && (
                      <div className="mt-2 text-[8px] font-bold text-blue-600 dark:text-cyan-400">98% MATCH</div>
                    )}
                  </div>
                ))}
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
};

export default Hero;