interface AppConfig {
  name: string;
  github: {
    title: string;
    url: string;
  };
  author: {
    name: string;
    url: string;
  };
}

export const appConfig: AppConfig = {
  name: 'Sample App',
  github: {
    title: 'Reflash',
    url: 'https://github.com/rchauhan9/reflash',
  },
  author: {
    name: 'rchauhan9',
    url: 'https://github.com/rchauhan9/',
  },
};
