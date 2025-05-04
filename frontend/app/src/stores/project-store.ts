import { create } from 'zustand';

interface ProjectStore {}

export const useProjectStore = create<ProjectStore>((set) => ({}));
