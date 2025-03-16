import { Outlet } from 'react-router-dom';
import { SidebarProvider, SidebarTrigger } from '@/components/ui/sidebar';
import { AppSidebar } from '@/components/layouts/Sidebar';
import PrivatePage from '@/pages/PrivatePage';

export function Applayout({ children }: { children: React.ReactNode }) {
  return (
    <>
      <PrivatePage>
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
            <div className='container px-4 md:px-8'>{/* <Footer /> */}</div>
            {children}
          </main>
        </SidebarProvider>
      </PrivatePage>
    </>
  );
}
