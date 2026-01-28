const Hero = () => {
  const isDark = false;

  return (
    <section className="relative pt-32 pb-20 overflow-hidden">
      {/* Background Decorative Diagonal Split */}
      <div className={`absolute inset-0 -z-10 transition-colors duration-500 ${
        isDark ? 'bg-slate-900' : 'bg-white'
      }`}>
        <div className={`absolute inset-0 opacity-10 transition-opacity duration-500 ${
          isDark ? 'bg-[radial-gradient(circle_at_bottom_right,_var(--tw-gradient-stops))] from-violet-500 via-transparent to-transparent' : ''
        }`} />
      </div>

      <div className="max-w-7xl mx-auto px-6 grid lg:grid-cols-2 gap-12 items-center">
        <div>
          <h1 className="text-5xl lg:text-7xl font-bold leading-[1.1] mb-6 tracking-tight text-slate-900">
            Stop Guessing. <br />
            <span className="text-blue-600">
              The Right Resume
            </span> for Every App.
          </h1>
          <p className="text-xl mb-10 max-w-lg leading-relaxed text-slate-600">
            Use AI to instantly match your best-fit resume version against any job description. Increase your interview callback rate with semantic precision.
          </p>
          <div className="flex flex-wrap gap-4">
            <button className="px-8 py-4 rounded-full font-bold text-lg transition-all transform hover:-translate-y-1 bg-blue-600 text-white hover:shadow-xl hover:shadow-blue-200">
              Try the Demo
            </button>
            <button className="px-8 py-4 rounded-full font-bold text-lg border border-slate-200 text-slate-600 hover:bg-slate-50 transition-all">
              Build Your Vault
            </button>
          </div>
        </div>

        {/* Hero Visual */}
        <div
          className="relative"
        >
          <div className="relative p-8 rounded-3xl overflow-hidden bg-slate-50 border border-slate-100">
            <div className="flex items-center justify-between gap-4 h-[300px]">
              {/* Job Description Side */}
              <div className="w-1/3 h-full p-4 rounded-xl flex flex-col gap-2 bg-white border border-blue-100">
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
                    stroke="#2563eb"
                    strokeWidth="2"
                    strokeDasharray="4 4"
                    className="animate-pulse"
                  />
                  <circle cx="10" cy="50" r="4" fill="#2563eb" />
                  <circle cx="190" cy="30" r="4" fill="#2563eb" />
                </svg>
              </div>

              {/* Resume Stack Side */}
              <div className="w-1/3 relative h-full flex flex-col gap-3 justify-center">
                {[1, 2, 3].map((i) => (
                  <div
                    key={i}
                    className={`h-24 p-3 rounded-lg border transition-all duration-500 ${
                      i === 1 
                        ? 'bg-white border-blue-400 shadow-lg'
                        : 'bg-white border-slate-100 opacity-40'
                    }`}
                  >
                    <div className="h-2 w-1/2 rounded mb-2 bg-blue-600/20"></div>
                    <div className="h-1.5 w-full bg-slate-500/10 rounded mb-1"></div>
                    <div className="h-1.5 w-3/4 bg-slate-500/10 rounded"></div>
                    {i === 1 && (
                      <div className="mt-2 text-[8px] font-bold text-blue-600">98% MATCH</div>
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