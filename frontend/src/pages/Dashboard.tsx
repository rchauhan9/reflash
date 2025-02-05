import { PageHeader, PageHeaderHeading } from "@/components/page-header";
import { Card, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import  SignInButton from '@/components/auth/SignInButton';
import PrivatePage from '@/pages/PrivatePage';

export default function Dashboard() {
    return (
        <PrivatePage>
            <PageHeader>
                <PageHeaderHeading>Dashboard</PageHeaderHeading>
            </PageHeader>
            <Card>
                <CardHeader>
                    <CardTitle>React Shadcn Starter</CardTitle>
                    <CardDescription>React + Vite + TypeScript template for building apps with shadcn/ui.</CardDescription>
                </CardHeader>
            </Card>
        </PrivatePage>
    )
}
