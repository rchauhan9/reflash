import { useClerk, useSignIn } from '@clerk/clerk-react';
import { FcGoogle } from 'react-icons/fc';

import { Button } from '@/components/ui/button';
import { Checkbox } from '@/components/ui/checkbox';
import { Input } from '@/components/ui/input';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

interface LoginProps {
  heading?: string;
  subheading?: string;
  logo: {
    url: string;
    src: string;
    alt: string;
  };
  loginText?: string;
  googleText?: string;
  signupText?: string;
  signupUrl?: string;
}

export default function LoginPage({
  heading = 'Reflash',
  subheading = 'Welcome back',
  logo = {
    url: 'https://www.shadcnblocks.com',
    src: 'https://www.shadcnblocks.com/images/block/block-1.svg',
    alt: 'logo',
  },
  loginText = 'Log in',
  googleText = 'Log in with Google',
  signupText = "Don't have an account?",
  signupUrl = '/signup',
}: LoginProps) {
  const navigation = useNavigate();
  const { signIn, isLoaded } = useSignIn();
  // Get the clerk object to use redirect methods for OAuth sign-in
  const clerk = useClerk();

  // State for user input and error messages
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [, setError] = useState<string | null>(null);

  // Handler for email/password login
  const handleEmailSignIn = async (event: React.FormEvent) => {
    event.preventDefault();
    if (!isLoaded) return; // Ensure the signIn instance is loaded
    try {
      // Create a sign-in attempt with the provided credentials
      const result = await signIn.create({
        identifier: email, // This is the user's email address
        password,
      });
      // After a successful sign-in attempt, set the session as active
      await clerk.setActive({ session: result.createdSessionId });
      navigation('/');
    } catch (err: any) {
      // eslint-disable-line @typescript-eslint/no-explicit-any
      setError(err.errors ? err.errors[0].message : err.toString());
    }
  };

  // Handler for Google OAuth sign-in
  const handleGoogleSignIn = async () => {
    try {
      // Redirects to Clerk's OAuth flow with Google as the provider
      await clerk.redirectToSignIn({
        signInOptions: {
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
                <div className='flex justify-between'>
                  <div className='flex items-center space-x-2'>
                    <Checkbox id='remember' className='border-muted-foreground' />
                    <label
                      htmlFor='remember'
                      className='text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70'
                    >
                      Remember me
                    </label>
                  </div>
                  <a href='#' className='text-sm text-primary hover:underline'>
                    Forgot password
                  </a>
                </div>
                <Button type='submit' className='mt-2 w-full' onClick={handleEmailSignIn}>
                  {loginText}
                </Button>
                <Button variant='outline' className='w-full' onClick={handleGoogleSignIn}>
                  <FcGoogle className='mr-2 size-5' />
                  {googleText}
                </Button>
              </div>
              <div className='mx-auto mt-8 flex justify-center gap-1 text-sm text-muted-foreground'>
                <p>{signupText}</p>
                <a href={signupUrl} className='font-medium text-primary'>
                  Sign up
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
