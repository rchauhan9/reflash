import axios from 'axios';


export const listProjects = async () => {
  const { data } = await axios.get('http://localhost:8001/study/projects');
  return data;
};

export const createProject = async (project: any) => {
  const { data } = await axios.post('http://localhost:8001/study/projects', {
      'name': project.name,
      'icon': project.icon,
      'description': project.description,
  });
  return data;
}