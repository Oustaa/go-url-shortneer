export const API = {
  baseURL: "/api/v1/",
  // urls
  getUrls: async () => {
    return await API.fetch("urls");
  },
  // auth
  login: async (login, password) => {
    return await API.send("auth/login", {
      login,
      password,
    });
  },
  createAccount: async (username, email, password) => {
    return await API.send("auth/create-account", {
      username,
      email,
      password,
    });
  },

  logIn: async (login, password) => {
    return await API.send("auth/login", {
      login,
      password,
    });
  },

  // helper functions
  send: async (serviceName, data) => {
    try {
      const response = await fetch(API.baseURL + serviceName, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: app.Store.jwt ? `Bearer ${app.Store.jwt}` : null,
        },
        body: JSON.stringify(data),
      });
      const result = await response.json();
      return result;
    } catch (e) {
      console.error(e);
    }
  },
  fetch: async (serviceName, args) => {
    try {
      const queryString = args ? new URLSearchParams(args).toString() : "";
      const response = await fetch(
        API.baseURL + serviceName + "?" + queryString,
        {
          headers: {
            Authorization: app.Store.jwt ? `Bearer ${app.Store.jwt}` : null,
          },
        },
      );
      const result = await response.json();
      return result;
    } catch (e) {
      console.error(e);
    }
  },
};
