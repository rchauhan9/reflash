import {
  QueryClient,
  QueryClientProvider,
} from '@tanstack/react-query';
import { ReactNode } from 'react';
import { ReactQueryDevtools } from '@tanstack/react-query-devtools'
export const queryClient = new QueryClient();

export const QueryProvider = ({ children } : {children: ReactNode}) => (
  <QueryClientProvider client={queryClient}>
    <ReactQueryDevtools initialIsOpen={false} />
    {children}
  </QueryClientProvider>
);
