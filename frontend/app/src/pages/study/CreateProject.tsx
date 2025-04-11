import { PageHeader, PageHeaderHeading } from '@/components/page-header';
import { Card, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import CreateProjectForm from '@/components/forms/CreateProjectForm';
import { cn } from '@/lib/utils';

export default function CreateProject() {
  return (
    <>
      {/* <div className="bg-red-400"> */}
      <PageHeader className='py-0 mx-4'>
        <PageHeaderHeading>Create Project</PageHeaderHeading>
      </PageHeader>
      <Card className='mt-2 mx-4 py-4'>
        <CreateProjectForm />
      </Card>
    </>
  );
}
