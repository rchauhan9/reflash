import { PageHeader, PageHeaderHeading } from '@/components/page-header';
import { Card, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { useUser } from '@clerk/clerk-react';
import { BuiltLineChart } from '@/components/line-chart';

export default function Home() {
  const { user } = useUser();

  return (
    <div>
      <PageHeader>
        <PageHeaderHeading>Welcome back, {user?.firstName}</PageHeaderHeading>
      </PageHeader>

      <div className='grid auto-rows-min gap-4 md:grid-cols-3'>
        <Card className='bg-muted/50'>
          <CardHeader>
            <CardTitle>Study Streak</CardTitle>
            <CardDescription>10 days 🔥</CardDescription>
          </CardHeader>
        </Card>
        <Card className='bg-muted/50'>
          <CardHeader>
            <CardTitle>Questions Answered</CardTitle>
            <CardDescription>{124}</CardDescription>
          </CardHeader>
        </Card>
        <Card className='bg-muted/50'>
          <CardHeader>
            <CardTitle>Mastery Level</CardTitle>
            <CardDescription>⭐️ Intermediate</CardDescription>
          </CardHeader>
        </Card>
      </div>
      <br />
      <BuiltLineChart />
    </div>
  );
}
