import { createProject, listProjects } from '@/api/study';
import { useMutation, useQuery } from '@tanstack/react-query';

export const useListProjects = () => useQuery({ queryKey: ['projects'], queryFn:listProjects})

export const useCreateProject = () => useMutation(
  {mutationFn: createProject}
)