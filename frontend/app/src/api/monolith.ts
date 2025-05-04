export const createMonolithRequestHeaders = () => {
  return {
    ...createAuthHeader(),
    ...createJSONContentTypeHeader(),
  };
};

export const createAuthHeader = () => {
  const token = getCookie('__session');
  return {
    Authorization: `Bearer ${token}`,
  };
};

export const createJSONContentTypeHeader = () => {
  return {
    'Content-Type': 'application/json',
  };
};

const getCookie = (name: string): string => {
  const cookieValue = document.cookie
    .split('; ')
    .find((row) => row.startsWith(name + '='))
    ?.split('=')[1];

  return cookieValue || '';
};
