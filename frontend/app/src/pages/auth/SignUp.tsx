import { FcGoogle } from 'react-icons/fc';
import { useState } from 'react';
import { Button } from '@/components/ui/button';
import { Input } from '@/components/ui/input';

import { useSignUp, useClerk } from '@clerk/clerk-react';

interface SignupProps {
  heading?: string;
  subheading?: string;
  logo: {
    url: string;
    src: string;
    alt: string;
  };
  signupText?: string;
  googleText?: string;
  loginText?: string;
  loginUrl?: string;
}

export default function SignupPage({
  heading = 'Reflash',
  subheading = 'Sign up for free.',
  logo = {
    url: 'https://www.shadcnblocks.com',
    src: 'https://shadcnblocks.com/images/block/block-1.svg',
    alt: 'logo',
  },
  googleText = 'Sign up with Google',
  signupText = 'Create an account',
  loginText = 'Already have an account?',
  loginUrl = '/login',
}: SignupProps) {
  const { signUp, isLoaded } = useSignUp();
  const clerk = useClerk();

  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [, setError] = useState<string | null>(null);

  const handleEmailSignUp = async (event: React.FormEvent) => {
    event.preventDefault();
    if (!isLoaded) return; // ensure the signUp instance is loaded
    try {
      // Create a new user with the provided credentials
      const result = await signUp.create({
        emailAddress: email,
        password: password,
      });
      console.log('User created:', result);
      // You can choose to automatically set the session if desired:
      // await clerk.setActive({ session: result.createdSessionId });
    } catch (err: any) {
      // eslint-disable-line @typescript-eslint/no-explicit-any
      setError(err.errors ? err.errors[0].message : err.toString());
    }
  };

  // Handler for Google OAuth signup
  const handleGoogleSignUp = async () => {
    try {
      // Redirects to Clerk’s OAuth flow for Google signup
      await clerk.redirectToSignUp({
        signUpOptions: {
          // Tell Clerk you want to use the Google OAuth strategy.
          oauthOptions: {
            provider: 'google',
          },
        },
      });
    } catch (err: any) {
      // eslint-disable-line @typescript-eslint/no-explicit-any
      setError(err.errors ? err.errors[0].message : err.toString());
    }
  };

  return (
    <section className='py-32'>
      <div className='container'>
        <div className='flex flex-col gap-4'>
          <div className='mx-auto w-full max-w-sm rounded-md p-6 shadow'>
            <div className='mb-6 flex flex-col items-center'>
              <a href={logo.url}>
                <img src={logo.src} alt={logo.alt} className='mb-7 h-10 w-auto' />
              </a>
              <p className='mb-2 text-2xl font-bold'>{heading}</p>
              <p className='text-muted-foreground'>{subheading}</p>
            </div>
            <div>
              <div className='grid gap-4'>
                <Input
                  type='email'
                  placeholder='Enter your email'
                  required
                  onChange={(e) => setEmail(e.target.value)}
                />
                <div>
                  <Input
                    type='password'
                    placeholder='Enter your password'
                    required
                    onChange={(e) => setPassword(e.target.value)}
                  />
                </div>
                <Button type='submit' className='mt-2 w-full' onClick={handleEmailSignUp}>
                  {signupText}
                </Button>
                <Button variant='outline' className='w-full' onClick={handleGoogleSignUp}>
                  <FcGoogle className='mr-2 size-5' />
                  {googleText}
                </Button>
              </div>
              <div className='mx-auto mt-8 flex justify-center gap-1 text-sm text-muted-foreground'>
                <p>{loginText}</p>
                <a href={loginUrl} className='font-medium text-primary'>
                  Log in
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}

export { SignupPage };
