// app/page.tsx
import { DataEntryForm } from '@/components/data-entry-form';
import { WrenchIcon, ActivityIcon, DatabaseIcon } from 'lucide-react';

export default function Home() {
  return (
    <main className="min-h-screen bg-[radial-gradient(ellipse_at_top_right,_var(--tw-gradient-stops))] from-blue-100 via-slate-50 to-purple-100 dark:from-slate-900 dark:via-blue-950 dark:to-slate-900 p-4 md:p-8 overflow-hidden">
      <div className="absolute inset-0 bg-grid-slate-100 [mask-image:radial-gradient(ellipse_at_center,white,transparent)] dark:bg-grid-slate-700/25 pointer-events-none" />
      <div className="mx-auto max-w-4xl relative">
        <div className="absolute top-[-120px] right-[-120px] w-[300px] h-[300px] bg-primary/20 rounded-full blur-3xl animate-pulse" />
        <div className="absolute bottom-[-120px] left-[-120px] w-[300px] h-[300px] bg-purple-500/20 rounded-full blur-3xl animate-pulse" />
        
        <div className="mb-12 text-center animate-fade-in relative">
          <div className="flex items-center justify-center gap-4 mb-8">
            <div className="relative group">
              <div className="absolute inset-0 rounded-xl bg-gradient-to-r from-primary/60 to-purple-500/60 blur-xl opacity-75 group-hover:opacity-100 transition-all duration-500" />
              <div className="relative bg-white dark:bg-slate-900 p-4 rounded-xl shadow-2xl">
                <WrenchIcon className="h-8 w-8 text-primary animate-pulse" />
              </div>
            </div>
            <div className="relative group">
              <div className="absolute inset-0 rounded-xl bg-gradient-to-r from-purple-500/60 to-blue-500/60 blur-xl opacity-75 group-hover:opacity-100 transition-all duration-500" />
              <div className="relative bg-white dark:bg-slate-900 p-4 rounded-xl shadow-2xl">
                <ActivityIcon className="h-8 w-8 text-purple-500 animate-pulse [animation-delay:200ms]" />
              </div>
            </div>
            <div className="relative group">
              <div className="absolute inset-0 rounded-xl bg-gradient-to-r from-blue-500/60 to-primary/60 blur-xl opacity-75 group-hover:opacity-100 transition-all duration-500" />
              <div className="relative bg-white dark:bg-slate-900 p-4 rounded-xl shadow-2xl">
                <DatabaseIcon className="h-8 w-8 text-blue-500 animate-pulse [animation-delay:400ms]" />
              </div>
            </div>
          </div>
          <div className="relative">
            <h1 className="text-5xl font-bold tracking-tight bg-clip-text text-transparent bg-gradient-to-r from-primary via-purple-500 to-blue-500 pb-2">
              SRE Data Management
            </h1>
            <div className="absolute inset-0 bg-gradient-to-r from-primary/20 via-purple-500/20 to-blue-500/20 blur-3xl -z-10" />
          </div>
          <p className="mt-4 text-muted-foreground text-lg max-w-2xl mx-auto leading-relaxed">
            Streamline your infrastructure updates with our centralized management system.
            Efficiently manage components, maintenance windows, and system updates.
          </p>
        </div>
        <DataEntryForm />
      </div>
    </main>
  );
}