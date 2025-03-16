import { useClerk } from '@clerk/clerk-react';
import { Button } from '@/components/ui/button';

export default function SignInButton() {
  const { openSignIn } = useClerk();

  return (
    <Button
      onClick={() => {
        console.log('sign in clicked');
        openSignIn({ forceRedirectUrl: 'https://caring-mastiff-86.clerk.accounts.dev' });
      }}
    >
      Sign in
    </Button>
  );
}
