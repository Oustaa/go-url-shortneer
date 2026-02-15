import { API } from "./services/API.js";
import { Router } from "./services/Router.js";
import Store from "./services/Store.js";

window.addEventListener("DOMContentLoaded", () => {
  app.Router.init();
});

window.app = {
  API,
  Router,
  Store,
  showError: (message = "There was an error.", goToHome = false) => {
    document.getElementById("alert-modal").showModal();
    document.querySelector("#alert-modal ul").innerHTML = message;
    if (goToHome) app.Router.go("/");
  },
  closeError: () => {
    document.getElementById("alert-modal").close();
  },
  async createAccount(event) {
    event.preventDefault();

    const formData = new FormData(event.target);

    const username = formData.get("username");
    const email = formData.get("email");
    const password = formData.get("password");
    const confirmPassword = formData.get("confirm-password");

    const errors = [];

    if (!username) {
      errors.push("Username is Required");
    } else if (username.trim() <= 4) {
      errors.push("Username is invalid, should be more than 4 letters");
    }

    if (!email) {
      errors.push("Email is Required");
    } else if (!email.includes("@")) {
      errors.push("Email is invalid, should contains @");
    }

    if (!password) {
      errors.push("Password is Required");
    } else if (password.length < 6) {
      errors.push("Password should be 8 charachter or long");
    }

    if (password !== confirmPassword) {
      errors.push("Password Confirmation should be the same as password");
    }

    if (errors.length !== 0) {
      const errorsList = errors.reduce((prev, err) => {
        console.log({ err, prev });
        return prev + `<li>${err}<li>`;
      }, "");

      this.showError(errorsList);
      return;
    }

    const responce = await API.createAccount(username, email, password);

    // if (response.success) {
    // }
    app.Store.jwt = response.jwt;
  },
};
