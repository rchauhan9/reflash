import { PageHeader, PageHeaderHeading } from '@/components/page-header';
import { Card, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { useListProjects } from '@/hooks/api/use-study';

export default function EditProject() {

  const { isPending, isError, data, error } = useListProjects();

  if (isPending) return <div>Loading...</div>;
  if (isError) return <div>Error: {error}</div>;

  return (
    <>
      <PageHeader>
        <PageHeaderHeading>Edit Project</PageHeaderHeading>
      </PageHeader>
      {data['study_projects'].map((project: any) => (
        <Card key={project.id}>
          <CardHeader>
            <CardTitle>{project.name}</CardTitle>
            <CardDescription>{project.icon}</CardDescription>
          </CardHeader>
        </Card>
      ))}
    </>
  );
}
