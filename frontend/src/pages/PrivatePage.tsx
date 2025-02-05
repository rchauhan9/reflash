import { SignedIn, useAuth } from '@clerk/clerk-react';
import { Navigate } from 'react-router';

export default function PrivatePage({ children }: { children: React.ReactNode }) {
  const { isLoaded, isSignedIn } = useAuth()
  console.log("isLoaded", isLoaded)
  console.log("isSignedIn", isSignedIn)

  if (isLoaded && !isSignedIn) {
    return <Navigate to="/login" />
  }

  return (
    <>
    <SignedIn>
      {children}
    </SignedIn>
    </>
  )
}