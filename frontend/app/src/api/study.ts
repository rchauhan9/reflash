import axios from 'axios';
import { createMonolithRequestHeaders } from './monolith';

export const listProjects = async () => {
  const { data } = await axios.get('http://localhost:8001/study/projects', {
    headers: createMonolithRequestHeaders(),
  });
  return data;
};

export const createProject = async (project: any) => {
  const { data } = await axios.post(
    'http://localhost:8001/study/projects',
    {
      name: project.name,
      icon: project.icon,
      description: project.description,
    },
    {
      headers: createMonolithRequestHeaders(),
    },
  );
  return data;
};
