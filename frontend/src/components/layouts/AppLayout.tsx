import { Outlet } from 'react-router-dom';
import { Header } from './Header';
import { Footer } from './Footer';
import { SidebarProvider, SidebarTrigger } from '@/components/ui/sidebar';
import { AppSidebar } from '@/components/layouts/Sidebar';

export function Applayout({ children }: { children: React.ReactNode }) {
  return (
    <>
      <SidebarProvider>
        <AppSidebar />
          <main>
            <SidebarTrigger />
            {/* <Header /> */}

            <div className='flex-grow flex flex-col'>
              <div className='container px-4 md:px-8 flex-grow flex flex-col'>
                <Outlet />
              </div>
            </div>
            <div className='container px-4 md:px-8'>
              {/* <Footer /> */}
            </div>
            {children}
          </main>
      </SidebarProvider>
    </>
  );
}
