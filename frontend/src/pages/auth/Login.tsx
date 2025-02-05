import { SignIn } from '@clerk/clerk-react';


export default function LoginPage()  {

  return (
    <>
      <SignIn signUpUrl="/signup" forceRedirectUrl="/" />
    </>
  )
}