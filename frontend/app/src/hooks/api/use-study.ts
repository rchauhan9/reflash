import { createProject, listProjects } from '@/api/study';
import { useMutation, useQuery } from '@tanstack/react-query';

export const useListProjects = () => {
  return useQuery({ queryKey: ['projects'], queryFn: listProjects });
};

export const useCreateProject = () => useMutation({ mutationFn: createProject });
