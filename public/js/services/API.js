export const API = {
  getUrls: async () => {
    return await fetch("/api/v1/urls");
  },
};
