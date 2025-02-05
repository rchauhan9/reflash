import { SignUp } from '@clerk/clerk-react';


export default function SignUpPage()  {

      return (
          <>
              <SignUp signInUrl="/login" forceRedirectUrl="/" />
          </>
      )
}